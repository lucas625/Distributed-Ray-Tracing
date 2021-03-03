package marshaller

import (
	"errors"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/path_tracing"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/screen"
)

// Controller is a class for controlling the marshaller of the application.
//
// Members:
// 	none
//
type Controller struct {}

// parsePixelScreenFromMap parses the pixel screen from a map.
//
// Parameters:
//  pathTracingData - The path tracing data.
//
// Returns:
// 	The pixel screen.
// 	An error.
//
func (*Controller) parsePixelScreenFromMap(pathTracingData map[string]interface{}) (*screen.Screen, error) {
	errorMessage := "unable to parse pixel screen"

	pixelScreenMap, found := pathTracingData["pixelScreen"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	pixelScreenMapParsed, parsed := pixelScreenMap.(map[string]interface{})
	if !parsed{
		return nil, errors.New(errorMessage)
	}

	width, found := pixelScreenMapParsed["width"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	widthParsed, parsed := width.(float64)
	if !parsed{
		return nil, errors.New(errorMessage)
	}

	height, found := pixelScreenMapParsed["height"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	heightParsed, parsed := height.(float64)
	if !parsed{
		return nil, errors.New(errorMessage)
	}

	pixelScreen, err := screen.Init(int(widthParsed), int(heightParsed))
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	return pixelScreen, nil
}

// ParsePathTracingInputsFromMap parses the inputs for a path tracing run.
//
// Parameters:
//  pathTracingData - The path tracing data.
//
// Returns:
// 	The PathTracer.
// 	The number of rays per pixel.
// 	The number recursions of each ray.
// 	The starting line index of the window of the screen to use the path tracing.
// 	The starting column index of the window of the screen to use the path tracing.
// 	The ending line index of the window of the screen to use the path tracing.
// 	The ending column index of the window of the screen to use the path tracing.
// 	An error.
//
func (controller *Controller) ParsePathTracingInputsFromMap(pathTracingData map[string]interface{}) (
	*path_tracing.PathTracer, int, int, int, int, int, int, error) {
	pixelScreen, err := controller.parsePixelScreenFromMap(pathTracingData)
	if err != nil {
		return nil, 0, 0, 0, 0, 0, 0, nil
	}
	pathTracer := path_tracing.Init(nil, pixelScreen, nil, nil)
	return pathTracer, 0, 0, 0, 0, 0, 0, nil
}
