package problem4

func Compose(f, g func(x int) int) func(x int) int {
	return func(x int) int {
		x = g(x)
		x = f(x)
		return x
	}
}

func AddOne(value int) int {
	return value + 1
}

func Square(num int) int {
	return num * num
}
