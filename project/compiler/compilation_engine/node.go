package compilationengine

import (
	"fmt"
	"io"
)

func writeIndent(w io.Writer, indent int) {
	for i := 0; i < indent; i++ {
		fmt.Fprint(w, "  ")
	}
}

type Node interface {
	XML(w io.Writer, indent int) string
}

type Class struct {
	ClassName *Identifier
}

func (c *Class) XML(w io.Writer, indent int) {
	writeIndent(w, indent)
	fmt.Fprintln(w, "<class>")

	c.ClassName.XML(w, indent+1)

	writeIndent(w, indent)
	fmt.Fprintln(w, "</class>")
}

type Identifier struct {
	Value string
}

func (i *Identifier) XML(w io.Writer, indent int) {
	writeIndent(w, indent)
	fmt.Fprintf(w, "<identifier> %s </identifier>\n", i.Value)
}
