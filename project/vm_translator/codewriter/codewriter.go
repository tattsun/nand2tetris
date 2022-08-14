package codewriter

import (
	"io"
	"vm_translator/parser"
)

type CodeWriter struct {
}

func NewCodeWriter(w io.Writer) *CodeWriter {
	return &CodeWriter{}
}

func (w *CodeWriter) SetFileName(fileName string) {
	panic("not implemented")
}

func (w *CodeWriter) WriteArithmetic(command parser.CommandType) error {
	panic("not implemented")
}

func (w *CodeWriter) WritePushPop(command parser.CommandType, segment string, index int64) error {
	panic("not implemented")
}
