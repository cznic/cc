// Copyright 2016 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/cznic/mathutil"
	"github.com/cznic/xc"
)

type (
	// StringLitID is the type of an Expression.Value representing the numeric
	// ID of a string literal.
	StringLitID int

	// LongStringLitID is the type of an Expression.Value representing the
	// numeric ID of a long string literal.
	LongStringLitID int
)

// ModelItem is a single item of a model.
//
// Note about StructAlign: To provide GCC ABI compatibility set, for example,
// Align of Double to 8 and StructAlign of Double to 4.
type ModelItem struct {
	Size        int         // Size of the entity in bytes.
	Align       int         // Alignment of the entity when it's not a struct field. Cannot be smaller than Size.
	StructAlign int         // Alignment of the entity when it's a struct field. Can be smaller than Size.
	More        interface{} // Optional user data.
}

// Model describes size and align requirements of predeclared types.
type Model struct {
	Items map[Kind]ModelItem

	BoolType       Type
	CharType       Type
	DoubleType     Type
	FloatType      Type
	IntType        Type
	LongDoubleType Type
	LongLongType   Type
	LongType       Type
	ShortType      Type
	UCharType      Type
	UIntType       Type
	ULongLongType  Type
	ULongType      Type
	UShortType     Type
	UintPtrType    Type
	VoidType       Type
	longStrType    Type
	ptrDiffType    Type
	sizeType       Type
	strType        Type

	initialized bool
}

func (m *Model) initialize(lx *lexer) {
	m.BoolType = m.makeType(lx, 0, tsBool)
	m.CharType = m.makeType(lx, 0, tsChar)
	m.DoubleType = m.makeType(lx, 0, tsDouble)
	m.FloatType = m.makeType(lx, 0, tsFloat)
	m.IntType = m.makeType(lx, 0, tsInt)
	m.LongDoubleType = m.makeType(lx, 0, tsLong, tsDouble)
	m.LongLongType = m.makeType(lx, 0, tsLong, tsLong)
	m.LongType = m.makeType(lx, 0, tsLong)
	m.ShortType = m.makeType(lx, 0, tsShort)
	m.UCharType = m.makeType(lx, 0, tsUnsigned, tsChar)
	m.UIntType = m.makeType(lx, 0, tsUnsigned)
	m.ULongLongType = m.makeType(lx, 0, tsUnsigned, tsLong, tsLong)
	m.ULongType = m.makeType(lx, 0, tsUnsigned, tsLong)
	m.UShortType = m.makeType(lx, 0, tsUnsigned, tsShort)
	m.UintPtrType = m.makeType(lx, 0, tsUintptr) // Pseudo type.
	m.VoidType = m.makeType(lx, 0, tsVoid)
	m.strType = m.makeType(lx, 0, tsChar).Pointer()
	m.initialized = true
}

func (m *Model) toInt(v interface{}) (interface{}, bool) {
	switch x := v.(type) {
	case byte, int8, int16, uint16, int32:
		return m.MustConvert(x, m.IntType), true
	case uint32:
		switch sz := m.Items[Int].Size; sz {
		case 4:
			return m.MustConvert(x, m.IntType), x <= math.MaxInt32
		case 8:
			return m.MustConvert(x, m.IntType), true
		default:
			panic(sz)
		}
	case int64:
		switch sz := m.Items[Int].Size; sz {
		case 4:
			return m.MustConvert(x, m.IntType), x <= math.MaxInt32
		case 8:
			return m.MustConvert(x, m.IntType), true
		default:
			panic(sz)
		}
	case uint64:
		switch sz := m.Items[Int].Size; sz {
		case 4:
			return m.MustConvert(x, m.IntType), x <= math.MaxInt32
		case 8:
			return m.MustConvert(x, m.IntType), x <= math.MaxInt64
		default:
			panic(sz)
		}
	default:
		panic(fmt.Errorf("%T", x))
	}
}

// sanityCheck reports model errors, if any.
func (m *Model) sanityCheck() error {
	if len(m.Items) == 0 {
		return fmt.Errorf("model has no items")
	}

	tab := map[Kind]struct {
		minSize, maxSize   int
		minAlign, maxAlign int
	}{
		Ptr:               {4, 8, 4, 8},
		UintPtr:           {4, 8, 4, 8},
		Void:              {0, 0, 1, 1},
		Char:              {1, 1, 1, 1},
		SChar:             {1, 1, 1, 1},
		UChar:             {1, 1, 1, 1},
		Short:             {2, 2, 2, 2},
		UShort:            {2, 2, 2, 2},
		Int:               {4, 4, 4, 4},
		UInt:              {4, 4, 4, 4},
		Long:              {4, 8, 4, 8},
		ULong:             {4, 8, 4, 8},
		LongLong:          {8, 8, 8, 8},
		ULongLong:         {8, 8, 8, 8},
		Float:             {4, 4, 4, 4},
		Double:            {8, 8, 8, 8},
		LongDouble:        {8, 8, 8, 8},
		Bool:              {1, 1, 1, 1},
		FloatComplex:      {8, 8, 8, 8},
		DoubleComplex:     {16, 16, 16, 16},
		LongDoubleComplex: {16, 16, 16, 16},
	}
	a := []int{}
	required := map[Kind]bool{}
	seen := map[Kind]bool{}
	for k := range tab {
		required[k] = true
		a = append(a, int(k))
	}
	sort.Ints(a)
	for k, v := range m.Items {
		if seen[k] {
			return fmt.Errorf("model has duplicate item: %s", k)
		}

		seen[k] = true
		if !required[k] {
			return fmt.Errorf("model has invalid type: %s: %#v", k, v)
		}

		for typ, t := range tab {
			if typ == k {
				if v.Size < t.minSize {
					return fmt.Errorf("size %d too small: %s", v.Size, k)
				}

				if v.Size > t.maxSize {
					return fmt.Errorf("size %d too big: %s", v.Size, k)
				}

				if v.Size != 0 && mathutil.PopCount(v.Size) != 1 {
					return fmt.Errorf("size %d is not a power of two: %s", v.Size, k)
				}

				if v.Align < t.minAlign {
					return fmt.Errorf("align %d too small: %s", v.Align, k)
				}

				if v.Align > t.maxAlign {
					return fmt.Errorf("align %d too big: %s", v.Align, k)
				}

				if v.Align < v.Size {
					return fmt.Errorf("align is smaller than size: %s", k)
				}

				if v.StructAlign < 1 {
					return fmt.Errorf("struct align %d too small: %s", v.StructAlign, k)
				}

				if v.StructAlign > t.maxAlign {
					return fmt.Errorf("struct align %d too big: %s", v.Align, k)
				}

				if mathutil.PopCount(v.Align) != 1 {
					return fmt.Errorf("align %d is not a power of two: %s", v.Align, k)
				}

				break
			}
		}
	}
	w := m.Items[Ptr].Size
	if m.Items[Short].Size < w &&
		m.Items[Int].Size < w &&
		m.Items[Long].Size < w &&
		m.Items[LongLong].Size < w {
		return fmt.Errorf("model has no integer type suitable for pointer difference and sizeof")
	}

	for _, typ := range a {
		if !seen[Kind(typ)] {
			return fmt.Errorf("model has no item for type %s", Kind(typ))
		}
	}

	if g, e := w, m.Items[UintPtr].Size; g != e {
		return fmt.Errorf("model uintptr has different sizes than ptr has")
	}
	return nil
}

// MustConvert returns v converted to the type of typ. If the conversion is
// impossible, the method panics.
//
// Conversion an integer type to any pointer type yields an uintptr.
func (m *Model) MustConvert(v interface{}, typ Type) interface{} {
	if typ.Kind() == Enum {
		typ = m.IntType
	}
	mi, ok := m.Items[typ.Kind()]
	if !ok && typ.Kind() != Function {
		panic(fmt.Errorf("internal error: no model item for %s, %s", typ, typ.Kind()))
	}

	w := mi.Size
	switch typ.Kind() {
	case Short:
		switch x := v.(type) {
		case int32:
			switch w {
			case 2:
				return int16(x)
			default:
				panic(w)
			}
		case int64:
			switch w {
			case 2:
				return int16(x)
			default:
				panic(w)
			}
		default:
			panic(fmt.Errorf("internal error %T", x))
		}
	case UShort:
		switch x := v.(type) {
		case uint16:
			switch w {
			case 2:
				return x
			default:
				panic(w)
			}
		case int32:
			switch w {
			case 2:
				return uint16(x)
			default:
				panic(w)
			}
		default:
			panic(fmt.Errorf("internal error %T", x))
		}
	case Int:
		switch x := v.(type) {
		case byte:
			switch w {
			case 4:
				return int32(x)
			default:
				panic(w)
			}
		case int32:
			switch w {
			case 4:
				return x
			default:
				panic(w)
			}
		case uint32:
			switch w {
			case 4:
				return int32(x)
			default:
				panic(w)
			}
		case int64:
			switch w {
			case 4:
				return int32(x)
			default:
				panic(w)
			}
		case uint64:
			switch w {
			case 4:
				return int32(x)
			default:
				panic(w)
			}
		default:
			panic(fmt.Errorf("internal error %T", x))
		}
	case UInt:
		switch x := v.(type) {
		case uint8:
			switch w {
			case 4:
				return uint32(x)
			default:
				panic(w)
			}
		case uint16:
			switch w {
			case 4:
				return uint32(x)
			default:
				panic(w)
			}
		case int32:
			switch w {
			case 4:
				return uint32(x)
			default:
				panic(w)
			}
		case uint32:
			switch w {
			case 4:
				return x
			default:
				panic(w)
			}
		case int64:
			switch w {
			case 4:
				return uint32(x)
			default:
				panic(w)
			}
		case uint64:
			switch w {
			case 4:
				return uint32(x)
			default:
				panic(w)
			}
		case uintptr:
			switch w {
			case 4:
				return uint32(x)
			default:
				panic(w)
			}
		case float64:
			switch w {
			case 4:
				return uint32(x)
			default:
				panic(w)
			}
		default:
			panic(fmt.Errorf("internal error %T", x))
		}
	case Long:
		switch x := v.(type) {
		case int32:
			switch w {
			case 8:
				return int64(x)
			default:
				panic(w)
			}
		case uint32:
			switch w {
			case 8:
				return int64(x)
			default:
				panic(w)
			}
		case int64:
			switch w {
			case 8:
				return x
			default:
				panic(w)
			}
		case uint64:
			switch w {
			case 4:
				return int32(x)
			case 8:
				return int64(x)
			default:
				panic(w)
			}
		default:
			panic(fmt.Errorf("internal error %T", x))
		}
	case LongLong:
		switch x := v.(type) {
		case int32:
			switch w {
			case 8:
				return int64(x)
			default:
				panic(w)
			}
		case int64:
			switch w {
			case 8:
				return x
			default:
				panic(w)
			}
		case uint64:
			switch w {
			case 8:
				return int64(x)
			default:
				panic(w)
			}
		default:
			panic(fmt.Errorf("internal error %T", x))
		}
	case ULong:
		switch x := v.(type) {
		case int:
			switch w {
			case 8:
				return uint64(x)
			default:
				panic(w)
			}
		case int32:
			switch w {
			case 8:
				return uint64(x)
			default:
				panic(w)
			}
		case int64:
			switch w {
			case 8:
				return uint64(x)
			default:
				panic(w)
			}
		case uint32:
			switch w {
			case 8:
				return uint64(x)
			default:
				panic(w)
			}
		case uint64:
			switch w {
			case 8:
				return x
			default:
				panic(w)
			}
		case uintptr:
			switch w {
			case 8:
				return uint64(x)
			default:
				panic(w)
			}
		default:
			panic(fmt.Errorf("internal error %T", x))
		}
	case ULongLong:
		switch x := v.(type) {
		case int32:
			switch w {
			case 8:
				return uint64(x)
			default:
				panic(w)
			}
		case int64:
			switch w {
			case 8:
				return uint64(x)
			default:
				panic(w)
			}
		case uint32:
			switch w {
			case 8:
				return uint64(x)
			default:
				panic(w)
			}
		case uint64:
			switch w {
			case 8:
				return x
			default:
				panic(w)
			}
		case uintptr:
			switch w {
			case 8:
				return uint64(x)
			default:
				panic(w)
			}
		default:
			panic(fmt.Errorf("internal error %T", x))
		}
	case Float:
		switch x := v.(type) {
		case float32:
			switch w {
			case 4:
				return float32(x)
			case 8:
				return float64(x)
			default:
				panic(w)
			}
		case float64:
			switch w {
			case 4:
				return float32(x)
			case 8:
				return float64(x)
			default:
				panic(w)
			}
		default:
			panic(fmt.Errorf("internal error %T", x))
		}
	case Double:
		switch x := v.(type) {
		case int32:
			switch w {
			case 8:
				return float64(x)
			default:
				panic(w)
			}
		case int64:
			switch w {
			case 8:
				return float64(x)
			default:
				panic(w)
			}
		case uint64:
			switch w {
			case 8:
				return float64(x)
			default:
				panic(w)
			}
		case float64:
			switch w {
			case 8:
				return x
			default:
				panic(w)
			}
		default:
			panic(fmt.Errorf("internal error %T", x))
		}
	case Ptr, Function:
		switch x := v.(type) {
		case int32:
			return uintptr(x)
		case uint32:
			return uintptr(x)
		case int64:
			return uintptr(x)
		case uint64:
			return uintptr(x)
		case uintptr:
			return x
		case StringLitID:
			return nil
		default:
			panic(fmt.Errorf("internal error %T", x))
		}
	case Void:
		return nil
	case Char, SChar:
		switch x := v.(type) {
		case int32:
			switch w {
			case 1:
				return int8(x)
			default:
				panic(w)
			}
		default:
			panic(fmt.Errorf("internal error %T", x))
		}
	case UChar:
		switch x := v.(type) {
		case uint8:
			switch w {
			case 1:
				return byte(x)
			default:
				panic(w)
			}
		case int32:
			switch w {
			case 1:
				return byte(x)
			default:
				panic(w)
			}
		default:
			panic(fmt.Errorf("internal error %T", x))
		}
	case UintPtr:
		switch x := v.(type) {
		case int32:
			switch w {
			case 4, 8:
				return uintptr(x)
			default:
				panic(w)
			}
		case uint32:
			switch w {
			case 4:
				return uintptr(x)
			default:
				panic(w)
			}
		case uint64:
			switch w {
			case 8:
				return uintptr(x)
			default:
				panic(w)
			}
		default:
			panic(fmt.Errorf("internal error %T", x))
		}
	case LongDouble:
		switch x := v.(type) {
		case float64:
			switch w {
			case 8:
				return x
			default:
				panic(w)
			}
		default:
			panic(fmt.Errorf("internal error %T", x))
		}
	default:
		panic(fmt.Errorf("internal error %s, %s", typ, typ.Kind()))
	}
}

func (m *Model) value2(v interface{}, typ Type) (interface{}, Type) {
	return m.MustConvert(v, typ), typ
}

func (m *Model) charConst(t xc.Token) (interface{}, Type) {
	s := string(t.S())
	typ := m.IntType
	var r rune
	switch t.Rune {
	case LONGCHARCONST:
		typ = m.LongType
		s = s[1:] // Remove leading 'L'.
		fallthrough
	case CHARCONST:
		s = s[1 : len(s)-1] // Remove outer 's.
		if len(s) == 1 {
			return rune(s[0]), m.IntType
		}

		runes := []rune(string(s))
		switch runes[0] {
		case '\\':
			r, _ = decodeEscapeSequence(runes)
		default:
			r = runes[0]
		}
	default:
		panic("internal error")
	}
	return r, typ
}

func (m *Model) getSizeType(lx *lexer, tok xc.Token) Type {
	if t := m.sizeType; t != nil {
		return t
	}

	b := lx.scope.Lookup(NSIdentifiers, xc.Dict.SID("size_t"))
	if b.Node == nil {
		w := m.Items[Ptr].Size
		switch {
		case m.Items[Short].Size >= w:
			return m.ShortType
		case m.Items[Int].Size >= w:
			return m.IntType
		case m.Items[Long].Size >= w:
			return m.LongType
		default:
			return m.LongLongType
		}
	}

	d := b.Node.(*DirectDeclarator)
	if !d.TopDeclarator().RawSpecifier().IsTypedef() {
		lx.report.Err(d.Pos(), "size_t is not a typedef name")
		m.sizeType = undefined
		return undefined
	}

	m.sizeType = b.Node.(*DirectDeclarator).top().declarator.Type
	return m.sizeType
}

func (m *Model) getPtrDiffType(lx *lexer, tok xc.Token) Type {
	if t := m.ptrDiffType; t != nil {
		return t
	}

	b := lx.scope.Lookup(NSIdentifiers, xc.Dict.SID("ptrdiff_t"))
	if b.Node == nil {
		w := m.Items[Ptr].Size
		switch {
		case m.Items[Short].Size >= w:
			return m.ShortType
		case m.Items[Int].Size >= w:
			return m.IntType
		case m.Items[Long].Size >= w:
			return m.LongType
		default:
			return m.LongLongType
		}
	}

	d := b.Node.(*DirectDeclarator)
	if !d.TopDeclarator().RawSpecifier().IsTypedef() {
		lx.report.Err(d.Pos(), "ptrdiff_t is not a typedef name")
		m.ptrDiffType = undefined
		return undefined
	}

	m.ptrDiffType = b.Node.(*DirectDeclarator).top().declarator.Type
	return m.ptrDiffType
}

func (m *Model) getLongStrType(lx *lexer, tok xc.Token) Type {
	if t := m.longStrType; t != nil {
		return t
	}

	b := lx.scope.Lookup(NSIdentifiers, xc.Dict.SID("wchar_t"))
	if b.Node == nil {
		lx.report.ErrTok(tok, "undefined: wchar_t")
		m.longStrType = undefined
		return m.longStrType
	}

	d := b.Node.(*DirectDeclarator)
	if !d.TopDeclarator().RawSpecifier().IsTypedef() {
		lx.report.Err(d.Pos(), "wchar_t is not a typedef name")
		m.longStrType = undefined
		return m.longStrType
	}

	m.longStrType = b.Node.(*DirectDeclarator).top().declarator.Type.Pointer()
	return m.longStrType
}

func (m *Model) strConst(lx *lexer, t xc.Token) (interface{}, Type) {
	s := t.S()
	typ := m.strType
	var buf bytes.Buffer
	switch t.Rune {
	case LONGSTRINGLITERAL:
		typ = m.getLongStrType(lx, t)
		s = s[1:] // Remove leading 'L'.
		fallthrough
	case STRINGLITERAL:
		s = s[1 : len(s)-1] // Remove outer "s.
		runes := []rune(string(s))
		for i := 0; i < len(runes); {
			switch r := runes[i]; {
			case r == '\\':
				r, n := decodeEscapeSequence(runes[i:])
				buf.WriteRune(r)
				i += n
			default:
				buf.WriteByte(byte(r))
				i++
			}
		}
	default:
		panic("internal error")
	}
	s = buf.Bytes()
	switch t.Rune {
	case LONGSTRINGLITERAL:
		return LongStringLitID(xc.Dict.ID(s)), typ
	case STRINGLITERAL:
		return StringLitID(xc.Dict.ID(s)), typ
	default:
		panic("internal error")
	}
}

func (m *Model) floatConst(lx *lexer, t xc.Token) (interface{}, Type) {
	s0 := t.S()
	s := s0
	switch s[len(s)-1] {
	case 'l', 'L', 'f', 'F':
		s = s[:len(s)-1]
	}
	f, err := strconv.ParseFloat(string(s), 64)
	if err != nil {
		lx.report.Err(t.Pos(), "invalid floating point constant")
		f = 0
	}
	switch s0[len(s0)-1] {
	case 'l', 'L':
		return m.value2(f, m.LongDoubleType)
	case 'f', 'F':
		return m.value2(f, m.FloatType)
	default:
		return m.value2(f, m.DoubleType)
	}
}

func (m *Model) intConst(lx *lexer, t xc.Token) (interface{}, Type) {
	const (
		l = 1 << iota
		ll
		u
	)
	k := 0
	s := t.S()
	i := len(s) - 1
more:
	switch c := s[i]; c {
	case 'u', 'U':
		k |= u
		i--
		goto more
	case 'l', 'L':
		if i > 0 && (s[i-1] == 'l' || s[i-1] == 'L') {
			k |= ll
			i -= 2
			goto more
		}

		k |= l
		i--
		goto more
	}
	n, err := strconv.ParseUint(string(s[:i+1]), 0, 64)
	if err != nil {
		lx.report.Err(t.Pos(), "invalid integer constant: %s", s)
	}

	switch k {
	case 0:
		switch b := mathutil.BitLenUint64(n); {
		case b < 32:
			return m.value2(n, m.IntType)
		case b < 33:
			return m.value2(n, m.UIntType)
		case b < 64:
			if m.Items[Long].Size == 8 {
				return m.value2(n, m.LongType)
			}

			return m.value2(n, m.LongLongType)
		default:
			if m.Items[ULong].Size == 8 {
				return m.value2(n, m.ULongType)
			}

			return m.value2(n, m.ULongLongType)
		}
	case l:
		return m.value2(n, m.LongType)
	case ll:
		return m.value2(n, m.LongLongType)
	case u:
		return m.value2(n, m.UIntType)
	case u | l:
		return m.value2(n, m.ULongType)
	case u | ll:
		return m.value2(n, m.ULongLongType)
	default:
		panic("internal error")
	}
}

func (m *Model) cBool(v bool) interface{} {
	if v {
		return m.MustConvert(int32(1), m.IntType)

	}
	return m.MustConvert(int32(0), m.IntType)
}

func (m *Model) binOp(lx *lexer, a, b operand) (interface{}, interface{}, Type) {
	av, at := a.eval(lx)
	bv, bt := b.eval(lx)
	t := m.binOpType(at, bt)
	if av == nil || bv == nil || t.Kind() == Undefined {
		return nil, nil, t
	}

	return m.MustConvert(av, t), m.MustConvert(bv, t), t
}

func (m *Model) binOpType(a, b Type) Type {
	ak := a.Kind()
	bk := b.Kind()
	if ak > bk {
		ak, bk = bk, ak
		a, b = b, a
	}
	switch ak {
	case Bool:
		switch bk {
		case Bool:
			return m.BoolType
		default:
			panic(bk)
		}
	case Char:
		switch bk {
		case Char, Int, Enum:
			return m.IntType
		case Long:
			return m.LongType
		case UChar, UInt:
			return m.UIntType
		case LongLong:
			return m.LongLongType
		case Double:
			return m.DoubleType
		default:
			panic(bk)
		}
	case SChar:
		switch bk {
		case SChar, Int:
			return m.IntType
		case UChar:
			return m.UIntType
		default:
			panic(bk)
		}
	case UChar:
		switch bk {
		case Array:
			return b.(*ctype).arrayDecay()
		case UChar, Short, UShort, Int, UInt:
			return m.UIntType
		case Long:
			return m.LongType
		case ULong:
			return m.ULongType
		case LongLong:
			return m.LongLongType
		case ULongLong:
			return m.ULongLongType
		default:
			panic(bk)
		}
	case Short:
		switch bk {
		case Short, Int:
			return m.IntType
		case UShort:
			return m.UShortType
		case UInt:
			return m.UIntType
		case Long:
			return m.LongType
		case ULong:
			return m.ULongType
		case ULongLong:
			return m.ULongLongType
		default:
			panic(bk)
		}
	case UShort:
		switch bk {
		case Array:
			return b.(*ctype).arrayDecay()
		case UShort, Int, UInt, Bool:
			return m.UIntType
		case Long:
			return m.LongType
		case ULong:
			return m.ULongType
		case LongLong, ULongLong:
			return m.ULongLongType
		case Float:
			return m.FloatType
		case Double:
			return m.DoubleType
		default:
			panic(bk)
		}
	case Int:
		switch bk {
		case Array:
			return b.(*ctype).arrayDecay()
		case Int, Enum, Bool:
			return m.IntType
		case UInt:
			return m.UIntType
		case Long:
			return m.LongType
		case ULong:
			return m.ULongType
		case LongLong:
			return m.LongLongType
		case ULongLong:
			return m.ULongLongType
		case Float:
			return m.FloatType
		case Double:
			return m.DoubleType
		default:
			panic(bk)
		}
	case UInt:
		switch bk {
		case Int, UInt, Enum, Bool:
			return m.UIntType
		case Long, ULong:
			return m.ULongType
		case LongLong, ULongLong:
			return m.ULongLongType
		case Float:
			return m.FloatType
		case Double:
			return m.DoubleType
		default:
			panic(bk)
		}
	case Long:
		switch bk {
		case Long, Bool, Enum:
			return m.LongType
		case ULong:
			return m.ULongType
		case LongLong:
			return m.LongLongType
		case ULongLong:
			return m.ULongLongType
		case Double:
			return m.DoubleType
		default:
			panic(bk)
		}
	case LongLong:
		switch bk {
		case LongLong:
			return m.LongLongType
		case ULongLong:
			return m.ULongLongType
		case Double:
			return m.DoubleType
		case LongDouble:
			return m.LongDoubleType
		default:
			panic(bk)
		}
	case ULong:
		switch bk {
		case ULong, Bool, Enum:
			return m.ULongType
		case LongLong, ULongLong:
			return m.ULongLongType
		case Double:
			return m.DoubleType
		default:
			panic(bk)
		}
	case ULongLong:
		switch bk {
		case ULongLong:
			return m.ULongLongType
		case Double:
			return m.DoubleType
		default:
			panic(bk)
		}
	case Float:
		switch bk {
		case Float:
			return m.FloatType
		case Double:
			return m.DoubleType
		case LongDouble:
			return m.LongDoubleType
		default:
			panic(bk)
		}
	case Double:
		switch bk {
		case Double:
			return m.DoubleType
		case LongDouble:
			return m.LongDoubleType
		default:
			panic(bk)
		}
	case LongDouble:
		switch bk {
		case LongDouble:
			return m.LongDoubleType
		default:
			panic(bk)
		}
	case Enum:
		switch bk {
		case Enum:
			return m.IntType
		default:
			panic(bk)
		}
	case Undefined:
		return undefined
	default:
		panic(fmt.Errorf("%v %v", ak, bk))
	}
}

func (m *Model) makeType(lx *lexer, attr int, ts ...int) Type {
	d := m.makeDeclarator(attr, ts...)
	return d.setFull(lx)
}

func (m *Model) makeDeclarator(attr int, ts ...int) *Declarator {
	s := &spec{attr, tsEncode(ts...)}
	d := &Declarator{specifier: s}
	dd := &DirectDeclarator{declarator: d, specifier: s}
	d.DirectDeclarator = dd
	return d
}

func (m *Model) checkArithmeticType(lx *lexer, a ...operand) (r bool) {
	r = true
	for _, v := range a {
		_, t := v.eval(lx)
		if !IsArithmeticType(t) {
			lx.report.Err(v.Pos(), "not an arithmetic type (have '%s')", t)
			r = false
		}
	}
	return r
}

func (m *Model) checkIntegerType(lx *lexer, a ...operand) (r bool) {
	r = true
	for _, v := range a {
		_, t := v.eval(lx)
		if !IsIntType(t) {
			lx.report.Err(v.Pos(), "not an integer type (have '%s')", t)
			r = false
		}
	}
	return r
}
