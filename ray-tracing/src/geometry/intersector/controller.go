package intersector

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/line"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point_repository"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/triangle"
)

// Controller is a class for controlling intersections.
//
// Members:
// 	none
//
type Controller struct {}

// IntersectLineTriangle is a function to intersect triangles.
//
// Parameters:
// 	intersectingLine - The line.
//  targetTriangle   - The target triangle.
//
// Returns:
//  The line t parameter (A + tV).
//  The barycentric coordinates at that point.
//  A flag checking if has intersection.
//
func (*Controller) IntersectLineTriangle(intersectingLine *line.Line, targetTriangle *triangle.Triangle,
	points *point_repository.PointRepository) (float64, []float64, bool) {
	//rayOrigin := line.Start
	//rayVector := line.Director
	//
	//EPSILON := 0.0000001;
	//
	//vertex0 := triang[0]
	//vertex1 := triang[1]
	//vertex2 := triang[2]


	bCoord := make([]float64, 3)
	//
	//
	//edge1 := ExtractVector(&vertex0, &vertex1)
	//edge2 := ExtractVector(&vertex0, &vertex2)
	//
	//h := utils.VectorCrossProduct(&rayVector, &edge2)
	//a := utils.DotProduct(&edge1, &h);
	//
	//if a > -EPSILON && a < EPSILON {
	//	return 0,bCoord, false;    // This ray is parallel to this triangle.
	//}
	//f := 1.0/a
	//
	//s := ExtractVector(&vertex0,&rayOrigin)
	//u := f * utils.DotProduct(&s, &h)
	//if u < 0.0 || u > 1.0 {
	//	return 0,bCoord,false
	//}
	//q := utils.VectorCrossProduct(&s,&edge1);
	//v := f * utils.DotProduct(&rayVector,&q);
	//if v < 0.0 || u + v > 1.0 {
	//	return 0,bCoord,false
	//}
	//// At this stage we can compute t to find out where the intersection point is on the line.
	//t := f * utils.DotProduct(&edge2, &q)
	//if t > EPSILON && t < 1/EPSILON {
	//	// outIntersectionPoint = rayOrigin + rayVector * t
	//	bCoord = []float64{1-u-v, u, v}
	//	return t,bCoord,true
	//}

	return 0,bCoord,false
}
