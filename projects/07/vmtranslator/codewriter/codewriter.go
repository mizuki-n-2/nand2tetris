package codewriter

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"vmtranslator/ast"
)

type CodeWriter struct {
	fileName    string
	assembly    []byte
	vmClassName string
}

func New() *CodeWriter {
	assembly := getInitAssembly()
	return &CodeWriter{assembly: []byte(assembly)}
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
	cw.vmClassName = strings.Split(path.Base(fileName), ".")[0]
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
	assembly += "M=M-1" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M-1" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "D=M-D" + "\r\n"

	// ランダムな数字を生成
	randomLabelNumber := strconv.Itoa(rand.Int())

	assembly += "@TRUE_" + randomLabelNumber + "\r\n"

	switch compareCommand {
	case ast.EQ:
		assembly += "D;JEQ" + "\r\n"
	case ast.GT:
		assembly += "D;JGT" + "\r\n"
	case ast.LT:
		assembly += "D;JLT" + "\r\n"
	}
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "M=0" + "\r\n"
	assembly += "@NEXT_" + randomLabelNumber + "\r\n"
	assembly += "0;JMP" + "\r\n"
	assembly += "(TRUE_" + randomLabelNumber + ")" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "M=-1" + "\r\n"
	assembly += "@NEXT_" + randomLabelNumber + "\r\n"
	assembly += "0;JMP" + "\r\n"
	assembly += "(NEXT_" + randomLabelNumber + ")" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M+1" + "\r\n"
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
	case ast.LOCAL, ast.ARGUMENT, ast.THIS, ast.THAT, ast.TEMP, ast.POINTER:
		assembly = cw.getPushMemoryAccessAssembly(segment, index)
	case ast.STATIC:
		assembly = cw.getPushStaticAssembly(index)
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

var TEMP_BASE_ADDRESS = 5
var POINTER_BASE_ADDRESS = 3

func (cw *CodeWriter) getPushMemoryAccessAssembly(segment ast.SegmentSymbol, index int) string {
	assembly := ""
	assembly += "@" + strconv.Itoa(index) + "\r\n"
	assembly += "D=A" + "\r\n"

	switch segment {
	case ast.LOCAL:
		assembly += "@LCL" + "\r\n"
		assembly += "A=M" + "\r\n"
	case ast.ARGUMENT:
		assembly += "@ARG" + "\r\n"
		assembly += "A=M" + "\r\n"
	case ast.THIS:
		assembly += "@THIS" + "\r\n"
		assembly += "A=M" + "\r\n"
	case ast.THAT:
		assembly += "@THAT" + "\r\n"
		assembly += "A=M" + "\r\n"
	case ast.TEMP:
		assembly += "@" + strconv.Itoa(TEMP_BASE_ADDRESS) + "\r\n"
	case ast.POINTER:
		assembly += "@" + strconv.Itoa(POINTER_BASE_ADDRESS) + "\r\n"
	}
	assembly += "A=D+A" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "M=D" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M+1" + "\r\n"
	return assembly
}

func (cw *CodeWriter) getPushStaticAssembly(index int) string {
	assembly := ""
	assembly += fmt.Sprintf("@%s.%d", cw.vmClassName, index) + "\r\n"
	assembly += "D=M" + "\r\n"
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
	case ast.LOCAL, ast.ARGUMENT, ast.THIS, ast.THAT, ast.TEMP, ast.POINTER:
		assembly = cw.getPopMemoryAccessAssembly(segment, index)
	case ast.STATIC:
		assembly = cw.getPopStaticAssembly(index)
	}

	return assembly
}

func (cw *CodeWriter) getPopMemoryAccessAssembly(segment ast.SegmentSymbol, index int) string {
	assembly := ""
	assembly += "@" + strconv.Itoa(index) + "\r\n"
	assembly += "D=A" + "\r\n"

	switch segment {
	case ast.LOCAL:
		assembly += "@LCL" + "\r\n"
		assembly += "A=M" + "\r\n"
	case ast.ARGUMENT:
		assembly += "@ARG" + "\r\n"
		assembly += "A=M" + "\r\n"
	case ast.THIS:
		assembly += "@THIS" + "\r\n"
		assembly += "A=M" + "\r\n"
	case ast.THAT:
		assembly += "@THAT" + "\r\n"
		assembly += "A=M" + "\r\n"
	case ast.TEMP:
		assembly += "@" + strconv.Itoa(TEMP_BASE_ADDRESS) + "\r\n"
	case ast.POINTER:
		assembly += "@" + strconv.Itoa(POINTER_BASE_ADDRESS) + "\r\n"
	}
	assembly += "D=D+A" + "\r\n"
	assembly += "@temp" + "\r\n"
	assembly += "M=D" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "A=A-1" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@temp" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "M=D" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M-1" + "\r\n"
	return assembly
}

func (cw *CodeWriter) getPopStaticAssembly(index int) string {
	assembly := ""
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "A=A-1" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += fmt.Sprintf("@%s.%d", cw.vmClassName, index) + "\r\n"
	assembly += "M=D" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M-1" + "\r\n"
	return assembly
}
