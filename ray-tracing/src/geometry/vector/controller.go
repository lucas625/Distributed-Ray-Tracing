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
	vectAux, _ := Init(vect.Dimension())
	for i := 0; i < vect.Dimension(); i++ {
		vectAux.SetCoordinates(i, scalar * vect.GetCoordinate(i))
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
	if !vect1.IsEqualDimension(vect2) {
		return nil, differentDimensionError(vect1, vect2)
	}

	firstMultipliedVector := ScalarMultiplication(vect1, scalar1)
	secondMultipliedVector := ScalarMultiplication(vect2, scalar2)
	
	resultingVector, _ := Init(vect1.Dimension())
	for i := 0; i < vect1.Dimension(); i++ {
		resultingVector.SetCoordinates(i, firstMultipliedVector.GetCoordinate(i) +
			secondMultipliedVector.GetCoordinate(i))
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
	if !vect1.IsEqualDimension(vect2) {
		return 0, differentDimensionError(vect1, vect2)
	}
	var totalSum float64
	for i := 0; i < vect1.Dimension(); i++ {
		totalSum += vect1.GetCoordinate(i) * vect2.GetCoordinate(i)
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
	if !vect1.IsEqualDimension(vect2) {
		return nil, differentDimensionError(vect1, vect2)
	}
	if vect1.Dimension() != 3 {
		return nil, non3DError(vect1)
	}

	i := (vect1.GetCoordinate(1) * vect2.GetCoordinate(2)) -
		(vect1.GetCoordinate(2) * vect2.GetCoordinate(1))
	j := (vect1.GetCoordinate(2) * vect2.GetCoordinate(0)) -
		(vect1.GetCoordinate(0) * vect2.GetCoordinate(2))
	k := (vect1.GetCoordinate(0) * vect2.GetCoordinate(1)) -
		(vect1.GetCoordinate(1) * vect2.GetCoordinate(0))

	newVector, _ := Init(3)
	newVector.SetCoordinates(0, i)
	newVector.SetCoordinates(1, j)
	newVector.SetCoordinates(2, k)

	return newVector, nil
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
		normalizedVector, _ = Init(vector.Dimension())
	}

	return normalizedVector
}

// ProjectOnVector is a function to project one vector on another.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	The resulting vector.
//  An error.
//
func ProjectOnVector(vect1, vect2 *Vector) (*Vector, error) {
	if !vect1.IsEqualDimension(vect2) {
		return nil, differentDimensionError(vect1, vect2)
	}
	topConstant, _ := DotProduct(vect1, vect2)
	bottomConstant, _ := DotProduct(vect2, vect2)
	return ScalarMultiplication(vect2, topConstant/bottomConstant), nil
}

// Orthogonalize is a function to orthogonalize two vectors.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	The resulting vector.
//  An error.
//
func Orthogonalize(vect1, vect2 *Vector) (*Vector, error) {
	vectAux, err := ProjectOnVector(vect1, vect2)
	if err != nil {
		return nil, err
	}
	return Sum(vect1, vectAux, 1, -1)
}

// IsOrthogonalVector is a function to check if two vectors are orthogonal to each other.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	If the vectors are orthogonal to each other.
//  An error.
//
func IsOrthogonalVector(vect1, vect2 *Vector) (bool, error) {
	dotProduct, err := DotProduct(vect1, vect2)
	if err != nil {
		return false, err
	}
	return dotProduct == 0, nil
}
