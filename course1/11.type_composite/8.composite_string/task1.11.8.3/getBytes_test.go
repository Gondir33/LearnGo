package main

import "testing"

func cmpBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func cmpRunes(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestGetBytes(t *testing.T) {
	res := getBytes("hello")
	exp := getBytes("hello")
	if cmpBytes(res, exp) == false {
		t.Errorf("unexpected")
	}
}

func TestGetRunes(t *testing.T) {
	res := getRunes("hello")
	exp := getRunes("hello")
	if cmpRunes(res, exp) == false {
		t.Errorf("unexpected")
	}
}
