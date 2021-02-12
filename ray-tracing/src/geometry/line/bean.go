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

// GetStartingPoint gets the starting point of the Line.
//
// Parameters:
// 	none
//
// Returns:
// 	The starting point of the Line.
//
func (line *Line) GetStartingPoint() *point.Point {
	return line.startingPoint
}

// GetVectorDirector gets the vector director of the Line.
//
// Parameters:
// 	none
//
// Returns:
// 	The vector director of the Line.
//
func (line *Line) GetVectorDirector() *vector.Vector {
	return line.vectorDirector
}

// IsEqual checks if a Line object is equal to another.
//
// Parameters:
// 	other - The other Line.
//
// Returns:
// 	If the lines are equal.
//
func (line *Line) IsEqual(other *Line) bool {
	return line.GetVectorDirector().IsEqual(other.GetVectorDirector()) &&
		line.GetStartingPoint().IsEqual(other.GetStartingPoint())
}

// Init checks if a Line object is equal to another.
//
// Parameters:
// 	startingPoint  - The starting point.
// 	vectorDirector - The vector director.
//
// Returns:
// 	A Line.
// 	An error.
//
func Init(startingPoint *point.Point, vectorDirector *vector.Vector) (*Line, error) {
	if startingPoint.Dimension() != vectorDirector.Dimension() {
		return nil, pointAndVectorIncompatibleDimensionError(startingPoint, vectorDirector)
	}
	return &Line{startingPoint: startingPoint, vectorDirector: vectorDirector}, nil
}
