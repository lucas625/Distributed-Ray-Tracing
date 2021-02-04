package plane

import (
	"errors"
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
)

// non3DPointError is the error where a point is not on R3 space.
//
// Parameters:
//	errorPoint - The point raising the error.
//
// Returns:
//  An Error.
//
func non3DPointError(errorPoint *point.Point) error {
	errorMessage := fmt.Sprintf("Point dimension is not 3: %d.", errorPoint.Dimension())
	return errors.New(errorMessage)
}
