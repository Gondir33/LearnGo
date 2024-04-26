package main

import "strings"

func ConcatenateStrings(sep string, str ...string) string {
	var evenStr, oddStr string
	i := 0
	evenStr = strings.TrimSuffix(sep, sep)
	for ; i < len(str)-2; i += 2 {
		evenStr += str[i] + sep
	}
	evenStr += str[i]
	i = 1
	for ; i < len(str)-2; i += 2 {
		oddStr += str[i] + sep
	}
	oddStr += str[i]
	return "even: " + evenStr + ", odd: " + oddStr
}
