package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var stdout bytes.Buffer
	stdout.ReadFrom(r)

	expected := "package main\n\ntype User struct {\n\tID\t\tint\n\tFirstName\tstring\n\tLastName\tstring\n\tUsername\tstring\n\tEmail\t\tstring\n\tAddress\t\tstring\n\tStatus\t\tint\n\tDeletedAt\tstring\n}\n"
	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
