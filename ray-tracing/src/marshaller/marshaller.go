package marshaller

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/path_tracing"
)

// Controller is a class for controlling the marshaller of the application.
//
// Members:
// 	none
//
type Controller struct {}

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
	screenMarshallerController := screenController{}
	pixelScreen, err := screenMarshallerController.parsePixelScreenFromMap(pathTracingData)
	if err != nil {
		return nil, 0, 0, 0, 0, 0, 0, err
	}
	cameraMarshallerController := cameraController{}
	sceneCamera, err := cameraMarshallerController.parseCameraFromMap(pathTracingData)
	if err != nil {
		return nil, 0, 0, 0, 0, 0, 0, err
	}
	pathTracer := path_tracing.Init(nil, pixelScreen, sceneCamera, nil)
	return pathTracer, 0, 0, 0, 0, 0, 0, nil
}
