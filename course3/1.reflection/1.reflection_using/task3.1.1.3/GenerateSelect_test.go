package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildSelect(t *testing.T) {
	d := NewDAO()
	s, _, err := d.BuildSelect("users", Condition{
		Equal: map[string]interface{}{
			"username": "test",
		},
		LimitOffset: &LimitOffset{
			Offset: 5,
			Limit:  3,
		},
		Order: []*Order{
			{
				Field: "id",
				Asc:   true,
			},
		},
	}, "id", "username")

	if err != nil {
		t.Errorf("don't work")
	}

	assert.Equal(t, "SELECT id, username FROM users WHERE username = $1 ORDER BY id ASC LIMIT 3 OFFSET 5", s)
}
