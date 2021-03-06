package marshaller

import (
	"errors"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
)

// parseVectorFromMap parses a vector from a map.
//
// Parameters:
//  vectorAsMap - The vector as a map.
//
// Returns:
// 	The vector.
// 	An error.
//
func (*Controller) parseVectorFromMap(vectorAsMap map[string]interface{}) (*vector.Vector, error) {
	errorMessage := "invalid vector"

	coordinatesAsInterface, found := vectorAsMap["coordinates"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	coordinateAsInterfaceList, parsed := coordinatesAsInterface.([]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}

	parsedVector, err := vector.Init(len(coordinateAsInterfaceList))
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	for coordinateIndex := 0; coordinateIndex < parsedVector.Dimension(); coordinateIndex++ {
		coordinate, parsed := coordinateAsInterfaceList[coordinateIndex].(float64)
		if !parsed {
			return nil, errors.New(errorMessage)
		}
		err = parsedVector.SetCoordinate(coordinateIndex, coordinate)
		if err != nil {
			return nil, errors.New(errorMessage)
		}
	}

	return parsedVector, nil
}
