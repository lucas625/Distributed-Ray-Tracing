package point

import (
	"errors"
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
)

// invalidDimensionError is the error where a Point has an invalid dimension.
//
// Parameters:
//	dimension - The dimension of the Point.
//
// Returns:
//  An Error.
//
func invalidDimensionError(dimension int) error {
	errorMessage := fmt.Sprintf("Invalid dimension for point: %d.", dimension)
	return errors.New(errorMessage)
}

// indexError is the error where we try to access an index out of the limits of the Point.
//
// Parameters:
//  point - The Point being accessed.
//	index - The index being accessed.
//
// Returns:
//  An Error.
//
func indexError(point *Point, index int) error {
	errorMessage := fmt.Sprintf(
		"Index out of limits of the point. Expected from 0 to %v and got %v.", point.Dimension(), index)
	return errors.New(errorMessage)
}

// differentDimensionsError is a function to get the error where two points do not have the same dimension.
//
// Parameters:
//	firstPoint  - The first Point.
//	secondPoint - The second Point.
//
// Returns:
//  An Error
//
func differentDimensionsError(firstPoint, secondPoint *Point) error {
	errorMessage := fmt.Sprintf(
		"Invalid dimension of point. Expected: %d and got: %d.\n", firstPoint.Dimension(), secondPoint.Dimension())
	return errors.New(errorMessage)
}

// pointAndVectorIncompatibleDimensionError is the error where a point has an incompatible dimension to a vector.
//
// Parameters:
//	startingPoint - The starting Point.
//  targetVector  - The vector to sum with the Point.
//
// Returns:
//  An Error
//
func pointAndVectorIncompatibleDimensionError(startingPoint *Point, targetVector *vector.Vector) error {
	errorMessage := fmt.Sprintf(
		"Incompatible dimension for point: %d and vector: %d.", startingPoint.Dimension(), targetVector.Dimension())
	return errors.New(errorMessage)
}
