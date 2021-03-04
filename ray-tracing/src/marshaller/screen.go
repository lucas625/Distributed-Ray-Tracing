package marshaller

import (
	"errors"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/screen"
)

// screenController is a class for controlling the marshaller of the screens.
//
// Members:
// 	none
//
type screenController struct {}

// parsePixelScreenFromMap parses the pixel screen from a map.
//
// Parameters:
//  pathTracingData - The path tracing data.
//
// Returns:
// 	The pixel screen.
// 	An error.
//
func (*screenController) parsePixelScreenFromMap(pathTracingData map[string]interface{}) (*screen.Screen, error) {
	errorMessage := "unable to parse pixel screen"

	pixelScreenMap, found := pathTracingData["pixelScreen"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	pixelScreenMapParsed, parsed := pixelScreenMap.(map[string]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}

	generalController := generalController{}

	width, err := generalController.parseFloatFromMap(pixelScreenMapParsed, "width")
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	height, err := generalController.parseFloatFromMap(pixelScreenMapParsed, "height")
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	pixelScreen, err := screen.Init(int(width), int(height))
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	return pixelScreen, nil
}
