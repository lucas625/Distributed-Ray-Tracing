package line

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
)

// Controller is a class for controlling lines.
//
// Members:
// 	none
//
type Controller struct {}

// ExtractLine is a function to extract a line from 2 points.
//
// Parameters:
// 	startingPoint - The starting point.
//  targetPoint   - The target point.
//
// Returns:
// 	A Line.
//  An error.
//
func (*Controller) ExtractLine(startingPoint, targetPoint *point.Point) (*Line, error) {
	pointController := point.Controller{}
	vectorDirector, err := pointController.ExtractVector(startingPoint, targetPoint)
	if err != nil {
		return nil, err
	}
	return Init(startingPoint, vectorDirector)
}

// FindPoint calculates the point on a line at a given parametric parameter.
//
// Parameters:
// 	line                - The Line.
// 	parametricParameter - The parametric parameter of the Line.
//
// Returns:
// 	The Point.
//  An error.
//
func (*Controller) FindPoint(line *Line, parametricParameter float64) (*point.Point, error) {
	vectorController := vector.Controller{}
	multipliedVector := vectorController.ScalarMultiplication(line.GetVectorDirector(), parametricParameter)
	pointController := point.Controller{}
	return pointController.SumWithVector(line.GetStartingPoint(), multipliedVector)
}
