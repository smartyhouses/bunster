package parser_test

import "github.com/yassinebenaid/bunny/ast"

var arithmeticsTests = []testCase{
	{`$((1))`, ast.Script{
		Statements: []ast.Statement{
			ast.Command{Name: ast.Arithmetic{
				Expr: ast.Word("1"),
			}},
		},
	}},
}
