package solution

import (
	"math/rand"
	"sort"
	"sync"
)

func createVector(n int) Vector {
	v := make(Vector, n)
	for i := range v {
		v[i] = rand.Float64()
	}
	return v
}

func createMatrix(n int) Matrix {
	m := make(Matrix, n)
	for i := range m {
		m[i] = createVector(n)
	}
	return m
}

func multiplyMatricesConcurrent(A, B Matrix, size int) Matrix {
	result := make(Matrix, size)
	for i := range result {
		result[i] = make(Vector, size)
	}

	var wg sync.WaitGroup
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < size; j++ {
				var numbers Vector
				for k := 0; k < size; k++ {
					numbers = append(numbers, A[i][k]*B[k][j])
				}
				sort.Float64s(numbers)
				sum := 0.0
				for _, num := range numbers {
					sum += num
				}
				result[i][j] = sum
			}
		}(i)
	}
	wg.Wait()
	return result
}

func multiplyVectorAndMatrixConcurrent(vector Vector, matrix Matrix, size int) Vector {
	result := make(Vector, size)

	var wg sync.WaitGroup
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func(i int) {
			defer wg.Done()
			var numbers Vector
			for j := 0; j < size; j++ {
				numbers = append(numbers, vector[j]*matrix[j][i])
			}
			sort.Float64s(numbers)
			sum := 0.0
			for _, num := range numbers {
				sum += num
			}
			result[i] = sum
		}(i)
	}
	wg.Wait()

	return result
}

func multiplyVectorAndScalarConcurrent(vector Vector, scalar float64, size int) Vector {
	result := make(Vector, size)
	var wg sync.WaitGroup
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func(i int) {
			defer wg.Done()
			result[i] = vector[i] * scalar
		}(i)
	}
	wg.Wait()

	return result
}

func multiplyMatrixAndScalarConcurrent(matrix Matrix, scalar float64, size int) Matrix {
	result := make(Matrix, size)
	for i := range result {
		result[i] = make(Vector, size)
	}

	var wg sync.WaitGroup
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < size; j++ {
				result[i][j] = matrix[i][j] * scalar
			}
		}(i)
	}
	wg.Wait()

	return result
}

func addVectorsConcurrent(A, B Vector, size int) Vector {
	result := make(Vector, size)

	var wg sync.WaitGroup
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func(i int) {
			defer wg.Done()
			result[i] = A[i] + B[i]
		}(i)
	}
	wg.Wait()

	return result
}
