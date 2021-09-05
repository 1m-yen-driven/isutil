package tags

import (
	"fmt"
	"go/types"
	"reflect"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

var Analyzer = &analysis.Analyzer{
	Name: "tags",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

const Doc = "tags finds all tags of a struct"

var StructPath string
var Key string

func init() {
	Analyzer.Flags.StringVar(&StructPath, "struct", "", "absolute path to the struct (e.g. example.com/pkg/name.StructName)")
	Analyzer.Flags.StringVar(&Key, "key", "", "key of tags to find")
}

func extractStruct(typ types.Type) (*types.Struct, error) {
	original := typ
	for {
		switch t := typ.(type) {
		case *types.Struct:
			return t, nil
		case *types.Named:
			typ = t.Underlying()
		case *types.Pointer:
			typ = t.Elem()
		default:
			return nil, fmt.Errorf("this type is not a struct: %v", original)
		}
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	if StructPath == "" {
		return nil, fmt.Errorf("-struct option is required")
	}
	if Key == "" {
		return nil, fmt.Errorf("-key option is required")
	}
	ssa := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	name := strings.TrimPrefix(StructPath, ssa.Pkg.Pkg.Path()+".")
	if name == StructPath || strings.Contains(name, "/") {
		// recursive analysis may encounter other packages
		return nil, nil
	}
	m, ok := ssa.Pkg.Members[name]
	if !ok {
		return nil, fmt.Errorf("struct not found in package %s", ssa.Pkg.Pkg.Path())
	}
	Struct, err := extractStruct(m.Type())
	if err != nil {
		return nil, err
	}
	for i := 0; i < Struct.NumFields(); i++ {
		tag := Struct.Tag(i)
		stag := reflect.StructTag(tag)
		if value, ok := stag.Lookup(Key); ok {
			fmt.Println(value)
		}
	}

	return nil, nil
}
