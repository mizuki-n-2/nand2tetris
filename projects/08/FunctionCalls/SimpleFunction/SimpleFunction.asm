@256
D=A
@SP
M=D
(SimpleFunction.test)
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
@0
D=A
@LCL
A=M
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
@1
D=A
@LCL
A=M
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
@SP
A=M
A=A-1
D=M
A=A-1
M=D+M
@SP
M=M-1
@SP
A=M
A=A-1
M=!M
@0
D=A
@ARG
A=M
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
@SP
A=M
A=A-1
D=M
A=A-1
M=D+M
@SP
M=M-1
@1
D=A
@ARG
A=M
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
@SP
A=M
A=A-1
D=M
A=A-1
M=M-D
@SP
M=M-1
@LCL
D=M
@FRAME
M=D
@5
D=A
@FRAME
A=M-D
D=M
@RET
M=D
@0
D=A
@ARG
A=M
D=D+A
@temp
M=D
@SP
A=M
A=A-1
D=M
@temp
A=M
M=D
@SP
M=M-1
@ARG
D=M
@SP
M=D+1
@1
D=A
@FRAME
A=M-D
D=M
@THAT
M=D
@2
D=A
@FRAME
A=M-D
D=M
@THIS
M=D
@3
D=A
@FRAME
A=M-D
D=M
@ARG
M=D
@4
D=A
@FRAME
A=M-D
D=M
@LCL
M=D
@RET
A=M
0;JMP
