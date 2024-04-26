package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var stdout bytes.Buffer
	stdout.ReadFrom(r)

	expected := "Hello, world!\n"
	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}

func main() {
	fmt.Println("Hello, world!")
}
