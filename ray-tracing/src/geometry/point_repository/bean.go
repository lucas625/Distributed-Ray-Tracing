package point_repository

import "github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"

// PointRepository is a class for a list of points.
//
// Members:
// 	points - List of points.
//
type PointRepository struct {
	points []*point.Point
}

// GetPoint gets a point of the PointRepository.
//
// Parameters:
// 	index - The index of the point.
//
// Returns:
// 	The point.
//  An error.
//
func (pointRepository *PointRepository) GetPoint(index int) (*point.Point, error) {
	if index < 0 || index >= pointRepository.NumberOfPoints() {
		return nil, indexError(pointRepository, index)
	}
	return pointRepository.points[index], nil
}

// NumberOfPoints gets the number of points on PointRepository.
//
// Parameters:
// 	none
//
// Returns:
// 	The number of points.
//
func (pointRepository *PointRepository) NumberOfPoints() int {
	return len(pointRepository.points)
}

// NumberOfPoints gets the number of points on PointRepository.
//
// Parameters:
// 	none
//
// Returns:
// 	The number of points.
//
func (pointRepository *PointRepository) PointsDimension() int {
	return pointRepository.points[0].Dimension()
}

// Init initializes a PointRepository.
//
// Parameters:
// 	points    - List of points.
//  dimension - The expected dimension for all points.
//
// Returns:
// 	The PointRepository.
//  An Error.
//
func Init(points []*point.Point, dimension int) (*PointRepository, error) {
	if len(points) == 0 {
		return nil, invalidSizeError(points)
	}
	for index := 0; index < len(points); index++ {
		if points[index].Dimension() != dimension {
			return nil, incompatibleDimensionError(points, dimension)
		}
	}
	return &PointRepository{points: points}, nil
}
