package marshaller

import (
	"errors"
)

// pathTracingParametersController is a class for controlling the marshaller of the path tracing parameters.
//
// Members:
// 	none
//
type pathTracingParametersController struct {}

// pathTracingParameters is a class for lines.
//
// Members:
// 	raysPerPixel      - The number of rays per pixel.
// 	recursions        - The number recursions of each ray.
// 	windowStartLine   - The starting line index of the window of the screen to use the path tracing.
// 	windowStartColumn - The starting column index of the window of the screen to use the path tracing.
// 	windowEndLine     - The ending line index of the window of the screen to use the path tracing.
// 	windowEndColumn   - The ending column index of the window of the screen to use the path tracing.
//
type pathTracingParameters struct {
	raysPerPixel int
	recursions int
	windowStartLine int
	windowStartColumn int
	windowEndLine int
	windowEndColumn int
}

// parsePathTracingParametersFromMap parses a point from a map.
//
// Parameters:
//  pathTracingData - The path tracing data.
//
// Returns:
// 	The path tracing parameters.
// 	An error.
//
func (*pathTracingParametersController) parsePathTracingParametersFromMap(pathTracingData map[string]interface{}) (
	*pathTracingParameters, error) {
	errorMessage := "invalid path tracing parameters"

	pathTracingParametersInterface, found := pathTracingData["pathTracingParameters"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	pathTracingParametersMap, parsed := pathTracingParametersInterface.(map[string]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}

	generalMarshallerController := generalController{}
	raysPerPixel, err := generalMarshallerController.parseFloatFromMap(pathTracingParametersMap, "raysPerPixel")
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	recursions, err := generalMarshallerController.parseFloatFromMap(pathTracingParametersMap, "recursions")
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	windowStartLine, err := generalMarshallerController.parseFloatFromMap(pathTracingParametersMap, "windowStartLine")
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	windowStartColumn, err := generalMarshallerController.parseFloatFromMap(pathTracingParametersMap, "windowStartColumn")
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	windowEndLine, err := generalMarshallerController.parseFloatFromMap(pathTracingParametersMap, "windowEndLine")
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	windowEndColumn, err := generalMarshallerController.parseFloatFromMap(pathTracingParametersMap, "windowEndColumn")
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	return &pathTracingParameters{
		raysPerPixel: int(raysPerPixel), recursions: int(recursions), windowStartLine: int(windowStartLine),
		windowStartColumn: int(windowStartColumn), windowEndLine: int(windowEndLine),
		windowEndColumn: int(windowEndColumn)}, nil
}

