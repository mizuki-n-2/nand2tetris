package main

import (
	"assembler/ast"
	"assembler/code"
	"assembler/parser"
	"assembler/symboltable"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	asmFile := flag.Args()[0]
	asmDirName := path.Dir(asmFile)
	asmFileName := path.Base(asmFile)
	hackFileName := fmt.Sprintf("%s.hack", strings.Split(asmFileName, ".")[0])
	hackFile := path.Join(asmDirName, hackFileName)

	asm, _ := ioutil.ReadFile(asmFile)
	stringArray := Assemble(string(asm))
	ioutil.WriteFile(hackFile, []byte(strings.Join(stringArray, "\r\n")), os.ModePerm)
}

func Assemble(input string) []string {
	symbolTable := symboltable.New()
	binaryArray := []string{}
	p := parser.New(input)

	currentCommandAddress := 0
	// 1回目のパス
	for p.HasMoreCommands() {
		switch p.CommandType() {
		case ast.A_COMMAND, ast.C_COMMAND:
			currentCommandAddress++
		case ast.L_COMMAND:
			symbolTable.AddEntry(p.Symbol(), currentCommandAddress)
		}

		p.Advance()
	}

	p.ResetParsePosition()

	variableCount := 0
	initialAddress := 16
	// 2回目のパス
	for p.HasMoreCommands() {
		switch p.CommandType() {
		case ast.A_COMMAND:
			symbol := p.Symbol()
			value, err := strconv.Atoi(symbol)
			if err != nil {
				if symbolTable.Contains(symbol) {
					value = symbolTable.GetAddress(symbol)
				} else {
					symbolTable.AddEntry(symbol, initialAddress+variableCount)
					value = initialAddress + variableCount
					variableCount++
				}
			}
			binary := fmt.Sprintf("%016b", value)
			binaryArray = append(binaryArray, binary)
		case ast.C_COMMAND:
			dest := code.Dest(p.Dest())
			comp := code.Comp(p.Comp())
			jump := code.Jump(p.Jump())
			binary := "111" + comp + dest + jump
			binaryArray = append(binaryArray, binary)
		}

		p.Advance()
	}

	return binaryArray
}
