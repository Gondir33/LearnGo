package main

import (
	"database/sql"
	"encoding/json"

	"github.com/Masterminds/squirrel"
)

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	UserID int    `json:"user_id"`
}

func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER,
		comments TEXT
	)`)
	return err
}
func InsertUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = prepareQuery("insert", "users", user).(squirrel.InsertBuilder).RunWith(db).Exec()
	return err
}
func SelectUser(userID int) (User, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	rows, err := prepareQuery("select", "users", User{ID: userID}).(squirrel.SelectBuilder).RunWith(db).Query()
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	var name string
	var age int
	var comments []Comment
	var bytes []byte
	rows.Next()
	err = rows.Scan(&userID, &name, &age, &bytes)
	if err != nil {
		return User{}, err
	}
	err = json.Unmarshal(bytes, &comments)
	if err != nil {
		return User{}, err
	}
	return User{userID, name, age, comments}, nil
}
func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = prepareQuery("update", "users", user).(squirrel.UpdateBuilder).RunWith(db).Exec()
	return err
}

func DeleteUser(userID int) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = prepareQuery("delete", "users", User{ID: userID}).(squirrel.DeleteBuilder).RunWith(db).Exec()
	return err
}

func prepareQuery(operation string, table string, user User) interface{} {
	for i := 0; i < len(user.Comments); i++ {
		user.Comments[i].UserID = user.ID
	}
	comments, err := json.Marshal(user.Comments)
	if err != nil {
		return err
	}
	if operation == "insert" {
		return squirrel.Insert(table).Columns("name", "age", "comments").Values(user.Name, user.Age, string(comments))
	} else if operation == "select" {
		return squirrel.Select("*").From(table).Where("id", user.ID)
	} else if operation == "update" {
		return squirrel.Update(table).Set("name", user.Name).Set("age", user.Age).Set("comments", string(comments))
	} else if operation == "delete" {
		return squirrel.Delete(table).Where("id", user.ID)
	} else {
		return nil
	}
}

/*
func main() {

	err := CreateUserTable()
	if err != nil {
		fmt.Printf("create %v\n", err)
		return
	}
	err = InsertUser(User{
		Name: "Andrew",
		Age:  18,
		Comments: []Comment{{
			ID:   1,
			Text: "u are the best",
		}, {
			ID:   2,
			Text: "u are suck",
		}},
	})
	if err != nil {
		fmt.Printf("insert %v\n", err)
		return
	}
	user, err := SelectUser(1)
	if err != nil {
		fmt.Printf("select %v\n", err)
		return
	}
	fmt.Println(user)
	err = UpdateUser(User{
		ID:   1,
		Name: "Grigoriy",
		Age:  21,
		Comments: []Comment{{
			ID:   1,
			Text: "u are asdasdasd",
		}, {
			ID:   2,
			Text: "u are zxczxcxzczxc",
		}},
	})
	if err != nil {
		fmt.Printf("update %v\n", err)
		return
	}
	user, err = SelectUser(1)
	if err != nil {
		fmt.Printf("select2 %v\n", err)
		return
	}
	fmt.Println(user)
	var asd sqlite3.SQLiteConn
	fmt.Println(asd)
}
*/
