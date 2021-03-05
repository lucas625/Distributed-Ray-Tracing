package marshaller

import (
	"errors"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/light"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/object"
)

// parseLightObjectFromMap parses a light object from a map.
//
// Parameters:
//  lightData - The light data.
//
// Returns:
// 	An object.
// 	An error.
//
func (controller *Controller) parseLightObjectFromMap(lightData map[string]interface{}) (*object.Object, error) {
	errorMessage := "unable to parse light object"

	lightObjectInterface, found := lightData["lightObject"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	lightObjectMap, parsed := lightObjectInterface.(map[string]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}

	lightObject, err := controller.parseObjectFromMap(lightObjectMap)
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	return lightObject, nil
}

// parseLightFromMap parses a light from a map.
//
// Parameters:
//  lightData - The light data.
//
// Returns:
// 	A light.
// 	An error.
//
func (controller *Controller) parseLightFromMap(lightData map[string]interface{}) (*light.Light, error) {
	errorMessage := "unable to parse light"

	lightIntensity, err := controller.parseFloatFromMap(lightData, "lightIntensity")
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	color, err := controller.parseFloatListFromMap(lightData, "color")
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	lightObject, err := controller.parseLightObjectFromMap(lightData)
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	parsedLight, err := light.Init(lightIntensity, lightObject, color)
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
func (controller *Controller) parseLightsFromMap(pathTracingData map[string]interface{}) ([]*light.Light, error) {
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
