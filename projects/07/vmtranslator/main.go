package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"vmtranslator/ast"
	"vmtranslator/codewriter"
	"vmtranslator/parser"
)

func main() {
	flag.Parse()
	// TODO: 入力をdirとし、その配下のvmファイル全てに適用するように変更
	vmFile := flag.Args()[0]
	vmDirName := path.Dir(vmFile)
	vmFileName := path.Base(vmFile)
	asmFileName := fmt.Sprintf("%s.asm", strings.Split(vmFileName, ".")[0])
	asmFile := path.Join(vmDirName, asmFileName)

	vm, _ := ioutil.ReadFile(vmFile)
	Translate(string(vm), asmFile)
}

func Translate(input string, asmFile string) {
	cw := codewriter.New()

	// TODO: 複数ファイル対応
	p := parser.New(input)

	cw.SetFileName(asmFile)

	for p.HasMoreCommands() {
		p.SetCurrentCommandArray()
		switch p.CommandType() {
		case ast.C_PUSH:
			cw.WritePushPop(ast.PUSH, ast.SegmentSymbol(p.Arg1()), p.Arg2())
		case ast.C_POP:
			cw.WritePushPop(ast.POP, ast.SegmentSymbol(p.Arg1()), p.Arg2())
		case ast.C_ARITHMETIC:
			cw.WriteArithmetic(ast.CommandSymbol(p.Arg1()))
		}

		p.Advance()
	}

	cw.Close()
}
