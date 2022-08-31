package ast

type CommandType string

const (
	A_COMMAND = "A_COMMAND"
	C_COMMAND = "C_COMMAND"
	L_COMMAND = "L_COMMAND"
)

type Command interface {
	Type() CommandType
}

type ACommand struct {
	Value string
}

func (ac *ACommand) Type() CommandType {
	return A_COMMAND
}

type CCommand struct {
	Dest string
	Comp string
	Jump string
}

func (cc *CCommand) Type() CommandType {
	return C_COMMAND
}

type LCommand struct {
	Value string
}

func (lc *LCommand) Type() CommandType {
	return L_COMMAND
}
