package main

import (
	"os"

	"github.com/gostaticanalysis/unuseparam"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	// on fix flag forcibly
	os.Args = append([]string{os.Args[0], "-fix"}, os.Args[1:]...)
	singlechecker.Main(unuseparam.Analyzer)
}
