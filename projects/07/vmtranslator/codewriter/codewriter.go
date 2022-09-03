package codewriter

import (
	"io/ioutil"
	"os"
	"strconv"
	"vmtranslator/ast"
)

type CodeWriter struct {
	fileName string
	assembly []byte
}

func New(fileName string) *CodeWriter {
	assembly := getInitAssembly()
	return &CodeWriter{fileName: fileName, assembly: []byte(assembly)}
}

func getInitAssembly() string {
	assembly := ""
	// SPに256を設定
	assembly += "@256" + "\r\n"
	assembly += "D=A" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=D" + "\r\n"
	return assembly
}

func (cw *CodeWriter) SetFileName(fileName string) {
	cw.fileName = fileName
}

func (cw *CodeWriter) WriteArithmetic(command ast.CommandSymbol) {
	var assembly string
	switch command {
	case ast.ADD:
		assembly = cw.getAddAssembly()
	case ast.SUB:
		assembly = cw.getSubAssembly()
	case ast.NEG:
		assembly = cw.getNegAssembly()
	case ast.EQ, ast.GT, ast.LT:
		assembly = cw.getCompareAssembly(command)
	case ast.AND:
		assembly = cw.getAndAssembly()
	case ast.OR:
		assembly = cw.getOrAssembly()
	case ast.NOT:
		assembly = cw.getNotAssembly()
	}

	cw.assembly = append(cw.assembly, []byte(assembly)...)
}

func (cw *CodeWriter) WritePushPop(command ast.CommandSymbol, segment ast.SegmentSymbol, index int) {
	var assembly string
	switch command {
	case ast.PUSH:
		assembly = cw.getPushAssembly(segment, index)
	case ast.POP:
		assembly = cw.getPopAssembly(segment, index)
	}

	cw.assembly = append(cw.assembly, []byte(assembly)...)
}

func (cw *CodeWriter) Close() {
	ioutil.WriteFile(cw.fileName, cw.assembly, os.ModePerm)
}

func (cw *CodeWriter) getAddAssembly() string {
	assembly := ""
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "A=A-1" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "A=A-1" + "\r\n"
	assembly += "M=D+M" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M-1" + "\r\n"
	return assembly
}

func (cw *CodeWriter) getSubAssembly() string {
	assembly := ""
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "A=A-1" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "A=A-1" + "\r\n"
	assembly += "M=M-D" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M-1" + "\r\n"
	return assembly
}

func (cw *CodeWriter) getNegAssembly() string {
	assembly := ""
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "A=A-1" + "\r\n"
	assembly += "M=-M" + "\r\n"
	return assembly
}

func (cw *CodeWriter) getCompareAssembly(compareCommand ast.CommandSymbol) string {
	assembly := ""
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "A=A-1" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "A=A-1" + "\r\n"
	assembly += "D=M-D" + "\r\n"
	assembly += "@TRUE" + "\r\n"
	switch compareCommand {
	case ast.EQ:
		assembly += "D;JEQ" + "\r\n"
	case ast.GT:
		assembly += "D;JGT" + "\r\n"
	case ast.LT:
		assembly += "D;JLT" + "\r\n"
	}
	assembly += "M=0" + "\r\n"
	assembly += "@NEXT" + "\r\n"
	assembly += "0;JMP" + "\r\n"
	assembly += "(TRUE)" + "\r\n"
	assembly += "M=-1" + "\r\n"
	assembly += "@NEXT" + "\r\n"
	assembly += "0;JMP" + "\r\n"
	assembly += "(NEXT)" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M-1" + "\r\n"
	return assembly
}

func (cw *CodeWriter) getAndAssembly() string {
	assembly := ""
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "A=A-1" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "A=A-1" + "\r\n"
	assembly += "M=D&M" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M-1" + "\r\n"
	return assembly
}

func (cw *CodeWriter) getOrAssembly() string {
	assembly := ""
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "A=A-1" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "A=A-1" + "\r\n"
	assembly += "M=D|M" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M-1" + "\r\n"
	return assembly
}

func (cw *CodeWriter) getNotAssembly() string {
	assembly := ""
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "A=A-1" + "\r\n"
	assembly += "M=!M" + "\r\n"
	return assembly
}

func (cw *CodeWriter) getPushAssembly(segment ast.SegmentSymbol, index int) string {
	var assembly string
	switch segment {
	case ast.CONSTANT:
		assembly = cw.getPushConstantAssembly(index)
	}

	return assembly
}

func (cw *CodeWriter) getPushConstantAssembly(index int) string {
	assembly := ""
	assembly += "@" + strconv.Itoa(index) + "\r\n"
	assembly += "D=A" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "M=D" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M+1" + "\r\n"
	return assembly
}

func (cw *CodeWriter) getPopAssembly(segment ast.SegmentSymbol, index int) string {
	var assembly string
	switch segment {

	}

	return assembly
}
