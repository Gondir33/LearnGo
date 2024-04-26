package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/brianvoe/gofakeit"
)

// Определение структуры пользователя
type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

type Tabler interface {
	TableName() string
}

func (u User) TableName() string {
	return "users"
}

// Интерфейс для генерации SQL-запросов
type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

// Интерфейс для генерации фейковых данных
type FakeDataGenerator interface {
	GenerateFakeUser() User
}

type SQLiteGenerator struct {
}

func (sqlite SQLiteGenerator) CreateTableSQL(table Tabler) string {
	res := "CREATE TABLE IF NOT EXISTS " + table.TableName() + " ("
	t := reflect.TypeOf(table)

	res += t.Field(0).Tag.Get("db_field") + " " + t.Field(0).Tag.Get("db_type") + ", "
	res += t.Field(1).Tag.Get("db_field") + " " + t.Field(1).Tag.Get("db_type") + ", "
	res += t.Field(2).Tag.Get("db_field") + " " + t.Field(2).Tag.Get("db_type") + ", "
	res += t.Field(3).Tag.Get("db_field") + " " + t.Field(3).Tag.Get("db_type") + ")"
	return res
}

func (sqlite SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	user := reflect.ValueOf(model)
	t := reflect.TypeOf(model)
	types := []string{
		t.Field(0).Tag.Get("db_field"),
		t.Field(1).Tag.Get("db_field"),
		t.Field(2).Tag.Get("db_field"),
		t.Field(3).Tag.Get("db_field"),
	}

	return fmt.Sprintf(`INSERT INTO %s (%v) VALUES (%v, '%v', '%v', '%v')`,
		model.TableName(), strings.Join(types, ", "), user.Field(0).Interface(), user.Field(1).Interface(), user.Field(2).Interface(), user.Field(3).Interface())
}

type GoFakeitGenerator struct {
}

func (fake GoFakeitGenerator) GenerateFakeUser() User {
	user := User{
		ID:        int(gofakeit.Uint16()),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
	return user
}

/*
func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}

	user := User{}
	sql := sqlGenerator.CreateTableSQL(user)
	fmt.Println(sql)

	for i := 0; i < 34; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(fakeUser)
		fmt.Println(query)
	}
}
*/
