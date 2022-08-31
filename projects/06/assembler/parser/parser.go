package parser

import (
	"assembler/ast"
	"strings"
)

type Parser struct {
	input             string
	commandStrList    []string
	currentCommandIdx int
	position          int
	readPosition      int
}

func New(input string) *Parser {
	return &Parser{
		input:             input,
		commandStrList:    strings.Split(input, "\r\n"),
		currentCommandIdx: 0,
		position:          0,
		readPosition:      0,
	}
}

func (p *Parser) HasMoreCommands() bool {
	return p.currentCommandIdx < len(p.commandStrList)
}

func (p *Parser) Advance() {
	p.currentCommandIdx++
	p.position = 0
	p.readPosition = 0
}

func (p *Parser) CommandType() ast.CommandType {
	if p.commandStrList[p.currentCommandIdx] == "" {
		return ""
	}

	switch p.commandStrList[p.currentCommandIdx][0] {
	case '@':
		return ast.A_COMMAND
	case '(':
		return ast.L_COMMAND
	case '0', '1', '-', '!', 'D', 'A', 'M':
		return ast.C_COMMAND
	default:
		return ""
	}
}

func (p *Parser) Symbol() string {
	switch p.CommandType() {
	case ast.A_COMMAND:
		aCommand := p.parseACommand()
		return aCommand.Value
	case ast.L_COMMAND:
		lCommand := p.parseLCommand()
		return lCommand.Value
	default:
		return ""
	}
}

func (p *Parser) parseACommand() *ast.ACommand {
	// read @
	p.readChar()
	value := ""
	for p.hasMoreChar() {
		value += string(p.commandStrList[p.currentCommandIdx][p.readPosition])
		p.readChar()
	}

	return &ast.ACommand{Value: value}
}

func (p *Parser) parseCCommand() *ast.CCommand {
	dest := ""
	if strings.Contains(p.commandStrList[p.currentCommandIdx], "=") {
		for p.hasMoreChar() {
			if p.peekCharIs('=') {
				break
			}
			dest += string(p.commandStrList[p.currentCommandIdx][p.readPosition])
			p.readChar()
		}

		// read =
		p.readChar()
	}

	comp := ""
	for p.hasMoreChar() {
		if p.peekCharIs(';') {
			break
		}
		comp += string(p.commandStrList[p.currentCommandIdx][p.readPosition])
		p.readChar()
	}

	// read ;
	p.readChar()

	jump := ""
	for p.hasMoreChar() {
		jump += string(p.commandStrList[p.currentCommandIdx][p.readPosition])
		p.readChar()
	}

	return &ast.CCommand{
		Dest: dest,
		Comp: comp,
		Jump: jump,
	}
}

func (p *Parser) parseLCommand() *ast.LCommand {
	// read (
	p.readChar()
	value := ""
	for p.hasMoreChar() {
		if p.peekCharIs(')') {
			break
		}
		value += string(p.commandStrList[p.currentCommandIdx][p.readPosition])
		p.readChar()
	}

	return &ast.LCommand{Value: value}
}

func (p *Parser) hasMoreChar() bool {
	return p.readPosition < len(p.commandStrList[p.currentCommandIdx])
}

func (p *Parser) peekCharIs(char byte) bool {
	return p.commandStrList[p.currentCommandIdx][p.readPosition] == char
}

func (p *Parser) readChar() {
	p.position = p.readPosition
	p.readPosition++
}

func (p *Parser) Dest() string {
	p.resetPosition()
	cCommand := p.parseCCommand()
	return cCommand.Dest
}

func (p *Parser) Comp() string {
	p.resetPosition()
	cCommand := p.parseCCommand()
	return cCommand.Comp
}

func (p *Parser) Jump() string {
	p.resetPosition()
	cCommand := p.parseCCommand()
	return cCommand.Jump
}

func (p *Parser) removeWhiteSpace() {
	p.commandStrList[p.currentCommandIdx] = strings.Replace(p.commandStrList[p.currentCommandIdx], string(byte(' ')), "", -1)
	p.commandStrList[p.currentCommandIdx] = strings.Replace(p.commandStrList[p.currentCommandIdx], string(byte('\t')), "", -1)
}

func (p *Parser) resetPosition() {
	p.position = 0
	p.readPosition = 0
}
