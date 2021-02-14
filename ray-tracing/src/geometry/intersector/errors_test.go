package intersector

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/line"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point_repository"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestIntersector_Non3DRayPointsError tests the non 3D ray or point repository error.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIntersector_Non3DRayPointsError(t *testing.T) {
	samplePoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	sampleVector, err := vector.Init(3)
	test_helpers.AssertNilError(t, err)

	repository, err := point_repository.Init([]*point.Point{samplePoint}, 3)
	test_helpers.AssertNilError(t, err)
	ray, err := line.Init(samplePoint, sampleVector)
	test_helpers.AssertNilError(t, err)

	expectedErrorMessage := fmt.Sprintf(
		"Non 3D ray or repository. Ray dimension: %d and point repository dimension: %d.",
		ray.Dimension(), repository.PointsDimension())
	err = non3DRayPointsError(ray, repository)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
