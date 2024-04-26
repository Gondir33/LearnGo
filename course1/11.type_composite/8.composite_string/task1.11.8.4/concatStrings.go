package main

func concatStrings(xs ...string) string {
	var res string

	for _, val := range xs {
		res += val
	}
	return res
}
