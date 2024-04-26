package main

import (
	"encoding/json"
	"os"
	"path"
)

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Text string `json:"text"`
}

func writeJSON(filePath string, data []User) error {
	err := os.MkdirAll(path.Dir(filePath), 777)
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, bytes, 777)
}

/*
func main() {
	user := []User{
		{
			Name: "Мария",
			Age:  22,
			Comments: []Comment{
				{"I ma"},
				{"stupid shit"},
			}},
		{
			Name: "Мария",
			Age:  22,
			Comments: []Comment{
				{"hekkow"},
				{"wrld"},
			}},
	}
	writeJSON("/home/gondir/go-kata/course1/13.popular_package/8.package_json/task1.13.8.3/tesst", user)
}
*/
