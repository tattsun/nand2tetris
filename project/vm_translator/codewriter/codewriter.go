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
	w     io.Writer
	state int
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

// Pops stack and set address to A
func (w *CodeWriter) popStack() {
	w.writeln("@%d", MemSP)
	w.writeln("M=M-1") // decr stack
	w.writeln("A=M")
}

// Pushes stack
func (w *CodeWriter) pushStack() {
	w.writeln("@%d", MemSP)
	w.writeln("M=M+1") // incr stack
}

func (w *CodeWriter) compare(cmp string) {
	w.popStack()
	w.writeln("D=M")
	w.popStack()
	w.writeln("D=M-D")

	w.writeln("@true_%d", w.state)
	w.writeln("D;%s", cmp)

	// if false
	w.writeln("@%d", MemSP)
	w.writeln("A=M")
	w.writeln("M=0")
	w.pushStack()
	w.writeln("@after_%d", w.state)
	w.writeln("0;JMP")

	// if true
	w.writeln("(true_%d)", w.state)
	w.writeln("@%d", MemSP)
	w.writeln("A=M")
	w.writeln("M=-1")
	w.pushStack()
	w.writeln("@after_%d", w.state)
	w.writeln("0;JMP")

	w.writeln("(after_%d)", w.state)

	w.state++
}

func (w *CodeWriter) WriteArithmetic(command string) error {
	w.writeln("// %s", command)
	if command == "add" {
		w.popStack()
		w.writeln("D=M") // get stack data
		w.popStack()
		w.writeln("M=D+M") // set stack to addded data
		w.pushStack()
	} else if command == "sub" {
		w.popStack()
		w.writeln("D=M") // get stack data
		w.popStack()
		w.writeln("M=M-D")
		w.pushStack()
	} else if command == "not" {
		w.popStack()
		w.writeln("M=!M")
		w.pushStack()
	} else if command == "and" {
		w.popStack()
		w.writeln("D=M") // get stack data
		w.popStack()
		w.writeln("M=D&M")
		w.pushStack()
	} else if command == "or" {
		w.popStack()
		w.writeln("D=M") // get stack data
		w.popStack()
		w.writeln("M=D|M")
		w.pushStack()
	} else if command == "neg" {
		w.popStack()
		w.writeln("M=!M")
		w.writeln("M=M+1")
		w.pushStack()
	} else if command == "eq" {
		w.compare("JEQ")
	} else if command == "lt" {
		w.compare("JLT")
	} else if command == "gt" {
		w.compare("JGT")
	} else {
		panic("not implemented")
	}

	return nil
}

func (w *CodeWriter) WritePushPop(command parser.CommandType, segment string, index int64) error {
	w.writeln("// %s:%d", segment, index)
	if command == parser.C_PUSH {
		if segment == "constant" {
			w.writeln("@%d", index)
			w.writeln("D=A")
			w.writeln("@%d", MemSP)
			w.writeln("A=M")
			w.writeln("M=D")
			w.pushStack()
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
