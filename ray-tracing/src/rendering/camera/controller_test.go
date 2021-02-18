package camera

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/utils/matrix"
	"math"
	"testing"
)

// TestController_WorldToCameraMatrix tests the build of the world to Camera matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_WorldToCameraMatrix(t *testing.T) {
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

	cameraController := Controller{}
	worldToCameraMatrix := cameraController.WorldToCameraMatrix(camera)

	matrixController := matrix.Controller{}
	expectedMatrix, _ := matrixController.BuildHomogeneousCoordinates(3)

	err = expectedMatrix.SetValue(0, 0, 0)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(0, 1, 0)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(0, 2, -1)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(0, 3, -3)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(1, 0,  math.Sqrt(2) / 2)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(1, 1,  (-1 * math.Sqrt(2)) / 2)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(1, 2,  0)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(1, 3,  -2)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(2, 0,  math.Sqrt(2) / 2)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(2, 1,  math.Sqrt(2) / 2)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(2, 2,  0)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(2, 3,  2)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(3, 0,  0)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(3, 1,  0)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(3, 2,  0)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(3, 3,  1)
	test_helpers.AssertNilError(t, err)

	test_helpers.AssertEqual(t, true, expectedMatrix.IsEqual(worldToCameraMatrix))
}

// TestController_CameraToWorldMatrix tests the build of the Camera to world matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_CameraToWorldMatrix(t *testing.T) {
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

	cameraController := Controller{}
	cameraToWorldMatrix := cameraController.CameraToWorldMatrix(camera)

	matrixController := matrix.Controller{}
	expectedMatrix, _ := matrixController.BuildHomogeneousCoordinates(3)

	err = expectedMatrix.SetValue(0, 0, 0)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(1, 0, 0)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(2, 0, -1)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(3, 0,  0)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(0, 1,  math.Sqrt(2) / 2)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(1, 1,  (-1 * math.Sqrt(2)) / 2)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(2, 1,  0)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(3, 1,  0)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(0, 2,  math.Sqrt(2) / 2)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(1, 2,  math.Sqrt(2) / 2)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(2, 2,  0)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(3, 2,  0)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(0, 3, 3)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(1, 3,  2)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(2, 3,  -2)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(3, 3,  1)
	test_helpers.AssertNilError(t, err)

	test_helpers.AssertEqual(t, true, expectedMatrix.IsEqual(cameraToWorldMatrix))
}
