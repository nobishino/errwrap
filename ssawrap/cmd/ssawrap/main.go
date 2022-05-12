package main

import (
	"ssawrap"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(ssawrap.Analyzer) }

