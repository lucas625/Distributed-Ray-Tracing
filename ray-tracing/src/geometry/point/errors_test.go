package point

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestPoint_invalidDimensionError tests the invalid dimension error of a Point.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_invalidDimensionError(t *testing.T) {
	dimension := -1
	expectedErrorMessage := fmt.Sprintf("Invalid dimension for point: %d.", dimension)

	err := invalidDimensionError(dimension)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestPoint_indexError tests the index error of a Point.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_indexError(t *testing.T) {
	dimension := 3
	index := -1
	point, err := Init(dimension)
	test_helpers.AssertNilError(t, err)
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the point. Expected from 0 to %v and got %v.", point.Dimension(), index)

	err = indexError(point, index)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
