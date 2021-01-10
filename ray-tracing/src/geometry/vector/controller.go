package vector

import (
	"math"
)

type Controller struct {}

// ScalarMultiplication is a function for Scalar Multiplication.
//
// Parameters:
// 	vector - The Vector.
// 	scalar - The constant for the multiplication.
//
// Returns:
// 	The resulting Vector.
//
func (_ *Controller) ScalarMultiplication(vector *Vector, scalar float64) *Vector {
	newVector, _ := Init(vector.Dimension())
	for i := 0; i < vector.Dimension(); i++ {
		newVector.SetCoordinate(i, scalar * vector.GetCoordinate(i))
	}
	return newVector
}

// Sum is a function to sum 2 vectors.
//
// Parameters:
// 	firstVector  - The first Vector.
// 	secondVector - The second Vector.
// 	firstScalar  - Constant multiplying the first Vector.
// 	secondScalar - Constant multiplying the second Vector.
//
// Returns:
// 	The resulting Vector.
//	An error.
//
func (controller *Controller) Sum(firstVector *Vector, secondVector *Vector, firstScalar, secondScalar float64) (*Vector, error) {
	if !firstVector.IsEqualDimension(secondVector) {
		return nil, differentDimensionError(firstVector, secondVector)
	}

	firstMultipliedVector := controller.ScalarMultiplication(firstVector, firstScalar)
	secondMultipliedVector := controller.ScalarMultiplication(secondVector, secondScalar)
	
	resultingVector, _ := Init(firstVector.Dimension())
	for i := 0; i < firstVector.Dimension(); i++ {
		resultingVector.SetCoordinate(i, firstMultipliedVector.GetCoordinate(i) +
			secondMultipliedVector.GetCoordinate(i))
	}
	return resultingVector, nil
}

// DotProduct is a function to dot product 2 vectors.
//
// Parameters:
// 	firstVector  - The first Vector.
// 	secondVector - The second Vector.
//
// Returns:
// 	The resulting sum.
//	An error.
//
func (_ *Controller) DotProduct(firstVector *Vector, secondVector *Vector) (float64, error) {
	if !firstVector.IsEqualDimension(secondVector) {
		return 0, differentDimensionError(firstVector, secondVector)
	}
	var totalSum float64
	for i := 0; i < firstVector.Dimension(); i++ {
		totalSum += firstVector.GetCoordinate(i) * secondVector.GetCoordinate(i)
	}
	return totalSum, nil
}

// CrossProduct is a function to calculate the cross product of two Vectors.
//
// Parameters:
// 	firstVector  - The first Vector.
//  secondVector - The second Vector.
//
// Returns:
// 	The resulting Vector.
//  An error.
//
func (_ *Controller) CrossProduct(firstVector, secondVector *Vector) (*Vector, error) {
	if !firstVector.IsEqualDimension(secondVector) {
		return nil, differentDimensionError(firstVector, secondVector)
	}
	if firstVector.Dimension() != 3 {
		return nil, non3DError(firstVector)
	}

	i := (firstVector.GetCoordinate(1) * secondVector.GetCoordinate(2)) -
		(firstVector.GetCoordinate(2) * secondVector.GetCoordinate(1))
	j := (firstVector.GetCoordinate(2) * secondVector.GetCoordinate(0)) -
		(firstVector.GetCoordinate(0) * secondVector.GetCoordinate(2))
	k := (firstVector.GetCoordinate(0) * secondVector.GetCoordinate(1)) -
		(firstVector.GetCoordinate(1) * secondVector.GetCoordinate(0))

	newVector, _ := Init(3)
	newVector.SetCoordinate(0, i)
	newVector.SetCoordinate(1, j)
	newVector.SetCoordinate(2, k)

	return newVector, nil
}

// Norm is a function to calculate the norm of a Vector.
//
// Parameters:
// 	vector - The Vector.
//
// Returns:
// 	The norm of the Vector.
//
func (controller *Controller) Norm(vector *Vector) float64 {
	dotProduct, _ := controller.DotProduct(vector, vector)
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
func (controller *Controller) Normalize(vector *Vector) *Vector {
	vectorNorm := controller.Norm(vector)
	var normalizedVector *Vector
	if vectorNorm != 0 {
		normalizedVector = controller.ScalarMultiplication(vector, 1/vectorNorm)
	} else {
		normalizedVector, _ = Init(vector.Dimension())
	}

	return normalizedVector
}

// ProjectOnVector is a function to project one Vector on another.
//
// Parameters:
// 	firstVector  - The first Vector.
// 	secondVector - The second Vector.
//
// Returns:
// 	The resulting Vector.
//  An error.
//
func (controller *Controller) ProjectOnVector(firstVector, secondVector *Vector) (*Vector, error) {
	if !firstVector.IsEqualDimension(secondVector) {
		return nil, differentDimensionError(firstVector, secondVector)
	}
	topConstant, _ := controller.DotProduct(firstVector, secondVector)
	bottomConstant, _ := controller.DotProduct(secondVector, secondVector)
	return controller.ScalarMultiplication(secondVector, topConstant/bottomConstant), nil
}

// Orthogonalize is a function to orthogonalize two vectors.
//
// Parameters:
// 	firstVector  - The first Vector.
// 	secondVector - The second Vector.
//
// Returns:
// 	The resulting Vector.
//  An error.
//
func (controller *Controller) Orthogonalize(firstVector, secondVector *Vector) (*Vector, error) {
	newVector, err := controller.ProjectOnVector(firstVector, secondVector)
	if err != nil {
		return nil, err
	}
	return controller.Sum(firstVector, newVector, 1, -1)
}

// IsOrthogonalVector is a function to check if two vectors are orthogonal to each other.
//
// Parameters:
// 	firstVector  - The first Vector.
// 	secondVector - The second Vector.
//
// Returns:
// 	If the vectors are orthogonal to each other.
//  An error.
//
func (controller *Controller) IsOrthogonalVector(firstVector, secondVector *Vector) (bool, error) {
	dotProduct, err := controller.DotProduct(firstVector, secondVector)
	if err != nil {
		return false, err
	}
	return dotProduct == 0, nil
}
