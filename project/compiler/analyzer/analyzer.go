package analyzer

import (
	"compiler/utils"
	"log"
)

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

	for i, path := range paths {
		log.Printf("%d: %s", i, path)
	}

	return nil
}
