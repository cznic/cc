// Copyright 2016 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generate.go
//go:generate golex -o trigraphs.go trigraphs.l
//go:generate golex -o scanner.go scanner.l
//go:generate stringer -type Kind
//go:generate stringer -type Linkage
//go:generate stringer -type Namespace
//go:generate stringer -type Scope
//go:generate go run generate.go -2

// Package cc is a C99 compiler front end.
//
// Links
//
// Referenced from elsewhere:
//
//  [0]: http://www.open-std.org/jtc1/sc22/wg14/www/docs/n1256.pdf
//  [1]: http://www.open-std.org/jtc1/sc22/wg14/www/docs/n1406.pdf
//  [2]: https://github.com/rsc/c2go/blob/fc8cbfad5a47373828c81c7a56cccab8b221d310/cc/cc.y
package cc

import (
	"bufio"
	"bytes"
	"fmt"
	"go/token"
	"os"
	"os/exec"
	"strings"

	"github.com/cznic/golex/lex"
	"github.com/cznic/mathutil"
	"github.com/cznic/xc"
)

const (
	fakeTime = "__TESTING_TIME__"

	gccPredefine = `
#define __attribute__(x)
#define __builtin_offsetof(type, member) ((size_t)(&((type *)0)->member))
#define __builtin_va_arg(ap, type) ( *( type* )ap )
#define __builtin_va_end(x)
#define __builtin_va_list void*
#define __builtin_va_start(x, y)
#define __extension__
#define __PRETTY_FUNCTION__ __func__

double __builtin_inff();
double __builtin_nanf(char *);
int __builtin_constant_p (exp);
int __builtin_ctz (unsigned int x);
int __builtin_ctzl (unsigned long);
int __builtin_ctzll (unsigned long long);
int __builtin_popcount (unsigned int x);
int __builtin_popcountl (unsigned long);
int __builtin_popcountll (unsigned long long);
long __builtin_expect(long, long);
long long strlen (const char*);
unsigned __builtin_bswap32 (unsigned x);
unsigned long long __builtin_bswap64 (unsigned long long x);
unsigned short __builtin_bswap16 (unsigned short x);
void __builtin_unreachable (void);
void __builtin_unwind_init();
void __builtin_va_copy(void* dest, void* src);
void* __builtin_alloca(int);
void* __builtin_memcpy(void *restrict dest, const void *restrict src, long long count);
void* memcpy(void *restrict dest, const void *restrict src, long long count);
void* memset(void*, int, long long);
`
)

// HostConfig executes HostCppConfig with the cpp argument set to "cpp". For
// more info please see the documentation of HostCppConfig.
func HostConfig(opts ...string) (predefined string, includePaths, sysIncludePaths []string, err error) {
	return HostCppConfig("cpp", opts...)
}

// HostCppConfig returns the system C preprocessor configuration, or an error,
// if any.  The configuration is obtained by running the cpp command. For the
// predefined macros list the '-dM' options is added. For the include paths
// lists, the option '-v' is added and the output is parsed to extract the
// "..." include and <...> include paths. To add any other options to cpp, list
// them in opts.
//
// The function relies on a POSIX compatible C preprocessor installed.
// Execution of HostConfig is not free, so caching the results is recommended
// whenever possible.
func HostCppConfig(cpp string, opts ...string) (predefined string, includePaths, sysIncludePaths []string, err error) {
	args := append(append([]string{"-dM"}, opts...), "/dev/null")
	pre, err := exec.Command(cpp, args...).Output()
	if err != nil {
		return "", nil, nil, err
	}

	args = append(append([]string{"-v"}, opts...), "/dev/null")
	out, err := exec.Command(cpp, args...).CombinedOutput()
	if err != nil {
		return "", nil, nil, err
	}

	a := strings.Split(string(out), "\n")
	for i := 0; i < len(a); {
		switch a[i] {
		case "#include \"...\" search starts here:":
		loop:
			for i = i + 1; i < len(a); {
				switch v := a[i]; {
				case strings.HasPrefix(v, "#") || v == "End of search list.":
					break loop
				default:
					includePaths = append(includePaths, strings.TrimSpace(v))
					i++
				}
			}
		case "#include <...> search starts here:":
			for i = i + 1; i < len(a); {
				switch v := a[i]; {
				case strings.HasPrefix(v, "#") || v == "End of search list.":
					return string(pre), includePaths, sysIncludePaths, nil
				default:
					sysIncludePaths = append(sysIncludePaths, strings.TrimSpace(v))
					i++
				}
			}
		default:
			i++
		}
	}
	return "", nil, nil, fmt.Errorf("failed parsing %s -v output", cpp)
}

type tweaks struct {
	comments                       map[token.Pos]int //
	devTest                        bool              //
	disablePredefinedLineMacro     bool              // __LINE__ will not expand.
	enableAlignof                  bool              //
	enableAlternateKeywords        bool              // __asm__ etc.
	enableAnonymousStructFields    bool              //
	enableAsm                      bool              //
	enableEmptyDeclarations        bool              // ; // C++11
	enableDefineOmitCommaBeforeDDD bool              // #define foo(a, b...)
	enableDlrInIdentifiers         bool              // foo$bar
	enableEmptyDefine              bool              // #define
	enableIncludeNext              bool              //
	enableWideBitFieldTypes        bool              // long long v : 2;
	enableNoreturn                 bool              //
	enableStaticAssert             bool              // _Static_assert
	enableTrigraphs                bool              // ??=define foo(bar)
	enableTypeof                   bool              //
	enableUndefExtraTokens         bool              // #undef foo(bar)
	enableWarnings                 bool              // #warning
	enableWideEnumValues           bool              // enum { v = X } for X wider than 32bits
	gccEmu                         bool              //
	preprocessOnly                 bool              //
}

func (t *tweaks) doGccEmu() *tweaks {
	t.enableAlignof = true
	t.enableAlternateKeywords = true
	t.enableAnonymousStructFields = true
	t.enableAsm = true
	t.enableDefineOmitCommaBeforeDDD = true
	t.enableDlrInIdentifiers = true
	t.enableEmptyDefine = true
	t.enableIncludeNext = true
	t.enableNoreturn = true
	t.enableStaticAssert = true
	t.enableTypeof = true
	t.enableUndefExtraTokens = true
	t.enableWarnings = false
	return t
}

func exampleAST(rule int, src string) interface{} {
	report := xc.NewReport()
	report.IgnoreErrors = true
	lx, err := newLexer(
		fmt.Sprintf("example%v.c", rule),
		len(src)+1, // Plus final injected NL
		bytes.NewBufferString(src),
		report,
		(&tweaks{gccEmu: true}).doGccEmu(),
	)
	lx.model = &Model{ // 64 bit
		Items: map[Kind]ModelItem{
			Ptr:               {8, 8, 8, nil},
			Void:              {0, 1, 1, nil},
			Char:              {1, 1, 1, nil},
			SChar:             {1, 1, 1, nil},
			UChar:             {1, 1, 1, nil},
			Short:             {2, 2, 2, nil},
			UShort:            {2, 2, 2, nil},
			Int:               {4, 4, 4, nil},
			UInt:              {4, 4, 4, nil},
			Long:              {8, 8, 8, nil},
			ULong:             {8, 8, 8, nil},
			LongLong:          {8, 8, 8, nil},
			ULongLong:         {8, 8, 8, nil},
			Float:             {4, 4, 4, nil},
			Double:            {8, 8, 8, nil},
			LongDouble:        {8, 8, 8, nil},
			Bool:              {1, 1, 1, nil},
			FloatComplex:      {8, 8, 8, nil},
			DoubleComplex:     {16, 16, 16, nil},
			LongDoubleComplex: {16, 16, 16, nil},
		},
		tweaks: &tweaks{},
	}

	lx.model.initialize(lx)
	if err != nil {
		panic(err)
	}

	lx.exampleRule = rule
	yyParse(lx)
	return lx.example
}

func ppParseString(fn, src string, report *xc.Report, tweaks *tweaks) (*PreprocessingFile, error) {
	sz := len(src)
	lx, err := newLexer(fn, int(sz)+1, bytes.NewBufferString(src), report, tweaks)
	if err != nil {
		return nil, err
	}

	lx.Unget(lex.NewChar(token.Pos(lx.File.Base()), PREPROCESSING_FILE))
	yyParse(lx)
	return lx.preprocessingFile, nil
}

func ppParse(fn string, report *xc.Report, tweaks *tweaks) (*PreprocessingFile, error) {
	o := xc.Files.Once(fn, func() interface{} {
		f, err := os.Open(fn)
		if err != nil {
			return err
		}

		defer f.Close()

		fi, err := os.Stat(fn)
		if err != nil {
			return err
		}

		sz := fi.Size()
		if sz > mathutil.MaxInt-1 {
			return fmt.Errorf("%s: file size too big: %v", fn, sz)
		}

		lx, err := newLexer(fn, int(sz)+1, bufio.NewReader(f), report, tweaks)
		if err != nil {
			return err
		}

		lx.Unget(lex.NewChar(token.Pos(lx.File.Base()), PREPROCESSING_FILE))
		if yyParse(lx) != 0 {
			return report.Errors(true)
		}

		return lx.preprocessingFile
	})
	switch r := o.Value(); x := r.(type) {
	case error:
		return nil, x
	case *PreprocessingFile:
		return x, nil
	default:
		panic("internal error")
	}
}

// Opt is a configuration/setup function that can be passed to the Parser
// function.
type Opt func(*lexer)

// KeepComments makes the parser keep comments.
func KeepComments() Opt {
	return func(l *lexer) { l.tweaks.comments = map[token.Pos]int{} }
}

// EnableAnonymousStructFields makes the parser accept non standard
//
//	struct {
//		int i;
//		struct {
//			int j;
//		};
//		int k;
//	};
func EnableAnonymousStructFields() Opt {
	return func(l *lexer) { l.tweaks.enableAnonymousStructFields = true }
}

// EnableEmptyDeclarations makes the parser accept non standard
//
//	; // C++11 empty declaration
func EnableEmptyDeclarations() Opt {
	return func(l *lexer) { l.tweaks.enableEmptyDeclarations = true }
}

// EnableIncludeNext makes the parser accept non standard
//
//	#include_next "foo.h"
func EnableIncludeNext() Opt {
	return func(l *lexer) { l.tweaks.enableIncludeNext = true }
}

// EnableDefineOmitCommaBeforeDDD makes the parser accept non standard
//
//	#define foo(a, b...) // Note the missing comma after identifier list.
func EnableDefineOmitCommaBeforeDDD() Opt {
	return func(l *lexer) { l.tweaks.enableDefineOmitCommaBeforeDDD = true }
}

// EnableAlternateKeywords makes the parser accept, for example, non standard
//
//	__asm__
//
// as an equvalent of keyowrd asm (which first hast be permitted by EnableAsm).
func EnableAlternateKeywords() Opt {
	return func(l *lexer) { l.tweaks.enableAlternateKeywords = true }
}

// EnableDlrInIdentifiers makes the parser accept non standard
//
//	int foo$bar
func EnableDlrInIdentifiers() Opt {
	return func(l *lexer) { l.tweaks.enableDlrInIdentifiers = true }
}

// EnableEmptyDefine makes the parser accept non standard
//
//	#define
func EnableEmptyDefine() Opt {
	return func(l *lexer) { l.tweaks.enableEmptyDefine = true }
}

// EnableUndefExtraTokens makes the parser accept non standard
//
//	#undef foo(bar)
func EnableUndefExtraTokens() Opt {
	return func(l *lexer) { l.tweaks.enableUndefExtraTokens = true }
}

// EnableWideEnumValues makes the parser accept non standard
//
//	enum { v = X }; for X wider than 32 bits.
func EnableWideEnumValues() Opt {
	return func(l *lexer) { l.tweaks.enableWideEnumValues = true }
}

// EnableWideBitFieldTypes makes the parser accept non standard bitfield
// types (i.e, long long and unsigned long long).
//
//	unsigned long long bits : 2;
func EnableWideBitFieldTypes() Opt {
	return func(l *lexer) { l.tweaks.enableWideBitFieldTypes = true }
}

// SysIncludePaths option configures where to search for system include files
// (eg. <name.h>). Multiple SysIncludePaths options may be used, the resulting
// search path list is produced by appending the option arguments in order of
// appearance.
func SysIncludePaths(paths []string) Opt {
	return func(l *lexer) {
		var err error
		if l.sysIncludePaths, err = dedupAbsPaths(append(l.sysIncludePaths, fromSlashes(paths)...)); err != nil {
			l.report.Err(0, "synIncludepaths option: %v", err)
		}
		l.sysIncludePaths = l.sysIncludePaths[:len(l.sysIncludePaths):len(l.sysIncludePaths)]
	}
}

// IncludePaths option configures where to search for include files (eg.
// "name.h").  Multiple IncludePaths options may be used, the resulting search
// path list is the produce by appending the option arguments in order of
// appearance.
func IncludePaths(paths []string) Opt {
	return func(l *lexer) {
		var err error
		if l.includePaths, err = dedupAbsPaths(append(l.includePaths, fromSlashes(paths)...)); err != nil {
			l.report.Err(0, "includepaths option: %v", err)
		}
		l.includePaths = l.includePaths[:len(l.includePaths):len(l.includePaths)]
	}
}

// YyDebug sets the parser debug level.
func YyDebug(n int) Opt {
	return func(*lexer) { yyDebug = n }
}

// Cpp registers a preprocessor hook function which is called for every line,
// or group of lines the preprocessor produces before it is consumed by the
// parser. The token slice must not be modified by the hook.
func Cpp(f func([]xc.Token)) Opt {
	return func(lx *lexer) { lx.cpp = f }
}

// ErrLimit limits the number of calls to the error reporting methods.  After
// the limit is reached, all errors are reported using log.Print and then
// log.Fatal() is called with a message about too many errors.  To disable
// error limit, set ErrLimit to value less or equal zero.  Default value is 10.
func ErrLimit(n int) Opt {
	return func(lx *lexer) { lx.report.ErrLimit = n }
}

// Trigraphs enables processing of trigraphs.
func Trigraphs() Opt { return func(lx *lexer) { lx.tweaks.enableTrigraphs = true } }

// EnableAsm enables recognizing the reserved word asm.
func EnableAsm() Opt { return func(lx *lexer) { lx.tweaks.enableAsm = true } }

// EnableNoreturn enables recognizing the reserved word _Noreturn.
func EnableNoreturn() Opt { return func(lx *lexer) { lx.tweaks.enableNoreturn = true } }

// EnableTypeOf enables recognizing the reserved word typeof.
func EnableTypeOf() Opt { return func(lx *lexer) { lx.tweaks.enableTypeof = true } }

// EnableAlignOf enables recognizing the reserved word _Alignof.
func EnableAlignOf() Opt { return func(lx *lexer) { lx.tweaks.enableAlignof = true } }

// EnableStaticAssert enables recognizing the reserved word _Static_assert.
func EnableStaticAssert() Opt { return func(lx *lexer) { lx.tweaks.enableStaticAssert = true } }

// CrashOnError is an debugging option.
func CrashOnError() Opt { return func(lx *lexer) { lx.report.PanicOnError = true } }

func disableWarnings() Opt      { return func(lx *lexer) { lx.tweaks.enableWarnings = false } }
func gccEmu() Opt               { return func(lx *lexer) { lx.tweaks.gccEmu = true } }
func getTweaks(dst *tweaks) Opt { return func(lx *lexer) { *dst = *lx.tweaks } }
func nopOpt() Opt               { return func(*lexer) {} }
func preprocessOnly() Opt       { return func(lx *lexer) { lx.tweaks.preprocessOnly = true } }

func devTest() Opt {
	return func(lx *lexer) { lx.tweaks.devTest = true }
}

func disablePredefinedLineMacro() Opt {
	return func(lx *lexer) { lx.tweaks.disablePredefinedLineMacro = true }
}

// Parse defines any macros in predefine. Then Parse preprocesses and parses
// the translation unit consisting of files in paths. The m communicates the
// scalar types model and opts allow to amend parser behavior. m cannot be
// reused and passed to Parse again.
func Parse(predefine string, paths []string, m *Model, opts ...Opt) (*TranslationUnit, error) {
	if m == nil {
		return nil, fmt.Errorf("invalid nil model passed")
	}

	if m.initialized {
		return nil, fmt.Errorf("invalid reused model passed")
	}

	fromSlashes(paths)
	report := xc.NewReport()
	lx0 := &lexer{tweaks: &tweaks{enableWarnings: true}, report: report}
	for _, opt := range opts {
		opt(lx0)
	}
	m.tweaks = lx0.tweaks
	if err := report.Errors(true); err != nil {
		return nil, err
	}

	if lx0.tweaks.devTest {
		predefine += fmt.Sprintf(`
#define __DATE__ %q
#define __TIME__ %q
`, xc.Dict.S(idTDate), fakeTime)
	}

	if t := lx0.tweaks; t.gccEmu {
		t.doGccEmu()
		predefine += gccPredefine
	}

	m.initialize(lx0)
	if err := m.sanityCheck(); err != nil {
		report.Err(0, "%s", err.Error())
		return nil, report.Errors(true)
	}

	tweaks := lx0.tweaks
	predefined, err := ppParseString("<predefine>", predefine, report, tweaks)
	if err != nil {
		return nil, err
	}

	ch := make(chan []xc.Token, 1000)
	macros := newMacros()
	stop := make(chan int, 1)
	go func() {
		defer close(ch)

		newPP(ch, lx0.includePaths, lx0.sysIncludePaths, macros, false, m, report, tweaks).preprocessingFile(predefined)
		for _, path := range paths {
			select {
			case <-stop:
				return
			default:
			}
			pf, err := ppParse(path, report, tweaks)
			if err != nil {
				report.Err(0, err.Error())
				return
			}

			newPP(ch, lx0.includePaths, lx0.sysIncludePaths, macros, true, m, report, tweaks).preprocessingFile(pf)
		}
	}()

	if err := report.Errors(true); err != nil { // Do not parse if preprocessing already failed.
		go func() {
			for range ch { // Drain.
			}
		}()
		stop <- 1
		return nil, err
	}

	lx := newSimpleLexer(lx0.cpp, report, tweaks)
	lx.ch = ch
	lx.state = lsTranslationUnit0
	lx.model = m
	if lx.tweaks.preprocessOnly {
		var lval yySymType
		for lval.Token.Rune != lex.RuneEOF {
			lx.Lex(&lval)
		}
		return nil, report.Errors(true)
	}

	yyParse(lx)
	stop <- 1
	for range ch { // Drain.
	}
	if tu := lx.translationUnit; tu != nil {
		tu.Macros = macros.macros()
		tu.Model = m
		tu.Comments = lx0.tweaks.comments
	}
	return lx.translationUnit, report.Errors(true)
}
