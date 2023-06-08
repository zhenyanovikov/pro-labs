package solution

import "math"

func Calculate(data *Data) (Matrix, Vector) {
	// Calculate MA = min(D + B) * MD * MT + MX * ME
	minDPlusB := math.Min(data.D[0]+data.B[0], data.D[1]+data.B[1])
	for i := 2; i < data.N; i++ {
		minDPlusB = math.Min(minDPlusB, data.D[i]+data.B[i])
	}
	MDMT := multiplyMatrices(data.MD, data.MT, data.N)
	MDMTScalar := multiplyMatrixAndScalar(MDMT, minDPlusB, data.N)
	MXME := multiplyMatrices(data.MX, data.ME, data.N)
	MA := multiplyMatrices(MDMTScalar, MXME, data.N)

	// Calculate C = B * MT + D * MX * A
	BMT := multiplyVectorAndMatrix(data.B, data.MT, data.N)
	DMX := multiplyVectorAndMatrix(data.D, data.MX, data.N)
	C := addVectors(BMT, multiplyVectorAndScalar(DMX, data.A, data.N), data.N)

	return MA, C
}

func CalculateConcurrently(data *Data) (Matrix, Vector) {
	// Calculate MA = min(D + B) * MD * MT + MX * ME
	minDPlusB := math.Min(data.D[0]+data.B[0], data.D[1]+data.B[1])
	for i := 2; i < data.N; i++ {
		minDPlusB = math.Min(minDPlusB, data.D[i]+data.B[i])
	}
	MDMT := multiplyMatricesConcurrent(data.MD, data.MT, data.N)
	MDMTScalar := multiplyMatrixAndScalarConcurrent(MDMT, minDPlusB, data.N)
	MXME := multiplyMatricesConcurrent(data.MX, data.ME, data.N)
	MA := multiplyMatricesConcurrent(MDMTScalar, MXME, data.N)

	// Calculate C = B * MT + D * MX * A
	BMT := multiplyVectorAndMatrixConcurrent(data.B, data.MT, data.N)
	DMX := multiplyVectorAndMatrixConcurrent(data.D, data.MX, data.N)
	C := addVectorsConcurrent(BMT, multiplyVectorAndScalarConcurrent(DMX, data.A, data.N), data.N)

	return MA, C
}
