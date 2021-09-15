package unuseparam

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "unuseparam finds a unused parameter but its name is not _"

var Analyzer = &analysis.Analyzer{
	Name: "unuseparam",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}

	ps := make(map[types.Object]*ast.Ident)
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			params(pass, n.Type, ps)
		case *ast.FuncLit:
			params(pass, n.Type, ps)
		}
	})

	unused(pass, ps)

	return nil, nil
}

func params(pass *analysis.Pass, typ *ast.FuncType, ps map[types.Object]*ast.Ident) {
	for _, field := range typ.Params.List {
		for _, n := range field.Names {
			obj := pass.TypesInfo.Defs[n]
			if obj != nil && n.Name != "_" {
				ps[obj] = n
			}
		}
	}
}

func unused(pass *analysis.Pass, ps map[types.Object]*ast.Ident) {
	for _, obj := range pass.TypesInfo.Uses {
		if ps[obj] != nil {
			ps[obj] = nil
		}
	}

	for _, id := range ps {
		if id == nil {
			continue
		}

		fix := analysis.SuggestedFix{
			Message: fmt.Sprintf(`%s is unused parameter, it can be changed to "_"`, id.Name),
			TextEdits: []analysis.TextEdit{{
				Pos:     id.Pos(),
				End:     id.End(),
				NewText: []byte("_"),
			}},
		}

		pass.Report(analysis.Diagnostic{
			Pos:            id.Pos(),
			End:            id.End(),
			Message:        fmt.Sprintf("%s is unused parameter", id.Name),
			SuggestedFixes: []analysis.SuggestedFix{fix},
		})
	}
}
