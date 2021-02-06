package triangle

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point_repository"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
)

// Controller is a class for controlling triangles.
//
// Members:
// 	none
//
type Controller struct {}

// AreaByTrianglePoints calculates the area of a Triangle by its 3 points.
//
// Parameters:
// 	firstPoint  - The first point of the Triangle.
// 	secondPoint - The second point of the Triangle.
// 	thirdPoint  - The third point of the Triangle.
//
// Returns:
//  The area of the Triangle
//  An error.
//
func (*Controller) AreaByTrianglePoints(firstPoint, secondPoint, thirdPoint *point.Point) (float64, error){
	pointController := point.Controller{}
	vectorFirstToSecondPoint, err := pointController.ExtractVector(firstPoint, secondPoint)
	if err != nil {
		return 0, err
	}

	vectorFirstToThirdPoint, err := pointController.ExtractVector(firstPoint, thirdPoint)
	if err != nil {
		return 0, err
	}

	vectorController := vector.Controller{}

	normalVector, err := vectorController.CrossProduct(vectorFirstToSecondPoint, vectorFirstToThirdPoint)
	if err != nil {
		return 0, err
	}

	triangleArea := vectorController.Norm(normalVector) / 2
	return triangleArea, nil
}

// FindBarycentricCoordinatesByPoint finds the barycentric coordinates of a point based on a Triangle.
// https://www.scratchapixel.com/lessons/3d-basic-rendering/ray-tracing-rendering-a-triangle/barycentric-coordinates
//
// Parameters:
// 	triangle    - The Triangle.
//  targetPoint - The target point.
//  repository  - The point repository.
//
// Returns:
//  The 3 barycentric coordinates.
//  An error.
//
func (controller *Controller) FindBarycentricCoordinatesByPoint(triangle *Triangle, targetPoint *point.Point,
	repository *point_repository.PointRepository) (float64, float64, float64, error) {
	triangleFirstPointIndex, _ := triangle.GetVertexIndex(0)
	triangleSecondPointIndex, _ := triangle.GetVertexIndex(1)
	triangleThirdPointIndex, _ := triangle.GetVertexIndex(2)

	triangleFirstPoint, err := repository.GetPoint(triangleFirstPointIndex)
	if err != nil {
		return 0, 0, 0, err
	}

	triangleSecondPoint, err := repository.GetPoint(triangleSecondPointIndex)
	if err != nil {
		return 0, 0, 0, err
	}

	triangleThirdPoint, err := repository.GetPoint(triangleThirdPointIndex)
	if err != nil {
		return 0, 0, 0, err
	}

	triangleArea, err := controller.AreaByTrianglePoints(triangleFirstPoint, triangleSecondPoint, triangleThirdPoint)
	if err != nil {
		return 0, 0, 0, err
	}

	triangleTargetSecondThirdArea, err := controller.AreaByTrianglePoints(
		targetPoint, triangleSecondPoint, triangleThirdPoint)
	if err != nil {
		return 0, 0, 0, err
	}

	triangleTargetThirdFirstArea, err := controller.AreaByTrianglePoints(
		targetPoint, triangleThirdPoint, triangleFirstPoint)
	if err != nil {
		return 0, 0, 0, err
	}

	alpha := triangleTargetSecondThirdArea / triangleArea
	beta := triangleTargetThirdFirstArea / triangleArea
	gama := (1 - alpha) - beta

	return alpha, beta, gama, nil
}
