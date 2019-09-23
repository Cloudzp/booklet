package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	astFile, err := parser.ParseFile(token.NewFileSet(), "golang/src/go-parser/parser_test.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Println("astFile error:", err)
		return
	}
	// 文件名称
	fmt.Println(astFile.Name)

	for _, decl := range astFile.Decls {
		if generalDeclaration, ok := decl.(*ast.GenDecl); ok {
			//fmt.Println("===",generalDeclaration.Doc)
			//结构体遍历
			for _, astSpec := range generalDeclaration.Specs {
				if typeSpec, ok := astSpec.(*ast.TypeSpec); ok {

					fmt.Println(fmt.Sprintf("typeName: %+v", typeSpec.Name))
					switch expr := typeSpec.Type.(type) {
					case *ast.StructType:
						for _, field := range expr.Fields.List {
							expr := field.Type.(ast.Expr)
							if astTypeInterface, ok := expr.(*ast.InterfaceType); ok {
								fmt.Println("interface", field.Names, astTypeInterface)
							}

						}
					}

				}
			}
		}
	}
}
