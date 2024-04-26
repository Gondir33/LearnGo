package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	fSet := token.NewFileSet()
	file, err := parser.ParseFile(fSet, "struct.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Inspect(file, func(n ast.Node) bool {
		if typeSpec, ok := n.(*ast.TypeSpec); ok {
			if _, ok := typeSpec.Type.(*ast.StructType); ok {
				if typeSpec.Name.Name == "MyStruct" {
					typeSpec.Name.Name = "User"
					printer.Fprint(os.Stdout, fSet, file)
					return false
				}
			}
		}
		return true
	})
}
