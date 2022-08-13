package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func Assemble(r io.Reader, w io.Writer) error {
	parser := NewParser(r)
	code := NewCode()

	for parser.HasMoreCommands() {
		if err := parser.Advance(); err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		switch parser.CommandType() {
		case A_COMMAND:
			w.Write([]byte("0"))
			num, err := strconv.Atoi(parser.Symbol())
			if err != nil {
				return fmt.Errorf("invalid number: symbol: %s", parser.Symbol())
			}
			w.Write([]byte(fmt.Sprintf("%015b", num)))
			w.Write([]byte(fmt.Sprintln()))
		case C_COMMAND:
			w.Write([]byte("111"))
			w.Write([]byte(code.Comp(parser.Comp())))
			w.Write([]byte(code.Dest(parser.Dest())))
			w.Write([]byte(code.Jump(parser.Jump())))
			w.Write([]byte(fmt.Sprintln()))
		default:
			return fmt.Errorf("unexpected command type: %s", parser.CommandType())
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("./assembler <input file>")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("failed to open file: %+v", err)
		os.Exit(1)
	}

	hackFilePath := strings.TrimSuffix(filepath.Base(os.Args[1]), filepath.Ext(os.Args[1]))
	w, err := os.Create(path.Join(filepath.Dir(os.Args[1]), hackFilePath+".hack"))
	if err != nil {
		fmt.Printf("failed to create file: %+v", err)
		os.Exit(1)
	}

	if err := Assemble(f, w); err != nil {
		fmt.Printf("failed to assemble: %+v", err)
		os.Exit(1)
	}
}
