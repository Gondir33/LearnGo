package main

func sortDescInt(args [8]int) [8]int {
	var tmp int
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 8; j++ {
			if args[i] <= args[j] {
				tmp = args[i]
				args[i] = args[j]
				args[j] = tmp
			}
		}
	}
	return args
}

func sortAscInt(args [8]int) [8]int {
	var tmp int
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 8; j++ {
			if args[i] >= args[j] {
				tmp = args[i]
				args[i] = args[j]
				args[j] = tmp
			}
		}
	}
	return args
}

func sortDescFloat(args [8]float64) [8]float64 {
	var tmp float64
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 8; j++ {
			if args[i] <= args[j] {
				tmp = args[i]
				args[i] = args[j]
				args[j] = tmp
			}
		}
	}
	return args
}

func sortAscFloat(args [8]float64) [8]float64 {
	var tmp float64
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 8; j++ {
			if args[i] >= args[j] {
				tmp = args[i]
				args[i] = args[j]
				args[j] = tmp
			}
		}
	}
	return args
}
