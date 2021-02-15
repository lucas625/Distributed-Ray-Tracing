package object

import "github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"

// Controller is a class for controlling objects.
//
// Members:
// 	none
//
type Controller struct {}

// GetCenter is a function to get the center of the bounding box of the object.
//
// Parameters:
// 	object - The Object.
//
// Returns:
//  The center point of the Object.
//
func (controller *Controller) GetCenter(object *Object) *point.Point {
	boundingBox := controller.GetBoundingBox(object)
	centerPoint, _ := point.Init(object.GetRepository().PointsDimension())
	for coordinateIndex := 0; coordinateIndex < centerPoint.Dimension(); coordinateIndex++ {
		centerCoordinateValue := (
			boundingBox[coordinateIndex] + boundingBox[coordinateIndex+centerPoint.Dimension()]) / 2
		_ = centerPoint.SetCoordinate(coordinateIndex, centerCoordinateValue)
	}
	return centerPoint
}

// GetBoundingBox is a function to get the bounding box of an Object.
//
// Parameters:
//  object - The Object.
//
// Returns:
//  [minX, minY, minZ, ..., maxX, maxY, maxZ, ...]
//
func (*Controller) GetBoundingBox(object *Object) []float64 {
	boundingBox := make([]float64, 6)
	repository := object.GetRepository()
	fistPoint, _ := repository.GetPoint(0)

	for coordinateIndex := 0; coordinateIndex < repository.PointsDimension(); coordinateIndex++ {
		coordinateValue, _ := fistPoint.GetCoordinate(coordinateIndex)
		boundingBox[coordinateIndex] = coordinateValue
		boundingBox[coordinateIndex+repository.PointsDimension()] = coordinateValue
	}
	for pointIndex := 0; pointIndex < repository.PointsDimension(); pointIndex++ {
		currentPoint, _ := repository.GetPoint(pointIndex)
		for coordinateIndex := 0; coordinateIndex < repository.PointsDimension(); coordinateIndex++ {
			coordinateValue, _ := currentPoint.GetCoordinate(coordinateIndex)
			if boundingBox[coordinateIndex] > coordinateValue {
				boundingBox[coordinateIndex] = coordinateValue
			}
			if boundingBox[coordinateIndex+repository.PointsDimension()] < coordinateValue {
				boundingBox[coordinateIndex+repository.PointsDimension()] = coordinateValue
			}
		}
	}
	return boundingBox
}
