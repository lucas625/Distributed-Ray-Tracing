package point

import (
	"errors"
	"fmt"
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
