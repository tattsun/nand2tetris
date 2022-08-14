package codewriter

import (
	"fmt"
	"io"
	"vm_translator/parser"
)

const (
	MemSP    = 0
	MemStack = 256
)

type CodeWriter struct {
	w io.Writer
}

func NewCodeWriter(w io.Writer) *CodeWriter {
	cw := &CodeWriter{w: w}

	cw.writeln("@256")
	cw.writeln("D=A")
	cw.writeln("@%d", MemSP)
	cw.writeln("M=D")

	return cw
}

// panic if write failed
func (w *CodeWriter) writeln(format string, args ...any) {
	_, err := w.w.Write([]byte(fmt.Sprintf(format+"\n", args...)))
	if err != nil {
		panic(err)
	}
}

func (w *CodeWriter) SetFileName(fileName string) {
	panic("not implemented")
}

func (w *CodeWriter) WriteArithmetic(command string) error {
	/*
		push constant 5

		// initialize SP
		@256
		D=A
		@0
		M=D

		// --- push constant 5
		@5
		D=A

		// set 5 to 256
		@0
		A=M
		M=D
		// incr stack pointer
		@0
		M=M+1

		// add 258
		@0
		// decr stack pointer
		M=M-1 // 257
		A=M
		D=M

		// desc stack pointer
		@0
		M=M-1 // 256
		A=M // a=256
		M=M+D

		//
	*/

	if command == "add" {
		w.writeln("@%d", MemSP)
		w.writeln("M=M-1") // decr stack
		w.writeln("A=M")
		w.writeln("D=M") // get stack data
		w.writeln("@%d", MemSP)
		w.writeln("M=M-1") // decr stack
		w.writeln("A=M")
		w.writeln("M=D+M") // set stack to addded data
		w.writeln("@%d", MemSP)
		w.writeln("M=M+1") // incr stack
	} else {
		panic("not implemented")
	}

	return nil
}

func (w *CodeWriter) WritePushPop(command parser.CommandType, segment string, index int64) error {
	if command == parser.C_PUSH {
		if segment == "constant" {
			w.writeln("@%d", index)
			w.writeln("D=A")
			w.writeln("@%d", MemSP)
			w.writeln("A=M")
			w.writeln("M=D")
			w.writeln("@%d", MemSP)
			w.writeln("M=M+1")
		} else {
			panic("not implemented")
		}
	} else if command == parser.C_POP {
		panic("not implemented")
	} else {
		return fmt.Errorf("expected PUSH/POP but got: %s", command)
	}

	return nil
}
