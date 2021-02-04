package plane

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
)

// Controller is a class for the Plane is controller.
//
// Members:
// 	none
//
type Controller struct {}

// extractPlaneNormalVectorFromR3Points is a function to extract a Plane is normal vector from 3 points in the R3 space.
//
// Parameters:
// 	firstPoint  - The first point.
//  secondPoint - The second point.
//  thirdPoint  - The third point.
//
// Returns:
// 	The normal vector.
// 	An error.
//
func (*Controller) extractPlaneNormalVectorFromR3Points(firstPoint, secondPoint, thirdPoint *point.Point) (
	*vector.Vector, error) {
	pointController := point.Controller{}

	if firstPoint.Dimension() != 3 || secondPoint.Dimension() != 3 || thirdPoint.Dimension() != 3{
		return nil, non3DPointsError(firstPoint, secondPoint, thirdPoint)
	}

	firstVector, _ := pointController.ExtractVector(firstPoint, secondPoint)
	secondVector, _ := pointController.ExtractVector(firstPoint, thirdPoint)

	vectorController := vector.Controller{}

	normalVector, _ := vectorController.CrossProduct(firstVector, secondVector)
	return vectorController.Normalize(normalVector), nil
}

// ExtractPlaneFromR3Points is a function to extract a Plane from 3 points in the R3 space.
//
// Parameters:
// 	firstPoint  - The first point.
//  secondPoint - The second point.
//  thirdPoint  - The third point.
//
// Returns:
// 	A Plane.
// 	An error.
//
func (controller *Controller) ExtractPlaneFromR3Points(firstPoint, secondPoint, thirdPoint *point.Point) (
	*Plane, error) {
	normalVector, err := controller.extractPlaneNormalVectorFromR3Points(firstPoint, secondPoint, thirdPoint)

	if err != nil {
		return nil, err
	}

	xCoefficient, _ := normalVector.GetCoordinate(0)
	yCoefficient, _ := normalVector.GetCoordinate(1)
	zCoefficient, _ := normalVector.GetCoordinate(2)

	startingX, _ := firstPoint.GetCoordinate(0)
	startingY, _ := firstPoint.GetCoordinate(1)
	startingZ, _ := firstPoint.GetCoordinate(2)

	isolatedTerm := -1 * ((xCoefficient * startingX) + (yCoefficient * startingY) + (zCoefficient * startingZ))

	return Init(xCoefficient, yCoefficient, zCoefficient, isolatedTerm), nil
}
