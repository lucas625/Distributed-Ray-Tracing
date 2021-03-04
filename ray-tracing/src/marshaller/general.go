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