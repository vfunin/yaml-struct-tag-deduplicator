package main

import (
	"github.com/vfunin/yaml-struct-tag-deduplicator/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
