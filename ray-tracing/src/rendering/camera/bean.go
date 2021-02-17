package camera

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
)

// Camera is a class for cameras.
//
// Members:
// 	position         - The position of the Camera.
// 	look             - Vector to were the Camera is looking.
//  up               - Vector head of the Camera.
//  right            - Side vector of the Camera.
//  fieldOfView      - The Camera is field of view in degrees.
//  distanceToScreen - Distance to the screen, also called Near.
//
type Camera struct {
	position         *point.Point
	look             *vector.Vector
	up               *vector.Vector
	right            *vector.Vector
	fieldOfView      float64
	distanceToScreen float64
}

// GetPosition gets the position of the Camera.
//
// Parameters:
// 	none
//
// Returns:
// 	The position of the Camera.
//
func (camera *Camera) GetPosition() *point.Point {
	return camera.position
}

// GetLook gets the look vector of the Camera.
//
// Parameters:
// 	none
//
// Returns:
// 	The look vector of the Camera.
//
func (camera *Camera) GetLook() *vector.Vector {
	return camera.look
}

// GetUp gets the up vector of the Camera.
//
// Parameters:
// 	none
//
// Returns:
// 	The up vector of the Camera.
//
func (camera *Camera) GetUp() *vector.Vector {
	return camera.up
}

// GetRight gets the right vector of the Camera.
//
// Parameters:
// 	none
//
// Returns:
// 	The right vector of the Camera.
//
func (camera *Camera) GetRight() *vector.Vector {
	return camera.right
}

// GetFieldOfView gets the field of view of the Camera.
//
// Parameters:
// 	none
//
// Returns:
// 	The field of view of the Camera.
//
func (camera *Camera) GetFieldOfView() float64 {
	return camera.fieldOfView
}

// GetDistanceToScreen gets the distance to screen of the Camera.
//
// Parameters:
// 	none
//
// Returns:
// 	The distance to screen of the Camera.
//
func (camera *Camera) GetDistanceToScreen() float64 {
	return camera.distanceToScreen
}

// SetLook sets the look vector of the Camera.
//
// Parameters:
// 	look - The new Camera is look vector.
//
// Returns:
// 	none
//
func (camera *Camera) SetLook(look *vector.Vector) {
	camera.look = look
}

// SetUp sets the up vector of the Camera.
//
// Parameters:
// 	up - The new Camera is up vector.
//
// Returns:
// 	none
//
func (camera *Camera) SetUp(up *vector.Vector) {
	camera.up = up
}

// SetRight sets the right vector of the Camera.
//
// Parameters:
// 	right - The new Camera is right vector.
//
// Returns:
// 	none
//
func (camera *Camera) SetRight(right *vector.Vector) {
	camera.right = right
}

// IsEqual checks if a Camera is equal to another.
//
// Parameters:
// 	other - The other Camera.
//
// Returns:
// 	If the cameras are equal.
//
func (camera *Camera) IsEqual(other *Camera) bool {
	return camera.GetPosition().IsEqual(other.GetPosition()) &&
		camera.GetLook().IsEqual(other.GetLook()) &&
		camera.GetUp().IsEqual(other.GetUp()) &&
		camera.GetRight().IsEqual(other.GetRight()) &&
		camera.GetFieldOfView() == other.GetFieldOfView() &&
		camera.GetDistanceToScreen() == other.GetDistanceToScreen()
}

// normalizeVectors normalizes the Camera vectors.
//
// Parameters:
//  none
//
// Returns:
// 	none
//
func (camera *Camera) normalizeVectors() {
	vectorController := vector.Controller{}
	camera.SetLook(vectorController.Normalize(camera.look))
	camera.SetUp(vectorController.Normalize(camera.up))
	camera.SetRight(vectorController.Normalize(camera.right))
}

// Init initializes a Camera.
//
// Parameters:
// 	position         - The position of the Camera.
// 	look             - Vector to were the Camera is looking.
//  up               - Vector head of the Camera.
//  right            - Side vector of the Camera.
//  fieldOfView      - The Camera is field of view in degrees.
//  distanceToScreen - Distance to the screen, also called Near.
//
// Returns:
// 	A Camera.
//  An error.
//
func Init(position *point.Point, look, up, right *vector.Vector, fieldOfView, distanceToScreen float64) (
	*Camera, error) {
	if position.Dimension() != 3 || look.Dimension() != 3 || up.Dimension() != 3 || right.Dimension() != 3 {
		return nil, non3DCameraError(position, look, up, right)
	}
	camera := &Camera{position: position, look: look, up: up, right: right, fieldOfView: fieldOfView,
		distanceToScreen: distanceToScreen}
	camera.normalizeVectors()
	return camera, nil
}
