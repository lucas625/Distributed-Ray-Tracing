package camera

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/utils/matrix"
)

// Controller is a class for controlling cameras.
//
// Members:
// 	none
//
type Controller struct {}

// WorldToCameraMatrix creates the homogeneous coordinates matrix that parses from world coordinates to Camera
// coordinates.
//
// Parameters:
// 	camera - A Camera.
//
// Returns:
// 	A matrix.
//
func (*Controller) WorldToCameraMatrix(camera *Camera) *matrix.Matrix {
	matrixController := matrix.Controller{}
	worldToCameraMatrix, _ := matrixController.BuildHomogeneousCoordinates(3)
	
	for coordinateIndex := 0; coordinateIndex < 3; coordinateIndex++ {
		lookCoordinate, _ := camera.GetLook().GetCoordinate(coordinateIndex)
		upCoordinate, _ := camera.GetUp().GetCoordinate(coordinateIndex)
		rightCoordinate, _ := camera.GetRight().GetCoordinate(coordinateIndex)
		positionCoordinate, _ := camera.GetPosition().GetCoordinate(coordinateIndex)
		_ = worldToCameraMatrix.SetValue(0, coordinateIndex, rightCoordinate)
		_ = worldToCameraMatrix.SetValue(1, coordinateIndex, upCoordinate)
		_ = worldToCameraMatrix.SetValue(2, coordinateIndex, lookCoordinate)
		_ = worldToCameraMatrix.SetValue(coordinateIndex, 3, -1 * positionCoordinate)
	}
	return worldToCameraMatrix
}

// CameraToWorldMatrix creates the homogeneous coordinates matrix that parses from Camera coordinates to world
// coordinates.
//
// Parameters:
// 	camera - A Camera.
//
// Returns:
// 	a Matrix.
//
func (*Controller) CameraToWorldMatrix(camera *Camera) *matrix.Matrix {
	matrixController := matrix.Controller{}
	cameraToWorldMatrix, _ := matrixController.BuildHomogeneousCoordinates(3)

	for coordinateIndex := 0; coordinateIndex < 3; coordinateIndex++ {
		lookCoordinate, _ := camera.GetLook().GetCoordinate(coordinateIndex)
		upCoordinate, _ := camera.GetUp().GetCoordinate(coordinateIndex)
		rightCoordinate, _ := camera.GetRight().GetCoordinate(coordinateIndex)
		positionCoordinate, _ := camera.GetPosition().GetCoordinate(coordinateIndex)
		_ = cameraToWorldMatrix.SetValue(coordinateIndex,0, rightCoordinate)
		_ = cameraToWorldMatrix.SetValue(coordinateIndex,1, upCoordinate)
		_ = cameraToWorldMatrix.SetValue(coordinateIndex,2, lookCoordinate)
		_ = cameraToWorldMatrix.SetValue(3, coordinateIndex, positionCoordinate)
	}
	return cameraToWorldMatrix
}
