// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2016 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on [0], 6.10. Substantial portions of expression AST size
// optimizations are from [2], license of which follows.

// ----------------------------------------------------------------------------

// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This grammar is derived from the C grammar in the 'ansitize'
// program, which carried this notice:
//
// Copyright (c) 2006 Russ Cox,
// 	Massachusetts Institute of Technology
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated
// documentation files (the "Software"), to deal in the
// Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute,
// sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall
// be included in all copies or substantial portions of the
// Software.
//
// The software is provided "as is", without warranty of any
// kind, express or implied, including but not limited to the
// warranties of merchantability, fitness for a particular
// purpose and noninfringement.  In no event shall the authors
// or copyright holders be liable for any claim, damages or
// other liability, whether in an action of contract, tort or
// otherwise, arising from, out of or in connection with the
// software or the use or other dealings in the software.

package cc

import __yyfmt__ "fmt"

import (
	"fmt"

	"github.com/cznic/golex/lex"
	"github.com/cznic/xc"
)

type yySymType struct {
	yys       int
	Token     xc.Token
	groupPart Node
	node      Node
	toks      PPTokenList
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault           = 57458
	yyEofCode           = 57344
	ADDASSIGN           = 57346
	ALIGNOF             = 57347
	ANDAND              = 57348
	ANDASSIGN           = 57349
	ARROW               = 57350
	ASM                 = 57351
	AUTO                = 57352
	BOOL                = 57353
	BREAK               = 57354
	CASE                = 57355
	CAST                = 57356
	CHAR                = 57357
	CHARCONST           = 57358
	COMPLEX             = 57359
	CONST               = 57360
	CONSTANT_EXPRESSION = 1048577
	CONTINUE            = 57361
	DDD                 = 57362
	DEC                 = 57363
	DEFAULT             = 57364
	DIVASSIGN           = 57365
	DO                  = 57366
	DOUBLE              = 57367
	ELSE                = 57368
	ENUM                = 57369
	EQ                  = 57370
	EXTERN              = 57371
	FLOAT               = 57372
	FLOATCONST          = 57373
	FOR                 = 57374
	GEQ                 = 57375
	GOTO                = 57376
	IDENTIFIER          = 57377
	IDENTIFIER_LPAREN   = 57378
	IDENTIFIER_NONREPL  = 57379
	IF                  = 57380
	INC                 = 57381
	INLINE              = 57382
	INT                 = 57383
	INTCONST            = 57384
	LEQ                 = 57385
	LONG                = 57386
	LONGCHARCONST       = 57387
	LONGSTRINGLITERAL   = 57388
	LSH                 = 57389
	LSHASSIGN           = 57390
	MODASSIGN           = 57391
	MULASSIGN           = 57392
	NEQ                 = 57393
	NOELSE              = 57394
	NORETURN            = 57395
	ORASSIGN            = 57396
	OROR                = 57397
	PPDEFINE            = 57398
	PPELIF              = 57399
	PPELSE              = 57400
	PPENDIF             = 57401
	PPERROR             = 57402
	PPHASH_NL           = 57403
	PPHEADER_NAME       = 57404
	PPIF                = 57405
	PPIFDEF             = 57406
	PPIFNDEF            = 57407
	PPINCLUDE           = 57408
	PPINCLUDE_NEXT      = 57409
	PPLINE              = 57410
	PPNONDIRECTIVE      = 57411
	PPNUMBER            = 57412
	PPOTHER             = 57413
	PPPASTE             = 57414
	PPPRAGMA            = 57415
	PPUNDEF             = 57416
	PREPROCESSING_FILE  = 1048576
	REGISTER            = 57417
	RESTRICT            = 57418
	RETURN              = 57419
	RSH                 = 57420
	RSHASSIGN           = 57421
	SHORT               = 57422
	SIGNED              = 57423
	SIZEOF              = 57424
	STATIC              = 57425
	STATIC_ASSERT       = 57426
	STRINGLITERAL       = 57427
	STRUCT              = 57428
	SUBASSIGN           = 57429
	SWITCH              = 57430
	TRANSLATION_UNIT    = 1048578
	TYPEDEF             = 57431
	TYPEDEFNAME         = 57432
	TYPEOF              = 57433
	UNARY               = 57434
	UNION               = 57435
	UNSIGNED            = 57436
	VOID                = 57437
	VOLATILE            = 57438
	WHILE               = 57439
	XORASSIGN           = 57440
	yyErrCode           = 57345

	yyMaxDepth = 200
	yyTabOfs   = -318
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (307x)
		42:      1,   // '*' (268x)
		57377:   2,   // IDENTIFIER (221x)
		38:      3,   // '&' (215x)
		43:      4,   // '+' (215x)
		45:      5,   // '-' (215x)
		57363:   6,   // DEC (215x)
		57381:   7,   // INC (215x)
		59:      8,   // ';' (210x)
		41:      9,   // ')' (196x)
		44:      10,  // ',' (182x)
		57427:   11,  // STRINGLITERAL (163x)
		91:      12,  // '[' (159x)
		33:      13,  // '!' (145x)
		126:     14,  // '~' (145x)
		57347:   15,  // ALIGNOF (145x)
		57358:   16,  // CHARCONST (145x)
		57373:   17,  // FLOATCONST (145x)
		57384:   18,  // INTCONST (145x)
		57387:   19,  // LONGCHARCONST (145x)
		57388:   20,  // LONGSTRINGLITERAL (145x)
		57424:   21,  // SIZEOF (145x)
		57438:   22,  // VOLATILE (136x)
		57360:   23,  // CONST (134x)
		57418:   24,  // RESTRICT (134x)
		57353:   25,  // BOOL (124x)
		57357:   26,  // CHAR (124x)
		57359:   27,  // COMPLEX (124x)
		57367:   28,  // DOUBLE (124x)
		57369:   29,  // ENUM (124x)
		57372:   30,  // FLOAT (124x)
		57383:   31,  // INT (124x)
		57386:   32,  // LONG (124x)
		57422:   33,  // SHORT (124x)
		57423:   34,  // SIGNED (124x)
		57428:   35,  // STRUCT (124x)
		57432:   36,  // TYPEDEFNAME (124x)
		57433:   37,  // TYPEOF (124x)
		57435:   38,  // UNION (124x)
		57436:   39,  // UNSIGNED (124x)
		57437:   40,  // VOID (124x)
		125:     41,  // '}' (122x)
		58:      42,  // ':' (119x)
		57425:   43,  // STATIC (115x)
		57352:   44,  // AUTO (109x)
		57371:   45,  // EXTERN (109x)
		57382:   46,  // INLINE (109x)
		57395:   47,  // NORETURN (109x)
		57417:   48,  // REGISTER (109x)
		57431:   49,  // TYPEDEF (109x)
		57344:   50,  // $end (102x)
		61:      51,  // '=' (90x)
		123:     52,  // '{' (86x)
		57501:   53,  // Expression (85x)
		93:      54,  // ']' (83x)
		46:      55,  // '.' (79x)
		57351:   56,  // ASM (78x)
		57426:   57,  // STATIC_ASSERT (75x)
		37:      58,  // '%' (71x)
		47:      59,  // '/' (71x)
		60:      60,  // '<' (71x)
		62:      61,  // '>' (71x)
		63:      62,  // '?' (71x)
		94:      63,  // '^' (71x)
		124:     64,  // '|' (71x)
		57346:   65,  // ADDASSIGN (71x)
		57348:   66,  // ANDAND (71x)
		57349:   67,  // ANDASSIGN (71x)
		57350:   68,  // ARROW (71x)
		57365:   69,  // DIVASSIGN (71x)
		57370:   70,  // EQ (71x)
		57375:   71,  // GEQ (71x)
		57385:   72,  // LEQ (71x)
		57389:   73,  // LSH (71x)
		57390:   74,  // LSHASSIGN (71x)
		57391:   75,  // MODASSIGN (71x)
		57392:   76,  // MULASSIGN (71x)
		57393:   77,  // NEQ (71x)
		57396:   78,  // ORASSIGN (71x)
		57397:   79,  // OROR (71x)
		57420:   80,  // RSH (71x)
		57421:   81,  // RSHASSIGN (71x)
		57429:   82,  // SUBASSIGN (71x)
		57440:   83,  // XORASSIGN (71x)
		10:      84,  // '\n' (58x)
		57376:   85,  // GOTO (53x)
		57413:   86,  // PPOTHER (52x)
		57439:   87,  // WHILE (51x)
		57354:   88,  // BREAK (50x)
		57355:   89,  // CASE (50x)
		57361:   90,  // CONTINUE (50x)
		57364:   91,  // DEFAULT (50x)
		57366:   92,  // DO (50x)
		57374:   93,  // FOR (50x)
		57380:   94,  // IF (50x)
		57419:   95,  // RETURN (50x)
		57430:   96,  // SWITCH (50x)
		57401:   97,  // PPENDIF (44x)
		57400:   98,  // PPELSE (40x)
		57399:   99,  // PPELIF (39x)
		57398:   100, // PPDEFINE (35x)
		57402:   101, // PPERROR (35x)
		57403:   102, // PPHASH_NL (35x)
		57405:   103, // PPIF (35x)
		57406:   104, // PPIFDEF (35x)
		57407:   105, // PPIFNDEF (35x)
		57408:   106, // PPINCLUDE (35x)
		57409:   107, // PPINCLUDE_NEXT (35x)
		57410:   108, // PPLINE (35x)
		57411:   109, // PPNONDIRECTIVE (35x)
		57415:   110, // PPPRAGMA (35x)
		57416:   111, // PPUNDEF (35x)
		57368:   112, // ELSE (30x)
		57553:   113, // TypeQualifier (28x)
		57502:   114, // ExpressionList (26x)
		57526:   115, // PPTokenList (22x)
		57528:   116, // PPTokens (22x)
		57497:   117, // EnumSpecifier (20x)
		57548:   118, // StructOrUnion (20x)
		57549:   119, // StructOrUnionSpecifier (20x)
		57556:   120, // TypeSpecifier (20x)
		57503:   121, // ExpressionListOpt (18x)
		57468:   122, // BasicAssemblerStatement (15x)
		57480:   123, // DeclarationSpecifiers (15x)
		57509:   124, // FunctionSpecifier (15x)
		57527:   125, // PPTokenListOpt (15x)
		57543:   126, // StorageClassSpecifier (15x)
		57466:   127, // AssemblerStatement (13x)
		57474:   128, // CompoundStatement (13x)
		57505:   129, // ExpressionStatement (12x)
		57523:   130, // IterationStatement (12x)
		57524:   131, // JumpStatement (12x)
		57525:   132, // LabeledStatement (12x)
		57537:   133, // SelectionStatement (12x)
		57541:   134, // Statement (12x)
		57533:   135, // Pointer (11x)
		57534:   136, // PointerOpt (10x)
		57542:   137, // StaticAssertDeclaration (9x)
		57476:   138, // ControlLine (8x)
		57482:   139, // Declarator (8x)
		57512:   140, // GroupPart (8x)
		57516:   141, // IfGroup (8x)
		57517:   142, // IfSection (8x)
		57550:   143, // TextLine (8x)
		57477:   144, // Declaration (7x)
		57452:   145, // $@4 (6x)
		57475:   146, // ConstantExpression (6x)
		57362:   147, // DDD (6x)
		57510:   148, // GroupList (6x)
		57511:   149, // GroupListOpt (5x)
		57536:   150, // ReplacementList (5x)
		57538:   151, // SpecifierQualifierList (5x)
		57554:   152, // TypeQualifierList (5x)
		57442:   153, // $@10 (4x)
		57459:   154, // AbstractDeclarator (4x)
		57464:   155, // AssemblerOperand (4x)
		57467:   156, // AssemblerSymbolicNameOpt (4x)
		57481:   157, // DeclarationSpecifiersOpt (4x)
		57486:   158, // Designator (4x)
		57529:   159, // ParameterDeclaration (4x)
		57552:   160, // TypeName (4x)
		57555:   161, // TypeQualifierListOpt (4x)
		57463:   162, // AssemblerInstructions (3x)
		57465:   163, // AssemblerOperands (3x)
		57473:   164, // CommaOpt (3x)
		57484:   165, // Designation (3x)
		57485:   166, // DesignationOpt (3x)
		57487:   167, // DesignatorList (3x)
		57504:   168, // ExpressionOpt (3x)
		57513:   169, // IdentifierList (3x)
		57518:   170, // InitDeclarator (3x)
		57521:   171, // Initializer (3x)
		57530:   172, // ParameterList (3x)
		57531:   173, // ParameterTypeList (3x)
		57443:   174, // $@11 (2x)
		57453:   175, // $@5 (2x)
		57460:   176, // AbstractDeclaratorOpt (2x)
		57469:   177, // BlockItem (2x)
		57472:   178, // Clobbers (2x)
		57483:   179, // DeclaratorOpt (2x)
		57488:   180, // DirectAbstractDeclarator (2x)
		57489:   181, // DirectAbstractDeclaratorOpt (2x)
		57490:   182, // DirectDeclarator (2x)
		57491:   183, // ElifGroup (2x)
		57498:   184, // EnumerationConstant (2x)
		57499:   185, // Enumerator (2x)
		57506:   186, // ExternalDeclaration (2x)
		57508:   187, // FunctionDefinition (2x)
		57514:   188, // IdentifierListOpt (2x)
		57515:   189, // IdentifierOpt (2x)
		57519:   190, // InitDeclaratorList (2x)
		57520:   191, // InitDeclaratorListOpt (2x)
		57522:   192, // InitializerList (2x)
		57532:   193, // ParameterTypeListOpt (2x)
		57539:   194, // SpecifierQualifierListOpt (2x)
		57544:   195, // StructDeclaration (2x)
		57546:   196, // StructDeclarator (2x)
		57557:   197, // VolatileOpt (2x)
		57441:   198, // $@1 (1x)
		57444:   199, // $@12 (1x)
		57445:   200, // $@13 (1x)
		57446:   201, // $@14 (1x)
		57447:   202, // $@15 (1x)
		57448:   203, // $@16 (1x)
		57449:   204, // $@17 (1x)
		57450:   205, // $@2 (1x)
		57451:   206, // $@3 (1x)
		57454:   207, // $@6 (1x)
		57455:   208, // $@7 (1x)
		57456:   209, // $@8 (1x)
		57457:   210, // $@9 (1x)
		57461:   211, // ArgumentExpressionList (1x)
		57462:   212, // ArgumentExpressionListOpt (1x)
		57470:   213, // BlockItemList (1x)
		57471:   214, // BlockItemListOpt (1x)
		1048577: 215, // CONSTANT_EXPRESSION (1x)
		57478:   216, // DeclarationList (1x)
		57479:   217, // DeclarationListOpt (1x)
		57492:   218, // ElifGroupList (1x)
		57493:   219, // ElifGroupListOpt (1x)
		57494:   220, // ElseGroup (1x)
		57495:   221, // ElseGroupOpt (1x)
		57496:   222, // EndifLine (1x)
		57500:   223, // EnumeratorList (1x)
		57507:   224, // FunctionBody (1x)
		57378:   225, // IDENTIFIER_LPAREN (1x)
		1048576: 226, // PREPROCESSING_FILE (1x)
		57535:   227, // PreprocessingFile (1x)
		57540:   228, // Start (1x)
		57545:   229, // StructDeclarationList (1x)
		57547:   230, // StructDeclaratorList (1x)
		1048578: 231, // TRANSLATION_UNIT (1x)
		57551:   232, // TranslationUnit (1x)
		57458:   233, // $default (0x)
		57356:   234, // CAST (0x)
		57345:   235, // error (0x)
		57379:   236, // IDENTIFIER_NONREPL (0x)
		57394:   237, // NOELSE (0x)
		57404:   238, // PPHEADER_NAME (0x)
		57412:   239, // PPNUMBER (0x)
		57414:   240, // PPPASTE (0x)
		57434:   241, // UNARY (0x)
	}

	yySymNames = []string{
		"'('",
		"'*'",
		"IDENTIFIER",
		"'&'",
		"'+'",
		"'-'",
		"DEC",
		"INC",
		"';'",
		"')'",
		"','",
		"STRINGLITERAL",
		"'['",
		"'!'",
		"'~'",
		"ALIGNOF",
		"CHARCONST",
		"FLOATCONST",
		"INTCONST",
		"LONGCHARCONST",
		"LONGSTRINGLITERAL",
		"SIZEOF",
		"VOLATILE",
		"CONST",
		"RESTRICT",
		"BOOL",
		"CHAR",
		"COMPLEX",
		"DOUBLE",
		"ENUM",
		"FLOAT",
		"INT",
		"LONG",
		"SHORT",
		"SIGNED",
		"STRUCT",
		"TYPEDEFNAME",
		"TYPEOF",
		"UNION",
		"UNSIGNED",
		"VOID",
		"'}'",
		"':'",
		"STATIC",
		"AUTO",
		"EXTERN",
		"INLINE",
		"NORETURN",
		"REGISTER",
		"TYPEDEF",
		"$end",
		"'='",
		"'{'",
		"Expression",
		"']'",
		"'.'",
		"ASM",
		"STATIC_ASSERT",
		"'%'",
		"'/'",
		"'<'",
		"'>'",
		"'?'",
		"'^'",
		"'|'",
		"ADDASSIGN",
		"ANDAND",
		"ANDASSIGN",
		"ARROW",
		"DIVASSIGN",
		"EQ",
		"GEQ",
		"LEQ",
		"LSH",
		"LSHASSIGN",
		"MODASSIGN",
		"MULASSIGN",
		"NEQ",
		"ORASSIGN",
		"OROR",
		"RSH",
		"RSHASSIGN",
		"SUBASSIGN",
		"XORASSIGN",
		"'\\n'",
		"GOTO",
		"PPOTHER",
		"WHILE",
		"BREAK",
		"CASE",
		"CONTINUE",
		"DEFAULT",
		"DO",
		"FOR",
		"IF",
		"RETURN",
		"SWITCH",
		"PPENDIF",
		"PPELSE",
		"PPELIF",
		"PPDEFINE",
		"PPERROR",
		"PPHASH_NL",
		"PPIF",
		"PPIFDEF",
		"PPIFNDEF",
		"PPINCLUDE",
		"PPINCLUDE_NEXT",
		"PPLINE",
		"PPNONDIRECTIVE",
		"PPPRAGMA",
		"PPUNDEF",
		"ELSE",
		"TypeQualifier",
		"ExpressionList",
		"PPTokenList",
		"PPTokens",
		"EnumSpecifier",
		"StructOrUnion",
		"StructOrUnionSpecifier",
		"TypeSpecifier",
		"ExpressionListOpt",
		"BasicAssemblerStatement",
		"DeclarationSpecifiers",
		"FunctionSpecifier",
		"PPTokenListOpt",
		"StorageClassSpecifier",
		"AssemblerStatement",
		"CompoundStatement",
		"ExpressionStatement",
		"IterationStatement",
		"JumpStatement",
		"LabeledStatement",
		"SelectionStatement",
		"Statement",
		"Pointer",
		"PointerOpt",
		"StaticAssertDeclaration",
		"ControlLine",
		"Declarator",
		"GroupPart",
		"IfGroup",
		"IfSection",
		"TextLine",
		"Declaration",
		"$@4",
		"ConstantExpression",
		"DDD",
		"GroupList",
		"GroupListOpt",
		"ReplacementList",
		"SpecifierQualifierList",
		"TypeQualifierList",
		"$@10",
		"AbstractDeclarator",
		"AssemblerOperand",
		"AssemblerSymbolicNameOpt",
		"DeclarationSpecifiersOpt",
		"Designator",
		"ParameterDeclaration",
		"TypeName",
		"TypeQualifierListOpt",
		"AssemblerInstructions",
		"AssemblerOperands",
		"CommaOpt",
		"Designation",
		"DesignationOpt",
		"DesignatorList",
		"ExpressionOpt",
		"IdentifierList",
		"InitDeclarator",
		"Initializer",
		"ParameterList",
		"ParameterTypeList",
		"$@11",
		"$@5",
		"AbstractDeclaratorOpt",
		"BlockItem",
		"Clobbers",
		"DeclaratorOpt",
		"DirectAbstractDeclarator",
		"DirectAbstractDeclaratorOpt",
		"DirectDeclarator",
		"ElifGroup",
		"EnumerationConstant",
		"Enumerator",
		"ExternalDeclaration",
		"FunctionDefinition",
		"IdentifierListOpt",
		"IdentifierOpt",
		"InitDeclaratorList",
		"InitDeclaratorListOpt",
		"InitializerList",
		"ParameterTypeListOpt",
		"SpecifierQualifierListOpt",
		"StructDeclaration",
		"StructDeclarator",
		"VolatileOpt",
		"$@1",
		"$@12",
		"$@13",
		"$@14",
		"$@15",
		"$@16",
		"$@17",
		"$@2",
		"$@3",
		"$@6",
		"$@7",
		"$@8",
		"$@9",
		"ArgumentExpressionList",
		"ArgumentExpressionListOpt",
		"BlockItemList",
		"BlockItemListOpt",
		"CONSTANT_EXPRESSION",
		"DeclarationList",
		"DeclarationListOpt",
		"ElifGroupList",
		"ElifGroupListOpt",
		"ElseGroup",
		"ElseGroupOpt",
		"EndifLine",
		"EnumeratorList",
		"FunctionBody",
		"IDENTIFIER_LPAREN",
		"PREPROCESSING_FILE",
		"PreprocessingFile",
		"Start",
		"StructDeclarationList",
		"StructDeclaratorList",
		"TRANSLATION_UNIT",
		"TranslationUnit",
		"$default",
		"CAST",
		"error",
		"IDENTIFIER_NONREPL",
		"NOELSE",
		"PPHEADER_NAME",
		"PPNUMBER",
		"PPPASTE",
		"UNARY",
	}

	yyTokenLiteralStrings = map[int]string{
		57377:   "identifier",
		57363:   "--",
		57381:   "++",
		57427:   "string literal",
		57347:   "_Alignof",
		57358:   "character constant",
		57373:   "floating-point constant",
		57384:   "integer constant",
		57387:   "long character constant",
		57388:   "long string constant",
		57424:   "sizeof",
		57438:   "volatile",
		57360:   "const",
		57418:   "restrict",
		57353:   "_Bool",
		57357:   "char",
		57359:   "_Complex",
		57367:   "double",
		57369:   "enum",
		57372:   "float",
		57383:   "int",
		57386:   "long",
		57422:   "short",
		57423:   "signed",
		57428:   "struct",
		57432:   "typedefname",
		57433:   "typeof",
		57435:   "union",
		57436:   "unsigned",
		57437:   "void",
		57425:   "static",
		57352:   "auto",
		57371:   "extern",
		57382:   "inline",
		57395:   "_Noreturn",
		57417:   "register",
		57431:   "typedef",
		57351:   "asm",
		57426:   "_Static_assert",
		57346:   "+=",
		57348:   "&&",
		57349:   "&=",
		57350:   "->",
		57365:   "/=",
		57370:   "==",
		57375:   ">=",
		57385:   "<=",
		57389:   "<<",
		57390:   "<<=",
		57391:   "%=",
		57392:   "*=",
		57393:   "!=",
		57396:   "|=",
		57397:   "||",
		57420:   ">>",
		57421:   ">>=",
		57429:   "-=",
		57440:   "^=",
		57376:   "goto",
		57413:   "ppother",
		57439:   "while",
		57354:   "break",
		57355:   "case",
		57361:   "continue",
		57364:   "default",
		57366:   "do",
		57374:   "for",
		57380:   "if",
		57419:   "return",
		57430:   "switch",
		57401:   "#endif",
		57400:   "#else",
		57399:   "#elif",
		57398:   "#define",
		57402:   "#error",
		57403:   "#",
		57405:   "#if",
		57406:   "#ifdef",
		57407:   "#ifndef",
		57408:   "#include",
		57409:   "#include_next",
		57410:   "#line",
		57411:   "#foo",
		57415:   "#pragma",
		57416:   "#undef",
		57368:   "else",
		57362:   "...",
		1048577: "constant expression prefix",
		57378:   "identifier immediatelly followed by '('",
		1048576: "preprocessing file prefix",
		1048578: "translation unit prefix",
		57379:   "non replaceable identifier",
		57404:   "header name",
		57412:   "preprocessing number",
		57414:   "##",
	}

	yyReductions = map[int]struct{ xsym, components int }{
		0:   {0, 1},
		1:   {198, 0},
		2:   {228, 3},
		3:   {205, 0},
		4:   {228, 3},
		5:   {206, 0},
		6:   {228, 3},
		7:   {184, 1},
		8:   {211, 1},
		9:   {211, 3},
		10:  {212, 0},
		11:  {212, 1},
		12:  {53, 1},
		13:  {53, 1},
		14:  {53, 1},
		15:  {53, 1},
		16:  {53, 1},
		17:  {53, 1},
		18:  {53, 1},
		19:  {53, 3},
		20:  {53, 4},
		21:  {53, 4},
		22:  {53, 3},
		23:  {53, 3},
		24:  {53, 2},
		25:  {53, 2},
		26:  {53, 7},
		27:  {53, 2},
		28:  {53, 2},
		29:  {53, 2},
		30:  {53, 2},
		31:  {53, 2},
		32:  {53, 2},
		33:  {53, 2},
		34:  {53, 2},
		35:  {53, 2},
		36:  {53, 4},
		37:  {53, 4},
		38:  {53, 3},
		39:  {53, 3},
		40:  {53, 3},
		41:  {53, 3},
		42:  {53, 3},
		43:  {53, 3},
		44:  {53, 3},
		45:  {53, 3},
		46:  {53, 3},
		47:  {53, 3},
		48:  {53, 3},
		49:  {53, 3},
		50:  {53, 3},
		51:  {53, 3},
		52:  {53, 3},
		53:  {53, 3},
		54:  {53, 3},
		55:  {53, 3},
		56:  {53, 5},
		57:  {53, 3},
		58:  {53, 3},
		59:  {53, 3},
		60:  {53, 3},
		61:  {53, 3},
		62:  {53, 3},
		63:  {53, 3},
		64:  {53, 3},
		65:  {53, 3},
		66:  {53, 3},
		67:  {53, 3},
		68:  {53, 4},
		69:  {168, 0},
		70:  {168, 1},
		71:  {114, 1},
		72:  {114, 3},
		73:  {121, 0},
		74:  {121, 1},
		75:  {145, 0},
		76:  {146, 2},
		77:  {144, 3},
		78:  {144, 1},
		79:  {123, 2},
		80:  {123, 2},
		81:  {123, 2},
		82:  {123, 2},
		83:  {157, 0},
		84:  {157, 1},
		85:  {190, 1},
		86:  {190, 3},
		87:  {191, 0},
		88:  {191, 1},
		89:  {170, 1},
		90:  {175, 0},
		91:  {170, 4},
		92:  {126, 1},
		93:  {126, 1},
		94:  {126, 1},
		95:  {126, 1},
		96:  {126, 1},
		97:  {120, 1},
		98:  {120, 1},
		99:  {120, 1},
		100: {120, 1},
		101: {120, 1},
		102: {120, 1},
		103: {120, 1},
		104: {120, 1},
		105: {120, 1},
		106: {120, 1},
		107: {120, 1},
		108: {120, 1},
		109: {120, 1},
		110: {120, 1},
		111: {120, 4},
		112: {120, 4},
		113: {207, 0},
		114: {208, 0},
		115: {119, 7},
		116: {119, 2},
		117: {118, 1},
		118: {118, 1},
		119: {229, 1},
		120: {229, 2},
		121: {195, 3},
		122: {195, 2},
		123: {195, 1},
		124: {151, 2},
		125: {151, 2},
		126: {194, 0},
		127: {194, 1},
		128: {230, 1},
		129: {230, 3},
		130: {196, 1},
		131: {196, 3},
		132: {164, 0},
		133: {164, 1},
		134: {209, 0},
		135: {117, 7},
		136: {117, 2},
		137: {223, 1},
		138: {223, 3},
		139: {185, 1},
		140: {185, 3},
		141: {113, 1},
		142: {113, 1},
		143: {113, 1},
		144: {124, 1},
		145: {124, 1},
		146: {139, 2},
		147: {179, 0},
		148: {179, 1},
		149: {182, 1},
		150: {182, 3},
		151: {182, 5},
		152: {182, 6},
		153: {182, 6},
		154: {182, 5},
		155: {210, 0},
		156: {182, 5},
		157: {182, 4},
		158: {135, 2},
		159: {135, 3},
		160: {136, 0},
		161: {136, 1},
		162: {152, 1},
		163: {152, 2},
		164: {161, 0},
		165: {161, 1},
		166: {173, 1},
		167: {173, 3},
		168: {193, 0},
		169: {193, 1},
		170: {172, 1},
		171: {172, 3},
		172: {159, 2},
		173: {159, 2},
		174: {169, 1},
		175: {169, 3},
		176: {188, 0},
		177: {188, 1},
		178: {189, 0},
		179: {189, 1},
		180: {153, 0},
		181: {160, 3},
		182: {154, 1},
		183: {154, 2},
		184: {176, 0},
		185: {176, 1},
		186: {180, 3},
		187: {180, 4},
		188: {180, 5},
		189: {180, 6},
		190: {180, 6},
		191: {180, 4},
		192: {174, 0},
		193: {180, 4},
		194: {199, 0},
		195: {180, 5},
		196: {181, 0},
		197: {181, 1},
		198: {171, 1},
		199: {171, 4},
		200: {192, 2},
		201: {192, 4},
		202: {165, 2},
		203: {166, 0},
		204: {166, 1},
		205: {167, 1},
		206: {167, 2},
		207: {158, 3},
		208: {158, 2},
		209: {134, 1},
		210: {134, 1},
		211: {134, 1},
		212: {134, 1},
		213: {134, 1},
		214: {134, 1},
		215: {134, 1},
		216: {132, 3},
		217: {132, 4},
		218: {132, 3},
		219: {200, 0},
		220: {128, 4},
		221: {213, 1},
		222: {213, 2},
		223: {214, 0},
		224: {214, 1},
		225: {177, 1},
		226: {177, 1},
		227: {129, 2},
		228: {133, 5},
		229: {133, 7},
		230: {133, 5},
		231: {130, 5},
		232: {130, 7},
		233: {130, 9},
		234: {130, 8},
		235: {131, 3},
		236: {131, 2},
		237: {131, 2},
		238: {131, 3},
		239: {232, 1},
		240: {232, 2},
		241: {186, 1},
		242: {186, 1},
		243: {186, 2},
		244: {186, 1},
		245: {201, 0},
		246: {187, 5},
		247: {202, 0},
		248: {224, 2},
		249: {203, 0},
		250: {224, 3},
		251: {216, 1},
		252: {216, 2},
		253: {217, 0},
		254: {204, 0},
		255: {217, 2},
		256: {162, 1},
		257: {162, 2},
		258: {122, 5},
		259: {197, 0},
		260: {197, 1},
		261: {155, 5},
		262: {163, 1},
		263: {163, 3},
		264: {156, 0},
		265: {156, 3},
		266: {178, 1},
		267: {178, 3},
		268: {127, 1},
		269: {127, 7},
		270: {127, 9},
		271: {127, 11},
		272: {127, 13},
		273: {127, 6},
		274: {137, 7},
		275: {227, 1},
		276: {148, 1},
		277: {148, 2},
		278: {149, 0},
		279: {149, 1},
		280: {140, 1},
		281: {140, 1},
		282: {140, 3},
		283: {140, 1},
		284: {142, 4},
		285: {141, 4},
		286: {141, 4},
		287: {141, 4},
		288: {218, 1},
		289: {218, 2},
		290: {219, 0},
		291: {219, 1},
		292: {183, 4},
		293: {220, 3},
		294: {221, 0},
		295: {221, 1},
		296: {222, 1},
		297: {138, 3},
		298: {138, 5},
		299: {138, 7},
		300: {138, 5},
		301: {138, 2},
		302: {138, 1},
		303: {138, 3},
		304: {138, 3},
		305: {138, 2},
		306: {138, 3},
		307: {138, 6},
		308: {138, 2},
		309: {138, 4},
		310: {138, 3},
		311: {143, 1},
		312: {150, 1},
		313: {115, 1},
		314: {125, 1},
		315: {125, 2},
		316: {116, 1},
		317: {116, 2},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 50}:   "invalid empty input",
		yyXError{564, -1}: "expected #endif",
		yyXError{566, -1}: "expected #endif",
		yyXError{1, -1}:   "expected $end",
		yyXError{484, -1}: "expected $end",
		yyXError{486, -1}: "expected $end",
		yyXError{32, -1}:  "expected '('",
		yyXError{48, -1}:  "expected '('",
		yyXError{74, -1}:  "expected '('",
		yyXError{284, -1}: "expected '('",
		yyXError{361, -1}: "expected '('",
		yyXError{382, -1}: "expected '('",
		yyXError{418, -1}: "expected '('",
		yyXError{419, -1}: "expected '('",
		yyXError{420, -1}: "expected '('",
		yyXError{422, -1}: "expected '('",
		yyXError{448, -1}: "expected '('",
		yyXError{53, -1}:  "expected ')'",
		yyXError{76, -1}:  "expected ')'",
		yyXError{83, -1}:  "expected ')'",
		yyXError{175, -1}: "expected ')'",
		yyXError{190, -1}: "expected ')'",
		yyXError{193, -1}: "expected ')'",
		yyXError{196, -1}: "expected ')'",
		yyXError{204, -1}: "expected ')'",
		yyXError{209, -1}: "expected ')'",
		yyXError{215, -1}: "expected ')'",
		yyXError{231, -1}: "expected ')'",
		yyXError{236, -1}: "expected ')'",
		yyXError{247, -1}: "expected ')'",
		yyXError{282, -1}: "expected ')'",
		yyXError{332, -1}: "expected ')'",
		yyXError{438, -1}: "expected ')'",
		yyXError{444, -1}: "expected ')'",
		yyXError{528, -1}: "expected ')'",
		yyXError{529, -1}: "expected ')'",
		yyXError{536, -1}: "expected ')'",
		yyXError{539, -1}: "expected ')'",
		yyXError{51, -1}:  "expected ','",
		yyXError{319, -1}: "expected ':'",
		yyXError{364, -1}: "expected ':'",
		yyXError{411, -1}: "expected ':'",
		yyXError{472, -1}: "expected ':'",
		yyXError{45, -1}:  "expected ';'",
		yyXError{54, -1}:  "expected ';'",
		yyXError{340, -1}: "expected ';'",
		yyXError{356, -1}: "expected ';'",
		yyXError{417, -1}: "expected ';'",
		yyXError{424, -1}: "expected ';'",
		yyXError{425, -1}: "expected ';'",
		yyXError{427, -1}: "expected ';'",
		yyXError{431, -1}: "expected ';'",
		yyXError{434, -1}: "expected ';'",
		yyXError{436, -1}: "expected ';'",
		yyXError{442, -1}: "expected ';'",
		yyXError{451, -1}: "expected ';'",
		yyXError{344, -1}: "expected '='",
		yyXError{88, -1}:  "expected '['",
		yyXError{508, -1}: "expected '\\n'",
		yyXError{512, -1}: "expected '\\n'",
		yyXError{516, -1}: "expected '\\n'",
		yyXError{519, -1}: "expected '\\n'",
		yyXError{521, -1}: "expected '\\n'",
		yyXError{543, -1}: "expected '\\n'",
		yyXError{548, -1}: "expected '\\n'",
		yyXError{551, -1}: "expected '\\n'",
		yyXError{558, -1}: "expected '\\n'",
		yyXError{563, -1}: "expected '\\n'",
		yyXError{569, -1}: "expected '\\n'",
		yyXError{94, -1}:  "expected ']'",
		yyXError{183, -1}: "expected ']'",
		yyXError{227, -1}: "expected ']'",
		yyXError{259, -1}: "expected ']'",
		yyXError{370, -1}: "expected ']'",
		yyXError{292, -1}: "expected '{'",
		yyXError{294, -1}: "expected '{'",
		yyXError{306, -1}: "expected '{'",
		yyXError{308, -1}: "expected '{'",
		yyXError{268, -1}: "expected '}'",
		yyXError{272, -1}: "expected '}'",
		yyXError{303, -1}: "expected '}'",
		yyXError{412, -1}: "expected '}'",
		yyXError{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		yyXError{203, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{87, -1}:  "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{286, -1}: "expected assembler instructions or string literal",
		yyXError{360, -1}: "expected assembler instructions or string literal",
		yyXError{362, -1}: "expected assembler instructions or string literal",
		yyXError{372, -1}: "expected assembler operand or one of ['[', string literal]",
		yyXError{387, -1}: "expected assembler operands or one of [')', '[', string literal]",
		yyXError{365, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{391, -1}: "expected assembler operands or one of ['[', string literal]",
		yyXError{355, -1}: "expected assembler statement or asm",
		yyXError{414, -1}: "expected block item or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{373, -1}: "expected clobbers or string literal",
		yyXError{394, -1}: "expected clobbers or string literal",
		yyXError{354, -1}: "expected compound statement or '{'",
		yyXError{3, -1}:   "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{49, -1}:  "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{256, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{300, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{323, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{410, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{483, -1}: "expected constant expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{346, -1}: "expected declaration list or one of [_Bool, _Complex, _Noreturn, _Static_assert, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{349, -1}: "expected declaration or one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{433, -1}: "expected declaration or optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{322, -1}: "expected declarator or one of ['(', '*', identifier]",
		yyXError{195, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{7, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{253, -1}: "expected designator or one of ['.', '=', '[']",
		yyXError{198, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		yyXError{84, -1}:  "expected direct abstract declarator or one of ['(', '[']",
		yyXError{320, -1}: "expected direct declarator or one of ['(', identifier]",
		yyXError{556, -1}: "expected elif group or one of [#elif, #else, #endif]",
		yyXError{562, -1}: "expected endif line or #endif",
		yyXError{493, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		yyXError{554, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		yyXError{295, -1}: "expected enumerator list or identifier",
		yyXError{302, -1}: "expected enumerator or one of ['}', identifier]",
		yyXError{99, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{123, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{449, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{453, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{457, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{461, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{63, -1}:  "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{244, -1}: "expected expression list or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{248, -1}: "expected expression or one of [!=, $end, %=, &&, &=, '!', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '{', '|', '}', '~', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal, |=, ||]",
		yyXError{91, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{226, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{283, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{50, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{65, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{66, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{67, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{68, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{69, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{70, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{71, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{72, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{73, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{97, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{105, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{106, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{107, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{108, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{109, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{110, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{111, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{112, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{113, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{114, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{115, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{116, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{117, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{118, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{119, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{120, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{121, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{122, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{124, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{125, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{126, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{127, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{128, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{129, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{130, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{131, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{132, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{133, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{134, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{149, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{150, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{177, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{184, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{220, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{223, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{383, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{95, -1}:  "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{218, -1}: "expected expression or optional type qualifier list or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{330, -1}: "expected expression or type name or one of ['!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{6, -1}:   "expected external declaration or one of [$end, ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{345, -1}: "expected function body or one of ['{', asm]",
		yyXError{352, -1}: "expected function body or one of ['{', asm]",
		yyXError{343, -1}: "expected function body or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{545, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{487, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{101, -1}: "expected identifier",
		yyXError{102, -1}: "expected identifier",
		yyXError{212, -1}: "expected identifier",
		yyXError{257, -1}: "expected identifier",
		yyXError{369, -1}: "expected identifier",
		yyXError{423, -1}: "expected identifier",
		yyXError{495, -1}: "expected identifier",
		yyXError{496, -1}: "expected identifier",
		yyXError{503, -1}: "expected identifier",
		yyXError{377, -1}: "expected identifier list or identifier",
		yyXError{525, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		yyXError{479, -1}: "expected init declarator or one of ['(', '*', identifier]",
		yyXError{250, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{264, -1}: "expected initializer list or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{252, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{270, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{477, -1}: "expected initializer or one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{267, -1}: "expected initializer or optional designation or one of ['!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{56, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{57, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{58, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{59, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{60, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{61, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{62, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{103, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{104, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{136, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{137, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{138, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{139, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{140, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{141, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{142, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{143, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{144, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{145, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{146, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{152, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{153, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{154, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{155, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{156, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{157, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{158, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{159, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{160, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{161, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{162, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{163, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{164, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{165, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{166, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{167, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{168, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{169, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{170, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{171, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{172, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{176, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{180, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{188, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{243, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{245, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{249, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{273, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{274, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{275, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{276, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{277, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{278, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{279, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{280, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{281, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{64, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{147, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{151, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{173, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{178, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{331, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{384, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{401, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{263, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{90, -1}:  "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{98, -1}:  "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{185, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{221, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{224, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		yyXError{488, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{489, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{490, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{492, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{499, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{505, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{507, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{510, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{513, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{515, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{517, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{518, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{520, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{522, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{523, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{526, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{531, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{532, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{534, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{538, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{541, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{542, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{547, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{567, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{568, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{570, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		yyXError{546, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{550, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{553, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{555, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{560, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{561, -1}: "expected one of [#elif, #else, #endif]",
		yyXError{469, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{8, -1}:   "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{55, -1}:  "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{481, -1}: "expected one of [$end, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{42, -1}:  "expected one of [$end, ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{43, -1}:  "expected one of [$end, ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{44, -1}:  "expected one of [$end, ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{46, -1}:  "expected one of [$end, ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{291, -1}: "expected one of [$end, ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{353, -1}: "expected one of [$end, ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{397, -1}: "expected one of [$end, ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{399, -1}: "expected one of [$end, ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{482, -1}: "expected one of [$end, ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{37, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{38, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{39, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{92, -1}:  "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{181, -1}: "expected one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{290, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{358, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{379, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{389, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{390, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{393, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{396, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{403, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{404, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{405, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{406, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{407, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{408, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{409, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{428, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{429, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{430, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{432, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{440, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{446, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{452, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{456, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{460, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{464, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{466, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{467, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{471, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{474, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{476, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{413, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{415, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{416, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{468, -1}: "expected one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{254, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{261, -1}: "expected one of ['!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{293, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{307, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{18, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{19, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{20, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{21, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{22, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{23, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{24, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{25, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{26, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{27, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{28, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{29, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{30, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{31, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{304, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{328, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{333, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{334, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{13, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{14, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{15, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{16, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{17, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{40, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{41, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{335, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{336, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{337, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{338, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{339, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		yyXError{240, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{241, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{242, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		yyXError{201, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{202, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{205, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{214, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{216, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{222, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{225, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{228, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{229, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{82, -1}:  "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{239, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		yyXError{86, -1}:  "expected one of ['(', ')', ',', '[']",
		yyXError{135, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{182, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{186, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{187, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{189, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{197, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{233, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{237, -1}: "expected one of ['(', ')', ',', '[']",
		yyXError{285, -1}: "expected one of ['(', goto]",
		yyXError{359, -1}: "expected one of ['(', goto]",
		yyXError{321, -1}: "expected one of ['(', identifier]",
		yyXError{367, -1}: "expected one of [')', ',', ':']",
		yyXError{374, -1}: "expected one of [')', ',', ':']",
		yyXError{380, -1}: "expected one of [')', ',', ':']",
		yyXError{381, -1}: "expected one of [')', ',', ':']",
		yyXError{385, -1}: "expected one of [')', ',', ':']",
		yyXError{388, -1}: "expected one of [')', ',', ':']",
		yyXError{392, -1}: "expected one of [')', ',', ':']",
		yyXError{402, -1}: "expected one of [')', ',', ';']",
		yyXError{210, -1}: "expected one of [')', ',', ...]",
		yyXError{213, -1}: "expected one of [')', ',', ...]",
		yyXError{527, -1}: "expected one of [')', ',', ...]",
		yyXError{85, -1}:  "expected one of [')', ',']",
		yyXError{174, -1}: "expected one of [')', ',']",
		yyXError{192, -1}: "expected one of [')', ',']",
		yyXError{194, -1}: "expected one of [')', ',']",
		yyXError{199, -1}: "expected one of [')', ',']",
		yyXError{200, -1}: "expected one of [')', ',']",
		yyXError{211, -1}: "expected one of [')', ',']",
		yyXError{232, -1}: "expected one of [')', ',']",
		yyXError{246, -1}: "expected one of [')', ',']",
		yyXError{378, -1}: "expected one of [')', ',']",
		yyXError{395, -1}: "expected one of [')', ',']",
		yyXError{450, -1}: "expected one of [')', ',']",
		yyXError{454, -1}: "expected one of [')', ',']",
		yyXError{458, -1}: "expected one of [')', ',']",
		yyXError{462, -1}: "expected one of [')', ',']",
		yyXError{287, -1}: "expected one of [')', ':', string literal]",
		yyXError{289, -1}: "expected one of [')', ':', string literal]",
		yyXError{386, -1}: "expected one of [')', ':', string literal]",
		yyXError{288, -1}: "expected one of [')', string literal]",
		yyXError{318, -1}: "expected one of [',', ':', ';']",
		yyXError{148, -1}: "expected one of [',', ':']",
		yyXError{368, -1}: "expected one of [',', ':']",
		yyXError{375, -1}: "expected one of [',', ':']",
		yyXError{351, -1}: "expected one of [',', ';', '=']",
		yyXError{269, -1}: "expected one of [',', ';', '}']",
		yyXError{315, -1}: "expected one of [',', ';']",
		yyXError{317, -1}: "expected one of [',', ';']",
		yyXError{324, -1}: "expected one of [',', ';']",
		yyXError{327, -1}: "expected one of [',', ';']",
		yyXError{341, -1}: "expected one of [',', ';']",
		yyXError{342, -1}: "expected one of [',', ';']",
		yyXError{478, -1}: "expected one of [',', ';']",
		yyXError{480, -1}: "expected one of [',', ';']",
		yyXError{296, -1}: "expected one of [',', '=', '}']",
		yyXError{299, -1}: "expected one of [',', '=', '}']",
		yyXError{179, -1}: "expected one of [',', ']']",
		yyXError{265, -1}: "expected one of [',', '}']",
		yyXError{271, -1}: "expected one of [',', '}']",
		yyXError{298, -1}: "expected one of [',', '}']",
		yyXError{301, -1}: "expected one of [',', '}']",
		yyXError{305, -1}: "expected one of [',', '}']",
		yyXError{255, -1}: "expected one of ['.', '=', '[']",
		yyXError{258, -1}: "expected one of ['.', '=', '[']",
		yyXError{260, -1}: "expected one of ['.', '=', '[']",
		yyXError{262, -1}: "expected one of ['.', '=', '[']",
		yyXError{363, -1}: "expected one of [':', string literal]",
		yyXError{497, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		yyXError{506, -1}: "expected one of ['\\n', ppother]",
		yyXError{509, -1}: "expected one of ['\\n', ppother]",
		yyXError{511, -1}: "expected one of ['\\n', ppother]",
		yyXError{348, -1}: "expected one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{350, -1}: "expected one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{34, -1}:  "expected one of ['{', identifier]",
		yyXError{35, -1}:  "expected one of ['{', identifier]",
		yyXError{312, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{314, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{316, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{325, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{329, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{535, -1}: "expected one of [..., identifier]",
		yyXError{80, -1}:  "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		yyXError{100, -1}: "expected optional argument expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{398, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{400, -1}: "expected optional block item list or one of ['!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		yyXError{251, -1}: "expected optional comma or one of [',', '}']",
		yyXError{266, -1}: "expected optional comma or one of [',', '}']",
		yyXError{297, -1}: "expected optional comma or one of [',', '}']",
		yyXError{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{11, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{12, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{437, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{443, -1}: "expected optional expression list or one of ['!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{426, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{435, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{441, -1}: "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{217, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		yyXError{206, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{89, -1}:  "expected optional expression or type qualifier list or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{93, -1}:  "expected optional expression or type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{544, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{549, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{552, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{559, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{565, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{207, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{33, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{36, -1}:  "expected optional identifier or one of ['{', identifier]",
		yyXError{347, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		yyXError{191, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{234, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{235, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{78, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{79, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{498, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{502, -1}: "expected optional token list or one of ['\\n', ppother]",
		yyXError{81, -1}:  "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		yyXError{357, -1}: "expected optional volatile or one of ['(', goto, volatile]",
		yyXError{47, -1}:  "expected optional volatile or one of ['(', volatile]",
		yyXError{230, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{208, -1}: "expected parameter type list or one of [_Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{238, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		yyXError{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{485, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		yyXError{524, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{530, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{533, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{537, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{540, -1}: "expected replacement list or one of ['\\n', ppother]",
		yyXError{77, -1}:  "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{421, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{439, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{445, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{455, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{459, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{463, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{465, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{470, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{473, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{475, -1}: "expected statement or one of ['!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		yyXError{52, -1}:  "expected string literal",
		yyXError{366, -1}: "expected string literal",
		yyXError{371, -1}: "expected string literal",
		yyXError{376, -1}: "expected string literal",
		yyXError{309, -1}: "expected struct declaration list or one of [_Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{310, -1}: "expected struct declaration list or one of [_Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{311, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{313, -1}: "expected struct declarator list or one of ['(', '*', ':', ';', identifier]",
		yyXError{326, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		yyXError{514, -1}: "expected token list or one of ['\\n', ppother]",
		yyXError{491, -1}: "expected token list or ppother",
		yyXError{494, -1}: "expected token list or ppother",
		yyXError{500, -1}: "expected token list or ppother",
		yyXError{501, -1}: "expected token list or ppother",
		yyXError{504, -1}: "expected token list or ppother",
		yyXError{557, -1}: "expected token list or ppother",
		yyXError{4, -1}:   "expected translation unit or one of [';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{5, -1}:   "expected translation unit or one of [';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{75, -1}:  "expected type name or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		yyXError{96, -1}:  "expected type qualifier or one of ['!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		yyXError{219, -1}: "expected type qualifier or one of ['!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		yyXError{447, -1}: "expected while",
		yyXError{3, 50}:   "unexpected EOF",
		yyXError{2, 50}:   "unexpected EOF",
		yyXError{4, 50}:   "unexpected EOF",
	}

	yyParseTab = [571][]uint16{
		// 0
		{215: 321, 226: 320, 228: 319, 231: 322},
		{50: 318},
		{84: 317, 86: 317, 100: 317, 317, 317, 317, 317, 317, 317, 317, 317, 317, 317, 317, 198: 803},
		{315, 315, 315, 315, 315, 315, 315, 315, 11: 315, 13: 315, 315, 315, 315, 315, 315, 315, 315, 315, 205: 801},
		{8: 313, 22: 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 43: 313, 313, 313, 313, 313, 313, 313, 56: 313, 313, 206: 323},
		// 5
		{8: 364, 22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 43: 333, 334, 332, 358, 359, 335, 331, 56: 365, 366, 113: 329, 117: 348, 351, 347, 328, 122: 363, 325, 330, 126: 327, 137: 326, 144: 362, 186: 360, 361, 232: 324},
		{8: 364, 22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 43: 333, 334, 332, 358, 359, 335, 331, 312, 56: 365, 366, 113: 329, 117: 348, 351, 347, 328, 122: 363, 325, 330, 126: 327, 137: 326, 144: 362, 186: 800, 361},
		{158, 399, 158, 8: 231, 135: 639, 638, 139: 661, 170: 659, 190: 660, 658},
		{240, 240, 240, 240, 240, 240, 240, 240, 240, 11: 240, 13: 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 43: 240, 240, 240, 240, 240, 240, 240, 240, 52: 240, 56: 240, 240, 85: 240, 87: 240, 240, 240, 240, 240, 240, 240, 240, 240, 240},
		{235, 235, 235, 8: 235, 235, 235, 12: 235, 22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 43: 333, 334, 332, 358, 359, 335, 331, 113: 329, 117: 348, 351, 347, 328, 123: 654, 330, 126: 327, 157: 657},
		// 10
		{235, 235, 235, 8: 235, 235, 235, 12: 235, 22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 43: 333, 334, 332, 358, 359, 335, 331, 113: 329, 117: 348, 351, 347, 328, 123: 654, 330, 126: 327, 157: 656},
		{235, 235, 235, 8: 235, 235, 235, 12: 235, 22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 43: 333, 334, 332, 358, 359, 335, 331, 113: 329, 117: 348, 351, 347, 328, 123: 654, 330, 126: 327, 157: 655},
		{235, 235, 235, 8: 235, 235, 235, 12: 235, 22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 43: 333, 334, 332, 358, 359, 335, 331, 113: 329, 117: 348, 351, 347, 328, 123: 654, 330, 126: 327, 157: 653},
		{226, 226, 226, 8: 226, 226, 226, 12: 226, 22: 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 43: 226, 226, 226, 226, 226, 226, 226},
		{225, 225, 225, 8: 225, 225, 225, 12: 225, 22: 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 43: 225, 225, 225, 225, 225, 225, 225},
		// 15
		{224, 224, 224, 8: 224, 224, 224, 12: 224, 22: 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 43: 224, 224, 224, 224, 224, 224, 224},
		{223, 223, 223, 8: 223, 223, 223, 12: 223, 22: 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 43: 223, 223, 223, 223, 223, 223, 223},
		{222, 222, 222, 8: 222, 222, 222, 12: 222, 22: 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 43: 222, 222, 222, 222, 222, 222, 222},
		{221, 221, 221, 8: 221, 221, 221, 12: 221, 22: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 42: 221, 221, 221, 221, 221, 221, 221, 221},
		{220, 220, 220, 8: 220, 220, 220, 12: 220, 22: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 42: 220, 220, 220, 220, 220, 220, 220, 220},
		// 20
		{219, 219, 219, 8: 219, 219, 219, 12: 219, 22: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 42: 219, 219, 219, 219, 219, 219, 219, 219},
		{218, 218, 218, 8: 218, 218, 218, 12: 218, 22: 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 42: 218, 218, 218, 218, 218, 218, 218, 218},
		{217, 217, 217, 8: 217, 217, 217, 12: 217, 22: 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 42: 217, 217, 217, 217, 217, 217, 217, 217},
		{216, 216, 216, 8: 216, 216, 216, 12: 216, 22: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 42: 216, 216, 216, 216, 216, 216, 216, 216},
		{215, 215, 215, 8: 215, 215, 215, 12: 215, 22: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 42: 215, 215, 215, 215, 215, 215, 215, 215},
		// 25
		{214, 214, 214, 8: 214, 214, 214, 12: 214, 22: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 42: 214, 214, 214, 214, 214, 214, 214, 214},
		{213, 213, 213, 8: 213, 213, 213, 12: 213, 22: 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 42: 213, 213, 213, 213, 213, 213, 213, 213},
		{212, 212, 212, 8: 212, 212, 212, 12: 212, 22: 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 42: 212, 212, 212, 212, 212, 212, 212, 212},
		{211, 211, 211, 8: 211, 211, 211, 12: 211, 22: 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 42: 211, 211, 211, 211, 211, 211, 211, 211},
		{210, 210, 210, 8: 210, 210, 210, 12: 210, 22: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 42: 210, 210, 210, 210, 210, 210, 210, 210},
		// 30
		{209, 209, 209, 8: 209, 209, 209, 12: 209, 22: 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 42: 209, 209, 209, 209, 209, 209, 209, 209},
		{208, 208, 208, 8: 208, 208, 208, 12: 208, 22: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 42: 208, 208, 208, 208, 208, 208, 208, 208},
		{648},
		{2: 625, 52: 140, 189: 624},
		{2: 201, 52: 201},
		// 35
		{2: 200, 52: 200},
		{2: 611, 52: 140, 189: 610},
		{177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 177, 42: 177, 177, 177, 177, 177, 177, 177, 177, 54: 177},
		{176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 42: 176, 176, 176, 176, 176, 176, 176, 176, 54: 176},
		{175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 42: 175, 175, 175, 175, 175, 175, 175, 175, 54: 175},
		// 40
		{174, 174, 174, 8: 174, 174, 174, 12: 174, 22: 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 43: 174, 174, 174, 174, 174, 174, 174},
		{173, 173, 173, 8: 173, 173, 173, 12: 173, 22: 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 43: 173, 173, 173, 173, 173, 173, 173},
		{8: 79, 22: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 43: 79, 79, 79, 79, 79, 79, 79, 79, 56: 79, 79},
		{8: 77, 22: 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 43: 77, 77, 77, 77, 77, 77, 77, 77, 56: 77, 77},
		{8: 76, 22: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 43: 76, 76, 76, 76, 76, 76, 76, 76, 56: 76, 76},
		// 45
		{8: 609},
		{8: 74, 22: 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 74, 43: 74, 74, 74, 74, 74, 74, 74, 74, 56: 74, 74},
		{59, 22: 603, 197: 602},
		{367},
		{243, 243, 243, 243, 243, 243, 243, 243, 11: 243, 13: 243, 243, 243, 243, 243, 243, 243, 243, 243, 145: 368, 369},
		// 50
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 382},
		{10: 370},
		{11: 371},
		{9: 372},
		{8: 373},
		// 55
		{44, 44, 44, 44, 44, 44, 44, 44, 44, 11: 44, 13: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 43: 44, 44, 44, 44, 44, 44, 44, 44, 52: 44, 56: 44, 44, 85: 44, 87: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44},
		{306, 306, 3: 306, 306, 306, 306, 306, 306, 306, 306, 12: 306, 41: 306, 306, 50: 306, 306, 54: 306, 306, 58: 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306},
		{305, 305, 3: 305, 305, 305, 305, 305, 305, 305, 305, 12: 305, 41: 305, 305, 50: 305, 305, 54: 305, 305, 58: 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305},
		{304, 304, 3: 304, 304, 304, 304, 304, 304, 304, 304, 12: 304, 41: 304, 304, 50: 304, 304, 54: 304, 304, 58: 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304},
		{303, 303, 3: 303, 303, 303, 303, 303, 303, 303, 303, 12: 303, 41: 303, 303, 50: 303, 303, 54: 303, 303, 58: 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303},
		// 60
		{302, 302, 3: 302, 302, 302, 302, 302, 302, 302, 302, 12: 302, 41: 302, 302, 50: 302, 302, 54: 302, 302, 58: 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302},
		{301, 301, 3: 301, 301, 301, 301, 301, 301, 301, 301, 12: 301, 41: 301, 301, 50: 301, 301, 54: 301, 301, 58: 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301},
		{300, 300, 3: 300, 300, 300, 300, 300, 300, 300, 300, 12: 300, 41: 300, 300, 50: 300, 300, 54: 300, 300, 58: 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 53: 465, 114: 564, 153: 395, 160: 600},
		{418, 423, 3: 436, 426, 427, 422, 421, 242, 10: 242, 12: 417, 41: 242, 242, 50: 242, 442, 54: 242, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		// 65
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 599},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 598},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 597},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 506},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 596},
		// 70
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 595},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 594},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 593},
		{562, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 563},
		{393},
		// 75
		{22: 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 153: 395, 160: 394},
		{9: 561},
		{22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 113: 397, 117: 348, 351, 347, 396, 151: 398},
		{192, 192, 192, 8: 192, 192, 12: 192, 22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 42: 192, 113: 397, 117: 348, 351, 347, 396, 151: 559, 194: 560},
		{192, 192, 192, 8: 192, 192, 12: 192, 22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 42: 192, 113: 397, 117: 348, 351, 347, 396, 151: 559, 194: 558},
		// 80
		{158, 399, 9: 134, 12: 158, 135: 400, 402, 154: 403, 176: 401},
		{154, 154, 154, 9: 154, 154, 12: 154, 22: 357, 355, 356, 113: 410, 152: 414, 161: 556},
		{157, 2: 157, 9: 136, 136, 12: 157},
		{9: 137},
		{405, 12: 122, 180: 404, 406},
		// 85
		{9: 133, 133},
		{552, 9: 135, 135, 12: 121},
		{158, 399, 9: 126, 12: 158, 22: 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 43: 126, 126, 126, 126, 126, 126, 126, 135: 400, 402, 154: 508, 174: 509},
		{12: 407},
		{381, 409, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 357, 355, 356, 43: 413, 53: 408, 249, 113: 410, 152: 411, 168: 412},
		// 90
		{418, 423, 3: 436, 426, 427, 422, 421, 12: 417, 51: 442, 54: 248, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 506, 507},
		{156, 156, 156, 156, 156, 156, 156, 156, 9: 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 43: 156, 54: 156},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 357, 355, 356, 43: 502, 53: 408, 249, 113: 499, 168: 501},
		{54: 500},
		// 95
		{154, 154, 154, 154, 154, 154, 154, 154, 11: 154, 13: 154, 154, 154, 154, 154, 154, 154, 154, 154, 357, 355, 356, 113: 410, 152: 414, 161: 415},
		{153, 153, 153, 153, 153, 153, 153, 153, 9: 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 357, 355, 356, 113: 499},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 416},
		{418, 423, 3: 436, 426, 427, 422, 421, 12: 417, 51: 442, 54: 453, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 465, 114: 497},
		// 100
		{381, 386, 374, 385, 387, 388, 384, 383, 9: 308, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 491, 211: 492, 493},
		{2: 490},
		{2: 489},
		{294, 294, 3: 294, 294, 294, 294, 294, 294, 294, 294, 12: 294, 41: 294, 294, 50: 294, 294, 54: 294, 294, 58: 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294},
		{293, 293, 3: 293, 293, 293, 293, 293, 293, 293, 293, 12: 293, 41: 293, 293, 50: 293, 293, 54: 293, 293, 58: 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293},
		// 105
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 488},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 487},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 486},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 485},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 484},
		// 110
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 483},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 482},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 481},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 480},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 479},
		// 115
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 478},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 477},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 476},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 475},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 474},
		// 120
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 473},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 472},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 471},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 465, 114: 466},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 464},
		// 125
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 463},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 462},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 461},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 460},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 459},
		// 130
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 458},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 457},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 456},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 455},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 454},
		// 135
		{129, 9: 129, 129, 12: 129},
		{418, 423, 3: 436, 426, 427, 422, 421, 251, 251, 251, 12: 417, 41: 251, 251, 50: 251, 442, 54: 251, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{418, 423, 3: 436, 426, 427, 422, 421, 252, 252, 252, 12: 417, 41: 252, 252, 50: 252, 442, 54: 252, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{418, 423, 3: 436, 426, 427, 422, 421, 253, 253, 253, 12: 417, 41: 253, 253, 50: 253, 442, 54: 253, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{418, 423, 3: 436, 426, 427, 422, 421, 254, 254, 254, 12: 417, 41: 254, 254, 50: 254, 442, 54: 254, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		// 140
		{418, 423, 3: 436, 426, 427, 422, 421, 255, 255, 255, 12: 417, 41: 255, 255, 50: 255, 442, 54: 255, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{418, 423, 3: 436, 426, 427, 422, 421, 256, 256, 256, 12: 417, 41: 256, 256, 50: 256, 442, 54: 256, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{418, 423, 3: 436, 426, 427, 422, 421, 257, 257, 257, 12: 417, 41: 257, 257, 50: 257, 442, 54: 257, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{418, 423, 3: 436, 426, 427, 422, 421, 258, 258, 258, 12: 417, 41: 258, 258, 50: 258, 442, 54: 258, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{418, 423, 3: 436, 426, 427, 422, 421, 259, 259, 259, 12: 417, 41: 259, 259, 50: 259, 442, 54: 259, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		// 145
		{418, 423, 3: 436, 426, 427, 422, 421, 260, 260, 260, 12: 417, 41: 260, 260, 50: 260, 442, 54: 260, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{418, 423, 3: 436, 426, 427, 422, 421, 261, 261, 261, 12: 417, 41: 261, 261, 50: 261, 442, 54: 261, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{418, 423, 3: 436, 426, 427, 422, 421, 247, 247, 247, 12: 417, 42: 247, 51: 442, 54: 247, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{10: 468, 42: 467},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 470},
		// 150
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 469},
		{418, 423, 3: 436, 426, 427, 422, 421, 246, 246, 246, 12: 417, 42: 246, 51: 442, 54: 246, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{418, 423, 3: 436, 426, 427, 422, 421, 262, 262, 262, 12: 417, 41: 262, 262, 50: 262, 262, 54: 262, 419, 58: 425, 424, 430, 431, 441, 437, 438, 262, 439, 262, 420, 262, 434, 433, 432, 428, 262, 262, 262, 435, 262, 440, 429, 262, 262, 262},
		{418, 423, 3: 436, 426, 427, 422, 421, 263, 263, 263, 12: 417, 41: 263, 263, 50: 263, 263, 54: 263, 419, 58: 425, 424, 430, 431, 263, 437, 438, 263, 439, 263, 420, 263, 434, 433, 432, 428, 263, 263, 263, 435, 263, 263, 429, 263, 263, 263},
		{418, 423, 3: 436, 426, 427, 422, 421, 264, 264, 264, 12: 417, 41: 264, 264, 50: 264, 264, 54: 264, 419, 58: 425, 424, 430, 431, 264, 437, 438, 264, 264, 264, 420, 264, 434, 433, 432, 428, 264, 264, 264, 435, 264, 264, 429, 264, 264, 264},
		// 155
		{418, 423, 3: 436, 426, 427, 422, 421, 265, 265, 265, 12: 417, 41: 265, 265, 50: 265, 265, 54: 265, 419, 58: 425, 424, 430, 431, 265, 437, 265, 265, 265, 265, 420, 265, 434, 433, 432, 428, 265, 265, 265, 435, 265, 265, 429, 265, 265, 265},
		{418, 423, 3: 436, 426, 427, 422, 421, 266, 266, 266, 12: 417, 41: 266, 266, 50: 266, 266, 54: 266, 419, 58: 425, 424, 430, 431, 266, 266, 266, 266, 266, 266, 420, 266, 434, 433, 432, 428, 266, 266, 266, 435, 266, 266, 429, 266, 266, 266},
		{418, 423, 3: 267, 426, 427, 422, 421, 267, 267, 267, 12: 417, 41: 267, 267, 50: 267, 267, 54: 267, 419, 58: 425, 424, 430, 431, 267, 267, 267, 267, 267, 267, 420, 267, 434, 433, 432, 428, 267, 267, 267, 435, 267, 267, 429, 267, 267, 267},
		{418, 423, 3: 268, 426, 427, 422, 421, 268, 268, 268, 12: 417, 41: 268, 268, 50: 268, 268, 54: 268, 419, 58: 425, 424, 430, 431, 268, 268, 268, 268, 268, 268, 420, 268, 268, 433, 432, 428, 268, 268, 268, 268, 268, 268, 429, 268, 268, 268},
		{418, 423, 3: 269, 426, 427, 422, 421, 269, 269, 269, 12: 417, 41: 269, 269, 50: 269, 269, 54: 269, 419, 58: 425, 424, 430, 431, 269, 269, 269, 269, 269, 269, 420, 269, 269, 433, 432, 428, 269, 269, 269, 269, 269, 269, 429, 269, 269, 269},
		// 160
		{418, 423, 3: 270, 426, 427, 422, 421, 270, 270, 270, 12: 417, 41: 270, 270, 50: 270, 270, 54: 270, 419, 58: 425, 424, 270, 270, 270, 270, 270, 270, 270, 270, 420, 270, 270, 270, 270, 428, 270, 270, 270, 270, 270, 270, 429, 270, 270, 270},
		{418, 423, 3: 271, 426, 427, 422, 421, 271, 271, 271, 12: 417, 41: 271, 271, 50: 271, 271, 54: 271, 419, 58: 425, 424, 271, 271, 271, 271, 271, 271, 271, 271, 420, 271, 271, 271, 271, 428, 271, 271, 271, 271, 271, 271, 429, 271, 271, 271},
		{418, 423, 3: 272, 426, 427, 422, 421, 272, 272, 272, 12: 417, 41: 272, 272, 50: 272, 272, 54: 272, 419, 58: 425, 424, 272, 272, 272, 272, 272, 272, 272, 272, 420, 272, 272, 272, 272, 428, 272, 272, 272, 272, 272, 272, 429, 272, 272, 272},
		{418, 423, 3: 273, 426, 427, 422, 421, 273, 273, 273, 12: 417, 41: 273, 273, 50: 273, 273, 54: 273, 419, 58: 425, 424, 273, 273, 273, 273, 273, 273, 273, 273, 420, 273, 273, 273, 273, 428, 273, 273, 273, 273, 273, 273, 429, 273, 273, 273},
		{418, 423, 3: 274, 426, 427, 422, 421, 274, 274, 274, 12: 417, 41: 274, 274, 50: 274, 274, 54: 274, 419, 58: 425, 424, 274, 274, 274, 274, 274, 274, 274, 274, 420, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274},
		// 165
		{418, 423, 3: 275, 426, 427, 422, 421, 275, 275, 275, 12: 417, 41: 275, 275, 50: 275, 275, 54: 275, 419, 58: 425, 424, 275, 275, 275, 275, 275, 275, 275, 275, 420, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275},
		{418, 423, 3: 276, 276, 276, 422, 421, 276, 276, 276, 12: 417, 41: 276, 276, 50: 276, 276, 54: 276, 419, 58: 425, 424, 276, 276, 276, 276, 276, 276, 276, 276, 420, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276},
		{418, 423, 3: 277, 277, 277, 422, 421, 277, 277, 277, 12: 417, 41: 277, 277, 50: 277, 277, 54: 277, 419, 58: 425, 424, 277, 277, 277, 277, 277, 277, 277, 277, 420, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277},
		{418, 278, 3: 278, 278, 278, 422, 421, 278, 278, 278, 12: 417, 41: 278, 278, 50: 278, 278, 54: 278, 419, 58: 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 420, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278, 278},
		{418, 279, 3: 279, 279, 279, 422, 421, 279, 279, 279, 12: 417, 41: 279, 279, 50: 279, 279, 54: 279, 419, 58: 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 420, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279, 279},
		// 170
		{418, 280, 3: 280, 280, 280, 422, 421, 280, 280, 280, 12: 417, 41: 280, 280, 50: 280, 280, 54: 280, 419, 58: 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 420, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280, 280},
		{295, 295, 3: 295, 295, 295, 295, 295, 295, 295, 295, 12: 295, 41: 295, 295, 50: 295, 295, 54: 295, 295, 58: 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295},
		{296, 296, 3: 296, 296, 296, 296, 296, 296, 296, 296, 12: 296, 41: 296, 296, 50: 296, 296, 54: 296, 296, 58: 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296},
		{418, 423, 3: 436, 426, 427, 422, 421, 9: 310, 310, 12: 417, 51: 442, 55: 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{9: 307, 495},
		// 175
		{9: 494},
		{297, 297, 3: 297, 297, 297, 297, 297, 297, 297, 297, 12: 297, 41: 297, 297, 50: 297, 297, 54: 297, 297, 58: 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 496},
		{418, 423, 3: 436, 426, 427, 422, 421, 9: 309, 309, 12: 417, 51: 442, 55: 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{10: 468, 54: 498},
		// 180
		{298, 298, 3: 298, 298, 298, 298, 298, 298, 298, 298, 12: 298, 41: 298, 298, 50: 298, 298, 54: 298, 298, 58: 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298},
		{155, 155, 155, 155, 155, 155, 155, 155, 9: 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 43: 155, 54: 155},
		{131, 9: 131, 131, 12: 131},
		{54: 505},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 503},
		// 185
		{418, 423, 3: 436, 426, 427, 422, 421, 12: 417, 51: 442, 54: 504, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{128, 9: 128, 128, 12: 128},
		{130, 9: 130, 130, 12: 130},
		{418, 288, 3: 288, 288, 288, 422, 421, 288, 288, 288, 12: 417, 41: 288, 288, 50: 288, 288, 54: 288, 419, 58: 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 420, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288},
		{127, 9: 127, 127, 12: 127},
		// 190
		{9: 551},
		{9: 150, 22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 43: 333, 334, 332, 358, 359, 335, 331, 113: 329, 117: 348, 351, 347, 328, 123: 513, 330, 126: 327, 159: 512, 172: 510, 511, 193: 514},
		{9: 152, 548},
		{9: 149},
		{9: 148, 148},
		// 195
		{158, 399, 158, 9: 134, 134, 12: 158, 135: 400, 516, 139: 517, 154: 403, 176: 518},
		{9: 515},
		{125, 9: 125, 125, 12: 125},
		{521, 2: 520, 12: 122, 180: 404, 406, 519},
		{9: 146, 146},
		// 200
		{9: 145, 145},
		{525, 8: 172, 172, 172, 12: 524, 22: 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 42: 172, 172, 172, 172, 172, 172, 172, 172, 51: 172, 172, 56: 172, 172},
		{169, 8: 169, 169, 169, 12: 169, 22: 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 42: 169, 169, 169, 169, 169, 169, 169, 169, 51: 169, 169, 56: 169, 169},
		{158, 399, 158, 9: 126, 12: 158, 22: 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 43: 126, 126, 126, 126, 126, 126, 126, 135: 400, 516, 139: 522, 154: 508, 174: 509},
		{9: 523},
		// 205
		{168, 8: 168, 168, 168, 12: 168, 22: 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 42: 168, 168, 168, 168, 168, 168, 168, 168, 51: 168, 168, 56: 168, 168},
		{154, 154, 154, 154, 154, 154, 154, 154, 11: 154, 13: 154, 154, 154, 154, 154, 154, 154, 154, 154, 357, 355, 356, 43: 536, 54: 154, 113: 410, 152: 537, 161: 535},
		{2: 528, 9: 142, 22: 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 43: 163, 163, 163, 163, 163, 163, 163, 169: 529, 188: 527, 210: 526},
		{22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 43: 333, 334, 332, 358, 359, 335, 331, 113: 329, 117: 348, 351, 347, 328, 123: 513, 330, 126: 327, 159: 512, 172: 510, 533},
		{9: 532},
		// 210
		{9: 144, 144, 147: 144},
		{9: 141, 530},
		{2: 531},
		{9: 143, 143, 147: 143},
		{161, 8: 161, 161, 161, 12: 161, 22: 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 42: 161, 161, 161, 161, 161, 161, 161, 161, 51: 161, 161, 56: 161, 161},
		// 215
		{9: 534},
		{162, 8: 162, 162, 162, 12: 162, 22: 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 42: 162, 162, 162, 162, 162, 162, 162, 162, 51: 162, 162, 56: 162, 162},
		{381, 544, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 408, 249, 168: 545},
		{154, 154, 154, 154, 154, 154, 154, 154, 11: 154, 13: 154, 154, 154, 154, 154, 154, 154, 154, 154, 357, 355, 356, 113: 410, 152: 414, 161: 541},
		{153, 153, 153, 153, 153, 153, 153, 153, 11: 153, 13: 153, 153, 153, 153, 153, 153, 153, 153, 153, 357, 355, 356, 43: 538, 54: 153, 113: 499},
		// 220
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 539},
		{418, 423, 3: 436, 426, 427, 422, 421, 12: 417, 51: 442, 54: 540, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{165, 8: 165, 165, 165, 12: 165, 22: 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 42: 165, 165, 165, 165, 165, 165, 165, 165, 51: 165, 165, 56: 165, 165},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 542},
		{418, 423, 3: 436, 426, 427, 422, 421, 12: 417, 51: 442, 54: 543, 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		// 225
		{166, 8: 166, 166, 166, 12: 166, 22: 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 42: 166, 166, 166, 166, 166, 166, 166, 166, 51: 166, 166, 56: 166, 166},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 506, 547},
		{54: 546},
		{167, 8: 167, 167, 167, 12: 167, 22: 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 42: 167, 167, 167, 167, 167, 167, 167, 167, 51: 167, 167, 56: 167, 167},
		{164, 8: 164, 164, 164, 12: 164, 22: 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 42: 164, 164, 164, 164, 164, 164, 164, 164, 51: 164, 164, 56: 164, 164},
		// 230
		{22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 43: 333, 334, 332, 358, 359, 335, 331, 113: 329, 117: 348, 351, 347, 328, 123: 513, 330, 126: 327, 147: 549, 159: 550},
		{9: 151},
		{9: 147, 147},
		{132, 9: 132, 132, 12: 132},
		{9: 124, 22: 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 43: 124, 124, 124, 124, 124, 124, 124, 199: 553},
		// 235
		{9: 150, 22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 43: 333, 334, 332, 358, 359, 335, 331, 113: 329, 117: 348, 351, 347, 328, 123: 513, 330, 126: 327, 159: 512, 172: 510, 511, 193: 554},
		{9: 555},
		{123, 9: 123, 123, 12: 123},
		{160, 399, 160, 9: 160, 160, 12: 160, 135: 557},
		{159, 2: 159, 9: 159, 159, 12: 159},
		// 240
		{193, 193, 193, 8: 193, 193, 12: 193, 42: 193},
		{191, 191, 191, 8: 191, 191, 12: 191, 42: 191},
		{194, 194, 194, 8: 194, 194, 12: 194, 42: 194},
		{250, 250, 3: 250, 250, 250, 250, 250, 250, 250, 250, 12: 250, 41: 250, 250, 50: 250, 250, 54: 250, 250, 58: 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250, 250},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 53: 465, 114: 564, 153: 395, 160: 565},
		// 245
		{418, 283, 3: 283, 283, 283, 422, 421, 283, 283, 283, 12: 417, 41: 283, 283, 50: 283, 283, 54: 283, 419, 58: 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 420, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283, 283},
		{9: 592, 468},
		{9: 566},
		{381, 282, 374, 282, 282, 282, 384, 383, 282, 282, 282, 380, 282, 390, 389, 392, 375, 376, 377, 378, 379, 391, 41: 282, 282, 50: 282, 282, 568, 567, 282, 282, 58: 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282, 282},
		{418, 281, 3: 281, 281, 281, 422, 421, 281, 281, 281, 12: 417, 41: 281, 281, 50: 281, 281, 54: 281, 419, 58: 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 420, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281, 281},
		// 250
		{115, 115, 115, 115, 115, 115, 115, 115, 11: 115, 574, 115, 115, 115, 115, 115, 115, 115, 115, 115, 52: 115, 55: 575, 158: 573, 165: 572, 570, 571, 192: 569},
		{10: 585, 41: 186, 164: 590},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 52: 582, 581, 171: 583},
		{12: 574, 51: 579, 55: 575, 158: 580},
		{114, 114, 114, 114, 114, 114, 114, 114, 11: 114, 13: 114, 114, 114, 114, 114, 114, 114, 114, 114, 52: 114},
		// 255
		{12: 113, 51: 113, 55: 113},
		{243, 243, 243, 243, 243, 243, 243, 243, 11: 243, 13: 243, 243, 243, 243, 243, 243, 243, 243, 243, 145: 368, 577},
		{2: 576},
		{12: 110, 51: 110, 55: 110},
		{54: 578},
		// 260
		{12: 111, 51: 111, 55: 111},
		{116, 116, 116, 116, 116, 116, 116, 116, 11: 116, 13: 116, 116, 116, 116, 116, 116, 116, 116, 116, 52: 116},
		{12: 112, 51: 112, 55: 112},
		{418, 423, 3: 436, 426, 427, 422, 421, 120, 10: 120, 12: 417, 41: 120, 51: 442, 55: 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{115, 115, 115, 115, 115, 115, 115, 115, 11: 115, 574, 115, 115, 115, 115, 115, 115, 115, 115, 115, 52: 115, 55: 575, 158: 573, 165: 572, 570, 571, 192: 584},
		// 265
		{10: 118, 41: 118},
		{10: 585, 41: 186, 164: 586},
		{115, 115, 115, 115, 115, 115, 115, 115, 11: 115, 574, 115, 115, 115, 115, 115, 115, 115, 115, 115, 41: 185, 52: 115, 55: 575, 158: 573, 165: 572, 588, 571},
		{41: 587},
		{8: 119, 10: 119, 41: 119},
		// 270
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 52: 582, 581, 171: 589},
		{10: 117, 41: 117},
		{41: 591},
		{292, 292, 3: 292, 292, 292, 292, 292, 292, 292, 292, 12: 292, 41: 292, 292, 50: 292, 292, 54: 292, 292, 58: 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292},
		{299, 299, 3: 299, 299, 299, 299, 299, 299, 299, 299, 12: 299, 41: 299, 299, 50: 299, 299, 54: 299, 299, 58: 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299},
		// 275
		{418, 284, 3: 284, 284, 284, 422, 421, 284, 284, 284, 12: 417, 41: 284, 284, 50: 284, 284, 54: 284, 419, 58: 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 420, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284},
		{418, 285, 3: 285, 285, 285, 422, 421, 285, 285, 285, 12: 417, 41: 285, 285, 50: 285, 285, 54: 285, 419, 58: 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 420, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285},
		{418, 286, 3: 286, 286, 286, 422, 421, 286, 286, 286, 12: 417, 41: 286, 286, 50: 286, 286, 54: 286, 419, 58: 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 420, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286},
		{418, 287, 3: 287, 287, 287, 422, 421, 287, 287, 287, 12: 417, 41: 287, 287, 50: 287, 287, 54: 287, 419, 58: 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 420, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287},
		{418, 289, 3: 289, 289, 289, 422, 421, 289, 289, 289, 12: 417, 41: 289, 289, 50: 289, 289, 54: 289, 419, 58: 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 420, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289},
		// 280
		{418, 290, 3: 290, 290, 290, 422, 421, 290, 290, 290, 12: 417, 41: 290, 290, 50: 290, 290, 54: 290, 419, 58: 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 420, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290},
		{418, 291, 3: 291, 291, 291, 422, 421, 291, 291, 291, 12: 417, 41: 291, 291, 50: 291, 291, 54: 291, 419, 58: 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 420, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291},
		{9: 601},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 52: 568, 567},
		{604},
		// 285
		{58, 85: 58},
		{11: 605, 162: 606},
		{9: 62, 11: 62, 42: 62},
		{9: 608, 11: 607},
		{9: 61, 11: 61, 42: 61},
		// 290
		{60, 60, 60, 60, 60, 60, 60, 60, 60, 11: 60, 13: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 43: 60, 60, 60, 60, 60, 60, 60, 52: 60, 56: 60, 60, 85: 60, 87: 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 112: 60},
		{8: 75, 22: 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 75, 43: 75, 75, 75, 75, 75, 75, 75, 75, 56: 75, 75},
		{52: 184, 209: 612},
		{182, 182, 182, 8: 182, 182, 182, 12: 182, 22: 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 42: 182, 182, 182, 182, 182, 182, 182, 182, 52: 139},
		{52: 613},
		// 295
		{2: 614, 184: 617, 616, 223: 615},
		{10: 311, 41: 311, 51: 311},
		{10: 620, 41: 186, 164: 621},
		{10: 181, 41: 181},
		{10: 179, 41: 179, 51: 618},
		// 300
		{243, 243, 243, 243, 243, 243, 243, 243, 11: 243, 13: 243, 243, 243, 243, 243, 243, 243, 243, 243, 145: 368, 619},
		{10: 178, 41: 178},
		{2: 614, 41: 185, 184: 617, 623},
		{41: 622},
		{183, 183, 183, 8: 183, 183, 183, 12: 183, 22: 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 42: 183, 183, 183, 183, 183, 183, 183, 183},
		// 305
		{10: 180, 41: 180},
		{52: 205, 207: 626},
		{202, 202, 202, 8: 202, 202, 202, 12: 202, 22: 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 42: 202, 202, 202, 202, 202, 202, 202, 202, 52: 139},
		{52: 627},
		{22: 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 57: 204, 208: 628},
		// 310
		{22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 57: 366, 113: 397, 117: 348, 351, 347, 396, 137: 632, 151: 631, 195: 630, 229: 629},
		{22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 646, 57: 366, 113: 397, 117: 348, 351, 347, 396, 137: 632, 151: 631, 195: 647},
		{22: 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 199, 57: 199},
		{158, 399, 158, 8: 634, 42: 171, 135: 639, 638, 139: 636, 179: 637, 196: 635, 230: 633},
		{22: 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 57: 195},
		// 315
		{8: 643, 10: 644},
		{22: 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 57: 196},
		{8: 190, 10: 190},
		{8: 188, 10: 188, 42: 170},
		{42: 641},
		// 320
		{640, 2: 520, 182: 519},
		{157, 2: 157},
		{158, 399, 158, 135: 639, 638, 139: 522},
		{243, 243, 243, 243, 243, 243, 243, 243, 11: 243, 13: 243, 243, 243, 243, 243, 243, 243, 243, 243, 145: 368, 642},
		{8: 187, 10: 187},
		// 325
		{22: 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 197, 57: 197},
		{158, 399, 158, 42: 171, 135: 639, 638, 139: 636, 179: 637, 196: 645},
		{8: 189, 10: 189},
		{203, 203, 203, 8: 203, 203, 203, 12: 203, 22: 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 42: 203, 203, 203, 203, 203, 203, 203, 203},
		{22: 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 198, 57: 198},
		// 330
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 53: 649, 153: 395, 160: 650},
		{418, 423, 3: 436, 426, 427, 422, 421, 9: 652, 12: 417, 51: 442, 55: 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		{9: 651},
		{206, 206, 206, 8: 206, 206, 206, 12: 206, 22: 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 42: 206, 206, 206, 206, 206, 206, 206, 206},
		{207, 207, 207, 8: 207, 207, 207, 12: 207, 22: 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 42: 207, 207, 207, 207, 207, 207, 207, 207},
		// 335
		{236, 236, 236, 8: 236, 236, 236, 12: 236},
		{234, 234, 234, 8: 234, 234, 234, 12: 234},
		{237, 237, 237, 8: 237, 237, 237, 12: 237},
		{238, 238, 238, 8: 238, 238, 238, 12: 238},
		{239, 239, 239, 8: 239, 239, 239, 12: 239},
		// 340
		{8: 799},
		{8: 233, 10: 233},
		{8: 230, 10: 797},
		{8: 229, 10: 229, 22: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 43: 64, 64, 64, 64, 64, 64, 64, 51: 228, 65, 56: 65, 64, 175: 662, 204: 664, 217: 663},
		{51: 795},
		// 345
		{52: 73, 56: 73, 201: 670},
		{22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 43: 333, 334, 332, 358, 359, 335, 331, 57: 366, 113: 329, 117: 348, 351, 347, 328, 123: 665, 330, 126: 327, 137: 326, 144: 666, 216: 667},
		{158, 399, 158, 8: 231, 135: 639, 638, 139: 669, 170: 659, 190: 660, 658},
		{22: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 43: 67, 67, 67, 67, 67, 67, 67, 52: 67, 56: 67, 67},
		{22: 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 43: 333, 334, 332, 358, 359, 335, 331, 52: 63, 56: 63, 366, 113: 329, 117: 348, 351, 347, 328, 123: 665, 330, 126: 327, 137: 326, 144: 668},
		// 350
		{22: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 43: 66, 66, 66, 66, 66, 66, 66, 52: 66, 56: 66, 66},
		{8: 229, 10: 229, 51: 228, 175: 662},
		{52: 71, 56: 69, 202: 672, 673, 224: 671},
		{8: 72, 22: 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 72, 43: 72, 72, 72, 72, 72, 72, 72, 72, 56: 72, 72},
		{52: 716, 128: 717},
		// 355
		{56: 675, 122: 676, 127: 674},
		{8: 715},
		{59, 22: 603, 85: 59, 197: 677},
		{50, 50, 50, 50, 50, 50, 50, 50, 50, 11: 50, 13: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 43: 50, 50, 50, 50, 50, 50, 50, 52: 50, 56: 50, 50, 85: 50, 87: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 112: 50},
		{678, 85: 679},
		// 360
		{11: 605, 162: 704},
		{680},
		{11: 605, 162: 681},
		{11: 607, 42: 682},
		{42: 683},
		// 365
		{11: 54, 687, 155: 685, 684, 163: 686},
		{11: 700},
		{9: 56, 56, 42: 56},
		{10: 690, 42: 691},
		{2: 688},
		// 370
		{54: 689},
		{11: 53},
		{11: 54, 687, 155: 699, 684},
		{11: 692, 178: 693},
		{9: 52, 52, 42: 52},
		// 375
		{10: 694, 42: 695},
		{11: 698},
		{2: 528, 169: 696},
		{9: 697, 530},
		{46, 46, 46, 46, 46, 46, 46, 46, 46, 11: 46, 13: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 43: 46, 46, 46, 46, 46, 46, 46, 52: 46, 56: 46, 46, 85: 46, 87: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 112: 46},
		// 380
		{9: 51, 51, 42: 51},
		{9: 55, 55, 42: 55},
		{701},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 702},
		{418, 423, 3: 436, 426, 427, 422, 421, 9: 703, 12: 417, 51: 442, 55: 419, 58: 425, 424, 430, 431, 441, 437, 438, 446, 439, 450, 420, 444, 434, 433, 432, 428, 448, 445, 443, 435, 452, 440, 429, 449, 447, 451},
		// 385
		{9: 57, 57, 42: 57},
		{9: 608, 11: 607, 42: 705},
		{9: 707, 11: 54, 687, 155: 685, 684, 163: 706},
		{9: 708, 690, 42: 709},
		{45, 45, 45, 45, 45, 45, 45, 45, 45, 11: 45, 13: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 43: 45, 45, 45, 45, 45, 45, 45, 52: 45, 56: 45, 45, 85: 45, 87: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 112: 45},
		// 390
		{49, 49, 49, 49, 49, 49, 49, 49, 49, 11: 49, 13: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 43: 49, 49, 49, 49, 49, 49, 49, 52: 49, 56: 49, 49, 85: 49, 87: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 112: 49},
		{11: 54, 687, 155: 685, 684, 163: 710},
		{9: 711, 690, 42: 712},
		{48, 48, 48, 48, 48, 48, 48, 48, 48, 11: 48, 13: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 43: 48, 48, 48, 48, 48, 48, 48, 52: 48, 56: 48, 48, 85: 48, 87: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 112: 48},
		{11: 692, 178: 713},
		// 395
		{9: 714, 694},
		{47, 47, 47, 47, 47, 47, 47, 47, 47, 11: 47, 13: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 43: 47, 47, 47, 47, 47, 47, 47, 52: 47, 56: 47, 47, 85: 47, 87: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 112: 47},
		{8: 68, 22: 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 43: 68, 68, 68, 68, 68, 68, 68, 68, 56: 68, 68},
		{99, 99, 99, 99, 99, 99, 99, 99, 99, 11: 99, 13: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 43: 99, 99, 99, 99, 99, 99, 99, 52: 99, 56: 99, 99, 85: 99, 87: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 200: 718},
		{8: 70, 22: 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 70, 43: 70, 70, 70, 70, 70, 70, 70, 70, 56: 70, 70},
		// 400
		{381, 386, 719, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 95, 43: 333, 334, 332, 358, 359, 335, 331, 52: 716, 465, 56: 675, 366, 85: 741, 87: 738, 743, 728, 742, 729, 739, 740, 736, 744, 737, 113: 329, 720, 117: 348, 351, 347, 328, 735, 676, 665, 330, 126: 327, 727, 722, 723, 725, 726, 721, 724, 734, 137: 326, 144: 733, 177: 731, 213: 732, 730},
		{306, 306, 3: 306, 306, 306, 306, 306, 306, 10: 306, 12: 306, 42: 793, 51: 306, 55: 306, 58: 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306},
		{8: 244, 244, 468},
		{109, 109, 109, 109, 109, 109, 109, 109, 109, 11: 109, 13: 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 43: 109, 109, 109, 109, 109, 109, 109, 52: 109, 56: 109, 109, 85: 109, 87: 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 112: 109},
		{108, 108, 108, 108, 108, 108, 108, 108, 108, 11: 108, 13: 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 43: 108, 108, 108, 108, 108, 108, 108, 52: 108, 56: 108, 108, 85: 108, 87: 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 112: 108},
		// 405
		{107, 107, 107, 107, 107, 107, 107, 107, 107, 11: 107, 13: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 43: 107, 107, 107, 107, 107, 107, 107, 52: 107, 56: 107, 107, 85: 107, 87: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 112: 107},
		{106, 106, 106, 106, 106, 106, 106, 106, 106, 11: 106, 13: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 43: 106, 106, 106, 106, 106, 106, 106, 52: 106, 56: 106, 106, 85: 106, 87: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 112: 106},
		{105, 105, 105, 105, 105, 105, 105, 105, 105, 11: 105, 13: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 43: 105, 105, 105, 105, 105, 105, 105, 52: 105, 56: 105, 105, 85: 105, 87: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 112: 105},
		{104, 104, 104, 104, 104, 104, 104, 104, 104, 11: 104, 13: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 43: 104, 104, 104, 104, 104, 104, 104, 52: 104, 56: 104, 104, 85: 104, 87: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 112: 104},
		{103, 103, 103, 103, 103, 103, 103, 103, 103, 11: 103, 13: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 43: 103, 103, 103, 103, 103, 103, 103, 52: 103, 56: 103, 103, 85: 103, 87: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 112: 103},
		// 410
		{243, 243, 243, 243, 243, 243, 243, 243, 11: 243, 13: 243, 243, 243, 243, 243, 243, 243, 243, 243, 145: 368, 790},
		{42: 788},
		{41: 787},
		{97, 97, 97, 97, 97, 97, 97, 97, 97, 11: 97, 13: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 43: 97, 97, 97, 97, 97, 97, 97, 52: 97, 56: 97, 97, 85: 97, 87: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97},
		{381, 386, 719, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 94, 43: 333, 334, 332, 358, 359, 335, 331, 52: 716, 465, 56: 675, 366, 85: 741, 87: 738, 743, 728, 742, 729, 739, 740, 736, 744, 737, 113: 329, 720, 117: 348, 351, 347, 328, 735, 676, 665, 330, 126: 327, 727, 722, 723, 725, 726, 721, 724, 734, 137: 326, 144: 733, 177: 786},
		// 415
		{93, 93, 93, 93, 93, 93, 93, 93, 93, 11: 93, 13: 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 43: 93, 93, 93, 93, 93, 93, 93, 52: 93, 56: 93, 93, 85: 93, 87: 93, 93, 93, 93, 93, 93, 93, 93, 93, 93},
		{92, 92, 92, 92, 92, 92, 92, 92, 92, 11: 92, 13: 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 43: 92, 92, 92, 92, 92, 92, 92, 52: 92, 56: 92, 92, 85: 92, 87: 92, 92, 92, 92, 92, 92, 92, 92, 92, 92},
		{8: 785},
		{779},
		{775},
		// 420
		{771},
		{381, 386, 719, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 52: 716, 465, 56: 675, 85: 741, 87: 738, 743, 728, 742, 729, 739, 740, 736, 744, 737, 114: 720, 121: 735, 676, 127: 727, 722, 723, 725, 726, 721, 724, 765},
		{751},
		{2: 749},
		{8: 748},
		// 425
		{8: 747},
		{381, 386, 374, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 465, 114: 720, 121: 745},
		{8: 746},
		{80, 80, 80, 80, 80, 80, 80, 80, 80, 11: 80, 13: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 43: 80, 80, 80, 80, 80, 80, 80, 52: 80, 56: 80, 80, 85: 80, 87: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 112: 80},
		{81, 81, 81, 81, 81, 81, 81, 81, 81, 11: 81, 13: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 43: 81, 81, 81, 81, 81, 81, 81, 52: 81, 56: 81, 81, 85: 81, 87: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 112: 81},
		// 430
		{82, 82, 82, 82, 82, 82, 82, 82, 82, 11: 82, 13: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 43: 82, 82, 82, 82, 82, 82, 82, 52: 82, 56: 82, 82, 85: 82, 87: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 112: 82},
		{8: 750},
		{83, 83, 83, 83, 83, 83, 83, 83, 83, 11: 83, 13: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 43: 83, 83, 83, 83, 83, 83, 83, 52: 83, 56: 83, 83, 85: 83, 87: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 112: 83},
		{381, 386, 374, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 357, 355, 356, 345, 337, 346, 342, 354, 341, 339, 340, 338, 343, 352, 349, 350, 353, 344, 336, 43: 333, 334, 332, 358, 359, 335, 331, 53: 465, 57: 366, 113: 329, 720, 117: 348, 351, 347, 328, 752, 123: 665, 330, 126: 327, 137: 326, 144: 753},
		{8: 759},
		// 435
		{381, 386, 374, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 465, 114: 720, 121: 754},
		{8: 755},
		{381, 386, 374, 385, 387, 388, 384, 383, 9: 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 465, 114: 720, 121: 756},
		{9: 757},
		{381, 386, 719, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 52: 716, 465, 56: 675, 85: 741, 87: 738, 743, 728, 742, 729, 739, 740, 736, 744, 737, 114: 720, 121: 735, 676, 127: 727, 722, 723, 725, 726, 721, 724, 758},
		// 440
		{84, 84, 84, 84, 84, 84, 84, 84, 84, 11: 84, 13: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 43: 84, 84, 84, 84, 84, 84, 84, 52: 84, 56: 84, 84, 85: 84, 87: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 112: 84},
		{381, 386, 374, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 465, 114: 720, 121: 760},
		{8: 761},
		{381, 386, 374, 385, 387, 388, 384, 383, 9: 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 465, 114: 720, 121: 762},
		{9: 763},
		// 445
		{381, 386, 719, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 52: 716, 465, 56: 675, 85: 741, 87: 738, 743, 728, 742, 729, 739, 740, 736, 744, 737, 114: 720, 121: 735, 676, 127: 727, 722, 723, 725, 726, 721, 724, 764},
		{85, 85, 85, 85, 85, 85, 85, 85, 85, 11: 85, 13: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 43: 85, 85, 85, 85, 85, 85, 85, 52: 85, 56: 85, 85, 85: 85, 87: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 112: 85},
		{87: 766},
		{767},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 465, 114: 768},
		// 450
		{9: 769, 468},
		{8: 770},
		{86, 86, 86, 86, 86, 86, 86, 86, 86, 11: 86, 13: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 43: 86, 86, 86, 86, 86, 86, 86, 52: 86, 56: 86, 86, 85: 86, 87: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 112: 86},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 465, 114: 772},
		{9: 773, 468},
		// 455
		{381, 386, 719, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 52: 716, 465, 56: 675, 85: 741, 87: 738, 743, 728, 742, 729, 739, 740, 736, 744, 737, 114: 720, 121: 735, 676, 127: 727, 722, 723, 725, 726, 721, 724, 774},
		{87, 87, 87, 87, 87, 87, 87, 87, 87, 11: 87, 13: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 43: 87, 87, 87, 87, 87, 87, 87, 52: 87, 56: 87, 87, 85: 87, 87: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 112: 87},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 465, 114: 776},
		{9: 777, 468},
		{381, 386, 719, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 52: 716, 465, 56: 675, 85: 741, 87: 738, 743, 728, 742, 729, 739, 740, 736, 744, 737, 114: 720, 121: 735, 676, 127: 727, 722, 723, 725, 726, 721, 724, 778},
		// 460
		{88, 88, 88, 88, 88, 88, 88, 88, 88, 11: 88, 13: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 43: 88, 88, 88, 88, 88, 88, 88, 52: 88, 56: 88, 88, 85: 88, 87: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 112: 88},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 53: 465, 114: 780},
		{9: 781, 468},
		{381, 386, 719, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 52: 716, 465, 56: 675, 85: 741, 87: 738, 743, 728, 742, 729, 739, 740, 736, 744, 737, 114: 720, 121: 735, 676, 127: 727, 722, 723, 725, 726, 721, 724, 782},
		{90, 90, 90, 90, 90, 90, 90, 90, 90, 11: 90, 13: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 43: 90, 90, 90, 90, 90, 90, 90, 52: 90, 56: 90, 90, 85: 90, 87: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 112: 783},
		// 465
		{381, 386, 719, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 52: 716, 465, 56: 675, 85: 741, 87: 738, 743, 728, 742, 729, 739, 740, 736, 744, 737, 114: 720, 121: 735, 676, 127: 727, 722, 723, 725, 726, 721, 724, 784},
		{89, 89, 89, 89, 89, 89, 89, 89, 89, 11: 89, 13: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 43: 89, 89, 89, 89, 89, 89, 89, 52: 89, 56: 89, 89, 85: 89, 87: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 112: 89},
		{91, 91, 91, 91, 91, 91, 91, 91, 91, 11: 91, 13: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 43: 91, 91, 91, 91, 91, 91, 91, 52: 91, 56: 91, 91, 85: 91, 87: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 112: 91},
		{96, 96, 96, 96, 96, 96, 96, 96, 96, 11: 96, 13: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 43: 96, 96, 96, 96, 96, 96, 96, 52: 96, 56: 96, 96, 85: 96, 87: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96},
		{98, 98, 98, 98, 98, 98, 98, 98, 98, 11: 98, 13: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 43: 98, 98, 98, 98, 98, 98, 98, 98, 52: 98, 56: 98, 98, 85: 98, 87: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 112: 98},
		// 470
		{381, 386, 719, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 52: 716, 465, 56: 675, 85: 741, 87: 738, 743, 728, 742, 729, 739, 740, 736, 744, 737, 114: 720, 121: 735, 676, 127: 727, 722, 723, 725, 726, 721, 724, 789},
		{100, 100, 100, 100, 100, 100, 100, 100, 100, 11: 100, 13: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 43: 100, 100, 100, 100, 100, 100, 100, 52: 100, 56: 100, 100, 85: 100, 87: 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 112: 100},
		{42: 791},
		{381, 386, 719, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 52: 716, 465, 56: 675, 85: 741, 87: 738, 743, 728, 742, 729, 739, 740, 736, 744, 737, 114: 720, 121: 735, 676, 127: 727, 722, 723, 725, 726, 721, 724, 792},
		{101, 101, 101, 101, 101, 101, 101, 101, 101, 11: 101, 13: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 43: 101, 101, 101, 101, 101, 101, 101, 52: 101, 56: 101, 101, 85: 101, 87: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 112: 101},
		// 475
		{381, 386, 719, 385, 387, 388, 384, 383, 245, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 52: 716, 465, 56: 675, 85: 741, 87: 738, 743, 728, 742, 729, 739, 740, 736, 744, 737, 114: 720, 121: 735, 676, 127: 727, 722, 723, 725, 726, 721, 724, 794},
		{102, 102, 102, 102, 102, 102, 102, 102, 102, 11: 102, 13: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 43: 102, 102, 102, 102, 102, 102, 102, 52: 102, 56: 102, 102, 85: 102, 87: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 112: 102},
		{381, 386, 374, 385, 387, 388, 384, 383, 11: 380, 13: 390, 389, 392, 375, 376, 377, 378, 379, 391, 52: 582, 581, 171: 796},
		{8: 227, 10: 227},
		{158, 399, 158, 135: 639, 638, 139: 669, 170: 798},
		// 480
		{8: 232, 10: 232},
		{241, 241, 241, 241, 241, 241, 241, 241, 241, 11: 241, 13: 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 43: 241, 241, 241, 241, 241, 241, 241, 241, 52: 241, 56: 241, 241, 85: 241, 87: 241, 241, 241, 241, 241, 241, 241, 241, 241, 241},
		{8: 78, 22: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 43: 78, 78, 78, 78, 78, 78, 78, 78, 56: 78, 78},
		{243, 243, 243, 243, 243, 243, 243, 243, 11: 243, 13: 243, 243, 243, 243, 243, 243, 243, 243, 243, 145: 368, 802},
		{50: 314},
		// 485
		{84: 825, 86: 827, 100: 815, 816, 817, 812, 813, 814, 818, 822, 819, 809, 820, 821, 115: 826, 824, 125: 823, 138: 807, 140: 806, 811, 808, 810, 148: 805, 227: 804},
		{50: 316},
		{50: 43, 84: 825, 86: 827, 100: 815, 816, 817, 812, 813, 814, 818, 822, 819, 809, 820, 821, 115: 826, 824, 125: 823, 138: 807, 140: 865, 811, 808, 810},
		{50: 42, 84: 42, 86: 42, 97: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{50: 38, 84: 38, 86: 38, 97: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		// 490
		{50: 37, 84: 37, 86: 37, 97: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		{86: 827, 115: 887, 824},
		{50: 35, 84: 35, 86: 35, 97: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		{97: 28, 28, 875, 183: 873, 218: 874, 872},
		{86: 827, 115: 869, 824},
		// 495
		{2: 866},
		{2: 861},
		{2: 842, 84: 844, 225: 843},
		{84: 825, 86: 827, 115: 826, 824, 125: 841},
		{50: 16, 84: 16, 86: 16, 97: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		// 500
		{86: 827, 115: 839, 824},
		{86: 827, 115: 837, 824},
		{84: 825, 86: 827, 115: 826, 824, 125: 836},
		{2: 832},
		{86: 827, 115: 830, 824},
		// 505
		{50: 7, 84: 7, 86: 7, 97: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{84: 5, 86: 829},
		{50: 4, 84: 4, 86: 4, 97: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{84: 828},
		{84: 2, 86: 2},
		// 510
		{50: 3, 84: 3, 86: 3, 97: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{84: 1, 86: 1},
		{84: 831},
		{50: 8, 84: 8, 86: 8, 97: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{84: 833, 86: 827, 115: 834, 824},
		// 515
		{50: 12, 84: 12, 86: 12, 97: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{84: 835},
		{50: 9, 84: 9, 86: 9, 97: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{50: 13, 84: 13, 86: 13, 97: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{84: 838},
		// 520
		{50: 14, 84: 14, 86: 14, 97: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{84: 840},
		{50: 15, 84: 15, 86: 15, 97: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{50: 17, 84: 17, 86: 17, 97: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{84: 825, 86: 827, 115: 826, 824, 125: 850, 150: 860},
		// 525
		{2: 528, 9: 142, 147: 846, 169: 845, 188: 847},
		{50: 10, 84: 10, 86: 10, 97: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{9: 141, 853, 147: 854},
		{9: 851},
		{9: 848},
		// 530
		{84: 825, 86: 827, 115: 826, 824, 125: 850, 150: 849},
		{50: 18, 84: 18, 86: 18, 97: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{50: 6, 84: 6, 86: 6, 97: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{84: 825, 86: 827, 115: 826, 824, 125: 850, 150: 852},
		{50: 20, 84: 20, 86: 20, 97: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		// 535
		{2: 531, 147: 857},
		{9: 855},
		{84: 825, 86: 827, 115: 826, 824, 125: 850, 150: 856},
		{50: 11, 84: 11, 86: 11, 97: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{9: 858},
		// 540
		{84: 825, 86: 827, 115: 826, 824, 125: 850, 150: 859},
		{50: 19, 84: 19, 86: 19, 97: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{50: 21, 84: 21, 86: 21, 97: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{84: 862},
		{84: 825, 86: 827, 97: 40, 40, 40, 815, 816, 817, 812, 813, 814, 818, 822, 819, 809, 820, 821, 115: 826, 824, 125: 823, 138: 807, 140: 806, 811, 808, 810, 148: 863, 864},
		// 545
		{84: 825, 86: 827, 97: 39, 39, 39, 815, 816, 817, 812, 813, 814, 818, 822, 819, 809, 820, 821, 115: 826, 824, 125: 823, 138: 807, 140: 865, 811, 808, 810},
		{97: 31, 31, 31},
		{50: 41, 84: 41, 86: 41, 97: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41},
		{84: 867},
		{84: 825, 86: 827, 97: 40, 40, 40, 815, 816, 817, 812, 813, 814, 818, 822, 819, 809, 820, 821, 115: 826, 824, 125: 823, 138: 807, 140: 806, 811, 808, 810, 148: 863, 868},
		// 550
		{97: 32, 32, 32},
		{84: 870},
		{84: 825, 86: 827, 97: 40, 40, 40, 815, 816, 817, 812, 813, 814, 818, 822, 819, 809, 820, 821, 115: 826, 824, 125: 823, 138: 807, 140: 806, 811, 808, 810, 148: 863, 871},
		{97: 33, 33, 33},
		{97: 24, 881, 220: 882, 880},
		// 555
		{97: 30, 30, 30},
		{97: 27, 27, 875, 183: 879},
		{86: 827, 115: 876, 824},
		{84: 877},
		{84: 825, 86: 827, 97: 40, 40, 40, 815, 816, 817, 812, 813, 814, 818, 822, 819, 809, 820, 821, 115: 826, 824, 125: 823, 138: 807, 140: 806, 811, 808, 810, 148: 863, 878},
		// 560
		{97: 26, 26, 26},
		{97: 29, 29, 29},
		{97: 886, 222: 885},
		{84: 883},
		{97: 23},
		// 565
		{84: 825, 86: 827, 97: 40, 100: 815, 816, 817, 812, 813, 814, 818, 822, 819, 809, 820, 821, 115: 826, 824, 125: 823, 138: 807, 140: 806, 811, 808, 810, 148: 863, 884},
		{97: 25},
		{50: 34, 84: 34, 86: 34, 97: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		{50: 22, 84: 22, 86: 22, 97: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{84: 888},
		// 570
		{50: 36, 84: 36, 86: 36, 97: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("'%c'", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), PrettyString(lval.Token): %v\n", yySymName(n), n, n, PrettyString(lval.Token))
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 235

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			lx := yylex.(*lexer)
			lx.preprocessingFile = nil
		}
	case 2:
		{
			lx := yylex.(*lexer)
			lx.preprocessingFile = yyS[yypt-0].node.(*PreprocessingFile)
		}
	case 3:
		{
			lx := yylex.(*lexer)
			lx.constantExpression = nil
		}
	case 4:
		{
			lx := yylex.(*lexer)
			lx.constantExpression = yyS[yypt-0].node.(*ConstantExpression)
		}
	case 5:
		{
			lx := yylex.(*lexer)
			lx.translationUnit = nil
		}
	case 6:
		{
			lx := yylex.(*lexer)
			if lx.report.Errors(false) == nil && lx.scope.kind != ScopeFile {
				panic("internal error")
			}

			lx.translationUnit = yyS[yypt-0].node.(*TranslationUnit).reverse()
			lx.translationUnit.Declarations = lx.scope
		}
	case 7:
		{
			yyVAL.node = &EnumerationConstant{
				Token: yyS[yypt-0].Token,
			}
		}
	case 8:
		{
			yyVAL.node = &ArgumentExpressionList{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 9:
		{
			yyVAL.node = &ArgumentExpressionList{
				Case: 1,
				ArgumentExpressionList: yyS[yypt-2].node.(*ArgumentExpressionList),
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 10:
		{
			yyVAL.node = (*ArgumentExpressionListOpt)(nil)
		}
	case 11:
		{
			yyVAL.node = &ArgumentExpressionListOpt{
				ArgumentExpressionList: yyS[yypt-0].node.(*ArgumentExpressionList).reverse(),
			}
		}
	case 12:
		{
			lx := yylex.(*lexer)
			lhs := &Expression{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.scope = lx.scope
		}
	case 13:
		{
			yyVAL.node = &Expression{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 14:
		{
			yyVAL.node = &Expression{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
		}
	case 15:
		{
			yyVAL.node = &Expression{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
		}
	case 16:
		{
			yyVAL.node = &Expression{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
		}
	case 17:
		{
			yyVAL.node = &Expression{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 18:
		{
			yyVAL.node = &Expression{
				Case:  6,
				Token: yyS[yypt-0].Token,
			}
		}
	case 19:
		{
			yyVAL.node = &Expression{
				Case:           7,
				Token:          yyS[yypt-2].Token,
				ExpressionList: yyS[yypt-1].node.(*ExpressionList).reverse(),
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 20:
		{
			yyVAL.node = &Expression{
				Case:           8,
				Expression:     yyS[yypt-3].node.(*Expression),
				Token:          yyS[yypt-2].Token,
				ExpressionList: yyS[yypt-1].node.(*ExpressionList).reverse(),
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 21:
		{
			lx := yylex.(*lexer)
			lhs := &Expression{
				Case:       9,
				Expression: yyS[yypt-3].node.(*Expression),
				Token:      yyS[yypt-2].Token,
				ArgumentExpressionListOpt: yyS[yypt-1].node.(*ArgumentExpressionListOpt),
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			o := lhs.ArgumentExpressionListOpt
			if o == nil {
				break
			}

			lhs.Expression.eval(lx)
			for l := o.ArgumentExpressionList; l != nil; l = l.ArgumentExpressionList {
				l.Expression.eval(lx)
			}
		}
	case 22:
		{
			yyVAL.node = &Expression{
				Case:       10,
				Expression: yyS[yypt-2].node.(*Expression),
				Token:      yyS[yypt-1].Token,
				Token2:     yyS[yypt-0].Token,
			}
		}
	case 23:
		{
			yyVAL.node = &Expression{
				Case:       11,
				Expression: yyS[yypt-2].node.(*Expression),
				Token:      yyS[yypt-1].Token,
				Token2:     yyS[yypt-0].Token,
			}
		}
	case 24:
		{
			yyVAL.node = &Expression{
				Case:       12,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 25:
		{
			yyVAL.node = &Expression{
				Case:       13,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 26:
		{
			yyVAL.node = &Expression{
				Case:            14,
				Token:           yyS[yypt-6].Token,
				TypeName:        yyS[yypt-5].node.(*TypeName),
				Token2:          yyS[yypt-4].Token,
				Token3:          yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token4:          yyS[yypt-0].Token,
			}
		}
	case 27:
		{
			yyVAL.node = &Expression{
				Case:       15,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 28:
		{
			yyVAL.node = &Expression{
				Case:       16,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 29:
		{
			yyVAL.node = &Expression{
				Case:       17,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 30:
		{
			yyVAL.node = &Expression{
				Case:       18,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 31:
		{
			yyVAL.node = &Expression{
				Case:       19,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 32:
		{
			yyVAL.node = &Expression{
				Case:       20,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 33:
		{
			yyVAL.node = &Expression{
				Case:       21,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 34:
		{
			yyVAL.node = &Expression{
				Case:       22,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 35:
		{
			yyVAL.node = &Expression{
				Case:       23,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 36:
		{
			yyVAL.node = &Expression{
				Case:     24,
				Token:    yyS[yypt-3].Token,
				Token2:   yyS[yypt-2].Token,
				TypeName: yyS[yypt-1].node.(*TypeName),
				Token3:   yyS[yypt-0].Token,
			}
		}
	case 37:
		{
			yyVAL.node = &Expression{
				Case:       25,
				Token:      yyS[yypt-3].Token,
				TypeName:   yyS[yypt-2].node.(*TypeName),
				Token2:     yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 38:
		{
			yyVAL.node = &Expression{
				Case:        26,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 39:
		{
			yyVAL.node = &Expression{
				Case:        27,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 40:
		{
			yyVAL.node = &Expression{
				Case:        28,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 41:
		{
			yyVAL.node = &Expression{
				Case:        29,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 42:
		{
			yyVAL.node = &Expression{
				Case:        30,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 43:
		{
			yyVAL.node = &Expression{
				Case:        31,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 44:
		{
			yyVAL.node = &Expression{
				Case:        32,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 45:
		{
			yyVAL.node = &Expression{
				Case:        33,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 46:
		{
			yyVAL.node = &Expression{
				Case:        34,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 47:
		{
			yyVAL.node = &Expression{
				Case:        35,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 48:
		{
			yyVAL.node = &Expression{
				Case:        36,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 49:
		{
			yyVAL.node = &Expression{
				Case:        37,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 50:
		{
			yyVAL.node = &Expression{
				Case:        38,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 51:
		{
			yyVAL.node = &Expression{
				Case:        39,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 52:
		{
			yyVAL.node = &Expression{
				Case:        40,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 53:
		{
			yyVAL.node = &Expression{
				Case:        41,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 54:
		{
			yyVAL.node = &Expression{
				Case:        42,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 55:
		{
			yyVAL.node = &Expression{
				Case:        43,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 56:
		{
			yyVAL.node = &Expression{
				Case:           44,
				Expression:     yyS[yypt-4].node.(*Expression),
				Token:          yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token2:         yyS[yypt-1].Token,
				Expression2:    yyS[yypt-0].node.(*Expression),
			}
		}
	case 57:
		{
			yyVAL.node = &Expression{
				Case:        45,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 58:
		{
			yyVAL.node = &Expression{
				Case:        46,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 59:
		{
			yyVAL.node = &Expression{
				Case:        47,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 60:
		{
			yyVAL.node = &Expression{
				Case:        48,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 61:
		{
			yyVAL.node = &Expression{
				Case:        49,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 62:
		{
			yyVAL.node = &Expression{
				Case:        50,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 63:
		{
			yyVAL.node = &Expression{
				Case:        51,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 64:
		{
			yyVAL.node = &Expression{
				Case:        52,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 65:
		{
			yyVAL.node = &Expression{
				Case:        53,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 66:
		{
			yyVAL.node = &Expression{
				Case:        54,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 67:
		{
			yyVAL.node = &Expression{
				Case:        55,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 68:
		{
			yyVAL.node = &Expression{
				Case:     56,
				Token:    yyS[yypt-3].Token,
				Token2:   yyS[yypt-2].Token,
				TypeName: yyS[yypt-1].node.(*TypeName),
				Token3:   yyS[yypt-0].Token,
			}
		}
	case 69:
		{
			yyVAL.node = (*ExpressionOpt)(nil)
		}
	case 70:
		{
			lx := yylex.(*lexer)
			lhs := &ExpressionOpt{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 71:
		{
			yyVAL.node = &ExpressionList{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 72:
		{
			yyVAL.node = &ExpressionList{
				Case:           1,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList),
				Token:          yyS[yypt-1].Token,
				Expression:     yyS[yypt-0].node.(*Expression),
			}
		}
	case 73:
		{
			yyVAL.node = (*ExpressionListOpt)(nil)
		}
	case 74:
		{
			lx := yylex.(*lexer)
			lhs := &ExpressionListOpt{
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 75:
		{
			lx := yylex.(*lexer)
			lx.constExprToks = []xc.Token{lx.last}
		}
	case 76:
		{
			lx := yylex.(*lexer)
			lhs := &ConstantExpression{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Value, lhs.Type = lhs.Expression.eval(lx)
			if lhs.Value == nil {
				lx.report.Err(lhs.Pos(), "not a constant expression")
			}
			l := lx.constExprToks
			lhs.toks = l[:len(l)-1]
			lx.constExprToks = nil
		}
	case 77:
		{
			lx := yylex.(*lexer)
			lhs := &Declaration{
				DeclarationSpecifiers: yyS[yypt-2].node.(*DeclarationSpecifiers),
				InitDeclaratorListOpt: yyS[yypt-1].node.(*InitDeclaratorListOpt),
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			ts := tsDecode(lhs.DeclarationSpecifiers.typeSpecifiers())
			ok := false
			for _, v := range ts {
				if v == tsStructSpecifier || v == tsUnionSpecifier {
					ok = true
					break
				}
			}
			if ok {
				s := lhs.DeclarationSpecifiers
				d := &Declarator{specifier: s}
				dd := &DirectDeclarator{
					Token:      xc.Token{Char: lex.NewChar(lhs.Pos(), 0)},
					declarator: d,
					idScope:    lx.scope,
					specifier:  s,
				}
				d.DirectDeclarator = dd
				d.setFull(lx)
				for l := lhs.DeclarationSpecifiers; l != nil; {
					ts := l.TypeSpecifier
					if ts != nil && ts.Case == 11 && ts.StructOrUnionSpecifier.Case == 0 { // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
						ts.StructOrUnionSpecifier.declarator = d
						break
					}

					if o := l.DeclarationSpecifiersOpt; o != nil {
						l = o.DeclarationSpecifiers
						continue
					}

					break
				}
			}

			o := lhs.InitDeclaratorListOpt
			if o != nil {
				break
			}

			s := lhs.DeclarationSpecifiers
			d := &Declarator{specifier: s}
			dd := &DirectDeclarator{
				Token:      xc.Token{Char: lex.NewChar(lhs.Pos(), 0)},
				declarator: d,
				idScope:    lx.scope,
				specifier:  s,
			}
			d.DirectDeclarator = dd
			d.setFull(lx)
			lhs.declarator = d
		}
	case 78:
		{
			yyVAL.node = &Declaration{
				Case: 1,
				StaticAssertDeclaration: yyS[yypt-0].node.(*StaticAssertDeclaration),
			}
		}
	case 79:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				StorageClassSpecifier:    yyS[yypt-1].node.(*StorageClassSpecifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.StorageClassSpecifier
			b := lhs.DeclarationSpecifiersOpt
			if b == nil {
				lhs.attr = a.attr
				break
			}

			if a.attr&b.attr != 0 {
				lx.report.Err(a.Pos(), "invalid storage class specifier")
				break
			}

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
			if lhs.StorageClassSpecifier.Case != 0 /* "typedef" */ && lhs.IsTypedef() {
				lx.report.Err(a.Pos(), "invalid storage class specifier")
			}
		}
	case 80:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				Case:                     1,
				TypeSpecifier:            yyS[yypt-1].node.(*TypeSpecifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.TypeSpecifier
			b := lhs.DeclarationSpecifiersOpt
			if b == nil {
				lhs.typeSpecifier = a.typeSpecifier
				break
			}

			lhs.attr = b.attr
			ts := tsEncode(append(tsDecode(a.typeSpecifier), tsDecode(b.typeSpecifier)...)...)
			if _, ok := tsValid[ts]; !ok {
				ts = tsEncode(tsInt)
				lx.report.Err(a.Pos(), "invalid type specifier")
			}
			lhs.typeSpecifier = ts
		}
	case 81:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				Case:                     2,
				TypeQualifier:            yyS[yypt-1].node.(*TypeQualifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.TypeQualifier
			b := lhs.DeclarationSpecifiersOpt
			if b == nil {
				lhs.attr = a.attr
				break
			}

			if a.attr&b.attr != 0 {
				lx.report.Err(a.Pos(), "invalid type qualifier")
				break
			}

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
			if lhs.IsTypedef() {
				lx.report.Err(a.Pos(), "invalid type qualifier")
			}
		}
	case 82:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				Case:                     3,
				FunctionSpecifier:        yyS[yypt-1].node.(*FunctionSpecifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.FunctionSpecifier
			b := lhs.DeclarationSpecifiersOpt
			if b == nil {
				lhs.attr = a.attr
				break
			}

			if a.attr&b.attr != 0 {
				lx.report.Err(a.Pos(), "invalid function specifier")
				break
			}

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
			if lhs.IsTypedef() {
				lx.report.Err(a.Pos(), "invalid function specifier")
			}
		}
	case 83:
		{
			yyVAL.node = (*DeclarationSpecifiersOpt)(nil)
		}
	case 84:
		{
			lhs := &DeclarationSpecifiersOpt{
				DeclarationSpecifiers: yyS[yypt-0].node.(*DeclarationSpecifiers),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.DeclarationSpecifiers.attr
			lhs.typeSpecifier = lhs.DeclarationSpecifiers.typeSpecifier
		}
	case 85:
		{
			yyVAL.node = &InitDeclaratorList{
				InitDeclarator: yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 86:
		{
			yyVAL.node = &InitDeclaratorList{
				Case:               1,
				InitDeclaratorList: yyS[yypt-2].node.(*InitDeclaratorList),
				Token:              yyS[yypt-1].Token,
				InitDeclarator:     yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 87:
		{
			yyVAL.node = (*InitDeclaratorListOpt)(nil)
		}
	case 88:
		{
			yyVAL.node = &InitDeclaratorListOpt{
				InitDeclaratorList: yyS[yypt-0].node.(*InitDeclaratorList).reverse(),
			}
		}
	case 89:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
		}
	case 90:
		{
			lx := yylex.(*lexer)
			d := yyS[yypt-0].node.(*Declarator)
			d.setFull(lx)
		}
	case 91:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Case:        1,
				Declarator:  yyS[yypt-3].node.(*Declarator),
				Token:       yyS[yypt-1].Token,
				Initializer: yyS[yypt-0].node.(*Initializer),
			}
			yyVAL.node = lhs
			switch i := lhs.Initializer; i.Case {
			case 0: // Expression
				s := i.Expression.Type
				d := lhs.Declarator.Type
				if !s.CanAssignTo(s) {
					lx.report.Err(i.Pos(), "incompatible types when initializing type '%s' using type ‘%s'", d, s)
				}
			case 1: // '{' InitializerList CommaOpt '}'
				limit := -1
				var checkType Type
				var mb []Member
				var incomplete bool
				d := lhs.Declarator
				k := d.Type.Kind()
				switch k {
				case Array:
					checkType = d.Type.Element()
					limit = d.Type.Elements()
				case Ptr:
					checkType = d.Type.Element()
					d.Type = d.Type.(*ctype).setElements(i.InitializerList.Len())
				case Struct, Union:
					mb, incomplete = d.Type.Members()
					if mb == nil {
						panic("internal error")
					}

					limit = len(mb)
					if k == Union {
						limit = 1
					}
				default:
					//dbg("", position(d.Pos()), d.Type.Kind())
					panic("TODO")
				}

				values := 0
				for l := i.InitializerList; l != nil; l = l.InitializerList {
					values++
					if l.DesignationOpt != nil {
						panic("TODO")
					}

					if incomplete {
						lx.report.Err(i.Pos(), "variable/field has initializer but incomplete type")
						break
					}

					l.Initializer.typeCheck(checkType, mb, values-1, limit, lx)
				}
			default:
				panic(i.Case)
			}
		}
	case 92:
		{
			lhs := &StorageClassSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saTypedef
		}
	case 93:
		{
			lhs := &StorageClassSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saExtern
		}
	case 94:
		{
			lhs := &StorageClassSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saStatic
		}
	case 95:
		{
			lhs := &StorageClassSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saAuto
		}
	case 96:
		{
			lhs := &StorageClassSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRegister
		}
	case 97:
		{
			lhs := &TypeSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsVoid)
		}
	case 98:
		{
			lhs := &TypeSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsChar)
		}
	case 99:
		{
			lhs := &TypeSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsShort)
		}
	case 100:
		{
			lhs := &TypeSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsInt)
		}
	case 101:
		{
			lhs := &TypeSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsLong)
		}
	case 102:
		{
			lhs := &TypeSpecifier{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsFloat)
		}
	case 103:
		{
			lhs := &TypeSpecifier{
				Case:  6,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsDouble)
		}
	case 104:
		{
			lhs := &TypeSpecifier{
				Case:  7,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsSigned)
		}
	case 105:
		{
			lhs := &TypeSpecifier{
				Case:  8,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsUnsigned)
		}
	case 106:
		{
			lhs := &TypeSpecifier{
				Case:  9,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsBool)
		}
	case 107:
		{
			lhs := &TypeSpecifier{
				Case:  10,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsComplex)
		}
	case 108:
		{
			lhs := &TypeSpecifier{
				Case: 11,
				StructOrUnionSpecifier: yyS[yypt-0].node.(*StructOrUnionSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(lhs.StructOrUnionSpecifier.typeSpecifiers())
		}
	case 109:
		{
			lhs := &TypeSpecifier{
				Case:          12,
				EnumSpecifier: yyS[yypt-0].node.(*EnumSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsEnumSpecifier)
		}
	case 110:
		{
			lx := yylex.(*lexer)
			lhs := &TypeSpecifier{
				Case:  13,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsTypedefName)
			_, lhs.scope = lx.scope.Lookup2(NSIdentifiers, lhs.Token.Val)
		}
	case 111:
		{
			lx := yylex.(*lexer)
			lhs := &TypeSpecifier{
				Case:       14,
				Token:      yyS[yypt-3].Token,
				Token2:     yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token3:     yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsTypeof)
			_, lhs.Type = lhs.Expression.eval(lx)
		}
	case 112:
		{
			lhs := &TypeSpecifier{
				Case:     15,
				Token:    yyS[yypt-3].Token,
				Token2:   yyS[yypt-2].Token,
				TypeName: yyS[yypt-1].node.(*TypeName),
				Token3:   yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsTypeof)
			lhs.Type = undefined
			if t := lhs.TypeName.Type; t != nil {
				lhs.Type = t
			}
		}
	case 113:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareStructTag(o.Token, lx.report)
			}
		}
	case 114:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeMembers)
			lx.scope.isUnion = yyS[yypt-3].node.(*StructOrUnion).Case == 1 // "union"
			lx.scope.prevStructDeclarator = nil
		}
	case 115:
		{
			lx := yylex.(*lexer)
			lhs := &StructOrUnionSpecifier{
				StructOrUnion: yyS[yypt-6].node.(*StructOrUnion),
				IdentifierOpt: yyS[yypt-5].node.(*IdentifierOpt),
				Token:         yyS[yypt-3].Token,
				StructDeclarationList: yyS[yypt-1].node.(*StructDeclarationList).reverse(),
				Token2:                yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			sc := lx.scope
			lhs.scope = sc
			if sc.bitOffset != 0 {
				finishBitField(lx)
			}

			i := 0
			var bt Type
			var d *Declarator
			for l := lhs.StructDeclarationList; l != nil; l = l.StructDeclarationList {
				for l := l.StructDeclaration.StructDeclaratorList; l != nil; l = l.StructDeclaratorList {
					switch sd := l.StructDeclarator; sd.Case {
					case 0: // Declarator
						d = sd.Declarator
					case 1: // DeclaratorOpt ':' ConstantExpression
						if o := sd.DeclaratorOpt; o != nil {
							x := o.Declarator
							if x.bitOffset == 0 {
								d = x
								bt = lx.scope.bitFieldTypes[i]
								i++
							}
							x.bitFieldType = bt
						}
					}
				}
			}
			lx.scope.bitFieldTypes = nil

			lhs.alignOf = sc.maxAlign
			switch {
			case sc.isUnion:
				lhs.sizeOf = align(sc.maxSize, sc.maxAlign)
			default:
				off := sc.offset
				lhs.sizeOf = align(sc.offset, sc.maxAlign)
				if d != nil {
					d.padding = lhs.sizeOf - off
				}
			}

			lx.popScope(lhs.Token2)
			if o := lhs.IdentifierOpt; o != nil {
				lx.scope.defineStructTag(o.Token, lhs, lx.report)
			}
		}
	case 116:
		{
			lx := yylex.(*lexer)
			lhs := &StructOrUnionSpecifier{
				Case:          1,
				StructOrUnion: yyS[yypt-1].node.(*StructOrUnion),
				Token:         yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lx.scope.declareStructTag(lhs.Token, lx.report)
			lhs.scope = lx.scope
		}
	case 117:
		{
			yyVAL.node = &StructOrUnion{
				Token: yyS[yypt-0].Token,
			}
		}
	case 118:
		{
			yyVAL.node = &StructOrUnion{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 119:
		{
			yyVAL.node = &StructDeclarationList{
				StructDeclaration: yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 120:
		{
			yyVAL.node = &StructDeclarationList{
				Case: 1,
				StructDeclarationList: yyS[yypt-1].node.(*StructDeclarationList),
				StructDeclaration:     yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 121:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclaration{
				SpecifierQualifierList: yyS[yypt-2].node.(*SpecifierQualifierList),
				StructDeclaratorList:   yyS[yypt-1].node.(*StructDeclaratorList).reverse(),
				Token:                  yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			s := lhs.SpecifierQualifierList
			if k := s.kind(); k != Struct && k != Union {
				break
			}

			d := &Declarator{specifier: s}
			dd := &DirectDeclarator{
				Token:      xc.Token{Char: lex.NewChar(lhs.Pos(), 0)},
				declarator: d,
				idScope:    lx.scope,
				specifier:  s,
			}
			d.DirectDeclarator = dd
			d.setFull(lx)
			for l := lhs.SpecifierQualifierList; l != nil; {
				ts := l.TypeSpecifier
				if ts != nil && ts.Case == 11 && ts.StructOrUnionSpecifier.Case == 0 { // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
					ts.StructOrUnionSpecifier.declarator = d
					break
				}

				if o := l.SpecifierQualifierListOpt; o != nil {
					l = o.SpecifierQualifierList
					continue
				}

				break
			}
		}
	case 122:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclaration{
				Case: 1,
				SpecifierQualifierList: yyS[yypt-1].node.(*SpecifierQualifierList),
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			s := lhs.SpecifierQualifierList
			if !lx.tweaks.enableAnonymousStructFields {
				lx.report.Err(lhs.Token.Pos(), "unnamed fields not allowed")
			} else if k := s.kind(); k != Struct && k != Union {
				lx.report.Err(lhs.Token.Pos(), "only unnamed structs and unions are allowed")
				break
			}

			d := &Declarator{specifier: s}
			dd := &DirectDeclarator{
				Token:      xc.Token{Char: lex.NewChar(lhs.Pos(), 0)},
				declarator: d,
				idScope:    lx.scope,
				specifier:  s,
			}
			d.DirectDeclarator = dd
			d.setFull(lx)
			for l := lhs.SpecifierQualifierList; l != nil; {
				ts := l.TypeSpecifier
				if ts != nil && ts.Case == 11 && ts.StructOrUnionSpecifier.Case == 0 { // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
					ts.StructOrUnionSpecifier.declarator = d
					break
				}

				if o := l.SpecifierQualifierListOpt; o != nil {
					l = o.SpecifierQualifierList
					continue
				}

				break
			}
		}
	case 123:
		{
			yyVAL.node = &StructDeclaration{
				Case: 2,
				StaticAssertDeclaration: yyS[yypt-0].node.(*StaticAssertDeclaration),
			}
		}
	case 124:
		{
			lx := yylex.(*lexer)
			lhs := &SpecifierQualifierList{
				TypeSpecifier:             yyS[yypt-1].node.(*TypeSpecifier),
				SpecifierQualifierListOpt: yyS[yypt-0].node.(*SpecifierQualifierListOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.TypeSpecifier
			b := lhs.SpecifierQualifierListOpt
			if b == nil {
				lhs.typeSpecifier = a.typeSpecifier
				break
			}

			lhs.attr = b.attr
			ts := tsEncode(append(tsDecode(a.typeSpecifier), tsDecode(b.typeSpecifier)...)...)
			if _, ok := tsValid[ts]; !ok {
				lx.report.Err(a.Pos(), "invalid type specifier")
				break
			}

			lhs.typeSpecifier = ts
		}
	case 125:
		{
			lx := yylex.(*lexer)
			lhs := &SpecifierQualifierList{
				Case:                      1,
				TypeQualifier:             yyS[yypt-1].node.(*TypeQualifier),
				SpecifierQualifierListOpt: yyS[yypt-0].node.(*SpecifierQualifierListOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.TypeQualifier
			b := lhs.SpecifierQualifierListOpt
			if b == nil {
				lhs.attr = a.attr
				break
			}

			if a.attr&b.attr != 0 {
				lx.report.Err(a.Pos(), "invalid type qualifier")
				break
			}

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
		}
	case 126:
		{
			yyVAL.node = (*SpecifierQualifierListOpt)(nil)
		}
	case 127:
		{
			lhs := &SpecifierQualifierListOpt{
				SpecifierQualifierList: yyS[yypt-0].node.(*SpecifierQualifierList),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.SpecifierQualifierList.attr
			lhs.typeSpecifier = lhs.SpecifierQualifierList.typeSpecifier
		}
	case 128:
		{
			yyVAL.node = &StructDeclaratorList{
				StructDeclarator: yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 129:
		{
			yyVAL.node = &StructDeclaratorList{
				Case:                 1,
				StructDeclaratorList: yyS[yypt-2].node.(*StructDeclaratorList),
				Token:                yyS[yypt-1].Token,
				StructDeclarator:     yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 130:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
			lhs.post(lx)
		}
	case 131:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Case:               1,
				DeclaratorOpt:      yyS[yypt-2].node.(*DeclaratorOpt),
				Token:              yyS[yypt-1].Token,
				ConstantExpression: yyS[yypt-0].node.(*ConstantExpression),
			}
			yyVAL.node = lhs
			m := lx.model
			e := lhs.ConstantExpression
			if e.Value == nil {
				e.Value, e.Type = m.value2(1, m.IntType)
			}
			if !IsIntType(e.Type) {
				lx.report.Err(e.Pos(), "bit field width not an integer (have '%s')", e.Type)
				e.Value, e.Type = m.value2(1, m.IntType)
			}
			if o := lhs.DeclaratorOpt; o != nil {
				o.Declarator.setFull(lx)
			}
			lhs.post(lx)
		}
	case 132:
		{
			yyVAL.node = (*CommaOpt)(nil)
		}
	case 133:
		{
			yyVAL.node = &CommaOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 134:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareEnumTag(o.Token, lx.report)
			}
			lx.iota = 0
		}
	case 135:
		{
			lx := yylex.(*lexer)
			lhs := &EnumSpecifier{
				Token:          yyS[yypt-6].Token,
				IdentifierOpt:  yyS[yypt-5].node.(*IdentifierOpt),
				Token2:         yyS[yypt-3].Token,
				EnumeratorList: yyS[yypt-2].node.(*EnumeratorList).reverse(),
				CommaOpt:       yyS[yypt-1].node.(*CommaOpt),
				Token3:         yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if o := lhs.IdentifierOpt; o != nil {
				lx.scope.defineEnumTag(o.Token, lhs, lx.report)
			}
		}
	case 136:
		{
			lx := yylex.(*lexer)
			lhs := &EnumSpecifier{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lx.scope.declareEnumTag(lhs.Token2, lx.report)
		}
	case 137:
		{
			yyVAL.node = &EnumeratorList{
				Enumerator: yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 138:
		{
			yyVAL.node = &EnumeratorList{
				Case:           1,
				EnumeratorList: yyS[yypt-2].node.(*EnumeratorList),
				Token:          yyS[yypt-1].Token,
				Enumerator:     yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 139:
		{
			lx := yylex.(*lexer)
			lhs := &Enumerator{
				EnumerationConstant: yyS[yypt-0].node.(*EnumerationConstant),
			}
			yyVAL.node = lhs
			m := lx.model
			v := m.MustConvert(lx.iota, m.IntType)
			lhs.enumVal = v
			lx.scope.defineEnumConst(lx, lhs.EnumerationConstant.Token, v)
		}
	case 140:
		{
			lx := yylex.(*lexer)
			lhs := &Enumerator{
				Case:                1,
				EnumerationConstant: yyS[yypt-2].node.(*EnumerationConstant),
				Token:               yyS[yypt-1].Token,
				ConstantExpression:  yyS[yypt-0].node.(*ConstantExpression),
			}
			yyVAL.node = lhs
			m := lx.model
			e := lhs.ConstantExpression
			var v interface{}
			// [0], 6.7.2.2
			// The expression that defines the value of an enumeration
			// constant shall be an integer constant expression that has a
			// value representable as an int.
			switch {
			case !IsIntType(e.Type):
				lx.report.Err(e.Pos(), "not an integer constant expression (have '%s')", e.Type)
				v = m.MustConvert(0, m.IntType)
			default:
				var ok bool
				if v, ok = m.enumValueToInt(e.Value, lx.tweaks); !ok {
					lx.report.Err(e.Pos(), "overflow in enumeration value: %v", e.Value)
				}
			}

			lhs.enumVal = v
			lx.scope.defineEnumConst(lx, lhs.EnumerationConstant.Token, v)
		}
	case 141:
		{
			lhs := &TypeQualifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saConst
		}
	case 142:
		{
			lhs := &TypeQualifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRestrict
		}
	case 143:
		{
			lhs := &TypeQualifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saVolatile
		}
	case 144:
		{
			lhs := &FunctionSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saInline
		}
	case 145:
		{
			lhs := &FunctionSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saNoreturn
		}
	case 146:
		{
			lx := yylex.(*lexer)
			lhs := &Declarator{
				PointerOpt:       yyS[yypt-1].node.(*PointerOpt),
				DirectDeclarator: yyS[yypt-0].node.(*DirectDeclarator),
			}
			yyVAL.node = lhs
			lhs.specifier = lx.scope.specifier
			lhs.DirectDeclarator.declarator = lhs
		}
	case 147:
		{
			yyVAL.node = (*DeclaratorOpt)(nil)
		}
	case 148:
		{
			yyVAL.node = &DeclaratorOpt{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
		}
	case 149:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.specifier = lx.scope.specifier
			lx.scope.declareIdentifier(lhs.Token, lhs, lx.report)
			lhs.idScope = lx.scope
		}
	case 150:
		{
			lhs := &DirectDeclarator{
				Case:       1,
				Token:      yyS[yypt-2].Token,
				Declarator: yyS[yypt-1].node.(*Declarator),
				Token2:     yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.Declarator.specifier = nil
			lhs.Declarator.DirectDeclarator.parent = lhs
		}
	case 151:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:                 2,
				DirectDeclarator:     yyS[yypt-4].node.(*DirectDeclarator),
				Token:                yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				ExpressionOpt:        yyS[yypt-1].node.(*ExpressionOpt),
				Token2:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.elements = -1
			if o := lhs.ExpressionOpt; o != nil {
				var err error
				if lhs.elements, err = elements(o.Expression.eval(lx)); err != nil {
					lx.report.Err(o.Expression.Pos(), "%s", err)
				}

			}
			lhs.DirectDeclarator.parent = lhs
		}
	case 152:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:                 3,
				DirectDeclarator:     yyS[yypt-5].node.(*DirectDeclarator),
				Token:                yyS[yypt-4].Token,
				Token2:               yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				Expression:           yyS[yypt-1].node.(*Expression),
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			var err error
			if lhs.elements, err = elements(lhs.Expression.eval(lx)); err != nil {
				lx.report.Err(lhs.Expression.Pos(), "%s", err)
			}
			lhs.DirectDeclarator.parent = lhs
		}
	case 153:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:              4,
				DirectDeclarator:  yyS[yypt-5].node.(*DirectDeclarator),
				Token:             yyS[yypt-4].Token,
				TypeQualifierList: yyS[yypt-3].node.(*TypeQualifierList).reverse(),
				Token2:            yyS[yypt-2].Token,
				Expression:        yyS[yypt-1].node.(*Expression),
				Token3:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			var err error
			if lhs.elements, err = elements(lhs.Expression.eval(lx)); err != nil {
				lx.report.Err(lhs.Expression.Pos(), "%s", err)
			}
			lhs.DirectDeclarator.parent = lhs
		}
	case 154:
		{
			lhs := &DirectDeclarator{
				Case:                 5,
				DirectDeclarator:     yyS[yypt-4].node.(*DirectDeclarator),
				Token:                yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				Token2:               yyS[yypt-1].Token,
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.DirectDeclarator.parent = lhs
			lhs.elements = -1
		}
	case 155:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 156:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:              6,
				DirectDeclarator:  yyS[yypt-4].node.(*DirectDeclarator),
				Token:             yyS[yypt-3].Token,
				ParameterTypeList: yyS[yypt-1].node.(*ParameterTypeList),
				Token2:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScope(lhs.Token2)
			lhs.DirectDeclarator.parent = lhs
		}
	case 157:
		{
			lhs := &DirectDeclarator{
				Case:              7,
				DirectDeclarator:  yyS[yypt-3].node.(*DirectDeclarator),
				Token:             yyS[yypt-2].Token,
				IdentifierListOpt: yyS[yypt-1].node.(*IdentifierListOpt),
				Token2:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.DirectDeclarator.parent = lhs
		}
	case 158:
		{
			yyVAL.node = &Pointer{
				Token:                yyS[yypt-1].Token,
				TypeQualifierListOpt: yyS[yypt-0].node.(*TypeQualifierListOpt),
			}
		}
	case 159:
		{
			yyVAL.node = &Pointer{
				Case:                 1,
				Token:                yyS[yypt-2].Token,
				TypeQualifierListOpt: yyS[yypt-1].node.(*TypeQualifierListOpt),
				Pointer:              yyS[yypt-0].node.(*Pointer),
			}
		}
	case 160:
		{
			yyVAL.node = (*PointerOpt)(nil)
		}
	case 161:
		{
			yyVAL.node = &PointerOpt{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
		}
	case 162:
		{
			lhs := &TypeQualifierList{
				TypeQualifier: yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.TypeQualifier.attr
		}
	case 163:
		{
			lx := yylex.(*lexer)
			lhs := &TypeQualifierList{
				Case:              1,
				TypeQualifierList: yyS[yypt-1].node.(*TypeQualifierList),
				TypeQualifier:     yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			a := lhs.TypeQualifierList
			b := lhs.TypeQualifier
			if a.attr&b.attr != 0 {
				lx.report.Err(b.Pos(), "invalid type qualifier")
				break
			}

			lhs.attr = a.attr | b.attr
		}
	case 164:
		{
			yyVAL.node = (*TypeQualifierListOpt)(nil)
		}
	case 165:
		{
			yyVAL.node = &TypeQualifierListOpt{
				TypeQualifierList: yyS[yypt-0].node.(*TypeQualifierList).reverse(),
			}
		}
	case 166:
		{
			lhs := &ParameterTypeList{
				ParameterList: yyS[yypt-0].node.(*ParameterList).reverse(),
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 167:
		{
			lhs := &ParameterTypeList{
				Case:          1,
				ParameterList: yyS[yypt-2].node.(*ParameterList).reverse(),
				Token:         yyS[yypt-1].Token,
				Token2:        yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 168:
		{
			yyVAL.node = (*ParameterTypeListOpt)(nil)
		}
	case 169:
		{
			yyVAL.node = &ParameterTypeListOpt{
				ParameterTypeList: yyS[yypt-0].node.(*ParameterTypeList),
			}
		}
	case 170:
		{
			yyVAL.node = &ParameterList{
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 171:
		{
			yyVAL.node = &ParameterList{
				Case:                 1,
				ParameterList:        yyS[yypt-2].node.(*ParameterList),
				Token:                yyS[yypt-1].Token,
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 172:
		{
			lx := yylex.(*lexer)
			lhs := &ParameterDeclaration{
				DeclarationSpecifiers: yyS[yypt-1].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
			lhs.declarator = lhs.Declarator
		}
	case 173:
		{
			lx := yylex.(*lexer)
			lhs := &ParameterDeclaration{
				Case: 1,
				DeclarationSpecifiers: yyS[yypt-1].node.(*DeclarationSpecifiers),
				AbstractDeclaratorOpt: yyS[yypt-0].node.(*AbstractDeclaratorOpt),
			}
			yyVAL.node = lhs
			if o := lhs.AbstractDeclaratorOpt; o != nil {
				lhs.declarator = o.AbstractDeclarator.declarator
				lhs.declarator.setFull(lx)
				break
			}

			d := &Declarator{
				specifier: lx.scope.specifier,
				DirectDeclarator: &DirectDeclarator{
					Case: 0, // IDENTIFIER
				},
			}
			d.DirectDeclarator.declarator = d
			lhs.declarator = d
			d.setFull(lx)
		}
	case 174:
		{
			yyVAL.node = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 175:
		{
			yyVAL.node = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 176:
		{
			yyVAL.node = (*IdentifierListOpt)(nil)
		}
	case 177:
		{
			yyVAL.node = &IdentifierListOpt{
				IdentifierList: yyS[yypt-0].node.(*IdentifierList).reverse(),
			}
		}
	case 178:
		{
			yyVAL.node = (*IdentifierOpt)(nil)
		}
	case 179:
		{
			yyVAL.node = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 180:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeBlock)
		}
	case 181:
		{
			lx := yylex.(*lexer)
			lhs := &TypeName{
				SpecifierQualifierList: yyS[yypt-1].node.(*SpecifierQualifierList),
				AbstractDeclaratorOpt:  yyS[yypt-0].node.(*AbstractDeclaratorOpt),
			}
			yyVAL.node = lhs
			if o := lhs.AbstractDeclaratorOpt; o != nil {
				lhs.declarator = o.AbstractDeclarator.declarator
			} else {
				d := &Declarator{
					specifier: lhs.SpecifierQualifierList,
					DirectDeclarator: &DirectDeclarator{
						Case:    0, // IDENTIFIER
						idScope: lx.scope,
					},
				}
				d.DirectDeclarator.declarator = d
				lhs.declarator = d
			}
			lhs.Type = lhs.declarator.setFull(lx)
			lx.popScope(xc.Token{})
		}
	case 182:
		{
			lx := yylex.(*lexer)
			lhs := &AbstractDeclarator{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
			yyVAL.node = lhs
			d := &Declarator{
				specifier: lx.scope.specifier,
				PointerOpt: &PointerOpt{
					Pointer: lhs.Pointer,
				},
				DirectDeclarator: &DirectDeclarator{
					Case:    0, // IDENTIFIER
					idScope: lx.scope,
				},
			}
			d.DirectDeclarator.declarator = d
			lhs.declarator = d
		}
	case 183:
		{
			lx := yylex.(*lexer)
			lhs := &AbstractDeclarator{
				Case:                     1,
				PointerOpt:               yyS[yypt-1].node.(*PointerOpt),
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
			yyVAL.node = lhs
			d := &Declarator{
				specifier:        lx.scope.specifier,
				PointerOpt:       lhs.PointerOpt,
				DirectDeclarator: lhs.DirectAbstractDeclarator.directDeclarator,
			}
			d.DirectDeclarator.declarator = d
			lhs.declarator = d
		}
	case 184:
		{
			yyVAL.node = (*AbstractDeclaratorOpt)(nil)
		}
	case 185:
		{
			yyVAL.node = &AbstractDeclaratorOpt{
				AbstractDeclarator: yyS[yypt-0].node.(*AbstractDeclarator),
			}
		}
	case 186:
		{
			lhs := &DirectAbstractDeclarator{
				Token:              yyS[yypt-2].Token,
				AbstractDeclarator: yyS[yypt-1].node.(*AbstractDeclarator),
				Token2:             yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.AbstractDeclarator.declarator.specifier = nil
			lhs.directDeclarator = &DirectDeclarator{
				Case:       1, // '(' Declarator ')'
				Declarator: lhs.AbstractDeclarator.declarator,
			}
			lhs.AbstractDeclarator.declarator.DirectDeclarator.parent = lhs.directDeclarator
		}
	case 187:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 1,
				DirectAbstractDeclaratorOpt: yyS[yypt-3].node.(*DirectAbstractDeclaratorOpt),
				Token:         yyS[yypt-2].Token,
				ExpressionOpt: yyS[yypt-1].node.(*ExpressionOpt),
				Token2:        yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if o := lhs.ExpressionOpt; o != nil {
				o.Expression.eval(lx)
			}
			var dd *DirectDeclarator
			switch o := lhs.DirectAbstractDeclaratorOpt; {
			case o == nil:
				dd = &DirectDeclarator{
					Case: 0, // IDENTIFIER
				}
			default:
				dd = o.DirectAbstractDeclarator.directDeclarator
			}
			lhs.directDeclarator = &DirectDeclarator{
				Case:             2, // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
				DirectDeclarator: dd,
				ExpressionOpt:    lhs.ExpressionOpt,
			}
			dd.parent = lhs.directDeclarator
		}
	case 188:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 2,
				DirectAbstractDeclaratorOpt: yyS[yypt-4].node.(*DirectAbstractDeclaratorOpt),
				Token:             yyS[yypt-3].Token,
				TypeQualifierList: yyS[yypt-2].node.(*TypeQualifierList).reverse(),
				ExpressionOpt:     yyS[yypt-1].node.(*ExpressionOpt),
				Token2:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if o := lhs.ExpressionOpt; o != nil {
				o.Expression.eval(lx)
			}
			var dd *DirectDeclarator
			switch o := lhs.DirectAbstractDeclaratorOpt; {
			case o == nil:
				dd = &DirectDeclarator{
					Case: 0, // IDENTIFIER
				}
			default:
				dd = o.DirectAbstractDeclarator.directDeclarator
			}
			lhs.directDeclarator = &DirectDeclarator{
				Case:                 2, // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
				DirectDeclarator:     dd,
				TypeQualifierListOpt: &TypeQualifierListOpt{lhs.TypeQualifierList},
				ExpressionOpt:        lhs.ExpressionOpt,
			}
			dd.parent = lhs.directDeclarator
		}
	case 189:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 3,
				DirectAbstractDeclaratorOpt: yyS[yypt-5].node.(*DirectAbstractDeclaratorOpt),
				Token:                yyS[yypt-4].Token,
				Token2:               yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				Expression:           yyS[yypt-1].node.(*Expression),
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
			var dd *DirectDeclarator
			switch o := lhs.DirectAbstractDeclaratorOpt; {
			case o == nil:
				dd = &DirectDeclarator{
					Case: 0, // IDENTIFIER
				}
			default:
				dd = o.DirectAbstractDeclarator.directDeclarator
			}
			lhs.directDeclarator = &DirectDeclarator{
				Case:                 2, // DirectDeclarator '[' "static" TypeQualifierListOpt Expression ']'
				DirectDeclarator:     dd,
				TypeQualifierListOpt: lhs.TypeQualifierListOpt,
				Expression:           lhs.Expression,
			}
			dd.parent = lhs.directDeclarator
		}
	case 190:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 4,
				DirectAbstractDeclaratorOpt: yyS[yypt-5].node.(*DirectAbstractDeclaratorOpt),
				Token:             yyS[yypt-4].Token,
				TypeQualifierList: yyS[yypt-3].node.(*TypeQualifierList).reverse(),
				Token2:            yyS[yypt-2].Token,
				Expression:        yyS[yypt-1].node.(*Expression),
				Token3:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
			var dd *DirectDeclarator
			switch o := lhs.DirectAbstractDeclaratorOpt; {
			case o == nil:
				dd = &DirectDeclarator{
					Case: 0, // IDENTIFIER
				}
			default:
				dd = o.DirectAbstractDeclarator.directDeclarator
			}
			lhs.directDeclarator = &DirectDeclarator{
				Case:              4, // DirectDeclarator '[' TypeQualifierList "static" Expression ']'
				DirectDeclarator:  dd,
				TypeQualifierList: lhs.TypeQualifierList,
				Expression:        lhs.Expression,
			}
			dd.parent = lhs.directDeclarator
		}
	case 191:
		{
			lhs := &DirectAbstractDeclarator{
				Case: 5,
				DirectAbstractDeclaratorOpt: yyS[yypt-3].node.(*DirectAbstractDeclaratorOpt),
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			var dd *DirectDeclarator
			switch o := lhs.DirectAbstractDeclaratorOpt; {
			case o == nil:
				dd = &DirectDeclarator{
					Case: 0, // IDENTIFIER
				}
			default:
				dd = o.DirectAbstractDeclarator.directDeclarator
			}
			lhs.directDeclarator = &DirectDeclarator{
				Case:             5, // DirectDeclarator '[' TypeQualifierListOpt '*' ']'
				DirectDeclarator: dd,
			}
			dd.parent = lhs.directDeclarator
		}
	case 192:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 193:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case:                 6,
				Token:                yyS[yypt-3].Token,
				ParameterTypeListOpt: yyS[yypt-1].node.(*ParameterTypeListOpt),
				Token2:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScope(lhs.Token2)
			switch o := lhs.ParameterTypeListOpt; {
			case o != nil:
				lhs.directDeclarator = &DirectDeclarator{
					Case: 6, // DirectDeclarator '(' ParameterTypeList ')'
					DirectDeclarator: &DirectDeclarator{
						Case: 0, // IDENTIFIER
					},
					ParameterTypeList: o.ParameterTypeList,
				}
			default:
				lhs.directDeclarator = &DirectDeclarator{
					Case: 7, // DirectDeclarator '(' IdentifierListOpt ')'
					DirectDeclarator: &DirectDeclarator{
						Case: 0, // IDENTIFIER
					},
				}
			}
			lhs.directDeclarator.DirectDeclarator.parent = lhs.directDeclarator
		}
	case 194:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 195:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 7,
				DirectAbstractDeclarator: yyS[yypt-4].node.(*DirectAbstractDeclarator),
				Token:                yyS[yypt-3].Token,
				ParameterTypeListOpt: yyS[yypt-1].node.(*ParameterTypeListOpt),
				Token2:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScope(lhs.Token2)
			switch o := lhs.ParameterTypeListOpt; {
			case o != nil:
				lhs.directDeclarator = &DirectDeclarator{
					Case:              6, // DirectDeclarator '(' ParameterTypeList ')'
					DirectDeclarator:  lhs.DirectAbstractDeclarator.directDeclarator,
					ParameterTypeList: o.ParameterTypeList,
				}
			default:
				lhs.directDeclarator = &DirectDeclarator{
					Case:             7, // DirectDeclarator '(' IdentifierListOpt ')'
					DirectDeclarator: lhs.DirectAbstractDeclarator.directDeclarator,
				}
			}
			lhs.directDeclarator.DirectDeclarator.parent = lhs.directDeclarator
		}
	case 196:
		{
			yyVAL.node = (*DirectAbstractDeclaratorOpt)(nil)
		}
	case 197:
		{
			yyVAL.node = &DirectAbstractDeclaratorOpt{
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
		}
	case 198:
		{
			lx := yylex.(*lexer)
			lhs := &Initializer{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 199:
		{
			yyVAL.node = &Initializer{
				Case:            1,
				Token:           yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token2:          yyS[yypt-0].Token,
			}
		}
	case 200:
		{
			yyVAL.node = &InitializerList{
				DesignationOpt: yyS[yypt-1].node.(*DesignationOpt),
				Initializer:    yyS[yypt-0].node.(*Initializer),
			}
		}
	case 201:
		{
			yyVAL.node = &InitializerList{
				Case:            1,
				InitializerList: yyS[yypt-3].node.(*InitializerList),
				Token:           yyS[yypt-2].Token,
				DesignationOpt:  yyS[yypt-1].node.(*DesignationOpt),
				Initializer:     yyS[yypt-0].node.(*Initializer),
			}
		}
	case 202:
		{
			yyVAL.node = &Designation{
				DesignatorList: yyS[yypt-1].node.(*DesignatorList).reverse(),
				Token:          yyS[yypt-0].Token,
			}
		}
	case 203:
		{
			yyVAL.node = (*DesignationOpt)(nil)
		}
	case 204:
		{
			yyVAL.node = &DesignationOpt{
				Designation: yyS[yypt-0].node.(*Designation),
			}
		}
	case 205:
		{
			yyVAL.node = &DesignatorList{
				Designator: yyS[yypt-0].node.(*Designator),
			}
		}
	case 206:
		{
			yyVAL.node = &DesignatorList{
				Case:           1,
				DesignatorList: yyS[yypt-1].node.(*DesignatorList),
				Designator:     yyS[yypt-0].node.(*Designator),
			}
		}
	case 207:
		{
			yyVAL.node = &Designator{
				Token:              yyS[yypt-2].Token,
				ConstantExpression: yyS[yypt-1].node.(*ConstantExpression),
				Token2:             yyS[yypt-0].Token,
			}
		}
	case 208:
		{
			yyVAL.node = &Designator{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 209:
		{
			yyVAL.node = &Statement{
				LabeledStatement: yyS[yypt-0].node.(*LabeledStatement),
			}
		}
	case 210:
		{
			yyVAL.node = &Statement{
				Case:              1,
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
		}
	case 211:
		{
			yyVAL.node = &Statement{
				Case:                2,
				ExpressionStatement: yyS[yypt-0].node.(*ExpressionStatement),
			}
		}
	case 212:
		{
			yyVAL.node = &Statement{
				Case:               3,
				SelectionStatement: yyS[yypt-0].node.(*SelectionStatement),
			}
		}
	case 213:
		{
			yyVAL.node = &Statement{
				Case:               4,
				IterationStatement: yyS[yypt-0].node.(*IterationStatement),
			}
		}
	case 214:
		{
			yyVAL.node = &Statement{
				Case:          5,
				JumpStatement: yyS[yypt-0].node.(*JumpStatement),
			}
		}
	case 215:
		{
			yyVAL.node = &Statement{
				Case:               6,
				AssemblerStatement: yyS[yypt-0].node.(*AssemblerStatement),
			}
		}
	case 216:
		{
			yyVAL.node = &LabeledStatement{
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 217:
		{
			yyVAL.node = &LabeledStatement{
				Case:               1,
				Token:              yyS[yypt-3].Token,
				ConstantExpression: yyS[yypt-2].node.(*ConstantExpression),
				Token2:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 218:
		{
			yyVAL.node = &LabeledStatement{
				Case:      2,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 219:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 220:
		{
			lx := yylex.(*lexer)
			lhs := &CompoundStatement{
				Token:            yyS[yypt-3].Token,
				BlockItemListOpt: yyS[yypt-1].node.(*BlockItemListOpt),
				Token2:           yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.scope = lx.scope
			lx.popScope(lhs.Token2)
		}
	case 221:
		{
			yyVAL.node = &BlockItemList{
				BlockItem: yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 222:
		{
			yyVAL.node = &BlockItemList{
				Case:          1,
				BlockItemList: yyS[yypt-1].node.(*BlockItemList),
				BlockItem:     yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 223:
		{
			yyVAL.node = (*BlockItemListOpt)(nil)
		}
	case 224:
		{
			yyVAL.node = &BlockItemListOpt{
				BlockItemList: yyS[yypt-0].node.(*BlockItemList).reverse(),
			}
		}
	case 225:
		{
			yyVAL.node = &BlockItem{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 226:
		{
			yyVAL.node = &BlockItem{
				Case:      1,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 227:
		{
			yyVAL.node = &ExpressionStatement{
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 228:
		{
			lx := yylex.(*lexer)
			lhs := &SelectionStatement{
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 229:
		{
			lx := yylex.(*lexer)
			lhs := &SelectionStatement{
				Case:           1,
				Token:          yyS[yypt-6].Token,
				Token2:         yyS[yypt-5].Token,
				ExpressionList: yyS[yypt-4].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-3].Token,
				Statement:      yyS[yypt-2].node.(*Statement),
				Token4:         yyS[yypt-1].Token,
				Statement2:     yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 230:
		{
			lx := yylex.(*lexer)
			lhs := &SelectionStatement{
				Case:           2,
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 231:
		{
			lx := yylex.(*lexer)
			lhs := &IterationStatement{
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 232:
		{
			lx := yylex.(*lexer)
			lhs := &IterationStatement{
				Case:           1,
				Token:          yyS[yypt-6].Token,
				Statement:      yyS[yypt-5].node.(*Statement),
				Token2:         yyS[yypt-4].Token,
				Token3:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token4:         yyS[yypt-1].Token,
				Token5:         yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 233:
		{
			yyVAL.node = &IterationStatement{
				Case:               2,
				Token:              yyS[yypt-8].Token,
				Token2:             yyS[yypt-7].Token,
				ExpressionListOpt:  yyS[yypt-6].node.(*ExpressionListOpt),
				Token3:             yyS[yypt-5].Token,
				ExpressionListOpt2: yyS[yypt-4].node.(*ExpressionListOpt),
				Token4:             yyS[yypt-3].Token,
				ExpressionListOpt3: yyS[yypt-2].node.(*ExpressionListOpt),
				Token5:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 234:
		{
			yyVAL.node = &IterationStatement{
				Case:               3,
				Token:              yyS[yypt-7].Token,
				Token2:             yyS[yypt-6].Token,
				Declaration:        yyS[yypt-5].node.(*Declaration),
				ExpressionListOpt:  yyS[yypt-4].node.(*ExpressionListOpt),
				Token3:             yyS[yypt-3].Token,
				ExpressionListOpt2: yyS[yypt-2].node.(*ExpressionListOpt),
				Token4:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 235:
		{
			yyVAL.node = &JumpStatement{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 236:
		{
			yyVAL.node = &JumpStatement{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 237:
		{
			yyVAL.node = &JumpStatement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 238:
		{
			yyVAL.node = &JumpStatement{
				Case:              3,
				Token:             yyS[yypt-2].Token,
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 239:
		{
			yyVAL.node = &TranslationUnit{
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 240:
		{
			yyVAL.node = &TranslationUnit{
				Case:                1,
				TranslationUnit:     yyS[yypt-1].node.(*TranslationUnit),
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 241:
		{
			yyVAL.node = &ExternalDeclaration{
				FunctionDefinition: yyS[yypt-0].node.(*FunctionDefinition),
			}
		}
	case 242:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:        1,
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 243:
		{
			yyVAL.node = &ExternalDeclaration{
				Case: 2,
				BasicAssemblerStatement: yyS[yypt-1].node.(*BasicAssemblerStatement),
				Token: yyS[yypt-0].Token,
			}
		}
	case 244:
		{
			lx := yylex.(*lexer)
			lhs := &ExternalDeclaration{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableEmptyDeclarations {
				lx.report.Err(lhs.Pos(), "C++11 empty declarations are illegal in C99.")
			}
		}
	case 245:
		{
			lx := yylex.(*lexer)
			d := yyS[yypt-1].node.(*Declarator)
			lx.scope.mergeScope = nil
			dlo := yyS[yypt-0].node.(*DeclarationListOpt)
			done := false
			for dd := d.DirectDeclarator.bottom(); !done && dd != nil; dd = dd.parent {
				switch dd.Case {
				case 6: // DirectDeclarator '(' ParameterTypeList ')'
					done = true
					lx.scope.mergeScope = dd.paramsScope
					if dlo != nil {
						lx.report.Err(dlo.Pos(), "declaration list not allowed in a function definition with parameter type list")
					}
				case 7: // DirectDeclarator '(' IdentifierListOpt ')'
					done = true
					ilo := dd.IdentifierListOpt
					if ilo != nil && dlo == nil {
						lx.report.Err(ilo.Pos(), "missing parameter declaration list")
						break
					}

					if ilo == nil {
						if dlo != nil {
							lx.report.Err(dlo.Pos(), "unexpected parameter declaration list")
						}
						break
					}

					// ilo != nil && dlo != nil
					lx.scope.mergeScope = dlo.paramsScope
					ilo.post(lx.report, dlo.DeclarationList)
				}
			}
			d.setFull(lx)
			if !done {
				lx.report.Err(d.Pos(), "declarator is not a function (have '%s': %v)", d.Type, d.Type.Kind())
			}
			lx.fnDeclarator = d
		}
	case 246:
		{
			yyVAL.node = &FunctionDefinition{
				DeclarationSpecifiers: yyS[yypt-4].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt:    yyS[yypt-2].node.(*DeclarationListOpt),
				FunctionBody:          yyS[yypt-0].node.(*FunctionBody),
			}
		}
	case 247:
		{
			lx := yylex.(*lexer)
			// Handle __func__, [0], 6.4.2.2.
			id, _ := lx.fnDeclarator.Identifier()
			lx.injectFunc = []xc.Token{
				{lex.Char{Rune: STATIC}, idStatic},
				{lex.Char{Rune: CONST}, idConst},
				{lex.Char{Rune: CHAR}, idChar},
				{lex.Char{Rune: IDENTIFIER}, idMagicFunc},
				{lex.Char{Rune: '['}, 0},
				{lex.Char{Rune: ']'}, 0},
				{lex.Char{Rune: '='}, 0},
				{lex.Char{Rune: STRINGLITERAL}, xc.Dict.SID(fmt.Sprintf("%q", xc.Dict.S(id)))},
				{lex.Char{Rune: ';'}, 0},
			}
		}
	case 248:
		{
			lhs := &FunctionBody{
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
			yyVAL.node = lhs
			lhs.scope = lhs.CompoundStatement.scope
		}
	case 249:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 250:
		{
			lx := yylex.(*lexer)
			lhs := &FunctionBody{
				Case:               1,
				AssemblerStatement: yyS[yypt-1].node.(*AssemblerStatement),
				Token:              yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.scope = lx.scope
			lx.popScope(lx.tokPrev)
		}
	case 251:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 252:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 253:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 254:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 255:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScopePos(lhs.Pos())
		}
	case 256:
		{
			yyVAL.node = &AssemblerInstructions{
				Token: yyS[yypt-0].Token,
			}
		}
	case 257:
		{
			yyVAL.node = &AssemblerInstructions{
				Case: 1,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions),
				Token: yyS[yypt-0].Token,
			}
		}
	case 258:
		{
			yyVAL.node = &BasicAssemblerStatement{
				Token:                 yyS[yypt-4].Token,
				VolatileOpt:           yyS[yypt-3].node.(*VolatileOpt),
				Token2:                yyS[yypt-2].Token,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-0].Token,
			}
		}
	case 259:
		{
			yyVAL.node = (*VolatileOpt)(nil)
		}
	case 260:
		{
			yyVAL.node = &VolatileOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 261:
		{
			yyVAL.node = &AssemblerOperand{
				AssemblerSymbolicNameOpt: yyS[yypt-4].node.(*AssemblerSymbolicNameOpt),
				Token:      yyS[yypt-3].Token,
				Token2:     yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token3:     yyS[yypt-0].Token,
			}
		}
	case 262:
		{
			yyVAL.node = &AssemblerOperands{
				AssemblerOperand: yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 263:
		{
			yyVAL.node = &AssemblerOperands{
				Case:              1,
				AssemblerOperands: yyS[yypt-2].node.(*AssemblerOperands),
				Token:             yyS[yypt-1].Token,
				AssemblerOperand:  yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 264:
		{
			yyVAL.node = (*AssemblerSymbolicNameOpt)(nil)
		}
	case 265:
		{
			yyVAL.node = &AssemblerSymbolicNameOpt{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 266:
		{
			yyVAL.node = &Clobbers{
				Token: yyS[yypt-0].Token,
			}
		}
	case 267:
		{
			yyVAL.node = &Clobbers{
				Case:     1,
				Clobbers: yyS[yypt-2].node.(*Clobbers),
				Token:    yyS[yypt-1].Token,
				Token2:   yyS[yypt-0].Token,
			}
		}
	case 268:
		{
			yyVAL.node = &AssemblerStatement{
				BasicAssemblerStatement: yyS[yypt-0].node.(*BasicAssemblerStatement),
			}
		}
	case 269:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  1,
				Token:                 yyS[yypt-6].Token,
				VolatileOpt:           yyS[yypt-5].node.(*VolatileOpt),
				Token2:                yyS[yypt-4].Token,
				AssemblerInstructions: yyS[yypt-3].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-2].Token,
				AssemblerOperands:     yyS[yypt-1].node.(*AssemblerOperands).reverse(),
				Token4:                yyS[yypt-0].Token,
			}
		}
	case 270:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  2,
				Token:                 yyS[yypt-8].Token,
				VolatileOpt:           yyS[yypt-7].node.(*VolatileOpt),
				Token2:                yyS[yypt-6].Token,
				AssemblerInstructions: yyS[yypt-5].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-4].Token,
				AssemblerOperands:     yyS[yypt-3].node.(*AssemblerOperands).reverse(),
				Token4:                yyS[yypt-2].Token,
				AssemblerOperands2:    yyS[yypt-1].node.(*AssemblerOperands).reverse(),
				Token5:                yyS[yypt-0].Token,
			}
		}
	case 271:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  3,
				Token:                 yyS[yypt-10].Token,
				VolatileOpt:           yyS[yypt-9].node.(*VolatileOpt),
				Token2:                yyS[yypt-8].Token,
				AssemblerInstructions: yyS[yypt-7].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-6].Token,
				AssemblerOperands:     yyS[yypt-5].node.(*AssemblerOperands).reverse(),
				Token4:                yyS[yypt-4].Token,
				AssemblerOperands2:    yyS[yypt-3].node.(*AssemblerOperands).reverse(),
				Token5:                yyS[yypt-2].Token,
				Clobbers:              yyS[yypt-1].node.(*Clobbers).reverse(),
				Token6:                yyS[yypt-0].Token,
			}
		}
	case 272:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  4,
				Token:                 yyS[yypt-12].Token,
				VolatileOpt:           yyS[yypt-11].node.(*VolatileOpt),
				Token2:                yyS[yypt-10].Token,
				Token3:                yyS[yypt-9].Token,
				AssemblerInstructions: yyS[yypt-8].node.(*AssemblerInstructions).reverse(),
				Token4:                yyS[yypt-7].Token,
				Token5:                yyS[yypt-6].Token,
				AssemblerOperands:     yyS[yypt-5].node.(*AssemblerOperands).reverse(),
				Token6:                yyS[yypt-4].Token,
				Clobbers:              yyS[yypt-3].node.(*Clobbers).reverse(),
				Token7:                yyS[yypt-2].Token,
				IdentifierList:        yyS[yypt-1].node.(*IdentifierList).reverse(),
				Token8:                yyS[yypt-0].Token,
			}
		}
	case 273:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  5,
				Token:                 yyS[yypt-5].Token,
				VolatileOpt:           yyS[yypt-4].node.(*VolatileOpt),
				Token2:                yyS[yypt-3].Token,
				AssemblerInstructions: yyS[yypt-2].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-1].Token,
				Token4:                yyS[yypt-0].Token,
			}
		}
	case 274:
		{
			lx := yylex.(*lexer)
			lhs := &StaticAssertDeclaration{
				Token:              yyS[yypt-6].Token,
				Token2:             yyS[yypt-5].Token,
				ConstantExpression: yyS[yypt-4].node.(*ConstantExpression),
				Token3:             yyS[yypt-3].Token,
				Token4:             yyS[yypt-2].Token,
				Token5:             yyS[yypt-1].Token,
				Token6:             yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			ce := lhs.ConstantExpression
			if ce.Type == nil || ce.Type.Kind() == Undefined || ce.Value == nil || !IsIntType(ce.Type) {
				lx.report.Err(ce.Pos(), "invalid static assert expression (have '%v')", ce.Type)
				break
			}

			if !isNonZero(ce.Value) {
				lx.report.ErrTok(lhs.Token, "%s", lhs.Token4.S())
			}
		}
	case 275:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 276:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 277:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 278:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 279:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 280:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 281:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 282:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 283:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 284:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 285:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 286:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 287:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 288:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 289:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 290:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 291:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 292:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 293:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 294:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 295:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 296:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 297:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 298:
		{
			yyVAL.node = &ControlLine{
				Case:            1,
				Token:           yyS[yypt-4].Token,
				Token2:          yyS[yypt-3].Token,
				Token3:          yyS[yypt-2].Token,
				Token4:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 299:
		{
			yyVAL.node = &ControlLine{
				Case:            2,
				Token:           yyS[yypt-6].Token,
				Token2:          yyS[yypt-5].Token,
				IdentifierList:  yyS[yypt-4].node.(*IdentifierList).reverse(),
				Token3:          yyS[yypt-3].Token,
				Token4:          yyS[yypt-2].Token,
				Token5:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 300:
		{
			yyVAL.node = &ControlLine{
				Case:              3,
				Token:             yyS[yypt-4].Token,
				Token2:            yyS[yypt-3].Token,
				IdentifierListOpt: yyS[yypt-2].node.(*IdentifierListOpt),
				Token3:            yyS[yypt-1].Token,
				ReplacementList:   yyS[yypt-0].toks,
			}
		}
	case 301:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 302:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 303:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 304:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 305:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 306:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 307:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:            10,
				Token:           yyS[yypt-5].Token,
				Token2:          yyS[yypt-4].Token,
				IdentifierList:  yyS[yypt-3].node.(*IdentifierList).reverse(),
				Token3:          yyS[yypt-2].Token,
				Token4:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableDefineOmitCommaBeforeDDD {
				lx.report.ErrTok(lhs.Token4, "missing comma before \"...\"")
			}
		}
	case 308:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:   11,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableEmptyDefine {
				lx.report.ErrTok(lhs.Token2, "expected identifier")
			}
		}
	case 309:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:        12,
				Token:       yyS[yypt-3].Token,
				Token2:      yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token3:      yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			toks := decodeTokens(lhs.PPTokenList, nil, false)
			if len(toks) == 0 {
				lhs.Case = 9 // PPUNDEF IDENTIFIER '\n'
				break
			}

			lx.report.ErrTok(toks[0], "extra tokens after #undef argument")
		}
	case 310:
		{
			yyVAL.node = &ControlLine{
				Case:        13,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 313:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 314:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
