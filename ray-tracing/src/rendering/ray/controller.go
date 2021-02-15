package ray

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/line"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point_repository"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/triangle"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
)

// Controller is a class for controlling ray intersections.
//
// Members:
// 	none
//
type Controller struct {}

// IntersectRayTriangle calculates the intersection between a ray and a triangle.
// https://www.scratchapixel.com/lessons/3d-basic-rendering/ray-tracing-rendering-a-triangle/moller-trumbore-ray-triangle-intersection
//
// Parameters:
// 	ray            - The line.
//  targetTriangle - The target triangle.
//  repository     - The point repository.
//
// Returns:
//  The line t parameter (A + tV).
//  The barycentric coordinates at that point.
//  A flag checking if has intersection.
//  An error.
//
func (*Controller) IntersectRayTriangle(ray *line.Line, targetTriangle *triangle.Triangle,
	repository *point_repository.PointRepository) (float64, []float64, bool, error) {

	if ray.Dimension() != 3 || repository.PointsDimension() != 3 {
		return 0, nil, false, non3DRayPointsError(ray, repository)
	}

	rayOrigin := ray.GetStartingPoint()
	rayVector := ray.GetVectorDirector()

	EPSILON := 0.0000001

	triangleController := triangle.Controller{}

	points, _ := triangleController.GetActualPoints(targetTriangle, repository)

	pointController := point.Controller{}

	firstEdge, _ := pointController.ExtractVector(points[0], points[1])
	secondEdge, _ := pointController.ExtractVector(points[0], points[2])

	vectorController := vector.Controller{}

	h, _ := vectorController.CrossProduct(rayVector, secondEdge)

	a, _ := vectorController.DotProduct(firstEdge, h)

	if a > -EPSILON && a < EPSILON {
		return 0, nil, false, nil    // This ray is parallel to this triangle.
	}
	f := 1.0/a

	s, _ := pointController.ExtractVector(points[0], rayOrigin)

	dotProductSH, _ := vectorController.DotProduct(s, h)

	secondBarycentricCoordinate := f * dotProductSH
	if secondBarycentricCoordinate < 0.0 || secondBarycentricCoordinate > 1.0 {
		return 0, nil, false, nil
	}

	q, _ := vectorController.CrossProduct(s, firstEdge)

	dotProductRayVectorQ, _ := vectorController.DotProduct(rayVector, q)

	thirdBarycentricCoordinate := f * dotProductRayVectorQ
	if thirdBarycentricCoordinate < 0.0 || secondBarycentricCoordinate+thirdBarycentricCoordinate > 1.0 {
		return 0, nil, false, nil
	}

	// At this stage we can compute lineParametricParameter to find out where the intersection point is on the line.
	dotProductSecondEdgeQ, _ := vectorController.DotProduct(secondEdge, q)

	lineParametricParameter := f * dotProductSecondEdgeQ
	if lineParametricParameter > EPSILON && lineParametricParameter < 1/EPSILON {
		firstBarycentricCoordinate := 1.0 - secondBarycentricCoordinate - thirdBarycentricCoordinate
		barycentricCoordinates := []float64{
			firstBarycentricCoordinate, secondBarycentricCoordinate, thirdBarycentricCoordinate}
		return lineParametricParameter, barycentricCoordinates, true, nil
	}

	return 0, nil, false, nil
}
