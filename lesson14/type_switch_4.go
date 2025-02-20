package main

import (
	"fmt"
	"log"
	"strings"

	"go/ast"
	"go/parser"
)

// výraz, který se má naparsovat
const source = `
1 + 2 * 3 + x + y * z - 1
`

// nový datový typ implementující rozhraní ast.Visitor
type visitor int

// implementace (jediné) funkce předepsané v rozhraní ast.Visitor
func (v visitor) Visit(n ast.Node) ast.Visitor {
	// dosáhli jsme koncového uzlu?
	if n == nil {
		return nil
	}

	// tisk pozice a typu uzlu
	fmt.Printf("%3d\t", v)
	var s string

	// převod uzlu do tisknutelné podoby
	switch x := n.(type) {
	case *ast.BasicLit:
		s = x.Value
	case *ast.Ident:
		s = x.Name
	case *ast.UnaryExpr:
		s = x.Op.String()
	case *ast.BinaryExpr:
		s = x.Op.String()
	}

	// tisk obsahu uzlu
	indent := strings.Repeat("  ", int(v))
	if s != "" {
		fmt.Printf("%s%s\n", indent, s)
	} else {
		fmt.Printf("%s%T\n", indent, n)
	}
	return v + 1
}

func main() {
	// konstrukce parseru a parsing zdrojového kódu
	f, err := parser.ParseExpr(source)
	if err != nil {
		log.Fatal(err)
	}

	// hodnota typu visitor
	var v visitor

	// zahájení průchodu abstraktním syntaktickým stromem
	ast.Walk(v, f)
}
