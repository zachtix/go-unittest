package functions

// Basic Functions

func Add(a int, b int) int {
	return a + b
}

func Factorial(n int) (result int) {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
