package plane

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestPoint_Non3DPointsError tests the error where a point is not on R3 space.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_Non3DPointsError(t *testing.T) {
	firstPoint, err := point.Init(2)
	secondPoint, err := point.Init(3)
	thirdPoint, err := point.Init(1)
	test_helpers.AssertNilError(t, err)
	expectedErrorMessage := fmt.Sprintf("Not all points is dimension is equal to 3: %d %d %d.", firstPoint.Dimension(),
		secondPoint.Dimension(), thirdPoint.Dimension())

	err = non3DPointsError(firstPoint, secondPoint, thirdPoint)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
