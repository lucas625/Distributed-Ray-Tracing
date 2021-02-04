package plane

import (
	"errors"
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
)

// non3DPointsError is the error where not all 3 point are on R3 space.
//
// Parameters:
// 	firstPoint  - The first point.
//  secondPoint - The second point.
//  thirdPoint  - The third point.
//
// Returns:
//  An error.
//
func non3DPointsError(firstPoint, secondPoint, thirdPoint *point.Point) error {
	errorMessage := fmt.Sprintf("Not all points is dimension is equal to 3: %d %d %d.", firstPoint.Dimension(),
		secondPoint.Dimension(), thirdPoint.Dimension())
	return errors.New(errorMessage)
}
