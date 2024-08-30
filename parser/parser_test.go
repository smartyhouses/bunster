package parser_test

import (
	"reflect"
	"testing"

	"github.com/yassinebenaid/godump"
	"github.com/yassinebenaid/nbs/ast"
	"github.com/yassinebenaid/nbs/lexer"
	"github.com/yassinebenaid/nbs/parser"
)

var dump = (&godump.Dumper{
	Theme:                   godump.DefaultTheme,
	ShowPrimitiveNamedTypes: true,
}).Sprintln

type testCase struct {
	input    string
	expected ast.Script
}

var testCases = []struct {
	label string
	cases []testCase
}{
	{"Simle Command calls", []testCase{
		{`git`, ast.Script{
			Statements: []ast.Node{
				ast.Command{Name: ast.Word("git")},
			},
		}},
		{`foo bar baz`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("foo"),
					Args: []ast.Node{
						ast.Word("bar"),
						ast.Word("baz"),
					},
				},
			},
		}},
		{`foo $bar $FOO_BAR_1234567890`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("foo"),
					Args: []ast.Node{
						ast.SimpleExpansion("bar"),
						ast.SimpleExpansion("FOO_BAR_1234567890"),
					},
				},
			},
		}},
		{`/usr/bin/foo bar baz`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("/usr/bin/foo"),
					Args: []ast.Node{
						ast.Word("bar"),
						ast.Word("baz"),
					},
				},
			},
		}},
		{`/usr/bin/foo-bar baz`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("/usr/bin/foo-bar"),
					Args: []ast.Node{
						ast.Word("baz"),
					},
				},
			},
		}},
	}},

	{"Strings", []testCase{
		{`cmd 'hello world'`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("cmd"),
					Args: []ast.Node{
						ast.Word("hello world"),
					},
				},
			},
		}},
		{`cmd 'if then else elif fi for in do done while until case esac function select trap return exit break continue declare local export readonly unset'`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("cmd"),
					Args: []ast.Node{
						ast.Word("if then else elif fi for in do done while until case esac function select trap return exit break continue declare local export readonly unset"),
					},
				},
			},
		}},
		{`cmd '+ - * / % %% = += -= *= /= == != < <= > >= =~ && || |'`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Word("cmd"),
					Args: []ast.Node{
						ast.Word("+ - * / % %% = += -= *= /= == != < <= > >= =~ && || |"),
					},
				},
			},
		}},
	}},

	{"Concatination", []testCase{
		{`/usr/bin/$BINARY_NAME --path=/home/$USER/dir --option -f --do=something $HOME$DIR_NAME$PKG_NAME/foo`, ast.Script{
			Statements: []ast.Node{
				ast.Command{
					Name: ast.Concatination{
						Nodes: []ast.Node{
							ast.Word("/usr/bin/"),
							ast.SimpleExpansion("BINARY_NAME"),
						},
					},
					Args: []ast.Node{
						ast.Concatination{
							Nodes: []ast.Node{
								ast.Word("--path=/home/"),
								ast.SimpleExpansion("USER"),
								ast.Word("/dir"),
							},
						},
						ast.Word("--option"),
						ast.Word("-f"),
						ast.Word("--do=something"),
						ast.Concatination{
							Nodes: []ast.Node{
								ast.SimpleExpansion("HOME"),
								ast.SimpleExpansion("DIR_NAME"),
								ast.SimpleExpansion("PKG_NAME"),
								ast.Word("/foo"),
							},
						},
					},
				},
			},
		}},
	}},
}

func TestParser(t *testing.T) {
	for _, group := range testCases {
		for i, tc := range group.cases {
			p := parser.New(
				lexer.New([]byte(tc.input)),
			)

			script := p.ParseScript()

			if !reflect.DeepEqual(script, tc.expected) {
				t.Fatalf("\nGroup: %sCase: %s\nwant:\n%s\ngot:\n%s", dump(group.label), dump(i), dump(tc.expected), dump(script))
			}
		}
	}
}
