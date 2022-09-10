package analyzer

import (
	compilationengine "compiler/compilation_engine"
	"compiler/tokenizer"
	"compiler/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func IsJackFile(path string) bool {
	return filepath.Ext(path) == ".jack"
}

func GetTokenFilePath(path string) string {
	dir := filepath.Dir(path)
	fileName := filepath.Base(path)
	fileBase := strings.TrimSuffix(fileName, ".jack")
	return filepath.Join(dir, fileBase+"T.xml")
}

type Analyzer struct {
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) Analyze(path string) error {
	paths, err := utils.GetFilePathsInDir(path)
	if err != nil {
		return err
	}

	for _, path := range paths {
		if IsJackFile(path) {
			f, err := os.Open(path)
			if err != nil {
				return errors.Wrapf(err, "failed to open file: %s", path)
			}
			defer f.Close()

			tf, err := os.Create(GetTokenFilePath(path))
			if err != nil {
				return errors.Wrapf(err, "failed to open file: %s", path)
			}
			defer tf.Close()

			compilationEngine := compilationengine.NewCompilationEngine(f, tf)
			class := compilationEngine.CompileClass()
			class.XML(tf, 0)
		}
	}

	return nil
}

func (a *Analyzer) Tokenize(path string) error {
	paths, err := utils.GetFilePathsInDir(path)
	if err != nil {
		return err
	}

	for _, path := range paths {
		if IsJackFile(path) {
			// Tokenize Test
			f, err := os.Open(path)
			if err != nil {
				return errors.Wrapf(err, "failed to open file: %s", path)
			}
			defer f.Close()

			// writer
			tf, err := os.Create(GetTokenFilePath(path))
			if err != nil {
				return errors.Wrapf(err, "failed to open file: %s", path)
			}
			defer tf.Close()

			fmt.Fprintln(tf, "<tokens>")

			t := tokenizer.NewTokenizer(f)
			for t.HasMoreTokens() {
				t.Advance()

				switch t.TokenType() {
				case "KEYWORD":
					fmt.Fprintf(tf, "<keyword> %s </keyword>", t.KeyWord())
				case "SYMBOL":
					fmt.Fprintf(tf, "<symbol> %s </symbol>", t.Symbol())
				case "IDENTIFIER":
					fmt.Fprintf(tf, "<identifier> %s </identifier>", t.Identifier())
				case "INT_CONST":
					fmt.Fprintf(tf, "<integerConstant> %d </integerConstant>", t.IntVal())
				case "STRING_CONST":
					fmt.Fprintf(tf, "<stringConstant> %s </stringConstant>", t.StringVal())
				}

				fmt.Fprintln(tf)
			}

			fmt.Fprintln(tf, "</tokens>")
		}
	}

	return nil
}
