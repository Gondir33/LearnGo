package main

import (
	"database/sql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getDB() (*sql.DB, error) {
	return sql.Open("sqlite3", "users.db")
}

func CreateUserTable() error {
	db, err := getDB()
	if err != nil {
		return nil
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	)`)
	return err
}

func InsertUser(user User) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", user.Name, user.Age)
	return err
}

func SelectUser(id int) (User, error) {
	db, err := getDB()
	if err != nil {
		return User{}, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	var name string
	var age int
	rows.Next()
	err = rows.Scan(&id, &name, &age)
	return User{id, name, age}, err
}

func UpdateUser(user User) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE users SET age = ?, name = ? WHERE id = ?", user.Age, user.Name, user.ID)
	return err
}

func DeleteUser(id int) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
