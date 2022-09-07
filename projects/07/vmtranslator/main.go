package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
	"vmtranslator/ast"
	"vmtranslator/codewriter"
	"vmtranslator/parser"
)

func main() {
	flag.Parse()
	vmFilePath := flag.Args()[0]
	vmDirName := path.Base(vmFilePath)

	asmFileName := fmt.Sprintf("%s.asm", strings.Split(vmDirName, ".")[0])
	asmFile := path.Join(vmFilePath, asmFileName)

	files, _ := ioutil.ReadDir(vmFilePath)
	var vmFiles []string
	for _, file := range files {
		fileName := file.Name()
		if isVmFile(fileName) {
			vmFiles = append(vmFiles, path.Join(vmFilePath, fileName))
		}
	}

	Translate(vmFiles, asmFile)
}

func isVmFile(fileName string) bool {
	return filepath.Ext(fileName) == ".vm"
}

func Translate(vmFiles []string, asmFile string) {
	cw := codewriter.New(asmFile)

	cw.WriteInit()

	for _, vmFile := range vmFiles {
		vmClassName := strings.Split(path.Base(vmFile), ".")[0]
		cw.SetVmClassName(vmClassName)

		input, _ := ioutil.ReadFile(vmFile)
		p := parser.New(string(input))

		for p.HasMoreCommands() {
			p.SetCurrentCommandArray()
			switch p.CommandType() {
			case ast.C_PUSH:
				cw.WritePushPop(ast.PUSH, ast.SegmentSymbol(p.Arg1()), p.Arg2())
			case ast.C_POP:
				cw.WritePushPop(ast.POP, ast.SegmentSymbol(p.Arg1()), p.Arg2())
			case ast.C_ARITHMETIC:
				cw.WriteArithmetic(ast.CommandSymbol(p.Arg1()))
			case ast.C_LABEL:
				cw.WriteLabel(p.Arg1())
			case ast.C_GOTO:
				cw.WriteGoto(p.Arg1())
			case ast.C_IF:
				cw.WriteIf(p.Arg1())
			case ast.C_CALL:
				cw.WriteCall(p.Arg1(), p.Arg2())
			case ast.C_RETURN:
				cw.WriteReturn()
			case ast.C_FUNCTION:
				cw.WriteFunction(p.Arg1(), p.Arg2())
			}

			p.Advance()
		}
	}

	cw.Close()
}
