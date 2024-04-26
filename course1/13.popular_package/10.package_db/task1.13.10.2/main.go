package main

import (
	"database/sql"
	"errors"

	"github.com/Masterminds/squirrel"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func openDB() (*sql.DB, error) {
	return sql.Open("sqlite3", "users.db")
}

// Создание таблицы пользователей
func CreateUserTable() error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		email TEXT
	)`)
	return err
}

// Вставка пользователя в таблицу
func InsertUser(user User) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("insert", "users", user)
	if err != nil {
		return err
	}

	_, err = db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// Выборка пользователя из таблицы
func SelectUser(userID int) (User, error) {
	db, err := openDB()
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	query, args, err := PrepareQuery("select", "users", User{ID: userID})
	if err != nil {
		return User{}, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	rows.Next()
	var username string
	var email string
	err = rows.Scan(&userID, &username, &email)
	if err != nil {
		return User{}, err
	}
	return User{userID, username, email}, err
}

// Обновление информации о пользователе
func UpdateUser(user User) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("update", "users", user)
	if err != nil {
		return nil
	}
	_, err = db.Exec(query, args...)
	return err
}

// Удаление пользователя из таблицы
func DeleteUser(userID int) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("delete", "users", User{ID: userID})
	if err != nil {
		return err
	}
	_, err = db.Exec(query, args...)
	return err
}

// Функция для подготовки запроса
func PrepareQuery(operation string, table string, user User) (string, []interface{}, error) {
	var query string
	var args []interface{}
	var err error

	if operation == "insert" {
		query, args, err = squirrel.Insert(table).Columns("username", "email").Values(user.Username, user.Email).ToSql()
	} else if operation == "select" {
		query, args, err = squirrel.Select("*").From(table).Where("id", user.ID).ToSql()
	} else if operation == "update" {
		query, args, err = squirrel.Update(table).Set("username", user.Username).Set("email", user.Email).Where("id", user.ID).ToSql()
	} else if operation == "delete" {
		query, args, err = squirrel.Delete(table).Where("id", user.ID).ToSql()
	} else {
		err = errors.New("not found command")
	}
	return query, args, err
}

/*
func main() {
	err := CreateUserTable()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = InsertUser(User{
		Username: "Andrew",
		Email:    "ya@yandex.ru",
	})
	if err != nil {
		fmt.Println("insert")
		return
	}
	user, err := SelectUser(1)
	if err != nil {
		fmt.Println("select")
		return
	}
	fmt.Println(user)
	err = UpdateUser(User{
		ID:       1,
		Username: "KOL",
		Email:    "ya@yandex.ru",
	})
	if err != nil {
		return
	}
	var asd sqlite3.SQLiteConn
	fmt.Println(asd)
}
*/
