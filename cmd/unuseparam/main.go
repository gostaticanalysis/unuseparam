package main

import (
	"github.com/gostaticanalysis/unuseparam"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(unuseparam.Analyzer) }
