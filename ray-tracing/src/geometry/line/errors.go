package line

import (
	"errors"
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
)

// pointAndVectorIncompatibleDimensionError is the error where a Line has an invalid dimension for the point or vector.
//
// Parameters:
//	startingPoint  - The starting point of the Line.
//	vectorDirector - The vector director of the Line.
//
// Returns:
//  An Error.
//
func pointAndVectorIncompatibleDimensionError(startingPoint *point.Point, vectorDirector *vector.Vector) error {
	errorMessage := fmt.Sprintf(
		"Incompatible dimension for point: %d and vector: %d.", startingPoint.Dimension(), vectorDirector.Dimension())
	return errors.New(errorMessage)
}
