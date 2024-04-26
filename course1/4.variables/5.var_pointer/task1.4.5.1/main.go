package main

func changeInt(a *int) {
	*a = 20
} // меняет значение переменной на 20

func changeFloat(b *float64) {
	*b = 6.28
} // меняет значение переменной на 6.28

func changeString(c *string) {
	*c = "Goodbye, world!"
} // меняет значение переменной на "Goodbye, world!"

func changeBool(d *bool) {
	*d = false
} // меняет значение переменной на false
