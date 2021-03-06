package marshaller

import (
	"errors"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/camera"
)

// parseCameraVectorFromMap parses the vector of the camera.
//
// Parameters:
//  cameraMap  - The camera as a map.
//  vectorName - The name of the vector.
//
// Returns:
// 	The vector.
// 	An error.
//
func (controller *Controller) parseCameraVectorFromMap(cameraMap map[string]interface{}, vectorName string) (
	*vector.Vector, error) {
	errorMessage := "unable to parse camera vector"

	vectorMap, found := cameraMap[vectorName]
	if !found {
		return nil, errors.New(errorMessage)
	}
	vectorMapParsed, parsed := vectorMap.(map[string]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}
	parsedVector, err := controller.parseVectorFromMap(vectorMapParsed)
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	return parsedVector, nil
}

// parseCameraVectorsFromMap parses the vectors of the camera.
//
// Parameters:
//  cameraMap  - The camera as a map.
//
// Returns:
// 	The look vector.
// 	The up vector.
// 	The right vector.
// 	An error.
//
func (controller *Controller) parseCameraVectorsFromMap(cameraMap map[string]interface{}) (
	*vector.Vector, *vector.Vector, *vector.Vector, error) {
	errorMessage := "unable to parse camera vectors"
	lookVector, err := controller.parseCameraVectorFromMap(cameraMap, "look")
	if err != nil {
		return nil, nil, nil, errors.New(errorMessage)
	}
	upVector, err := controller.parseCameraVectorFromMap(cameraMap, "up")
	if err != nil {
		return nil, nil, nil, errors.New(errorMessage)
	}
	rightVector, err := controller.parseCameraVectorFromMap(cameraMap, "right")
	if err != nil {
		return nil, nil, nil, errors.New(errorMessage)
	}
	return lookVector, upVector, rightVector, nil
}

// parseCameraPositionFromMap parses the a vector of the camera.
//
// Parameters:
//  cameraMap - The camera as a map.
//
// Returns:
// 	The vector.
// 	An error.
//
func (controller *Controller) parseCameraPositionFromMap(cameraMap map[string]interface{}) (
	*point.Point, error) {
	errorMessage := "unable to parse camera position"

	vectorMap, found := cameraMap["position"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	vectorMapParsed, parsed := vectorMap.(map[string]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}
	parsedPoint, err := controller.parsePointFromMap(vectorMapParsed)
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	return parsedPoint, nil
}

// parseCameraFromMap parses the scene camera from a map.
//
// Parameters:
//  pathTracingData - The path tracing data.
//
// Returns:
// 	The scene camera.
// 	An error.
//
func (controller *Controller) parseCameraFromMap(pathTracingData map[string]interface{}) (*camera.Camera, error) {
	errorMessage := "unable to parse scene is camera"

	sceneCameraMap, found := pathTracingData["sceneCamera"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	sceneCameraMapParsed, parsed := sceneCameraMap.(map[string]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}

	look, up, right, err := controller.parseCameraVectorsFromMap(sceneCameraMapParsed)
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	position, err := controller.parseCameraPositionFromMap(sceneCameraMapParsed)
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	fieldOfView, err := controller.parseFloatFromMap(sceneCameraMapParsed, "fieldOfView")
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	distanceToScreen, err := controller.parseFloatFromMap(
		sceneCameraMapParsed, "distanceToScreen")
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	sceneCamera, err := camera.Init(position, look, up, right, fieldOfView, distanceToScreen)
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	return sceneCamera, nil
}
