package intersector

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/line"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point_repository"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/triangle"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
)

// Controller is a class for controlling intersections.
//
// Members:
// 	none
//
type Controller struct {}

// IntersectLineTriangle is a function to intersect triangles.
// https://www.scratchapixel.com/lessons/3d-basic-rendering/ray-tracing-rendering-a-triangle/moller-trumbore-ray-triangle-intersection
//
// Parameters:
// 	intersectingLine - The line.
//  targetTriangle   - The target triangle.
//  repository       - The point repository.
//
// Returns:
//  The line t parameter (A + tV).
//  The barycentric coordinates at that point.
//  A flag checking if has intersection.
//  An error.
//
func (*Controller) IntersectLineTriangle(intersectingLine *line.Line, targetTriangle *triangle.Triangle,
	repository *point_repository.PointRepository) (float64, []float64, bool, error) {
	rayOrigin := intersectingLine.GetStartingPoint()
	rayVector := intersectingLine.GetVectorDirector()

	EPSILON := 0.0000001

	triangleController := triangle.Controller{}

	points, _ := triangleController.GetActualPoints(targetTriangle, repository)

	pointController := point.Controller{}

	firstEdge, _ := pointController.ExtractVector(points[0], points[1])
	secondEdge, _ := pointController.ExtractVector(points[0], points[2])

	vectorController := vector.Controller{}

	h, err := vectorController.CrossProduct(rayVector, secondEdge)
	if err != nil {
		return 0, nil, false, err
	}

	a, _ := vectorController.DotProduct(firstEdge, h)

	if a > -EPSILON && a < EPSILON {
		return 0, nil, false, nil    // This ray is parallel to this triangle.
	}
	f := 1.0/a

	s, _ := pointController.ExtractVector(points[0], rayOrigin)

	dotProductSH, _ := vectorController.DotProduct(s, h)

	u := f * dotProductSH
	if u < 0.0 || u > 1.0 {
		return 0, nil, false, nil
	}

	q, _ := vectorController.CrossProduct(s, firstEdge)

	dotProductRayVectorQ, _ := vectorController.DotProduct(rayVector, q)

	v := f * dotProductRayVectorQ

	if v < 0.0 || u + v > 1.0 {
		return 0, nil, false, nil
	}

	// At this stage we can compute t to find out where the intersection point is on the line.
	dotProductSecondEdgeQ, _ := vectorController.DotProduct(secondEdge, q)

	t := f * dotProductSecondEdgeQ
	if t > EPSILON && t < 1/EPSILON {
		// outIntersectionPoint = rayOrigin + rayVector * t
		barycentricCoordinates := []float64{1-u-v, u, v}
		return t, barycentricCoordinates, true, nil
	}

	return 0, nil, false, nil
}
