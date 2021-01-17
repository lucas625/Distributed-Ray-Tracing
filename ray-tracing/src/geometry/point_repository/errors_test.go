package point_repository

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestPoint_indexError tests the index error of a Point.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_indexError(t *testing.T) {
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
