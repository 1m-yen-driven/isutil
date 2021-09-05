package structs

import (
	"fmt"
	"go/types"
	"sort"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	ssaPkg "golang.org/x/tools/go/ssa"
)

var Analyzer = &analysis.Analyzer{
	Name: "structs",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

const Doc = "structs finds all structs in a package"

func isStruct(typ types.Type) bool {
	for {
		switch t := typ.(type) {
		case *types.Struct:
			return true
		case *types.Named:
			typ = t.Underlying()
		case *types.Pointer:
			typ = t.Elem()
		default:
			return false
		}
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	ssa := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)

	structs := make([]string, 0)
	for _, m := range ssa.Pkg.Members {
		if typ, ok := m.(*ssaPkg.Type); ok {
			if isStruct(typ.Type()) {
				structs = append(structs, m.String())
			}
		}
	}
	sort.Strings(structs)
	for _, s := range structs {
		fmt.Println(s)
	}

	return nil, nil
}
