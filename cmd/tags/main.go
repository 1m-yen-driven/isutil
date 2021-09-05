package main

import (
	"github.com/1m-yen-driven/isutil/tags"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(tags.Analyzer)
}
