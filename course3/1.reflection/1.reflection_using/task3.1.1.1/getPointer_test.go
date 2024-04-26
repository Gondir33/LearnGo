package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleGetFieldsPointers(t *testing.T) {
	user := User{
		ID:       1,
		Username: "JohnDoe",
		Email:    "johndoe@example.com",
		Address:  "123 Main St",
		Status:   1,
		Delete:   "yes",
	}
	pointers := SimpleGetFieldsPointers(&user)
	assert.Equal(t, &user.Username, pointers[0])
	assert.Equal(t, &user.Email, pointers[1])
	assert.Equal(t, &user.Address, pointers[2])
	assert.Equal(t, &user.Status, pointers[3])
}
