// Code generated by goyacc -o compiler/parser.go compiler/parser.go.y. DO NOT EDIT.

//line compiler/parser.go.y:2
// プログラムのヘッダを指定
package compiler

import __yyfmt__ "fmt"

//line compiler/parser.go.y:3

import (
	"fmt"
)

var driver *Driver

//line compiler/parser.go.y:13
type yySymType struct {
	yys  int
	ival int
	fval float32
	sval string

	node INode
}

const INUM = 57346
const FNUM = 57347
const IDENTIFIER = 57348
const STRING_LITERAL = 57349
const PLUS = 57350
const MINUS = 57351
const ASTARISK = 57352
const SLASH = 57353
const PERCENT = 57354
const P_EQ = 57355
const M_EQ = 57356
const A_EQ = 57357
const S_EQ = 57358
const EQ = 57359
const NEQ = 57360
const GT = 57361
const GE = 57362
const LT = 57363
const LE = 57364
const AND = 57365
const OR = 57366
const INCR = 57367
const DECR = 57368
const ASSIGN = 57369
const VAR = 57370
const INT = 57371
const FLOAT = 57372
const STRING = 57373
const VOID = 57374
const IF = 57375
const ELSE = 57376
const SWITCH = 57377
const CASE = 57378
const DEFAULT = 57379
const FALLTHROUGH = 57380
const FOR = 57381
const BREAK = 57382
const CONTINUE = 57383
const FUNC = 57384
const RETURN = 57385
const IMPORT = 57386
const TYPE = 57387
const STRUCT = 57388
const EOL = 57389
const NEG = 57390

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"INUM",
	"FNUM",
	"IDENTIFIER",
	"STRING_LITERAL",
	"PLUS",
	"MINUS",
	"ASTARISK",
	"SLASH",
	"PERCENT",
	"P_EQ",
	"M_EQ",
	"A_EQ",
	"S_EQ",
	"EQ",
	"NEQ",
	"GT",
	"GE",
	"LT",
	"LE",
	"AND",
	"OR",
	"INCR",
	"DECR",
	"ASSIGN",
	"VAR",
	"INT",
	"FLOAT",
	"STRING",
	"VOID",
	"IF",
	"ELSE",
	"SWITCH",
	"CASE",
	"DEFAULT",
	"FALLTHROUGH",
	"FOR",
	"BREAK",
	"CONTINUE",
	"FUNC",
	"RETURN",
	"IMPORT",
	"TYPE",
	"STRUCT",
	"EOL",
	"'('",
	"')'",
	"NEG",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line compiler/parser.go.y:105

func Parse(filename string, source string) int {

	driver = new(Driver)
	driver.Init(filename)

	// パース処理
	lexer := &Lexer{src: source, position: 0, readPosition: 0, line: 1, filename: filename, driver: driver}
	yyParse(lexer)

	fmt.Println("Parse End.")

	// パース結果出力
	driver.Dump()

	return 0
}

// 外部用
func GetErrorCount() int {
	return driver.err.errorCount
}
func GetWarningCount() int {
	return driver.err.warningCount
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 43,
	17, 0,
	18, 0,
	19, 0,
	20, 0,
	21, 0,
	22, 0,
	-2, 21,
	-1, 44,
	17, 0,
	18, 0,
	19, 0,
	20, 0,
	21, 0,
	22, 0,
	-2, 22,
	-1, 45,
	17, 0,
	18, 0,
	19, 0,
	20, 0,
	21, 0,
	22, 0,
	-2, 23,
	-1, 46,
	17, 0,
	18, 0,
	19, 0,
	20, 0,
	21, 0,
	22, 0,
	-2, 24,
	-1, 47,
	17, 0,
	18, 0,
	19, 0,
	20, 0,
	21, 0,
	22, 0,
	-2, 25,
	-1, 48,
	17, 0,
	18, 0,
	19, 0,
	20, 0,
	21, 0,
	22, 0,
	-2, 26,
}

const yyPrivate = 57344

const yyLast = 180

var yyAct = [...]int{
	5, 32, 15, 54, 58, 55, 56, 57, 17, 18,
	34, 36, 12, 13, 35, 14, 33, 9, 37, 1,
	38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
	48, 49, 50, 6, 51, 19, 20, 21, 22, 23,
	8, 4, 2, 53, 24, 25, 26, 27, 28, 29,
	30, 31, 17, 18, 0, 59, 10, 0, 0, 60,
	0, 0, 0, 0, 0, 0, 0, 0, 19, 20,
	21, 22, 23, 0, 0, 0, 52, 24, 25, 26,
	27, 28, 29, 30, 31, 17, 18, 12, 13, 7,
	14, 0, 9, 0, 0, 19, 20, 21, 22, 23,
	21, 22, 23, 0, 0, 0, 0, 16, 0, 0,
	0, 11, 17, 18, 0, 17, 18, 0, 0, 0,
	0, 0, 0, 19, 20, 21, 22, 23, 0, 0,
	3, 10, 24, 25, 26, 27, 28, 29, 30, 31,
	17, 18, 19, 20, 21, 22, 23, 0, 0, 0,
	0, 24, 25, 26, 27, 28, 29, 30, 0, 17,
	18, 19, 20, 21, 22, 23, 0, 0, 0, 0,
	24, 25, 26, 27, 28, 29, 0, 0, 17, 18,
}

var yyPact = [...]int{
	-1000, 83, -1000, -1000, -45, 60, -46, -11, -1000, 8,
	8, 12, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, -1000, 8, -1000, -1000, 27, -24, 90, 90,
	-17, -17, -17, 87, 87, 87, 87, 87, 87, 153,
	134, 115, -1000, -23, 8, -1000, -1000, -1000, 8, 115,
	115,
}

var yyPgo = [...]int{
	0, 43, 42, 0, 41, 40, 33, 19,
}

var yyR1 = [...]int{
	0, 7, 7, 2, 2, 2, 2, 6, 6, 6,
	4, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	5, 5, 5, 1, 1, 1,
}

var yyR2 = [...]int{
	0, 0, 2, 1, 2, 2, 2, 3, 5, 4,
	3, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	1, 1, 1, 1, 1, 1,
}

var yyChk = [...]int{
	-1000, -7, -2, 47, -4, -3, -6, 6, -5, 9,
	48, 28, 4, 5, 7, 47, 47, 25, 26, 8,
	9, 10, 11, 12, 17, 18, 19, 20, 21, 22,
	23, 24, 47, 27, -3, 6, -3, 6, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, 49, -1, 27, 29, 30, 31, 27, -3,
	-3,
}

var yyDef = [...]int{
	1, -2, 2, 3, 0, 0, 0, 12, 11, 0,
	0, 0, 30, 31, 32, 4, 5, 14, 15, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 6, 0, 13, 12, 0, 0, 16, 17,
	18, 19, 20, -2, -2, -2, -2, -2, -2, 27,
	28, 10, 29, 7, 0, 33, 34, 35, 0, 9,
	8,
}

var yyTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	48, 49,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 50,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
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
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line compiler/parser.go.y:54
		{
			if yyDollar[2].node != nil {
				yyDollar[2].node.Push()
			}
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line compiler/parser.go.y:57
		{
			yyVAL.node = nil
			driver.lineno++
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
//line compiler/parser.go.y:58
		{
			yyVAL.node = yyDollar[1].node
			driver.lineno++
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line compiler/parser.go.y:59
		{
			yyVAL.node = yyDollar[1].node
			driver.lineno++
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line compiler/parser.go.y:60
		{
			yyVAL.node = yyDollar[1].node
			driver.lineno++
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:63
		{
			yyVAL.node = nil
			driver.variableTable.DefineInLocal(yyDollar[2].sval, yyDollar[3].ival)
		}
	case 8:
		yyDollar = yyS[yypt-5 : yypt+1]
//line compiler/parser.go.y:64
		{
			yyVAL.node = driver.variableTable.DefineInLocalWithAssign(yyDollar[2].sval, yyDollar[3].ival, yyDollar[5].node)
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
//line compiler/parser.go.y:65
		{
			yyVAL.node = driver.variableTable.DefineInLocalWithAssignAutoType(yyDollar[2].sval, yyDollar[4].node)
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:69
		{
			varNode := MakeValueNode(yyDollar[1].sval, driver)
			yyVAL.node = &AssignNode{Node: Node{left: varNode, right: yyDollar[3].node, driver: driver}}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line compiler/parser.go.y:76
		{
			yyVAL.node = MakeValueNode(yyDollar[1].sval, driver)
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
//line compiler/parser.go.y:77
		{
			yyVAL.node = &Node{left: yyDollar[2].node, right: nil, op: OP_NOT, driver: driver}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
//line compiler/parser.go.y:78
		{
			yyVAL.node = &Node{left: yyDollar[1].node, right: nil, op: OP_INCR, driver: driver}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
//line compiler/parser.go.y:79
		{
			yyVAL.node = &Node{left: yyDollar[1].node, right: nil, op: OP_DECR, driver: driver}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:80
		{
			yyVAL.node = &Node{left: yyDollar[1].node, right: yyDollar[3].node, op: OP_ADD, driver: driver}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:81
		{
			yyVAL.node = &Node{left: yyDollar[1].node, right: yyDollar[3].node, op: OP_SUB, driver: driver}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:82
		{
			yyVAL.node = &Node{left: yyDollar[1].node, right: yyDollar[3].node, op: OP_MUL, driver: driver}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:83
		{
			yyVAL.node = &Node{left: yyDollar[1].node, right: yyDollar[3].node, op: OP_DIV, driver: driver}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:84
		{
			yyVAL.node = &Node{left: yyDollar[1].node, right: yyDollar[3].node, op: OP_MOD, driver: driver}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:85
		{
			yyVAL.node = &Node{left: yyDollar[1].node, right: yyDollar[3].node, op: OP_EQUAL, driver: driver}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:86
		{
			yyVAL.node = &Node{left: yyDollar[1].node, right: yyDollar[3].node, op: OP_NEQ, driver: driver}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:87
		{
			yyVAL.node = &Node{left: yyDollar[1].node, right: yyDollar[3].node, op: OP_GT, driver: driver}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:88
		{
			yyVAL.node = &Node{left: yyDollar[1].node, right: yyDollar[3].node, op: OP_GE, driver: driver}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:89
		{
			yyVAL.node = &Node{left: yyDollar[1].node, right: yyDollar[3].node, op: OP_LT, driver: driver}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:90
		{
			yyVAL.node = &Node{left: yyDollar[1].node, right: yyDollar[3].node, op: OP_LE, driver: driver}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:91
		{
			yyVAL.node = &Node{left: yyDollar[1].node, right: yyDollar[3].node, op: OP_AND, driver: driver}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:92
		{
			yyVAL.node = &Node{left: yyDollar[1].node, right: yyDollar[3].node, op: OP_OR, driver: driver}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line compiler/parser.go.y:93
		{
			yyVAL.node = yyDollar[2].node
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line compiler/parser.go.y:96
		{
			yyVAL.node = &Node{op: OP_INTEGER, driver: driver, ival: yyDollar[1].ival}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
//line compiler/parser.go.y:97
		{
			yyVAL.node = &Node{op: OP_FLOAT, driver: driver, fval: yyDollar[1].fval}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line compiler/parser.go.y:98
		{
			yyVAL.node = &Node{op: OP_STRING, driver: driver, sval: yyDollar[1].sval}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line compiler/parser.go.y:101
		{
			yyVAL.ival = TYPE_INTEGER
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line compiler/parser.go.y:102
		{
			yyVAL.ival = TYPE_FLOAT
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line compiler/parser.go.y:103
		{
			yyVAL.ival = TYPE_STRING
		}
	}
	goto yystack /* stack new state and value */
}
