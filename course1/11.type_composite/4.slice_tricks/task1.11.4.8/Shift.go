package main

// Функция должна возвращать первый элемент сдвинутого среза и сам сдвинутый срез
func Shift(xs []int) (int, []int) {
	if len(xs) == 0 {
		return 0, xs
	}
	return xs[0], append(xs[len(xs)-1:], xs[:len(xs)-1]...)
}
