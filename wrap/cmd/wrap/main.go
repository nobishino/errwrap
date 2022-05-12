package main

import (
	"wrap"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(wrap.Analyzer) }

