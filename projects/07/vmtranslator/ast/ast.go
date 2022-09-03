package ast

type CommandType string

const (
	C_ARITHMETIC = "C_ARITHMETIC"
	C_PUSH       = "C_PUSH"
	C_POP        = "C_POP"
	C_LABEL      = "C_LABEL"
	C_GOTO       = "C_GOTO"
	C_IF         = "C_IF"
	C_FUNCTION   = "C_FUNCTION"
	C_RETURN     = "C_RETURN"
	C_CALL       = "C_CALL"
)

type CommandSymbol string

const (
	PUSH     = "push"
	POP      = "pop"
	LABEL    = "label"
	GOTO     = "goto"
	IF_GOTO  = "if-goto"
	CALL     = "call"
	FUNCTION = "function"
	RETURN   = "return"
	ADD      = "add"
	SUB      = "sub"
	NEG      = "neg"
	EQ       = "eq"
	GT       = "gt"
	LT       = "lt"
	AND      = "and"
	OR       = "or"
	NOT      = "not"
)

type SegmentSymbol string

const (
	ARGUMENT = "argument"
	LOCAL    = "local"
	STATIC   = "static"
	CONSTANT = "constant"
	THIS     = "this"
	THAT     = "that"
	POINTER  = "pointer"
	TEMP     = "temp"
)
