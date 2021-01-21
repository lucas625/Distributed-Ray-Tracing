package point_repository

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/utils/matrix"
)

// Controller is a class for controlling pointRepositories.
//
// Members:
// 	none
//
type Controller struct {}

// ToHomogeneousCoordinates creates a matrix with all points in homogeneous coordinates.
//
// Parameters:
// 	pointRepository - The target points repository.
//
// Returns:
//  The corresponding matrix.
//
func (*Controller) ToHomogeneousCoordinates(pointRepository *PointRepository) *matrix.Matrix {
	homogenousCoordinatesPointsMatrix, _ := matrix.Init(pointRepository.PointsDimension() + 1,
		pointRepository.NumberOfPoints())
	pointController := point.Controller{}

	for index := 0; index < pointRepository.NumberOfPoints(); index++ {
		currentPoint, _ := pointRepository.GetPoint(index)
		pointAsHomogeneousCoordinatesMatrix := pointController.ToHomogeneousCoordinates(currentPoint)
		for pointCoordinateIndex := 0; pointCoordinateIndex < pointAsHomogeneousCoordinatesMatrix.Lines();
		pointCoordinateIndex++ {
			coordinateValue, _ := pointAsHomogeneousCoordinatesMatrix.GetValue(pointCoordinateIndex, 0)
			homogenousCoordinatesPointsMatrix.SetValue(pointCoordinateIndex, index, coordinateValue)
		}
	}
	return homogenousCoordinatesPointsMatrix
}

// FromMatrix parses e a Matrix into a PointRepository (a point is a column, removes the homogeneous coordinate).
//
// Parameters:
// 	matrix - A matrix in homogeneous coordinates.
//
// Returns:
//  The points on the matrix as a PointRepository.
//  An error.
//
func (*Controller) FromMatrix(pointsAsMatrix *matrix.Matrix) (*PointRepository, error) {
	points := make([]*point.Point, pointsAsMatrix.Columns())
	dimension := pointsAsMatrix.Lines() - 1

	for pointIndex := 0; pointIndex < pointsAsMatrix.Columns(); pointIndex++ {
		currentPoint, _ := point.Init(dimension)
		for coordinateIndex := 0; coordinateIndex < dimension; coordinateIndex++ {
			coordinate, _ := pointsAsMatrix.GetValue(coordinateIndex, pointIndex)
			currentPoint.SetCoordinate(coordinateIndex, coordinate)
		}
		points[pointIndex] = currentPoint
	}
	return Init(points, dimension)
}

// MultiplyByMatrix multiplies all points on the PointRepository by a matrix.
//
// Parameters:
// 	pointRepository   - The PointRepository..
//  multiplyingMatrix - The multiplying matrix.
//
// Returns:
// 	The multiplied points as a PointRepository.
//
func (controller *Controller) MultiplyByMatrix(pointRepository *PointRepository, multiplyingMatrix *matrix.Matrix) (
	*PointRepository, error) {
	pointMatrix := controller.ToHomogeneousCoordinates(pointRepository)
	matrixController := matrix.Controller{}
	resultingMatrix, err := matrixController.MultiplyMatrix(multiplyingMatrix, pointMatrix)
	if err != nil {
		return nil, err
	}
	return controller.FromMatrix(resultingMatrix)
}
