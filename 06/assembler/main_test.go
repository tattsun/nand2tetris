package main

import (
	"bytes"
	"testing"
)

func TestAssemble(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{name: "A_COMMAND", args: args{`@2
	@3
	`}, wantW: `0000000000000010
0000000000000011
`, wantErr: false},
		{name: "A_COMMAND with comment", args: args{`@2 //hoge
			// fuga
 //piyo
@3
`}, wantW: `0000000000000010
0000000000000011
`, wantErr: false},
		{name: "C_COMMAND", args: args{`@2
M=D;JGT
@10
M=A
`}, wantW: `0000000000000010
1110001100001001
0000000000001010
1110110000001000
`, wantErr: false},
		{name: "Add", args: args{`// This file is part of www.nand2tetris.org
		// and the book "The Elements of Computing Systems"
		// by Nisan and Schocken, MIT Press.
		// File name: projects/06/add/Add.asm
		
		// Computes R0 = 2 + 3  (R0 refers to RAM[0])
		
		@2
		D=A
		@3
		D=D+A
		@0
		M=D
		`}, wantW: `0000000000000010
1110110000010000
0000000000000011
1110000010010000
0000000000000000
1110001100001000
`, wantErr: false},
		{name: "Max", args: args{`// This file is part of www.nand2tetris.org
		// and the book "The Elements of Computing Systems"
		// by Nisan and Schocken, MIT Press.
		// File name: projects/06/max/MaxL.asm
		
		// Symbol-less version of the Max.asm program.
		
		@0
		D=M
		@1
		D=D-M
		@10
		D;JGT
		@1
		D=M
		@12
		0;JMP
		@0
		D=M
		@2
		M=D
		@14
		0;JMP
		`}, wantW: `0000000000000000
1111110000010000
0000000000000001
1111010011010000
0000000000001010
1110001100000001
0000000000000001
1111110000010000
0000000000001100
1110101010000111
0000000000000000
1111110000010000
0000000000000010
1110001100001000
0000000000001110
1110101010000111
`, wantErr: false},
		{name: "Symbol simple", args: args{`   @R0
		D=M              // D = first number
		@R1
		D=D-M            // D = first number - second number
		@OUTPUT_FIRST
		D;JGT            // if D>0 (first is greater) goto output_first
		@R1
		D=M              // D = second number
		@OUTPUT_D
		0;JMP            // goto output_d
	 (OUTPUT_FIRST)
		@R0             
		D=M              // D = first number
	 (OUTPUT_D)
		@R2
		M=D              // M[2] = D (greatest number)
	 (INFINITE_LOOP)
		@INFINITE_LOOP
		0;JMP            // infinite loop
`}, wantW: `0000000000000000
1111110000010000
0000000000000001
1111010011010000
0000000000001010
1110001100000001
0000000000000001
1111110000010000
0000000000001100
1110101010000111
0000000000000000
1111110000010000
0000000000000010
1110001100001000
0000000000001110
1110101010000111
`, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			r := bytes.NewBufferString(tt.args.input)
			if err := Assemble(r, w); (err != nil) != tt.wantErr {
				t.Errorf("Assemble() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Assemble() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
