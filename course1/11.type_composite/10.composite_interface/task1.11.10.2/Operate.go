package main

func Operate(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} { // реализуй меня
	return f(i...)
}

func Concat(xs ...interface{}) interface{} { // реализуй меня для string
	var res string
	for _, value := range xs {
		res += value.(string)
	}
	return res
}

func Sum(xs ...interface{}) interface{} { // реализуй меня для int и float64
	if len(xs) == 0 {
		return 0
	}
	switch xs[0].(type) {
	default:
		return 0
	case int:
		var res int
		for _, value := range xs {
			res = res + value.(int)
		}
		return res
	case float64:
		var res float64
		for _, value := range xs {
			res += value.(float64)
		}
		return res
	}

}
