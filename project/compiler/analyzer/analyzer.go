package analyzer

import (
	"compiler/tokenizer"
	"compiler/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func IsJackFile(path string) bool {
	return filepath.Ext(path) == ".jack"
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
			log.Printf("---------- START: %s", path)

			// Tokenize Test
			f, err := os.Open(path)
			if err != nil {
				return errors.Wrapf(err, "failed to open file: %s", path)
			}

			t := tokenizer.NewTokenizer(f)
			for t.HasMoreTokens() {
				t.Advance()

				switch t.TokenType() {
				case "KEYWORD":
					fmt.Printf("<keyword>%s</keyword>", t.KeyWord())
				case "SYMBOL":
					fmt.Printf("<symbol>%c</symbol>", t.Symbol())
				case "IDENTIFIER":
					fmt.Printf("<identifier>%s</identifier>", t.Identifier())
				case "INT_CONST":
					fmt.Printf("<integerConstant>%d</integerConstant>", t.IntVal())
				case "STRING_CONST":
					fmt.Printf("<stringConstant>%s</stringConstant>", t.StringVal())
				}

				fmt.Println()
			}

			log.Printf("---------- END: %s", path)
		}
	}

	return nil
}
