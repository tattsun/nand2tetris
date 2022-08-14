package main

import (
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

	log.Println(p.Advance())
	log.Println(cw)
}
