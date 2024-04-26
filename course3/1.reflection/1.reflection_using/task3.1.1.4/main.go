package main

import (
	"context"
	"database/sql"
	"fmt"
	"test/dao"
	"test/migrator"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int    `db:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db:"last_name" db_type:"VARCHAR(100)"`
	Username  string `db:"username" db_type:"VARCHAR(100)"`
	Email     string `db:"email" db_type:"VARCHAR(100)"`
	Address   string `db:"address" db_type:"VARCHAR(100)"`
	Status    int    `db:"status" db_type:"INT"`
	DeletedAt string `db:"deleted_at" db_type:"VARCHAR(100)"`
}

func (u User) TableName() string {
	return "users"
}

func main() {
	db, err := sql.Open("sqlite3", "mydao.db")
	if err != nil {
		panic(err)
	}

	dbx := sqlx.NewDb(db, "sqlite3")
	d := dao.NewDAO(dbx)

	var generator migrator.SQLiteGenerator
	m := migrator.NewMigrator(db, &generator)
	err = m.Migrate(&User{})
	if err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		user := GenerateFakeUser()
		err = d.Create(context.Background(), &user)
	}

	users := make([]User, 1)

	err = d.List(context.Background(), &users, &users[0], dao.Condition{
		LimitOffset: &dao.LimitOffset{
			Offset: 0,
			Limit:  3,
		},
		Equal: map[string]interface{}{
			"first_name": "Vilma",
		},
	})

	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Println(user)
	}
}

func GenerateFakeUser() User {
	return User{
		ID:        gofakeit.Number(1000, 9999),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
}
