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
func ScalarMultiplication(vect *Vector, scalar float64) Vector {
	vectAux := Init(len(vect.Coordinates))
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
//
func Sum(vect1 *Vector, vect2 *Vector, scalar1, scalar2 float64) Vector {
	if !IsEqualDimension(vect1, vect2) {
		differentDimensionError(vect1, vect2)
	}

	multipliedVect1 := ScalarMultiplication(vect1, scalar1)
	multipliedVect2 := ScalarMultiplication(vect2, scalar2)

	resultingVector := Init(len(vect1.Coordinates))
	for i := 0; i < len(vect1.Coordinates); i++ {
		resultingVector.Coordinates[i] = multipliedVect1.Coordinates[i] + multipliedVect2.Coordinates[i]
	}
	return resultingVector
}

// DotProduct is a function to dot product 2 vectors.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	The resulting sum.
//
func DotProduct(vect1 *Vector, vect2 *Vector) float64 {
	if !IsEqualDimension(vect1, vect2) {
		differentDimensionError(vect1, vect2)
	}
	var totalSum float64
	for i := 0; i < len(vect1.Coordinates); i++ {
		totalSum += vect1.Coordinates[i] * vect2.Coordinates[i]
	}
	return totalSum
}

// CrossProduct is a function to calculate the cross product of two Vectors.
//
// Parameters:
// 	vect1 - The first Vector.
//  vect2 - The second Vector.
//
// Returns:
// 	The resulting vector.
//
func CrossProduct(vect1, vect2 *Vector) Vector {
	if !IsEqualDimension(vect1, vect2) {
		differentDimensionError(vect1, vect2)
	}
	if len(vect1.Coordinates) != 3 {
		non3DError(vect1)
	}

	i := (vect1.Coordinates[1] * vect2.Coordinates[2]) - (vect1.Coordinates[2] * vect2.Coordinates[1])
	j := (vect1.Coordinates[2] * vect2.Coordinates[0]) - (vect1.Coordinates[0] * vect2.Coordinates[2])
	k := (vect1.Coordinates[0] * vect2.Coordinates[1]) - (vect1.Coordinates[1] * vect2.Coordinates[0])

	return Vector{Coordinates: []float64{i, j, k}}	 
}

// VectorNorm is a function to calculate the norm of a Vector.
//
// Parameters:
// 	vect - The Vector.
//
// Returns:
// 	the norm of the Vector.
//
func VectorNorm(vect *Vector) float64 {
	return math.Sqrt(DotProduct(vect, vect))
}

// NormalizeVector is a function to normalize a Vector.
//
// Parameters:
// 	vect - The Vector.
//
// Returns:
// 	the normalized Vector.
//
func NormalizeVector(vect *Vector) Vector {
	return ScalarMultiplication(vect, 1/VectorNorm(vect))
}
