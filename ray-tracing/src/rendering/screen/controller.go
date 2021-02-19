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
//  pixelLineIndex   - Y position of the pixel.
// 	pixelColumnIndex - X position of the pixel.
//  additionalY      - The additional value to the pixel coordinate on y [0,1).
//  additionalX      - The additional value to the pixel coordinate on x [0,1).
//  cameraToWorld    - The matrix from camera to world.
//  screen           - The Screen that has the pixel.
//  targetCamera     - The camera of the scene.
//
// Returns:
// 	a Vector.
//
func (*Controller) BuildRayVectorDirectorToPixel(pixelLineIndex, pixelColumnIndex int, additionalY, additionalX float64,
	cameraToWorld *matrix.Matrix, screen *Screen, targetCamera *camera.Camera) (
	*vector.Vector, error) {
	if pixelLineIndex >= screen.GetHeight() || pixelLineIndex < 0 || pixelColumnIndex >= screen.GetWidth() ||
		pixelColumnIndex < 0 {
		return nil, pixelIndexError(screen, pixelLineIndex, pixelColumnIndex)
	}

	if additionalX < 0 || additionalX >= 1 || additionalY < 0 || additionalY >= 1 {
		return nil, pixelExtraValueError(additionalY, additionalX)
	}

	aspectRatio := float64(screen.GetWidth()) / float64(screen.GetHeight())
	alpha := (targetCamera.GetFieldOfView() / 2) * math.Pi / 180.0

	vectorDirectorXOnCameraCoordinates := (2*(float64(pixelColumnIndex)+additionalX)/float64(screen.GetWidth()) - 1) *
		aspectRatio * math.Tan(alpha)
	vectorDirectorYOnCameraCoordinates := (1 - 2*(float64(pixelLineIndex)+additionalY)/float64(screen.GetHeight())) *
		math.Tan(alpha)

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
	normalizedVectorDirector := vectorController.Normalize(vectorDirectorOnWorldCoordinates)

	return normalizedVectorDirector, nil
}
