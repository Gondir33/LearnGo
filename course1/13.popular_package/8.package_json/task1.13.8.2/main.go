package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Text string `json:"text"`
}

func getUsersFromJSON(data []byte) ([]User, error) {
	Users := make([]User, 0, 10)
	err := json.Unmarshal(data, &Users)
	return Users, err
}

func main() {
	jsonData := []byte(`[
		{
			"name": "John",
			"age": 30,
			"comments": [
				{"text": "Great post!"},
				{"text": "I agree"}
			]
		},
		{
			"name": "Alice",
			"age": 25,
			"comments": [
				{"text": "Nice article"}
			]
		}
	]`)

	users, err := getUsersFromJSON(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, user := range users {
		fmt.Println("Name:", user.Name)
		fmt.Println("Age:", user.Age)
		fmt.Println("Comments:")
		for _, comment := range user.Comments {
			fmt.Println("- ", comment.Text)
		}
		fmt.Println()
	}
}
