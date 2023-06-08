package solution

import "sort"

func multiplyMatrices(A, B Matrix, size int) Matrix {
	result := make(Matrix, size)
	for i := range result {
		result[i] = make(Vector, size)
	}

	for i := 0; i < size; i++ {
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
	}
	return result
}

func multiplyVectorAndMatrix(vector Vector, matrix Matrix, size int) Vector {
	result := make(Vector, size)

	for i := 0; i < size; i++ {
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
	}
	return result
}

func multiplyVectorAndScalar(vector Vector, scalar float64, size int) Vector {
	result := make(Vector, size)

	for i := 0; i < size; i++ {
		result[i] = vector[i] * scalar
	}
	return result
}

func multiplyMatrixAndScalar(matrix Matrix, scalar float64, size int) Matrix {
	result := make(Matrix, size)
	for i := range result {
		result[i] = make(Vector, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			result[i][j] = matrix[i][j] * scalar
		}
	}

	return result
}

func addVectors(A, B Vector, size int) Vector {
	result := make(Vector, size)

	for i := 0; i < size; i++ {
		result[i] = A[i] + B[i]
	}
	return result
}
