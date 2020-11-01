package main

import (
	"github.com/rauljordan/static-analysis/writefile"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(writefile.Analyzer)
}
