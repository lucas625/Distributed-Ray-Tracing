package marshaller

import (
	"errors"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/light"
)

// lightController is a class for controlling the marshaller of lights.
//
// Members:
// 	none
//
type lightController struct {}

// parseLightFromMap parses a light from a map.
//
// Parameters:
//  lightData - The light data.
//
// Returns:
// 	A light.
// 	An error.
//
func (controller *lightController) parseLightFromMap(lightData map[string]interface{}) (*light.Light, error) {
	errorMessage := "unable to parse light"

	generalMarshallerController := generalController{}

	lightIntensity, err := generalMarshallerController.parseFloatFromMap(lightData, "lightIntensity")
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	color, err := generalMarshallerController.parseColorFromMap(lightData)
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	parsedLight, err := light.Init(lightIntensity, nil, color)
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	return parsedLight, nil
}

// parseLightsFromMap parses lights from map.
//
// Parameters:
//  pathTracingData - The path tracing data.
//
// Returns:
// 	The list of lights.
// 	An error.
//
func (controller *lightController) parseLightsFromMap(pathTracingData map[string]interface{}) ([]*light.Light, error) {
	errorMessage := "unable to parse lights"

	lightsInterface, found := pathTracingData["lights"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	lightsInterfaceList, parsed := lightsInterface.([]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}

	lights := make([]*light.Light, len(lightsInterfaceList))
	for lightIndex := 0; lightIndex < len(lightsInterfaceList); lightIndex++ {
		lightMap, parsed := lightsInterfaceList[lightIndex].(map[string]interface{})
		if !parsed {
			return nil, errors.New(errorMessage)
		}
		currentLight, err := controller.parseLightFromMap(lightMap)
		if err != nil {
			return nil, errors.New(errorMessage)
		}
		lights[lightIndex] = currentLight
	}

	return lights, nil
}
