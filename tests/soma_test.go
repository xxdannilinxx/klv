package tests

import (
	"testing"
)

type SomaTable struct {
	a      int
	b      int
	result int
}

func TestSoma(t *testing.T) {

	data := []SomaTable{
		{1, 2, 3},
		{2, 2, 4},
	}

	for _, valor := range data {

		total := Soma(valor.a, valor.b)

		if total != valor.result {
			t.Errorf("Erro encontrado.")
		}

	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(33)
	}
}
