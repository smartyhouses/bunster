package parser_test

import "github.com/yassinebenaid/bunny/ast"

var caseTests = []testCase{
	{`case foo in bar) cmd; esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Word("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
					},
				},
			},
		},
	}},
	{`case foo
	in
		bar) cmd
	esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Word("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
					},
				},
			},
		},
	}},
	{`case foo
	in
		bar )
			cmd "arg" arg
			cmd arg 'arg'
			cmd arg 'arg'
	esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Word("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						},
					},
				},
			},
		},
	}},
	{`case foo
	in
		bar)
			cmd "arg" arg
			cmd arg 'arg'
			cmd arg 'arg';;
		baz)
			cmd "arg" arg
			cmd arg 'arg'
			cmd arg 'arg';&
		boo)
			cmd "arg" arg
			cmd arg 'arg'
			cmd arg 'arg';;&
		fab)
			cmd "arg" arg
			cmd arg 'arg'
			cmd arg 'arg'
	esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Word("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						},
						Terminator: ";;",
					},
					{
						Patterns: []ast.Expression{ast.Word("baz")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						},
						Terminator: ";&",
					},
					{
						Patterns: []ast.Expression{ast.Word("boo")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						},
						Terminator: ";;&",
					},
					{
						Patterns: []ast.Expression{ast.Word("fab")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						},
					},
				},
			},
		},
	}},
	{`case $foo in
		bar|'foo'|$var ) cmd "arg" arg;;
		bar    |   'foo'   |   $var   ) cmd "arg" arg;;
	esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Var("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{
							ast.Word("bar"),
							ast.Word("foo"),
							ast.Var("var"),
						},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						},
						Terminator: ";;",
					},
					{
						Patterns: []ast.Expression{
							ast.Word("bar"),
							ast.Word("foo"),
							ast.Var("var"),
						},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd"), Args: []ast.Expression{ast.Word("arg"), ast.Word("arg")}},
						},
						Terminator: ";;",
					},
				},
			},
		},
	}},
	{`case $foo in
		(bar) cmd;;
		(bar | 'foo') cmd;;
	esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Var("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
						Terminator: ";;",
					},
					{
						Patterns: []ast.Expression{
							ast.Word("bar"),
							ast.Word("foo"),
						},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
						Terminator: ";;",
					},
				},
			},
		},
	}},
	{`case $foo in (bar) cmd;; (bar | 'foo') cmd;; esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Var("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
						Terminator: ";;",
					},
					{
						Patterns: []ast.Expression{
							ast.Word("bar"),
							ast.Word("foo"),
						},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
						Terminator: ";;",
					},
				},
			},
		},
	}},
	{`case $foo in bar) cmd;; esac & case $foo in bar) cmd;; esac & cmd`, ast.Script{
		Statements: []ast.Statement{
			ast.BackgroundConstruction{
				Statement: ast.Case{
					Word: ast.Var("foo"),
					Cases: []ast.CaseItem{
						{
							Patterns: []ast.Expression{ast.Word("bar")},
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
							},
							Terminator: ";;",
						},
					},
				},
			},
			ast.BackgroundConstruction{
				Statement: ast.Case{
					Word: ast.Var("foo"),
					Cases: []ast.CaseItem{
						{
							Patterns: []ast.Expression{ast.Word("bar")},
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
							},
							Terminator: ";;",
						},
					},
				},
			},
			ast.Command{Name: ast.Word("cmd")},
		},
	}},
	{`case $foo in bar) cmd;; esac | case $foo in bar) cmd;; esac |& cmd`, ast.Script{
		Statements: []ast.Statement{
			ast.Pipeline{
				ast.PipelineCommand{
					Command: ast.Case{
						Word: ast.Var("foo"),
						Cases: []ast.CaseItem{
							{
								Patterns: []ast.Expression{ast.Word("bar")},
								Body: []ast.Statement{
									ast.Command{Name: ast.Word("cmd")},
								},
								Terminator: ";;",
							},
						},
					},
				},
				ast.PipelineCommand{
					Command: ast.Case{
						Word: ast.Var("foo"),
						Cases: []ast.CaseItem{
							{
								Patterns: []ast.Expression{ast.Word("bar")},
								Body: []ast.Statement{
									ast.Command{Name: ast.Word("cmd")},
								},
								Terminator: ";;",
							},
						},
					},
				},
				ast.PipelineCommand{
					Stderr:  true,
					Command: ast.Command{Name: ast.Word("cmd")},
				},
			},
		},
	}},
	{`case $foo in bar) cmd;; esac || case $foo in bar) cmd;; esac`, ast.Script{
		Statements: []ast.Statement{
			ast.BinaryConstruction{
				Left: ast.Case{
					Word: ast.Var("foo"),
					Cases: []ast.CaseItem{
						{
							Patterns: []ast.Expression{ast.Word("bar")},
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
							},
							Terminator: ";;",
						},
					},
				},
				Operator: "||",
				Right: ast.Case{
					Word: ast.Var("foo"),
					Cases: []ast.CaseItem{
						{
							Patterns: []ast.Expression{ast.Word("bar")},
							Body: []ast.Statement{
								ast.Command{Name: ast.Word("cmd")},
							},
							Terminator: ";;",
						},
					},
				},
			},
		},
	}},
	{`case $foo in
		bar)
			case $foo in
				bar)
					cmd;;
			esac;;
	esac`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Var("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Case{
								Word: ast.Var("foo"),
								Cases: []ast.CaseItem{
									{
										Patterns: []ast.Expression{ast.Word("bar")},
										Body: []ast.Statement{
											ast.Command{Name: ast.Word("cmd")},
										},
										Terminator: ";;",
									},
								},
							},
						},
						Terminator: ";;",
					},
				},
			},
		},
	}},

	{` case $foo in
		bar)
			cmd;;
	esac >output.txt <input.txt 2>error.txt >&3 \
	 	>>output.txt <<<input.txt 2>>error.txt &>all.txt &>>all.txt <&4 5<&6`, ast.Script{
		Statements: []ast.Statement{
			ast.Case{
				Word: ast.Var("foo"),
				Cases: []ast.CaseItem{
					{
						Patterns: []ast.Expression{ast.Word("bar")},
						Body: []ast.Statement{
							ast.Command{Name: ast.Word("cmd")},
						},
						Terminator: ";;",
					},
				},
				Redirections: []ast.Redirection{
					{Src: "1", Method: ">", Dst: ast.Word("output.txt")},
					{Src: "0", Method: "<", Dst: ast.Word("input.txt")},
					{Src: "2", Method: ">", Dst: ast.Word("error.txt")},
					{Src: "1", Method: ">&", Dst: ast.Word("3")},
					{Src: "1", Method: ">>", Dst: ast.Word("output.txt")},
					{Src: "0", Method: "<<<", Dst: ast.Word("input.txt")},
					{Src: "2", Method: ">>", Dst: ast.Word("error.txt")},
					{Method: "&>", Dst: ast.Word("all.txt")},
					{Method: "&>>", Dst: ast.Word("all.txt")},
					{Src: "0", Method: "<&", Dst: ast.Word("4")},
					{Src: "5", Method: "<&", Dst: ast.Word("6")},
				},
			},
		},
	}},
}
