package migrator

import (
	"petstore/internal/infrastructure/db/tabler"
	"reflect"
)

type SQLGenerator interface {
	CreateTableSQL(table tabler.Tabler) string
}

type SQLiteGenerator struct{}

func (sg *SQLiteGenerator) CreateTableSQL(table tabler.Tabler) string {
	res := "CREATE TABLE IF NOT EXISTS " + table.TableName() + " ("
	t := reflect.TypeOf(table).Elem()

	i := 0
	for ; i < t.NumField()-1; i++ {
		res += t.Field(i).Tag.Get("db") + " " + t.Field(i).Tag.Get("db_type") + ", "
	}
	res += t.Field(i).Tag.Get("db") + " " + t.Field(i).Tag.Get("db_type") + ")"
	return res
}
