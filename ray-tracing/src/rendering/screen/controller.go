package screen

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/camera"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/utils/matrix"
	"math"
)

// Controller is a class for controlling screens.
//
// Members:
// 	none
//
type Controller struct {}

// BuildRayVectorDirectorToPixel builds a ray is vector director to a pixel, on world coordinates.
//
// Parameters:
//  pixelLineIndex    - Y position of the pixel.
// 	pixelColumnIndex  - X position of the pixel.
//  pixelLineOffset   - The additional value to the pixel coordinate on y [0,1).
//  pixelColumnOffset - The additional value to the pixel coordinate on x [0,1).
//  cameraToWorld     - The matrix from camera to world.
//  screen            - The Screen that has the pixel.
//  targetCamera      - The camera of the scene.
//
// Returns:
// 	A vector.
//
func (*Controller) BuildRayVectorDirectorToPixel(pixelLineIndex, pixelColumnIndex int, pixelLineOffset,
	pixelColumnOffset float64, cameraToWorld *matrix.Matrix, screen *Screen, targetCamera *camera.Camera) (
	*vector.Vector, error) {
	if pixelLineIndex >= screen.GetHeight() || pixelLineIndex < 0 || pixelColumnIndex >= screen.GetWidth() ||
		pixelColumnIndex < 0 {
		return nil, pixelIndexError(screen, pixelLineIndex, pixelColumnIndex)
	}

	if pixelColumnOffset < 0 || pixelColumnOffset > 1 || pixelLineOffset < 0 || pixelLineOffset > 1 {
		return nil, pixelExtraValueError(pixelLineOffset, pixelColumnOffset)
	}

	aspectRatio := float64(screen.GetWidth()) / float64(screen.GetHeight())
	alpha := (targetCamera.GetFieldOfView() / 2) * math.Pi / 180.0

	vectorDirectorXOnCameraCoordinates := (2*(float64(pixelColumnIndex)+pixelColumnOffset)/
		float64(screen.GetWidth()) - 1) * aspectRatio * math.Tan(alpha)
	vectorDirectorYOnCameraCoordinates := (1 - 2*(float64(pixelLineIndex)+pixelLineOffset)/
		float64(screen.GetHeight())) * math.Tan(alpha)

	vectorDirectorOnCameraCoordinates, _ := vector.Init(3)

	_ = vectorDirectorOnCameraCoordinates.SetCoordinate(0, vectorDirectorXOnCameraCoordinates)
	_ = vectorDirectorOnCameraCoordinates.SetCoordinate(1, vectorDirectorYOnCameraCoordinates)
	_ = vectorDirectorOnCameraCoordinates.SetCoordinate(2, targetCamera.GetDistanceToScreen())

	vectorController := vector.Controller{}
	vectorMatrix := vectorController.ToHomogeneousCoordinates(vectorDirectorOnCameraCoordinates)

	matrixController := matrix.Controller{}
	vectorMatrixOnWorldCoordinates, _ := matrixController.MultiplyMatrix(cameraToWorld, vectorMatrix)

	vectorDirectorOnWorldCoordinates, _ := vector.Init(3)
	for coordinateIndex := 0; coordinateIndex < 3; coordinateIndex++ {
		coordinate, _ := vectorMatrixOnWorldCoordinates.GetValue(coordinateIndex, 0)
		_ = vectorDirectorOnWorldCoordinates.SetCoordinate(coordinateIndex, coordinate)
	}

	return vectorDirectorOnWorldCoordinates, nil
}
