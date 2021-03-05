package marshaller

import (
	"errors"
)

// parseColorFromMap parses a color from a map.
//
// Parameters:
//  mapContainingColor - The map that contains the color.
//
// Returns:
// 	The float.
// 	An error.
//
func (*Controller) parseColorFromMap(mapContainingColor map[string]interface{}) ([]float64, error) {
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
func (*Controller) parseFloatFromMap(mapContainingFloat map[string]interface{}, floatName string) (float64, error) {
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

// parseStringFromMap parses a string from a map.
//
// Parameters:
//  mapContainingString - The map that contains the string.
//  stringName          - The name of the string.
//
// Returns:
// 	The string.
// 	An error.
//
func (*Controller) parseStringFromMap(mapContainingString map[string]interface{}, stringName string) (string, error) {
	errorMessage := "unable to parse string"

	stringInterface, found := mapContainingString[stringName]
	if !found {
		return "", errors.New(errorMessage)
	}
	stringParsed, parsed := stringInterface.(string)
	if !parsed {
		return "", errors.New(errorMessage)
	}

	return stringParsed, nil
}
