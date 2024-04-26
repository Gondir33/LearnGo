package main

import (
	"reflect"
)

type User struct {
	ID       int    `db:"id" db_ops:"create"`
	Username string `db:"username" db_ops:"create,update"`
	Email    string `db:"email" db_ops:"create,update"`
	Address  string `db:"address" db_ops:"update"`
	Status   int    `db:"status" db_ops:"create,update"`
	Delete   string `db:"delete" db_ops:"delete"`
}

func SimpleGetFieldsPointers(u interface{}) []interface{} {
	val := reflect.ValueOf(u)
	res := make([]interface{}, 0, 4)
	res = append(res, val.Elem().Field(1).Addr().Interface())
	res = append(res, val.Elem().Field(2).Addr().Interface())
	res = append(res, val.Elem().Field(3).Addr().Interface())
	res = append(res, val.Elem().Field(4).Addr().Interface())

	return res
}

/*
func main() {
	user := User{
		ID:       1,
		Username: "JohnDoe",
		Email:    "johndoe@example.com",
		Address:  "123 Main St",
		Status:   1,
		Delete:   "yes",
	}

	pointers := SimpleGetFieldsPointers(&user)
	fmt.Println(pointers)
}
*/
