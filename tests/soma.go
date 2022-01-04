// Isso é um este
package tests

// Funcoa de somar
func Soma(a int, b int) int {
	return a + b
}

func Fibonacci(v int) int {
	if v < 2 {
		return v
	} else {
		return Fibonacci(v-1) + Fibonacci(v-2)
	}
}
