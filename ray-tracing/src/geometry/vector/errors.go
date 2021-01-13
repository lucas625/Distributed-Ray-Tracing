package vector

import (
	"errors"
	"fmt"
)

// differentDimensionError is a function to get the error where two vectors do not have the same dimension.
//
// Parameters:
//	firstVector  - The first vector.
//	secondVector - The second vector.
//
// Returns:
//  An Error
//
func differentDimensionError(firstVector, secondVector *Vector) error {
	errorMessage := fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n",
		firstVector.Dimension(),
		secondVector.Dimension())
	return errors.New(errorMessage)
}

// non3DError is a function to get the error where a vector is not 3D.
//
// Parameters:
//  vector - The vector.
//
// Returns:
//  An Error.
//
func non3DError(vector *Vector) error {
	errorMessage := fmt.Sprintf("Invalid dimension of vector. Expected 3D and got %d.", vector.Dimension())
	return errors.New(errorMessage)
}

// negativeDimensionError is a function to get the error where a vector has negative dimension.
//
// Parameters:
// 	dimension - The dimension of the vector.
//
// Returns:
//  An Error.
//
func negativeDimensionError(dimension int) error {
	errorMessage := fmt.Sprintf("Invalid vector size %d.", dimension)
	return errors.New(errorMessage)
}

// indexError is the error where we try to access an index out of the limits of the Vector.
//
// Parameters:
//	vector - The Vector.
//	index  - The index.
//
// Returns:
//  An Error.
//
func indexError(vector *Vector, index int) error {
	errorMessage := fmt.Sprintf(
		"Index out of limits of the vector. Expected from 0 to %v and got %v.", vector.Dimension(), index)
	return errors.New(errorMessage)
}
