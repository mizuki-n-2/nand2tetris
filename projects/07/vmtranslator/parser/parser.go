package parser

import (
	"strconv"
	"strings"
	"vmtranslator/ast"
)

type Parser struct {
	input               string
	commandStrList      []string
	currentCommandIdx   int
	currentCommandArray []string
}

func New(input string) *Parser {
	commandStrList := strings.Split(input, "\r\n")
	return &Parser{
		input:               input,
		commandStrList:      commandStrList,
		currentCommandIdx:   0,
		currentCommandArray: strings.Split(commandStrList[0], " "),
	}
}

func (p *Parser) HasMoreCommands() bool {
	return p.currentCommandIdx < len(p.commandStrList)
}

func (p *Parser) Advance() {
	p.currentCommandIdx++
}

func (p *Parser) SetCurrentCommandArray() {
	p.currentCommandArray = strings.Split(p.commandStrList[p.currentCommandIdx], " ")
}

func (p *Parser) CommandType() ast.CommandType {
	if p.commandStrList[p.currentCommandIdx] == "" {
		return ""
	}

	switch p.currentCommandArray[0] {
	case ast.ADD, ast.SUB, ast.NEG, ast.EQ, ast.GT, ast.LT, ast.AND, ast.OR, ast.NOT:
		return ast.C_ARITHMETIC
	case ast.PUSH:
		return ast.C_PUSH
	case ast.POP:
		return ast.C_POP
	case ast.LABEL:
		return ast.C_LABEL
	case ast.GOTO:
		return ast.C_GOTO
	case ast.IF_GOTO:
		return ast.C_IF
	case ast.FUNCTION:
		return ast.C_FUNCTION
	case ast.CALL:
		return ast.C_CALL
	case ast.RETURN:
		return ast.C_RETURN
	default:
		return ""
	}
}

func (p *Parser) removeWhiteSpace() {
	p.commandStrList[p.currentCommandIdx] = strings.Replace(p.commandStrList[p.currentCommandIdx], string(byte(' ')), "", -1)
	p.commandStrList[p.currentCommandIdx] = strings.Replace(p.commandStrList[p.currentCommandIdx], string(byte('\t')), "", -1)
}

func (p *Parser) Arg1() string {
	switch p.CommandType() {
	case ast.C_ARITHMETIC:
		return p.currentCommandArray[0]
	default:
		return p.currentCommandArray[1]
	}
}

func (p *Parser) Arg2() int {
	arg2, _ := strconv.Atoi(p.currentCommandArray[2])
	return arg2
}
