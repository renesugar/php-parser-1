package expr_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/name"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestInstanceOf(t *testing.T) {
	src := `<? $a instanceof Foo;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.InstanceOf{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Class: &name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestInstanceOfRelative(t *testing.T) {
	src := `<? $a instanceof namespace\Foo;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.InstanceOf{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Class: &name.Relative{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestInstanceOfFullyQualified(t *testing.T) {
	src := `<? $a instanceof \Foo;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.InstanceOf{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					Class: &name.FullyQualified{
						Parts: []node.Node{
							&name.NamePart{Value: "Foo"},
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}
