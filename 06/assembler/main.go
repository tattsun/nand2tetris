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

func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func Assemble(r io.Reader, w io.Writer) error {
	parser := NewParser(r)
	code := NewCode()
	romsymtable := NewSymbolTable()
	romsymtable.AddEntry("SP", 0x0)
	romsymtable.AddEntry("LCL", 0x1)
	romsymtable.AddEntry("ARG", 0x2)
	romsymtable.AddEntry("THIS", 0x3)
	romsymtable.AddEntry("THAT", 0x4)
	romsymtable.AddEntry("R0", 0x0)
	romsymtable.AddEntry("R1", 0x1)
	romsymtable.AddEntry("R2", 0x2)
	romsymtable.AddEntry("R3", 0x3)
	romsymtable.AddEntry("R4", 0x4)
	romsymtable.AddEntry("R5", 0x5)
	romsymtable.AddEntry("R6", 0x6)
	romsymtable.AddEntry("R7", 0x7)
	romsymtable.AddEntry("R8", 0x8)
	romsymtable.AddEntry("R9", 0x9)
	romsymtable.AddEntry("R10", 0xa)
	romsymtable.AddEntry("R11", 0xb)
	romsymtable.AddEntry("R12", 0xc)
	romsymtable.AddEntry("R13", 0xd)
	romsymtable.AddEntry("R14", 0xe)
	romsymtable.AddEntry("R15", 0xf)
	romsymtable.AddEntry("SCREEN", 0x4000)
	romsymtable.AddEntry("KBD", 0x6000)

	ramsymtable := NewSymbolTable()

	program := make([]*token, 0)

	memoryOffset := int64(0x0010)
	cnt := int64(0)

	fmt.Println("1 PATH ------------------------------")

	for parser.HasMoreCommands() {
		if err := parser.Advance(); err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		if parser.currentToken.CommandType == L_COMMAND {
			romsymtable.AddEntry(parser.currentToken.Symbol, cnt)
			continue
		} else {
			cnt++
		}

		program = append(program, parser.currentToken)

		fmt.Println(parser.currentToken)
	}

	fmt.Println(romsymtable)
	fmt.Println(ramsymtable)

	fmt.Println("2 PATH ------------------------------")

	for _, t := range program {
		if t.CommandType == A_COMMAND {
			if !isNumber(t.Symbol) {

				if !romsymtable.Contains(t.Symbol) && !ramsymtable.Contains(t.Symbol) {
					ramsymtable.AddEntry(t.Symbol, memoryOffset)
					memoryOffset++
				}
			}
		}

		if t.CommandType == A_COMMAND {
			if !isNumber(t.Symbol) {
				if !romsymtable.Contains(t.Symbol) && !ramsymtable.Contains(t.Symbol) {
					fmt.Printf("symbol not found: %s", t.Symbol)
					os.Exit(1)
				}

				var addr int64

				if romsymtable.Contains(t.Symbol) {
					addr = romsymtable.GetAddress(t.Symbol)
				}

				if ramsymtable.Contains(t.Symbol) {
					addr = ramsymtable.GetAddress(t.Symbol)
				}

				t.Symbol = strconv.FormatInt(addr, 10)
			}
		}

		fmt.Println(t)

		switch t.CommandType {
		case A_COMMAND:
			w.Write([]byte("0"))
			num, err := strconv.Atoi(t.Symbol)
			if err != nil {
				return fmt.Errorf("invalid number: symbol: %s", t.Symbol)
			}
			w.Write([]byte(fmt.Sprintf("%015b", num)))
			w.Write([]byte(fmt.Sprintln()))
		case C_COMMAND:
			w.Write([]byte("111"))
			w.Write([]byte(code.Comp(t.Comp)))
			w.Write([]byte(code.Dest(t.Dest)))
			w.Write([]byte(code.Jump(t.Jump)))
			w.Write([]byte(fmt.Sprintln()))
		default:
			return fmt.Errorf("unexpected command type: %s", t.CommandType)
		}
	}

	fmt.Println("--- DONE ---")

	// for parser.HasMoreCommands() {
	// 	if err := parser.Advance(); err == io.EOF {
	// 		return nil
	// 	} else if err != nil {
	// 		return err
	// 	}

	// 	fmt.Println(parser.currentToken)

	// 	switch parser.CommandType() {
	// 	case A_COMMAND:
	// 		w.Write([]byte("0"))
	// 		num, err := strconv.Atoi(parser.Symbol())
	// 		if err != nil {
	// 			return fmt.Errorf("invalid number: symbol: %s", parser.Symbol())
	// 		}
	// 		w.Write([]byte(fmt.Sprintf("%015b", num)))
	// 		w.Write([]byte(fmt.Sprintln()))
	// 	case C_COMMAND:
	// 		w.Write([]byte("111"))
	// 		w.Write([]byte(code.Comp(parser.Comp())))
	// 		w.Write([]byte(code.Dest(parser.Dest())))
	// 		w.Write([]byte(code.Jump(parser.Jump())))
	// 		w.Write([]byte(fmt.Sprintln()))
	// 	default:
	// 		return fmt.Errorf("unexpected command type: %s", parser.CommandType())
	// 	}
	// }

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
