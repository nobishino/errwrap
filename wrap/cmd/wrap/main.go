package main

import (
	"github.com/nobishino/errwrap/wrap"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(wrap.Analyzer) }
