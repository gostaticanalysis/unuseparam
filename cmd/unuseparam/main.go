package main

import (
	"unuseparam"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(unuseparam.Analyzer) }

