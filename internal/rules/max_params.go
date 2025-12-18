package rules

import (
	"go/ast"
	"go/token"
)

const maxParams int = 5

func CheckMaxParams(
	f *ast.File,
	fset *token.FileSet,
	out []Issue,
) []Issue {
	ast.Inspect(f, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		params := fn.Type.Params
		if params == nil || len(params.List) < maxParams {
			return true
		}

		out = append(out, Issue{
			Pos:     fset.Position(fn.Pos()),
			Message: "functions must have up to 5 parameters",

			// NOTE: isso aqui ja caia no unsafe, pq cortar os params quebraria o
			// codigo
			Fix: func() {
				params.List = params.List[:maxParams]
			},
		})

		return true
	})

	return out
}
