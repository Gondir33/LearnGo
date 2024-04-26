package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestALL(t *testing.T) {
	want := make(map[string]bool)
	want[ITERATIVE] = false
	want[RECURSIVE] = true
	got := compareWhichFactorialIsFaster()
	assert.Equal(t, want, got)
}
