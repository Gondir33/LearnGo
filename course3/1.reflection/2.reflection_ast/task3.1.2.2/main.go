package main

import (
	"fmt"
	"go/token"
	"os"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/decorator/resolver/goast"
	east "gitlab.com/ptflp/goast"
)

func main() {
	// читаем файл в []byte
	data, err := os.ReadFile("models.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	// создаем декоратор decorator.NewDecoratorWithImports
	dec := decorator.NewDecoratorWithImports(token.NewFileSet(), "models", goast.New())

	// создаем файл *dest.File с помощью decorator.Parse
	f, err := dec.Parse(data)
	if err != nil {
		panic(err)
	}

	tableNameBody := func(name string) *dst.BlockStmt {
		return &dst.BlockStmt{
			List: []dst.Stmt{
				&dst.ReturnStmt{
					Results: []dst.Expr{
						&dst.BasicLit{
							Kind:  token.STRING,
							Value: fmt.Sprintf(`"%s"`, name),
						},
					},
				},
			},
		}
	}

	// создаем метод TableName, который возвращает строку, используя east.Method
	east.AddMethod(f, "User", east.Method{
		Name:         "TableName",
		Receiver:     "u",
		ReceiverType: "User",
		Arguments:    nil,
		Return: []east.Param{{
			Name:      "",
			Type:      "string",
			IsPointer: false,
		}},
		Body: tableNameBody("users"),
	})

	// добавляем метод в структуру User и Address с помощью east.AddMethod
	east.AddMethod(f, "Address", east.Method{
		Name:         "TableName",
		Receiver:     "a",
		ReceiverType: "Address",
		Arguments:    nil,
		Return: []east.Param{{
			Name:      "",
			Type:      "string",
			IsPointer: false,
		}},
		Body: tableNameBody("address"),
	})

	// получаем структуры из файла с помощью east.GetStructs
	structs := east.GetStructs(f)

	// добавляем теги в структуры с помощью east.ModifyStructs, east.AddDBTags, east.AddDBTypeTags
	east.ModifyStructs(structs, func(s *east.Struct) error {
		err = east.AddDBTags(s)
		if err != nil {
			return err
		}

		return east.AddDBTypeTags(s)
	})

	// синхронизируем код с измененными структурами с помощью east.SyncStructs
	err = east.SyncStructs(f, structs)
	if err != nil {
		panic(err)
	}

	// сохранить результат в файл с помощью
	res, err := east.PrintAST(f)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("models_change.go")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(res)
	if err != nil {
		panic(err)
	}
}
