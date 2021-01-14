package point

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/utils/matrix"
)

// Controller is a class for controlling points.
//
// Members:
// 	none
//
type Controller struct {}

// ExtractVector is a function to extract a vector between two points.
//
// Parameters:
// 	startingPoint - The starting Point.
//  targetPoint   - The target Point.
//
// Returns:
// 	A Vector.
//  An Error.
//
func (*Controller) ExtractVector(startingPoint, targetPoint *Point) (*vector.Vector, error) {
	if startingPoint.Dimension() != targetPoint.Dimension() {
		return nil, differentDimensionsError(startingPoint, targetPoint)
	}
	extractedVector, _ := vector.Init(startingPoint.Dimension())
	for index := 0; index < startingPoint.Dimension(); index++ {
		startingPointCoordinate, _ := startingPoint.GetCoordinate(index)
		targetPointCoordinate, _ := targetPoint.GetCoordinate(index)
		extractedVector.SetCoordinate(index, targetPointCoordinate - startingPointCoordinate)
	}
	return extractedVector, nil
}

// ToHomogeneousCoordinates adds the extra 1 coordinate and converts the Point to a 1 column Matrix.
//
// Parameters:
// 	point - The Point.
//
// Returns:
// 	a Matrix.
//
func (*Controller) ToHomogeneousCoordinates(point *Point) *matrix.Matrix {
	pointAsMatrix, _ := matrix.Init(point.Dimension() + 1, 1)
	for index := 0; index < point.Dimension(); index++ {
		pointCoordinate, _ := point.GetCoordinate(index)
		pointAsMatrix.SetValue(index, 0, pointCoordinate)
	}
	pointAsMatrix.SetValue(point.Dimension(), 0, 1)
	return pointAsMatrix
}
