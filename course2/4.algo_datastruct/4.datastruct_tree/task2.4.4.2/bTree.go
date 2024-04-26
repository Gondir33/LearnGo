package main

import (
	"reflect"

	"github.com/google/btree"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func (u User) Less(than btree.Item) bool {
	return u.ID < than.(User).ID
}

type BTree struct {
	tree *btree.BTree
}

func NewBTree(degree int) *BTree {
	btree := &BTree{tree: btree.New(degree)}
	return btree
}

func (bt *BTree) Insert(user User) {
	bt.tree.ReplaceOrInsert(user)
}

func (bt *BTree) Search(id int) *User {
	item := bt.tree.Get(User{ID: id})
	if item == nil {
		return nil
	}
	re := reflect.ValueOf(item)

	return &User{ID: int(re.Field(0).Int()), Name: re.Field(1).String(), Age: int(re.Field(2).Int())}
}

/*
func main() {
	bt := NewBTree(2)
	users := make([]User, 0, 50)
	for i := 0; i < 50; i++ {
		users = append(users, User{ID: i, Name: gofakeit.Name(), Age: gofakeit.IntRange(0, 120)})
	}

	for _, user := range users {
		bt.Insert(user)
	}

	if user := bt.Search(49); user != nil {
		fmt.Printf("Найден пользователь: %v\n", *user)
	} else {
		fmt.Println("Пользователь не найден")
	}
}
*/
