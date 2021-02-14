package intersector

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/line"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point_repository"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/triangle"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"reflect"
	"testing"
)

// TestController_IntersectRayTriangle tests the intersection between a ray and a triangle.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func buildSamplePointRepository(t *testing.T) *point_repository.PointRepository {
	dimension := 3

	firstPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = firstPoint.SetCoordinate(0, 2)
	test_helpers.AssertNilError(t, err)

	secondPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = secondPoint.SetCoordinate(1, 2)
	test_helpers.AssertNilError(t, err)

	thirdPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = thirdPoint.SetCoordinate(2, 2)
	test_helpers.AssertNilError(t, err)

	repository, err := point_repository.Init([]*point.Point{firstPoint, secondPoint, thirdPoint}, dimension)
	test_helpers.AssertNilError(t, err)
	return repository
}

// TestController_IntersectRayTriangle tests the intersection between a ray and a triangle.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_IntersectRayTriangle(t *testing.T) {
	dimension := 3

	repository := buildSamplePointRepository(t)

	targetTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)

	rayStartingPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = rayStartingPoint.SetCoordinate(0, 3)
	test_helpers.AssertNilError(t, err)
	err = rayStartingPoint.SetCoordinate(1, 3)
	test_helpers.AssertNilError(t, err)

	rayVectorDirector, err := vector.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = rayVectorDirector.SetCoordinate(0, -1)
	test_helpers.AssertNilError(t, err)
	err = rayVectorDirector.SetCoordinate(1, -1)
	test_helpers.AssertNilError(t, err)

	ray, err := line.Init(rayStartingPoint, rayVectorDirector)
	test_helpers.AssertNilError(t, err)

	intersectorController := Controller{}
	parametricParameter, barycentricCoordinates, hasIntersection, err := intersectorController.IntersectRayTriangle(
		ray, targetTriangle, repository)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, 2.0, parametricParameter)
	test_helpers.AssertEqual(t, true, reflect.DeepEqual([]float64{0.5, 0.5, 0}, barycentricCoordinates))
	test_helpers.AssertEqual(t, true, hasIntersection)
}

// TestController_IntersectRayTriangle_NegativeParametricParameter tests the intersection between a ray and a triangle
// when the intersection happens on a negative parametric parameter.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_IntersectRayTriangle_NegativeParametricParameter(t *testing.T) {
	dimension := 3

	repository := buildSamplePointRepository(t)

	targetTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)

	rayStartingPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)

	rayVectorDirector, err := vector.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = rayVectorDirector.SetCoordinate(0, -1)
	test_helpers.AssertNilError(t, err)
	err = rayVectorDirector.SetCoordinate(1, -1)
	test_helpers.AssertNilError(t, err)
	err = rayVectorDirector.SetCoordinate(2, -1)
	test_helpers.AssertNilError(t, err)

	ray, err := line.Init(rayStartingPoint, rayVectorDirector)
	test_helpers.AssertNilError(t, err)

	intersectorController := Controller{}
	parametricParameter, barycentricCoordinates, hasIntersection, err := intersectorController.IntersectRayTriangle(
		ray, targetTriangle, repository)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, 0.0, parametricParameter)
	test_helpers.AssertEqual(t, true, reflect.DeepEqual([]float64(nil), barycentricCoordinates))
	test_helpers.AssertEqual(t, false, hasIntersection)
}

// TestController_IntersectRayTriangle_Non3D tests the intersection between a ray and a triangle when the ray or the
// repository are no on the third dimension.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_IntersectRayTriangle_Non3D(t *testing.T) {
	dimension := 2
	expectedErrorMessage := fmt.Sprintf(
		"Non 3D ray or repository. Ray dimension: %d and point repository dimension: %d.",2, 3)

	repository := buildSamplePointRepository(t)

	targetTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)

	rayStartingPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)

	rayVectorDirector, err := vector.Init(dimension)
	test_helpers.AssertNilError(t, err)

	ray, err := line.Init(rayStartingPoint, rayVectorDirector)
	test_helpers.AssertNilError(t, err)

	intersectorController := Controller{}
	_, _, _, err = intersectorController.IntersectRayTriangle(ray, targetTriangle, repository)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestController_IntersectRayTriangle_NegativeSecondBarycentricCoordinate tests the intersection between a ray and a
// triangle.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_IntersectRayTriangle_NegativeSecondBarycentricCoordinate(t *testing.T) {
	dimension := 3

	repository := buildSamplePointRepository(t)

	targetTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)

	rayStartingPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = rayStartingPoint.SetCoordinate(0, 4)
	test_helpers.AssertNilError(t, err)
	err = rayStartingPoint.SetCoordinate(1, 3)
	test_helpers.AssertNilError(t, err)

	rayVectorDirector, err := vector.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = rayVectorDirector.SetCoordinate(1, -3)
	test_helpers.AssertNilError(t, err)

	ray, err := line.Init(rayStartingPoint, rayVectorDirector)
	test_helpers.AssertNilError(t, err)

	intersectorController := Controller{}
	parametricParameter, barycentricCoordinates, hasIntersection, err := intersectorController.IntersectRayTriangle(
		ray, targetTriangle, repository)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, 0.0, parametricParameter)
	test_helpers.AssertEqual(t, true, reflect.DeepEqual([]float64(nil), barycentricCoordinates))
	test_helpers.AssertEqual(t, false, hasIntersection)
}

// TestController_IntersectRayTriangle_BiggerThan1ThirdBarycentricCoordinate tests the intersection between a ray and a
// triangle.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_IntersectRayTriangle_BiggerThan1ThirdBarycentricCoordinate(t *testing.T) {
	dimension := 3

	repository := buildSamplePointRepository(t)

	targetTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)

	rayStartingPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = rayStartingPoint.SetCoordinate(0, 4)
	test_helpers.AssertNilError(t, err)
	err = rayStartingPoint.SetCoordinate(2, 4)
	test_helpers.AssertNilError(t, err)

	rayVectorDirector, err := vector.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = rayVectorDirector.SetCoordinate(0, -2)
	test_helpers.AssertNilError(t, err)

	ray, err := line.Init(rayStartingPoint, rayVectorDirector)
	test_helpers.AssertNilError(t, err)

	intersectorController := Controller{}
	parametricParameter, barycentricCoordinates, hasIntersection, err := intersectorController.IntersectRayTriangle(
		ray, targetTriangle, repository)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, 0.0, parametricParameter)
	test_helpers.AssertEqual(t, true, reflect.DeepEqual([]float64(nil), barycentricCoordinates))
	test_helpers.AssertEqual(t, false, hasIntersection)
}

// TestController_IntersectRayTriangle_RayParallelToTriangle tests the intersection between a ray and a triangle.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_IntersectRayTriangle_RayParallelToTriangle(t *testing.T) {
	dimension := 3

	repository := buildSamplePointRepository(t)

	targetTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)

	rayStartingPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = rayStartingPoint.SetCoordinate(0, 3)
	test_helpers.AssertNilError(t, err)

	rayVectorDirector, err := vector.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = rayVectorDirector.SetCoordinate(0, -3)
	test_helpers.AssertNilError(t, err)
	err = rayVectorDirector.SetCoordinate(1, 3)
	test_helpers.AssertNilError(t, err)

	ray, err := line.Init(rayStartingPoint, rayVectorDirector)
	test_helpers.AssertNilError(t, err)

	intersectorController := Controller{}
	parametricParameter, barycentricCoordinates, hasIntersection, err := intersectorController.IntersectRayTriangle(
		ray, targetTriangle, repository)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, 0.0, parametricParameter)
	test_helpers.AssertEqual(t, true, reflect.DeepEqual([]float64(nil), barycentricCoordinates))
	test_helpers.AssertEqual(t, false, hasIntersection)
}
