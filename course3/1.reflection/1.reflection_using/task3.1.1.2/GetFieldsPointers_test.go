package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFildsPointers(t *testing.T) {
	user := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
	}
	pointers := GetFieldsPointers(&user, FilterByFields(0, 1, 2))
	assert.Equal(t, &user.ID, pointers[0])
	assert.Equal(t, &user.Username, pointers[1])
	assert.Equal(t, &user.Email, pointers[2])
	filterTag := map[string]func(value string) bool{
		"db": func(value string) bool {
			values := []string{"username", "address", "status"}

			for _, v := range values {
				if v == value {
					return true
				}
			}

			return false
		},
	}
	pointers = GetFieldsPointers(&user, FilterByTags(filterTag))
	assert.Equal(t, &user.Username, pointers[0])
	assert.Equal(t, &user.Address, pointers[1])
	assert.Equal(t, &user.Status, pointers[2])
}
