package solution

import "testing"

func BenchmarkCalculate(b *testing.B) {
	n := 100

	data := GetData(n)

	for i := 0; i < b.N; i++ {
		Calculate(data)
	}
}

func BenchmarkCalculateConcurrently(b *testing.B) {
	n := 100

	data := GetData(n)

	for i := 0; i < b.N; i++ {
		CalculateConcurrently(data)
	}
}
