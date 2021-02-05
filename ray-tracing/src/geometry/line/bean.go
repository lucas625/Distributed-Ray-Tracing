package line

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
)

// Line is a class for lines.
//
// Members:
// 	startingPoint  - The starting point.
// 	vectorDirector - The vector director.
//
type Line struct {
	startingPoint *point.Point
	vectorDirector *vector.Vector
}
