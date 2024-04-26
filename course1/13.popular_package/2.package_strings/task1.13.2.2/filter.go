package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name     string
	Comments []Comment
}

type Comment struct {
	Message string
}

func main() {
	users := []User{
		{
			Name: "Betty",
			Comments: []Comment{
				{Message: "good Comment 1"},
				{Message: "BaD CoMmEnT 2"},
				{Message: "Bad Comment 3"},
				{Message: "Use camelCase please"},
			},
		},
		{
			Name: "Jhon",
			Comments: []Comment{
				{Message: "Good Comment 1"},
				{Message: "Good Comment 2"},
				{Message: "Good Comment 3"},
				{Message: "Bad Comments 4"},
			},
		},
	}

	users = FilterComments(users)
	fmt.Println(users)
}

func FilterComments(users []User) []User {
	for i := 0; i < len(users); i++ {
		for j := 0; j < len(users[i].Comments); j++ {
			if IsBadComment(users[i].Comments[j].Message) == true {
				users[i].Comments = append(users[i].Comments[:j], users[i].Comments[j+1:]...)
				j--
			}
		}
	}
	return users
}

func IsBadComment(comment string) bool {
	return strings.Contains(strings.ToLower(comment), "bad")
}

func GetBadComments(users []User) []Comment {
	comments := make([]Comment, 0)
	for i := 0; i < len(users); i++ {
		for j := 0; j < len(users[i].Comments); j++ {
			if IsBadComment(users[i].Comments[j].Message) == true {
				comments = append(comments, users[i].Comments[j])
			}
		}
	}
	return comments
}

func GetGoodComments(users []User) []Comment {
	comments := make([]Comment, 0)
	for i := 0; i < len(users); i++ {
		for j := 0; j < len(users[i].Comments); j++ {
			if IsBadComment(users[i].Comments[j].Message) == false {
				comments = append(comments, users[i].Comments[j])
			}
		}
	}
	return comments
}
