package point_repository

import (
	"errors"
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
)

// indexError is the error where we try to access an index out of the limits of the PointRepository.
//
// Parameters:
//  pointRepository - The PointRepository being accessed.
//	index           - The index being accessed.
//
// Returns:
//  An Error.
//
func indexError(pointRepository *PointRepository, index int) error {
	errorMessage := fmt.Sprintf(
		"Index out of limits of the point repository. Expected from 0 to %v and got %v.",
		pointRepository.NumberOfPoints(), index)
	return errors.New(errorMessage)
}

// incompatibleDimensionError is the error where we try to set points with dimensions different from expected to the
// PointRepository.
//
// Parameters:
//  points    - The list of points.
//	dimension - The expected dimension for points.
//
// Returns:
//  An Error.
//
func incompatibleDimensionError(points []*point.Point, dimension int) error {
	errorMessage := fmt.Sprintf("Not all points have %v dimensions. Points: %v.", dimension, points)
	return errors.New(errorMessage)
}

// invalidSizeError is the error where we try to instantiate a PointRepository without points.
//
// Parameters:
//  points    - The list of points.
//
// Returns:
//  An Error.
//
func invalidSizeError(points []*point.Point) error {
	errorMessage := fmt.Sprintf("Invalid points list: %v. There must be at least one point.", points)
	return errors.New(errorMessage)
}
