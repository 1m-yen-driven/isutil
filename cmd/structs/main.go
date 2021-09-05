package main

import (
	"github.com/1m-yen-driven/isutil/structs"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(structs.Analyzer)
}
