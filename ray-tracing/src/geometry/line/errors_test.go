package line

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestLine_PointAndVectorIncompatibleDimensionError tests the point and vector incompatible dimension error.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestLine_PointAndVectorIncompatibleDimensionError(t *testing.T) {
	samplePoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	sampleVector, err := vector.Init(2)
	expectedErrorMessage := fmt.Sprintf(
		"Incompatible dimension for point: %d and vector: %d.", samplePoint.Dimension(), sampleVector.Dimension())
	err = pointAndVectorIncompatibleDimensionError(samplePoint, sampleVector)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
