package main

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func Test_All(t *testing.T) {
	bt := NewBTree(2)
	users := make([]User, 0, 50)
	for i := 0; i < 50; i++ {
		users = append(users, User{ID: i, Name: gofakeit.Name(), Age: gofakeit.IntRange(0, 120)})
	}

	for _, user := range users {
		bt.Insert(user)
	}

	got := bt.Search(49)
	assert.Equal(t, users[49], *got)

	got = bt.Search(120)
	assert.Nil(t, got)
}
