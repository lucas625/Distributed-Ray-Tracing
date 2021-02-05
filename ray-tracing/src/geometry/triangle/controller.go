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

// ADD HERE A FUNCTION TO CALCULATE A TRIANGLE AREA

// FindBarycentricCoordinatesByPoint is a function to find the barycentric coordinates of a point based on a Triangle.
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
func (*Controller) FindBarycentricCoordinatesByPoint(triangle *Triangle, targetPoint *point.Point,
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

	pointController := point.Controller{}

	vectorFirstToSecondPoint := pointController.ExtractVector(triangleFirstPoint, triangleSecondPoint)
	vectorFirstToThirdPoint := pointController.ExtractVector(triangleFirstPoint, triangleThirdPoint)

	vectorTargetToFirstPoint := pointController.ExtractVector(targetPoint, triangleFirstPoint)
	vectorTargetToSecondPoint := pointController.ExtractVector(targetPoint, triangleSecondPoint)
	ectorTargetToThirdPoint := pointController.ExtractVector(targetPoint, triangleThirdPoint)

	normal := utils.VectorCrossProduct(&AB, &AC)
	AreaABC := utils.VectorNorm(&normal) / 2

	normalA := utils.VectorCrossProduct(&PB, &PC)
	normalB := utils.VectorCrossProduct(&PC, &PA)

	AreaA := utils.VectorNorm(&normalA) / 2
	AreaB := utils.VectorNorm(&normalB) / 2

	alpha := AreaA / AreaABC
	beta := AreaB / AreaABC
	gama := (1 - alpha) - beta

	coordinates := []float64{alpha, beta, gama}
	return coordinates
}
