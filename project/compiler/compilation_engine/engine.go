package compilationengine

import "io"

type CompilationEngine struct {
	r io.Reader
	w io.Writer
}

func NewCompilationEngine(r io.Reader, w io.Writer) *CompilationEngine {
	return &CompilationEngine{
		r: r,
		w: w,
	}
}

func (e *CompilationEngine) CompileClass() {
	panic("not implemented")
}

func (e *CompilationEngine) CompileClassVarDec() {
	panic("not implemented")
}

func (e *CompilationEngine) CompileSubroutine() {
	panic("not implemented")
}

func (e *CompilationEngine) CompileParameterList() {
	panic("not implemented")
}

func (e *CompilationEngine) CompileVarDec() {
	panic("not implemented")
}

func (e *CompilationEngine) CompileStatements() {
	panic("not implemented")
}

func (e *CompilationEngine) CompileDo() {
	panic("not implemented")
}

func (e *CompilationEngine) CompileLet() {
	panic("not implemented")

}

func (e *CompilationEngine) CompileWhile() {
	panic("not implemented")
}

func (e *CompilationEngine) CompileReturn() {
	panic("not implemented")
}

func (e *CompilationEngine) CompileIf() {
	panic("not implemented")
}

func (e *CompilationEngine) CompileExpression() {
	panic("not implemented")
}

func (e *CompilationEngine) CompileTerm() {
	panic("not implemented")
}

func (e *CompilationEngine) CompileExpressionList() {
	panic("not implemented")
}
