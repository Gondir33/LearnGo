package main

import (
	"encoding/json"
)

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Text string `json:"text"`
}

func getJSON(data []User) (string, error) {
	bytes, err := json.Marshal(data)
	return string(bytes), err
}

/*

func main() {
	user := []User{
		{
			Name: "Мария",
			Age:  22,
			Comments: []Comment{
				Comment{"I ma"},
				Comment{"stupid shit"},
			}},
		{
			Name: "Мария",
			Age:  22,
			Comments: []Comment{
				Comment{"hekkow"},
				Comment{"wrld"},
			}},
	}
	fmt.Println(getJSON(user))
}
*/
