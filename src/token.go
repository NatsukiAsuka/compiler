package token

import "strconv"

type Token int

const (
	ILLEGAL Token = iota
	COMMENT
	EOF

	literals_beg
	INTEGERLITERAL
	FLOATINGPOINTLITERAL
	BOOLEANLITERAL
	CHARACTERLITERAL
	STRINGLITERAL
	NULLLITERAL
	literals_end

	separators_beg
	LPAREN          // (
	LBRACK          // [
	LBRACE          // {
	COMMA           // ,
	PERIOD          // .
	RPAREN          // )
	RBRACK          // ]
	RBRACE          // }
	SEMICOLON       // ;
	separators_end

	operator_beg
	ADD           // +
	SUB           // -
	MUL           // *
	QUO           // /
	REM           // %
	COLON         // :
	TERNARY       // ?
	AND           // &
	OR            // |
	XOR           // ^
	SHL           // <<
	SHR           // >>
	SHUR          // >>>
	ADD_ASSIGN    // +=
	SUB_ASSIGN    // -=
	MUL_ASSIGN    // *=
	QUO_ASSIGN    // /=
	REM_ASSIGN    // %=
	AND_ASSIGN    // &=
	OR_ASSIGN     // |=
	XOR_ASSIGN    // ^=
	SHL_ASSIGN    // <<=
	SHUL_ASSIGN   // <<<=
	SHR_ASSIGN    // >>=
	SHUR_ASSIGN   // >>>=
	LAND          // &&
	LOR           // ||
	INC           // ++
	DEC           // --
	EQL           // ==
	LSS           // <
	GTR           // >
	ASSIGN        // =
	NOT           // !
	NEQ           // !=
	LEQ           // <=
	GEQ           // >=
	operator_end

	keyword_beg
	ABSTRACT
	CONTINUE
	FOR
	NEW
	SWITCH
	ASSERT
	DEFAULT
	IF
	PACKAGE
	SYNCHRONIZED
	BOOLEAN
	DO
	GOTO
	PRIVATE
	THIS
	BREAK
	DOUBLE
	IMPLEMENTS
	PROTECTED
	THROW
	BYTE
	ELSE
	IMPORT
	PUBLIC
	THROWS
	CASE
	ENUM
	INSTANCEOF
	RETURN
	TRANSIENT
	CATCH
	EXTENDS
	INT
	SHORT
	TRY
	CHAR
	FINAL
	INTERFACE
	STATIC
	VOID
	CLASS
	FINALLY
	LONG
	STRICTFP
	VOLATILE
	CONST
	FLOAT
	NATIVE
	SUPER
	WHILE
	keyword_end
)

var tokens = [...]string{
	ILLEGAL:              "ILLEGAL",
	COMMENT:              "COMMENT",
	EOF:                  "EOF",
	INTEGERLITERAL:       "INTEGERLITERAL",
	FLOATINGPOINTLITERAL: "FLOATINGPOINTLITERAL",
	BOOLEANLITERAL:       "BOOLEANLITERAL",
	CHARACTERLITERAL:     "CHARACTERLITERAL",
	STRINGLITERAL:        "STRINGLITERAL",
	NULLLITERAL:          "NULLLITERAL",
	LPAREN:               "(",
	LBRACK:               "[",
	LBRACE:               "{",
	COMMA:                ",",
	PERIOD:               ".",
	RPAREN:               ")",
	RBRACK:               "]",
	RBRACE:               "}",
	SEMICOLON:            ";",
	ADD:                  "+",
	SUB:                  "-",
	MUL:                  "*",
	QUO:                  "/",
	REM:                  "%",
	COLON:                ":",
	TERNARY:              "?",
	AND:                  "&",
	OR:                   "|",
	XOR:                  "^",
	SHL:                  "<<",
	SHR:                  ">>",
	SHUR:                 ">>>",
	ADD_ASSIGN:           "+=",
	SUB_ASSIGN:           "-=",
	MUL_ASSIGN:           "*=",
	QUO_ASSIGN:           "/=",
	REM_ASSIGN:           "%=",
	AND_ASSIGN:           "&=",
	OR_ASSIGN:            "|=",
	XOR_ASSIGN:           "^=",
	SHL_ASSIGN:           "<<=",
	SHUL_ASSIGN:          "<<<=",
	SHR_ASSIGN:           ">>=",
	SHUR_ASSIGN:          ">>>=",
	LAND:                 "&&",
	LOR:                  "||",
	INC:                  "++",
	DEC:                  "--",
	EQL:                  "==",
	LSS:                  "<",
	GTR:                  ">",
	ASSIGN:               "=",
	NOT:                  "!",
	NEQ:                  "!=",
	LEQ:                  "<=",
	GEQ:                  ">=",
	ABSTRACT:             "ABSTRACT",
	CONTINUE:             "CONTINUE",
	FOR:                  "FOR",
	NEW:                  "NEW",
	SWITCH:               "SWITCH",
	ASSERT:               "ASSERT",
	DEFAULT:              "DEFAULT",
	IF:                   "IF",
	PACKAGE:              "PACKAGE",
	SYNCHRONIZED:         "SYNCHRONIZED",
	BOOLEAN:              "BOOLEAN",
	DO:                   "DO",
	GOTO:                 "GOTO",
	PRIVATE:              "PRIVATE",
	THIS:                 "THIS",
	BREAK:                "BREAK",
	DOUBLE:               "DOUBLE",
	IMPLEMENTS:           "IMPLEMENTS",
	PROTECTED:            "PROTECTED",
	THROW:                "THROW",
	BYTE:                 "BYTE",
	ELSE:                 "ELSE",
	IMPORT:               "IMPORT",
	PUBLIC:               "PUBLIC",
	THROWS:               "THROWS",
	CASE:                 "CASE",
	ENUM:                 "ENUM",
	INSTANCEOF:           "INSTANCEOF",
	RETURN:               "RETURN",
	TRANSIENT:            "TRANSIENT",
	CATCH:                "CATCH",
	EXTENDS:              "EXTENDS",
	INT:                  "INT",
	SHORT:                "SHORT",
	TRY:                  "TRY",
	CHAR:                 "CHAR",
	FINAL:                "FINAL",
	INTERFACE:            "INTERFACE",
	STATIC:               "STATIC",
	VOID:                 "VOID",
	CLASS:                "CLASS",
	FINALLY:              "FINALLY",
	LONG:                 "LONG",
	STRICTFP:             "STRICTFP",
	VOLATILE:             "VOLATILE",
	CONST:                "CONST",
	FLOAT:                "FLOAT",
	NATIVE:               "NATIVE",
	SUPER:                "SUPER",
	WHILE:                "WHILE",
}

func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token)
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

func Lookup(id string) Token {
	if tok, isKeyword := keywords[id]; isKeyword {
		return tok
	}
	return STRINGLITERAL
}

func (tok Token) IsLiteral() bool { return literals_beg < tok && tok < literals_end }

func (tok Token) IsOperator() bool { return operator_beg < tok && tok < operator_end }

func (tok Token) IsKeyword() bool { return keyword_beg < tok && tok < keyword_end }

func (tok Token) IsSeparator() bool { return separators_beg < tok && tok < separators_end }

const (
	LowestPrec  = 0
	Unary       = 13
	Postfix     = 14
	HighestPrec = 999
)

func (op Token) Precedence() int {
	switch op {
	case ADD_ASSIGN, SUB_ASSIGN, MUL_ASSIGN, QUO_ASSIGN, REM_ASSIGN, AND_ASSIGN, OR_ASSIGN, XOR_ASSIGN, SHL_ASSIGN, SHUL_ASSIGN, SHR_ASSIGN, SHUR_ASSIGN, ASSIGN:
		return 1
	case TERNARY:
		return 2
	case LOR:
		return 3
	case LAND:
		return 4
	case OR:
		return 5
	case XOR:
		return 6
	case AND:
		return 7
	case EQL, NEQ:
		return 8
	case LEQ, GEQ, LSS, GTR, INSTANCEOF:
		return 9
	case SHL, SHR, SHUR:
		return 10
	case ADD, SUB:
		return 11
	case QUO, REM, MUL:
		return 12
	}
	return LowestPrec
}