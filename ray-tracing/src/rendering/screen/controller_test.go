package screen

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/camera"
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

// TestController_BuildRayVectorDirectorToPixel tests the build of a ray is vector director to a pixel, on world
// coordinates.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_BuildRayVectorDirectorToPixel(t *testing.T) {
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

	screenCamera, err := camera.Init(cameraPoint, lookVector, upVector, rightVector, fieldOfView, distanceToScreen)
	test_helpers.AssertNilError(t, err)

	cameraController := camera.Controller{}
	cameraToWorldMatrix := cameraController.CameraToWorldMatrix(screenCamera)

	screen, err := Init(5, 5)
	test_helpers.AssertNilError(t, err)

	screenController := Controller{}
	rayVectorDirector, err := screenController.BuildRayVectorDirectorToPixel(
		2, 2, 0.5, 0.5, cameraToWorldMatrix, screen, screenCamera)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, true, rayVectorDirector != nil)
}

// TestController_BuildRayVectorDirectorToPixel_PixelIndexError tests the build of a ray is vector director to a pixel,
// on world coordinates when we try to access a pixel out of the limits of the Screen.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_BuildRayVectorDirectorToPixel_PixelIndexError(t *testing.T) {
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

	screenCamera, err := camera.Init(cameraPoint, lookVector, upVector, rightVector, fieldOfView, distanceToScreen)
	test_helpers.AssertNilError(t, err)

	cameraController := camera.Controller{}
	cameraToWorldMatrix := cameraController.CameraToWorldMatrix(screenCamera)

	screen, err := Init(5, 5)
	test_helpers.AssertNilError(t, err)

	screenController := Controller{}
	_, err = screenController.BuildRayVectorDirectorToPixel(
		5, 2, 0.5, 0.5, cameraToWorldMatrix, screen, screenCamera)
	test_helpers.AssertNotNilError(t, err)
	expectedErrorMessage := fmt.Sprintf(
		"Pixel out of limits of the screen. Expected from 0 0 to %v %v, and got %v %v.",
		screen.GetHeight(), screen.GetWidth(), 5, 2)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestController_BuildRayVectorDirectorToPixel_PixelExtraValueError tests the build of a ray is vector director to a
// pixel, on world coordinates when we try to add an invalid value to a pixel.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_BuildRayVectorDirectorToPixel_PixelExtraValueError(t *testing.T) {
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

	screenCamera, err := camera.Init(cameraPoint, lookVector, upVector, rightVector, fieldOfView, distanceToScreen)
	test_helpers.AssertNilError(t, err)

	cameraController := camera.Controller{}
	cameraToWorldMatrix := cameraController.CameraToWorldMatrix(screenCamera)

	screen, err := Init(5, 5)
	test_helpers.AssertNilError(t, err)

	screenController := Controller{}
	_, err = screenController.BuildRayVectorDirectorToPixel(
		2, 2, -1, -1, cameraToWorldMatrix, screen, screenCamera)
	test_helpers.AssertNotNilError(t, err)
	expectedErrorMessage := fmt.Sprintf("Pixel extra value error. Expected [0,1], and got %v %v.",
		-1, -1)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
