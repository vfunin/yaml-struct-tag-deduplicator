package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{ //nolint:exhaustivestruct
	Name: "goyamldeduplicator",
	Doc:  "Checks that yaml struct tags are not repeated",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	values := map[string]struct{}{}
	inspect := func(node ast.Node) bool {
		sType, ok := node.(*ast.StructType)
		if !ok {
			return true
		}

		for _, f := range sType.Fields.List {
			if f.Tag == nil {
				continue
			}

			v := f.Tag.Value
			if v == "" {
				continue
			}

			_, ok = values[v]
			if !ok {
				values[v] = struct{}{}

				continue
			}

			pass.Reportf(node.Pos(), "found double declaration: '%s'", v)

			return true
		}

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}

	return nil, nil
}
