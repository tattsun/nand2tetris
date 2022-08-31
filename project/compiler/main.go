package main

import (
	"compiler/analyzer"
	"log"
	"os"
)

func main() {
	analyzer := analyzer.NewAnalyzer()
	if err := analyzer.Analyze(os.Args[1]); err != nil {
		log.Fatal(err)
	}
}
