package marshaller

import (
	"encoding/json"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/color_matrix"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/path_tracing"
)

// Controller is a class for controlling the marshaller of the application.
//
// Members:
// 	none
//
type Controller struct {}

// ColorMatrixToJson parses a color matrix to JSON.
//
// Parameters:
//  pathTracingData - The path tracing data.
//
// Returns:
// 	The list of objects.
// 	An error.
//
func (controller *Controller) ColorMatrixToJson(colorMatrix *color_matrix.ColorMatrix) ([]byte, error) {
	dtoColorMatrix := ColorMatrixDTO{Colors: colorMatrix.GetColors()}
	return json.Marshal(dtoColorMatrix)
}
// ParsePathTracingFromMap parses the inputs for a path tracing run.
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
func (controller *Controller) ParsePathTracingFromMap(pathTracingData map[string]interface{}) (
	*path_tracing.PathTracer, int, int, int, int, int, int, error) {

	pathTracingParametersInstance, err := controller.parsePathTracingParametersFromMap(
		pathTracingData)
	if err != nil {
		return nil, 0, 0, 0, 0, 0, 0, err
	}

	pixelScreen, err := controller.parsePixelScreenFromMap(pathTracingData)
	if err != nil {
		return nil, 0, 0, 0, 0, 0, 0, err
	}

	sceneCamera, err := controller.parseCameraFromMap(pathTracingData)
	if err != nil {
		return nil, 0, 0, 0, 0, 0, 0, err
	}

	lights, err := controller.parseLightsFromMap(pathTracingData)
	if err != nil {
		return nil, 0, 0, 0, 0, 0, 0, err
	}

	objects, err := controller.parseObjectsFromMap(pathTracingData)
	if err != nil {
		return nil, 0, 0, 0, 0, 0, 0, err
	}

	pathTracer := path_tracing.Init(objects, pixelScreen, sceneCamera, lights)
	return pathTracer, pathTracingParametersInstance.raysPerPixel, pathTracingParametersInstance.recursions,
	pathTracingParametersInstance.windowStartLine, pathTracingParametersInstance.windowStartColumn,
	pathTracingParametersInstance.windowEndLine, pathTracingParametersInstance.windowEndColumn, nil
}
