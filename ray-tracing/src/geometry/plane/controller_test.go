package plane

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestController_ExtractPlaneNormalVectorFromR3Points tests the extraction of the Plane is normal vector from 3 points
// in the R3.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_ExtractPlaneNormalVectorFromR3Points(t *testing.T) {
	controller := Controller{}
	firstPoint, secondPoint, thirdPoint := buildSamplePoints(t)
	expectedNormalVector, err := vector.Init(3)
	test_helpers.AssertNilError(t, err)
	err = expectedNormalVector.SetCoordinate(0, -1)
	test_helpers.AssertNilError(t, err)

	extractedNormalVector, err := controller.extractPlaneNormalVectorFromR3Points(firstPoint, secondPoint, thirdPoint)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, true, expectedNormalVector.IsEqual(extractedNormalVector))
}

// TestController_ExtractPlaneNormalVectorFromR3Points_Non3DPoints tests the extraction of the Plane is normal vector
// from 3 points in the R3 when not all of them are on R3.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_ExtractPlaneNormalVectorFromR3Points_Non3DPoints(t *testing.T) {
	controller := Controller{}
	firstPoint, secondPoint, thirdPoint := buildSamplePoints(t)
	thirdPoint, err := point.Init(2)
	test_helpers.AssertNilError(t, err)
	expectedErrorMessage := fmt.Sprintf("Not all points is dimension is equal to 3: %d %d %d.", firstPoint.Dimension(),
		secondPoint.Dimension(), thirdPoint.Dimension())

	_, err = controller.extractPlaneNormalVectorFromR3Points(firstPoint, secondPoint, thirdPoint)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestController_ExtractPlaneFromR3Points tests the extraction of a Plane from 3 points in the R3.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_ExtractPlaneFromR3Points(t *testing.T) {
	controller := Controller{}
	firstPoint, secondPoint, thirdPoint := buildSamplePoints(t)
	expectedPlane := Init(-1, 0, 0, 0)

	plane, err := controller.ExtractPlaneFromR3Points(firstPoint, secondPoint, thirdPoint)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, true, expectedPlane.IsEqual(plane))
}

// TestController_ExtractPlaneFromR3Points_Non3DPoints tests the extraction of a Plane from 3 points in the R3 when
// not all of them are on R3.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_ExtractPlaneFromR3Points_Non3DPoints(t *testing.T) {
	controller := Controller{}
	firstPoint, secondPoint, thirdPoint := buildSamplePoints(t)
	thirdPoint, err := point.Init(2)
	test_helpers.AssertNilError(t, err)
	expectedErrorMessage := fmt.Sprintf("Not all points is dimension is equal to 3: %d %d %d.", firstPoint.Dimension(),
		secondPoint.Dimension(), thirdPoint.Dimension())

	_, err = controller.ExtractPlaneFromR3Points(firstPoint, secondPoint, thirdPoint)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// buildSamplePoints builds sample points for the tests.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func buildSamplePoints(t *testing.T) (*point.Point, *point.Point, *point.Point){
	firstPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)

	secondPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	err = secondPoint.SetCoordinate(2, 2)
	test_helpers.AssertNilError(t, err)

	thirdPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	err = thirdPoint.SetCoordinate(1, 2)
	test_helpers.AssertNilError(t, err)
	err = thirdPoint.SetCoordinate(2, 2)
	test_helpers.AssertNilError(t, err)
	return firstPoint, secondPoint, thirdPoint
}
