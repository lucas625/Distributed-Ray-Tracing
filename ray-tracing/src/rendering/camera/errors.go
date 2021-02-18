package camera

import (
	"errors"
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
)

// non3DCameraError is the error where the camera is not on the third dimension.
//
// Parameters:
// 	position - The position of the Camera.
// 	look     - Vector to were the Camera is looking.
//  up       - Vector head of the Camera.
//  right    - Side vector of the Camera.
//
// Returns:
//  An Error.
//
func non3DCameraError(position *point.Point, look, up, right *vector.Vector) error {
	errorMessage := fmt.Sprintf("Non 3D camera. Dimensions: position %dD, look %dD, up %dD, right %dD.",
		position.Dimension(), look.Dimension(), up.Dimension(), right.Dimension())
	return errors.New(errorMessage)
}
