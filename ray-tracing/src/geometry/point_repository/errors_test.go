package point_repository

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestPoint_IndexError tests the index error of a PointRepository.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_IndexError(t *testing.T) {
	firstPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	index := -1
	pointRepository, err := Init([]*point.Point{firstPoint}, 3)
	test_helpers.AssertNilError(t, err)
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the point repository. Expected from 0 to %v and got %v.",
		len(pointRepository.points), index)

	err = indexError(pointRepository, index)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestPoint_IncompatibleDimensionError tests the incompatible dimension error of a PointRepository.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_IncompatibleDimensionError(t *testing.T) {
	firstPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	points := []*point.Point{firstPoint}
	dimension := 3
	expectedErrorMessage := fmt.Sprintf("Not all points have %v dimensions. Points: %v.", dimension, points)

	err = incompatibleDimensionError(points, dimension)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestPoint_InvalidSizeError tests the invalid size error of a PointRepository.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_InvalidSizeError(t *testing.T) {
	var points []*point.Point
	expectedErrorMessage := fmt.Sprintf("Invalid points list: %v. There must be at least one point.", points)

	err := invalidSizeError(points)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
