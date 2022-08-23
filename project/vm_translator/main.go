package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"vm_translator/codewriter"
	"vm_translator/parser"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("usage: vm_translator <.vm filepath or directory path>")
	}

	// TODO: Handle multiple files
	vmFilePath := os.Args[1]
	fileNameBase := strings.TrimSuffix(vmFilePath, filepath.Ext(vmFilePath))
	asmFilePath := fileNameBase + ".asm"

	vmFile, err := os.Open(vmFilePath)
	if err != nil {
		log.Fatalf("failed to open %s", vmFilePath)
	}
	defer vmFile.Close()

	asmFile, err := os.Create(asmFilePath)
	if err != nil {
		log.Fatalf("failed to create %s", asmFilePath)
	}
	defer asmFile.Close()

	// Do translate
	p := parser.NewParser(vmFile)
	cw := codewriter.NewCodeWriter(asmFile)
	cw.SetFileName(filepath.Base(vmFilePath))

	writeLine := func() error {
		var err error
		if p.CommandType() == parser.C_ARITHMETIC {
			err = cw.WriteArithmetic(p.Arg1())
		} else if p.CommandType() == parser.C_PUSH || p.CommandType() == parser.C_POP {
			err = cw.WritePushPop(p.CommandType(), p.Arg1(), p.Arg2())
		} else {
			return fmt.Errorf("unexpected command: %s", p.CommandType())
		}

		if err != nil {
			return fmt.Errorf("failed to write command: %s, %s, %d, %+v", p.CommandType(), p.Arg1(), p.Arg2(), err)
		}
		return nil
	}

	err = p.Advance()
	if err != nil {
		log.Fatalln(err)
	}

	for err == nil {
		err = writeLine()
		if err != nil {
			continue
		}

		err = p.Advance()
	}

	if err == nil || err == io.EOF {
		os.Exit(0)
	} else {
		log.Fatalln(err)
	}
}
