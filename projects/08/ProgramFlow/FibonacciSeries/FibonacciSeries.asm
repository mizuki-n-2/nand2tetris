@256
D=A
@SP
M=D
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
@1
D=A
@3
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
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
@0
D=A
@THAT
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
@1
D=A
@SP
A=M
M=D
@SP
M=M+1
@1
D=A
@THAT
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
@2
D=A
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
(MAIN_LOOP_START)
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
M=M-1
A=M
D=M
@COMPUTE_ELEMENT
D;JNE
@END_PROGRAM
0;JMP
(COMPUTE_ELEMENT)
@0
D=A
@THAT
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
@THAT
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
@2
D=A
@THAT
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
@1
D=A
@3
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
@1
D=A
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
@3
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
@1
D=A
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
@MAIN_LOOP_START
0;JMP
(END_PROGRAM)
