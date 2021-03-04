package marshaller

import (
	"errors"
)

// generalController is a class for controlling the marshaller of the general types.
//
// Members:
// 	none
//
type generalController struct {}

// parseColorFromMap parses a color from a map.
//
// Parameters:
//  mapContainingColor - The map that contains the color.
//
// Returns:
// 	The float.
// 	An error.
//
func (*generalController) parseColorFromMap(mapContainingColor map[string]interface{}) ([]float64, error) {
	errorMessage := "unable to parse color"

	colorInterface, found := mapContainingColor["color"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	colorInterfaceList, parsed := colorInterface.([]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}

	if len(colorInterfaceList) != 3 {
		return nil, errors.New(errorMessage)
	}

	color := make([]float64, 3)
	for index := 0; index < 3; index++ {
		coordinate, parsed := colorInterfaceList[index].(float64)
		if !parsed {
			return nil, errors.New(errorMessage)
		}
		color[index] = coordinate
	}

	return color, nil
}

// parseFloatFromMap parses a float from a map.
//
// Parameters:
//  mapContainingFloat - The map that contains the float.
//  floatName          - The name of the float.
//
// Returns:
// 	The float.
// 	An error.
//
func (*generalController) parseFloatFromMap(mapContainingFloat map[string]interface{}, floatName string) (
	float64, error) {
	errorMessage := "unable to parse float"

	floatInterface, found := mapContainingFloat[floatName]
	if !found {
		return 0, errors.New(errorMessage)
	}
	floatParsed, parsed := floatInterface.(float64)
	if !parsed {
		return 0, errors.New(errorMessage)
	}

	return floatParsed, nil
}
