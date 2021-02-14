package triangle

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point_repository"
)

// Controller is a class for controlling triangles.
//
// Members:
// 	none
//
type Controller struct {}

// getActualPoint gets the actual point of a Triangle.
//
// Parameters:
// 	triangle   - The target Triangle.
// 	repository - The point repository.
// 	index      - The index of the vertex on the Triangle vertices indexes.
//
// Returns:
// 	The point.
//  An error.
//
func (*Controller) getActualPoint(triangle *Triangle, repository *point_repository.PointRepository, index int) (
	*point.Point, error) {
	pointIndex, err := triangle.GetVertexIndex(index)
	if err != nil {
		return nil, err
	}
	return repository.GetPoint(pointIndex)
}

// GetActualPoints gets the actual points of a Triangle.
//
// Parameters:
// 	triangle   - The target Triangle.
// 	repository - The point repository.
//
// Returns:
// 	The points.
//  An error.
//
func (controller *Controller) GetActualPoints(triangle *Triangle, repository *point_repository.PointRepository) (
	[]*point.Point, error) {
	firstPoint, err := controller.getActualPoint(triangle, repository, 0)
	if err != nil {
		return nil, err
	}

	secondPoint, err := controller.getActualPoint(triangle, repository, 1)
	if err != nil {
		return nil, err
	}

	thirdPoint, err := controller.getActualPoint(triangle, repository, 2)
	if err != nil {
		return nil, err
	}

	return []*point.Point{firstPoint, secondPoint, thirdPoint}, nil
}
