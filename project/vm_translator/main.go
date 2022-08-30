package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"vm_translator/codewriter"
	"vm_translator/parser"
)

func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatalf("failed to get file info: %s", err)
	}

	return fileInfo.IsDir()
}

func readDir(dirPath string) []string {
	files := []string{}
	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("failed to read directory: %s", err)
	}

	for _, file := range fs {
		if filepath.Ext(file.Name()) == ".vm" {
			files = append(files, filepath.Join(dirPath, file.Name()))
		}
	}

	return files
}

func getAsmFilePath(path string) string {
	if isDirectory(path) {
		return filepath.Join(path, filepath.Base(path)+".asm")
	} else {
		return strings.TrimSuffix(path, filepath.Ext(path)) + ".asm"
	}
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("usage: vm_translator <.vm filepath or directory path>")
	}

	files := []string{os.Args[1]}
	if isDirectory(os.Args[1]) {
		files = readDir(os.Args[1])
	}

	asmFilePath := getAsmFilePath(os.Args[1])

	asmFile, err := os.Create(asmFilePath)
	if err != nil {
		log.Fatalf("failed to create %s", asmFilePath)
	}
	defer asmFile.Close()

	cw := codewriter.NewCodeWriter(asmFile)

	for _, file := range files {
		vmFile, err := os.Open(file)
		if err != nil {
			log.Fatalf("failed to open %s", file)
		}
		defer vmFile.Close()

		p := parser.NewParser(vmFile)
		cw.SetFileName(filepath.Base(file))

		writeLine := func() error {
			var err error
			if p.CommandType() == parser.C_ARITHMETIC {
				err = cw.WriteArithmetic(p.Arg1())
			} else if p.CommandType() == parser.C_PUSH || p.CommandType() == parser.C_POP {
				err = cw.WritePushPop(p.CommandType(), p.Arg1(), p.Arg2())
			} else if p.CommandType() == parser.C_LABEL {
				err = cw.WriteLabel(p.Arg1())
			} else if p.CommandType() == parser.C_GOTO {
				err = cw.WriteGoto(p.Arg1())
			} else if p.CommandType() == parser.C_IF {
				err = cw.WriteIf(p.Arg1())
			} else if p.CommandType() == parser.C_FUNCTION {
				err = cw.WriteFunction(p.Arg1(), p.Arg2())
			} else if p.CommandType() == parser.C_RETURN {
				err = cw.WriteReturn()
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
			continue
		} else {
			log.Fatalln(err)
		}
	}
}
