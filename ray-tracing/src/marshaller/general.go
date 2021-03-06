package marshaller

import (
	"errors"
)

// parseFloatListFromMap parses a float list from a map.
//
// Parameters:
//  mapContainingFloats - The map that contains the floats.
//  listName            - The name of the list of floats.
//
// Returns:
// 	The float list.
// 	An error.
//
func (*Controller) parseFloatListFromMap(mapContainingFloats map[string]interface{}, listName string) (
	[]float64, error) {
	errorMessage := "unable to parse floats"

	floatListInterface, found := mapContainingFloats[listName]
	if !found {
		return nil, errors.New(errorMessage)
	}
	floatListInterfaceList, parsed := floatListInterface.([]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}

	floats := make([]float64, len(floatListInterfaceList))
	for index := 0; index < len(floatListInterfaceList); index++ {
		currentFloat, parsed := floatListInterfaceList[index].(float64)
		if !parsed {
			return nil, errors.New(errorMessage)
		}
		floats[index] = currentFloat
	}

	return floats, nil
}

// parseIntListFromMap parses a int list from a map.
//
// Parameters:
//  mapContainingIntegers - The map that contains the integers.
//  listName              - The name of the list of integers.
//
// Returns:
// 	The int list.
// 	An error.
//
func (*Controller) parseIntListFromMap(mapContainingIntegers map[string]interface{}, listName string) ([]int, error) {
	errorMessage := "unable to parse floats"

	intListInterface, found := mapContainingIntegers[listName]
	if !found {
		return nil, errors.New(errorMessage)
	}
	intListInterfaceList, parsed := intListInterface.([]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}

	integers := make([]int, len(intListInterfaceList))
	for index := 0; index < len(intListInterfaceList); index++ {
		currentFloat, parsed := intListInterfaceList[index].(float64)
		if !parsed {
			return nil, errors.New(errorMessage)
		}
		integers[index] = int(currentFloat)
	}

	return integers, nil
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
