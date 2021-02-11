package triangle

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point_repository"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// buildSamplePoint builds a sample point for testing.
//
// Parameters:
//  t       - Test instance.
//  numbers - The value of all coordinates of the triangle.
//
// Returns:
//  none
//
func buildSamplePoint (t *testing.T, numbers float64) *point.Point {
	samplePoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)

	for index := 0; index < 3; index++ {
		err = samplePoint.SetCoordinate(index, numbers)
		test_helpers.AssertNilError(t, err)
	}
	return samplePoint
}

// buildSamplePoints builds 3 sample points for testing.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func buildSamplePoints (t *testing.T) []*point.Point {
	return []*point.Point{
		buildSamplePoint(t, 1), buildSamplePoint(t, 2), buildSamplePoint(t, 3)}
}

// TestController_GetActualPoint tests the get actual point of a Triangle.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_GetActualPoint(t *testing.T) {
	points := buildSamplePoints(t)
	repository, err := point_repository.Init(points, 3)
	test_helpers.AssertNilError(t, err)
	triangle, err := Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)
	triangleController := Controller{}
	receivedPoint, err := triangleController.getActualPoint(triangle, repository, 0)
	test_helpers.AssertNilError(t, err)
	expectedPoint := buildSamplePoint(t, 1)
	test_helpers.AssertEqual(t, true, expectedPoint.IsEqual(receivedPoint))
}

// TestController_GetActualPoint_OutOfTriangle tests the get actual point of a Triangle when the index is out of the
// triangle.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_GetActualPoint_OutOfTriangle(t *testing.T) {
	points := buildSamplePoints(t)
	repository, err := point_repository.Init(points, 3)
	test_helpers.AssertNilError(t, err)
	triangle, err := Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)
	triangleController := Controller{}
	_, err = triangleController.getActualPoint(triangle, repository, 4)
	test_helpers.AssertNotNilError(t, err)
	expectedErrorMessage := fmt.Sprintf("Index out of limits of the triangle. Expected from 0 to 2 and got %v.", 4)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestController_GetActualPoint_OutOfTriangle tests the get actual point of a Triangle when the index is out of the
// triangle.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_GetActualPoint_OutOfPointRepository(t *testing.T) {
	points := buildSamplePoints(t)
	repository, err := point_repository.Init(points, 3)
	test_helpers.AssertNilError(t, err)
	triangle, err := Init([]int{4, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)
	triangleController := Controller{}
	_, err = triangleController.getActualPoint(triangle, repository, 0)
	test_helpers.AssertNotNilError(t, err)
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the point repository. Expected from 0 to %v and got %v.",
		repository.NumberOfPoints(), 4)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
