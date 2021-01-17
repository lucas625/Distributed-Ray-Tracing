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
// 	vertices - a Vertices.
//
// Returns:
//  the corresponding matrix.
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

//// MatrixToVertices is a function to parse a Matrix into a Vertices (a point is a column, removes the homogeneous coord).
////
//// Parameters:
//// 	matrix - a Matrix.
////
//// Returns:
////  the corresponding Vertices.
////
//func MatrixToVertices(matrix *utils.Matrix) Vertices {
//	points := make([]Point, len(matrix.Values[0]))
//	for j := 0; j < len(matrix.Values[0]); j++ {
//		pointAux := InitPoint(len(matrix.Values) - 1)
//		for i := 0; i < len(matrix.Values)-1; i++ {
//			pointAux.Coordinates[i] = matrix.Values[i][j]
//		}
//		points[j] = pointAux
//	}
//	return InitVertices(points)
//}
//
//// MultVertices is a function to multiply all Vertices by a matrix.
////
//// Parameters:
//// 	vertices - a Vertices.
////  matrix   - the multiplying matrix.
////
//// Returns:
//// 	the Vertices multiplied by the matrix.
////
//func MultVertices(vertices *Vertices, matrix *utils.Matrix) Vertices {
//	pointMatrix := VerticesToHomogeneousCoord(vertices)
//	maux := utils.MultMatrix(matrix, &pointMatrix)
//	vertAux := MatrixToVertices(&maux)
//	return vertAux
//}
