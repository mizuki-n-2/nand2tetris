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

func New(fileName string) *CodeWriter {
	return &CodeWriter{fileName: fileName, assembly: []byte{}}
}

func (cw *CodeWriter) WriteInit() {
	assembly := getInitAssembly()
	cw.assembly = []byte(assembly)
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

func (cw *CodeWriter) SetVmClassName() {
	cw.vmClassName = strings.Split(path.Base(cw.fileName), ".")[0]
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

func (cw *CodeWriter) WriteLabel(label string) {
	assembly := cw.getLabelAssembly(label)
	cw.assembly = append(cw.assembly, []byte(assembly)...)
}

func (cw *CodeWriter) WriteGoto(label string) {
	assembly := cw.getGotoAssembly(label)
	cw.assembly = append(cw.assembly, []byte(assembly)...)
}

func (cw *CodeWriter) WriteIf(label string) {
	assembly := cw.getIfGotoAssembly(label)
	cw.assembly = append(cw.assembly, []byte(assembly)...)
}

func (cw *CodeWriter) WriteCall(functionName string, numArgs int) {
	assembly := cw.getCallAssembly(functionName, numArgs)
	cw.assembly = append(cw.assembly, []byte(assembly)...)
}

func (cw *CodeWriter) WriteReturn() {
	assembly := cw.getReturnAssembly()
	cw.assembly = append(cw.assembly, []byte(assembly)...)
}

func (cw *CodeWriter) WriteFunction(functionName string, numLocals int) {
	assembly := cw.getFunctionAssembly(functionName, numLocals)
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

func (cw *CodeWriter) getLabelAssembly(label string) string {
	assembly := ""
	assembly += "(" + label + ")" + "\r\n"
	return assembly
}

func (cw *CodeWriter) getGotoAssembly(label string) string {
	assembly := ""
	assembly += "@" + label + "\r\n"
	assembly += "0;JMP" + "\r\n"
	return assembly
}

func (cw *CodeWriter) getIfGotoAssembly(label string) string {
	assembly := ""
	assembly += "@SP" + "\r\n"
	assembly += "M=M-1" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@" + label + "\r\n"
	assembly += "D;JNE" + "\r\n"
	return assembly
}

func (cw *CodeWriter) getCallAssembly(functionName string, numArgs int) string {
	assembly := ""

	// ランダムな数字を生成
	randomLabelNumber := strconv.Itoa(rand.Int())

	returnAddressLabel := "RETURN_" + randomLabelNumber

	// push return-address
	assembly += "@" + returnAddressLabel + "\r\n"
	assembly += "D=A" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "M=D" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M+1" + "\r\n"
	// push LCL
	assembly += "@LCL" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "M=D" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M+1" + "\r\n"
	// push ARG
	assembly += "@ARG" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "M=D" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M+1" + "\r\n"
	// push THIS
	assembly += "@THIS" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "M=D" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M+1" + "\r\n"
	// push THAT
	assembly += "@THAT" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "M=D" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=M+1" + "\r\n"
	// ARG=SP-n-5 (n=numArgs)
	assembly += "@" + strconv.Itoa(numArgs+5) + "\r\n"
	assembly += "D=A" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "D=M-D" + "\r\n"
	assembly += "@ARG" + "\r\n"
	assembly += "M=D" + "\r\n"
	// LCL=SP
	assembly += "@SP" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@LCL" + "\r\n"
	assembly += "M=D" + "\r\n"
	// goto f (f=functionName)
	assembly += cw.getGotoAssembly(functionName)
	// (return-address)
	assembly += cw.getLabelAssembly(returnAddressLabel)

	return assembly
}

func (cw *CodeWriter) getReturnAssembly() string {
	assembly := ""
	// FRAME=LCL
	assembly += "@LCL" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@FRAME" + "\r\n"
	assembly += "M=D" + "\r\n"
	// RET=*(FRAME-5)
	assembly += "@5" + "\r\n"
	assembly += "D=A" + "\r\n"
	assembly += "@FRAME" + "\r\n"
	assembly += "A=M-D" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@RET" + "\r\n"
	assembly += "M=D" + "\r\n"
	// *ARG=pop()
	assembly += cw.getPopMemoryAccessAssembly(ast.ARGUMENT, 0)
	// SP=ARG+1
	assembly += "@ARG" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@SP" + "\r\n"
	assembly += "M=D+1" + "\r\n"
	// THAT=*(FRAME-1)
	assembly += "@1" + "\r\n"
	assembly += "D=A" + "\r\n"
	assembly += "@FRAME" + "\r\n"
	assembly += "A=M-D" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@THAT" + "\r\n"
	assembly += "M=D" + "\r\n"
	// THIS=*(FRAME-2)
	assembly += "@2" + "\r\n"
	assembly += "D=A" + "\r\n"
	assembly += "@FRAME" + "\r\n"
	assembly += "A=M-D" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@THIS" + "\r\n"
	assembly += "M=D" + "\r\n"
	// ARG=*(FRAME-3)
	assembly += "@3" + "\r\n"
	assembly += "D=A" + "\r\n"
	assembly += "@FRAME" + "\r\n"
	assembly += "A=M-D" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@ARG" + "\r\n"
	assembly += "M=D" + "\r\n"
	// LCL=*(FRAME-4)
	assembly += "@4" + "\r\n"
	assembly += "D=A" + "\r\n"
	assembly += "@FRAME" + "\r\n"
	assembly += "A=M-D" + "\r\n"
	assembly += "D=M" + "\r\n"
	assembly += "@LCL" + "\r\n"
	assembly += "M=D" + "\r\n"
	// goto RET
	assembly += "@RET" + "\r\n"
	assembly += "A=M" + "\r\n"
	assembly += "0;JMP" + "\r\n"
	return assembly
}

func (cw *CodeWriter) getFunctionAssembly(functionName string, numLocals int) string {
	assembly := ""
	assembly += cw.getLabelAssembly(functionName)

	for i := 0; i < numLocals; i++ {
		assembly += cw.getPushConstantAssembly(0)
	}

	return assembly
}
