package vector

import (
	"math"
)

// ScalarMultiplication is a function for Scalar Multiplication.
//
// Parameters:
// 	vect   - The vector.
// 	scalar - The constant for the multiplication.
//
// Returns:
// 	The resulting vector.
//
func ScalarMultiplication(vect *Vector, scalar float64) *Vector {
	vectAux, _ := Init(len(vect.Coordinates))
	for i := 0; i < len(vect.Coordinates); i++ {
		vectAux.Coordinates[i] = scalar * vect.Coordinates[i]
	}
	return vectAux
}

// Sum is a function to sum 2 vectors.
//
// Parameters:
// 	vect1   - The first vector.
// 	vect2   - The second vector.
// 	scalar1 - Constant multiplying the first vector.
// 	scalar2 - Constant multiplying the second vector.
//
// Returns:
// 	The resulting vector.
//	An error.
//
func Sum(vect1 *Vector, vect2 *Vector, scalar1, scalar2 float64) (*Vector, error) {
	if !IsEqualDimension(vect1, vect2) {
		return nil, differentDimensionError(vect1, vect2)
	}

	firstMultipliedVector := ScalarMultiplication(vect1, scalar1)
	secondMultipliedVector := ScalarMultiplication(vect2, scalar2)
	
	resultingVector, _ := Init(len(vect1.Coordinates))
	for i := 0; i < len(vect1.Coordinates); i++ {
		resultingVector.Coordinates[i] = firstMultipliedVector.Coordinates[i] + secondMultipliedVector.Coordinates[i]
	}
	return resultingVector, nil
}

// DotProduct is a function to dot product 2 vectors.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	The resulting sum.
//	An error.
//
func DotProduct(vect1 *Vector, vect2 *Vector) (float64, error) {
	if !IsEqualDimension(vect1, vect2) {
		return 0, differentDimensionError(vect1, vect2)
	}
	var totalSum float64
	for i := 0; i < len(vect1.Coordinates); i++ {
		totalSum += vect1.Coordinates[i] * vect2.Coordinates[i]
	}
	return totalSum, nil
}

// CrossProduct is a function to calculate the cross product of two Vectors.
//
// Parameters:
// 	vect1 - The first Vector.
//  vect2 - The second Vector.
//
// Returns:
// 	The resulting vector.
//  An error.
//
func CrossProduct(vect1, vect2 *Vector) (*Vector, error) {
	if !IsEqualDimension(vect1, vect2) {
		return nil, differentDimensionError(vect1, vect2)
	}
	if len(vect1.Coordinates) != 3 {
		return nil, non3DError(vect1)
	}

	i := (vect1.Coordinates[1] * vect2.Coordinates[2]) - (vect1.Coordinates[2] * vect2.Coordinates[1])
	j := (vect1.Coordinates[2] * vect2.Coordinates[0]) - (vect1.Coordinates[0] * vect2.Coordinates[2])
	k := (vect1.Coordinates[0] * vect2.Coordinates[1]) - (vect1.Coordinates[1] * vect2.Coordinates[0])

	return &Vector{Coordinates: []float64{i, j, k}}, nil
}

// Norm is a function to calculate the norm of a Vector.
//
// Parameters:
// 	vect - The Vector.
//
// Returns:
// 	The norm of the Vector.
//
func Norm(vect *Vector) float64 {
	dotProduct, _ := DotProduct(vect, vect)
	return math.Sqrt(dotProduct)
}

// Normalize is a function to normalize a Vector.
//
// Parameters:
// 	vector - The Vector.
//
// Returns:
// 	The normalized Vector.
//
func Normalize(vector *Vector) *Vector {
	vectorNorm := Norm(vector)
	var normalizedVector *Vector
	if vectorNorm != 0 {
		normalizedVector = ScalarMultiplication(vector, 1/vectorNorm)
	} else {
		normalizedVector, _ = Init(len(vector.Coordinates))
	}

	return normalizedVector
}
