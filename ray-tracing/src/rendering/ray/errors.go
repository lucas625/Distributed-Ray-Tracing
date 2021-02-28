package ray

import (
	"errors"
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/line"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/point_repository"
)

// non3DRayPointsError is the error where a ray or the points are not on the third dimension.
//
// Parameters:
//	ray        - The ray.
//	repository - The repository of points.
//
// Returns:
//  An Error.
//
func non3DRayPointsError(ray *line.Line, repository *point_repository.PointRepository) error {
	errorMessage := fmt.Sprintf(
		"Non 3D ray or repository. Ray dimension: %d and point repository dimension: %d.",
		ray.Dimension(), repository.PointsDimension())
	return errors.New(errorMessage)
}
