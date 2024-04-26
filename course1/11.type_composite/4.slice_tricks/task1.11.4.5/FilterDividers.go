package main

func FilterDividers(xs []int, divider int) []int {
	res := make([]int, 0, 0)
	for i := 0; i < len(xs); i++ {
		if xs[i]%divider == 0 {
			res = append(res, xs[i])
		}
	}
	return res
}
