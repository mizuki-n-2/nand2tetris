package main

import (
	"assembler/ast"
	"assembler/code"
	"assembler/parser"
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
	binaryArray := []string{}
	p := parser.New(input)

	for p.HasMoreCommands() {
		switch p.CommandType() {
		case ast.A_COMMAND:
			symbol := p.Symbol()
			value, _ := strconv.Atoi(symbol)
			binary := fmt.Sprintf("%016b", value)
			binaryArray = append(binaryArray, binary)
		case ast.C_COMMAND:
			dest := code.Dest(p.Dest())
			comp := code.Comp(p.Comp())
			jump := code.Jump(p.Jump())
			binary := "111" + comp + dest + jump
			binaryArray = append(binaryArray, binary)
		case ast.L_COMMAND:
		default:
		}

		p.Advance()
	}

	return binaryArray
}
