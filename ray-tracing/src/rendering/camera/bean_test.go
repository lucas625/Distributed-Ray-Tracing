package camera

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// buildCameraVectors builds the Camera vectors for testing.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  A list of vectors.
//
func buildCameraVectors(t *testing.T) (*vector.Vector, *vector.Vector, *vector.Vector) {
	dimension := 3

	lookVector, err := vector.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = lookVector.SetCoordinate(0, 1)
	test_helpers.AssertNilError(t, err)
	err = lookVector.SetCoordinate(1, 1)
	test_helpers.AssertNilError(t, err)

	upVector, err := vector.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = upVector.SetCoordinate(0, 1)
	test_helpers.AssertNilError(t, err)
	err = upVector.SetCoordinate(1, -1)
	test_helpers.AssertNilError(t, err)

	rightVector, err := vector.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = rightVector.SetCoordinate(2, -1)
	test_helpers.AssertNilError(t, err)

	return lookVector, upVector, rightVector
}

// TestCamera_Init tests the instantiation of a Camera.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestCamera_Init(t *testing.T) {
	lookVector, upVector, rightVector := buildCameraVectors(t)

	cameraPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	err = cameraPoint.SetCoordinate(0, 3)
	test_helpers.AssertNilError(t, err)
	err = cameraPoint.SetCoordinate(1, 2)
	test_helpers.AssertNilError(t, err)
	err = cameraPoint.SetCoordinate(2, -2)
	test_helpers.AssertNilError(t, err)

	fieldOfView := 50.0
	distanceToScreen := 1.0

	camera, err := Init(cameraPoint, lookVector, upVector, rightVector, fieldOfView, distanceToScreen)
	test_helpers.AssertNilError(t, err)

	vectorController := vector.Controller{}

	normalizedLookVector := vectorController.Normalize(lookVector)
	normalizedUpVector := vectorController.Normalize(upVector)
	normalizedRightVector := vectorController.Normalize(rightVector)
	expectedCamera := Camera{position: cameraPoint, look: normalizedLookVector, up: normalizedUpVector,
		right: normalizedRightVector, fieldOfView: fieldOfView, distanceToScreen: distanceToScreen}
	test_helpers.AssertEqual(t, true, expectedCamera.IsEqual(camera))
}

// TestCamera_Init_Non3DCameraError tests the instantiation of a Camera.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestCamera_Init_Non3DCameraError(t *testing.T) {
	lookVector, upVector, rightVector := buildCameraVectors(t)

	cameraPoint, err := point.Init(2)
	test_helpers.AssertNilError(t, err)

	fieldOfView := 50.0
	distanceToScreen := 1.0

	_, err = Init(cameraPoint, lookVector, upVector, rightVector, fieldOfView, distanceToScreen)
	test_helpers.AssertNotNilError(t, err)
	expectedErrorMessage := fmt.Sprintf("Non 3D camera. Dimensions: position %dD, look %dD, up %dD, right %dD.",
		cameraPoint.Dimension(), lookVector.Dimension(), upVector.Dimension(), rightVector.Dimension())
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
