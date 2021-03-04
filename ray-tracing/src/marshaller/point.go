package marshaller

import (
	"errors"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
)

// pointController is a class for controlling the marshaller of vectors.
//
// Members:
// 	none
//
type pointController struct {}

// parsePointFromMap parses a point from a map.
//
// Parameters:
//  pointAsMap - The point as a map.
//
// Returns:
// 	The point.
// 	An error.
//
func (*pointController) parsePointFromMap(pointAsMap map[string]interface{}) (*point.Point, error) {
	errorMessage := "invalid point"

	coordinatesAsInterface, found := pointAsMap["coordinates"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	coordinateAsInterfaceList, parsed := coordinatesAsInterface.([]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}

	parsedPoint, err := point.Init(len(coordinateAsInterfaceList))
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	for coordinateIndex := 0; coordinateIndex < parsedPoint.Dimension(); coordinateIndex++ {
		coordinate, parsed := coordinateAsInterfaceList[coordinateIndex].(float64)
		if !parsed {
			return nil, errors.New(errorMessage)
		}
		err = parsedPoint.SetCoordinate(coordinateIndex, coordinate)
		if err != nil {
			return nil, errors.New(errorMessage)
		}
	}

	return parsedPoint, nil
}
