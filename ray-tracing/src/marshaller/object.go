package marshaller

import (
	"errors"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/object"
)

// parseLightCharacteristicsFromMap parses an object is light characteristics from a map.
//
// Parameters:
//  objectData - The object data.
//
// Returns:
//  RGB for the color of the object.
//  Percentage of specular rays.
//  How much reflections rays get distorted.
//  Percentage of transmission rays.
//  Percentage of diffuse rays.
// 	An error.
//
func (controller *Controller) parseLightCharacteristicsFromMap(objectData map[string]interface{}) (
	[]float64, float64, float64, float64, float64, error) {
	errorMessage := "unable to parse light characteristics"

	lightCharacteristicsInterface, found := objectData["lightCharacteristics"]
	if !found {
		return nil, 0, 0, 0, 0, errors.New(errorMessage)
	}
	lightsInterfaceListMap, parsed := lightCharacteristicsInterface.(map[string]interface{})
	if !parsed {
		return nil, 0, 0, 0, 0, errors.New(errorMessage)
	}

	color, err := controller.parseColorFromMap(lightsInterfaceListMap)
	if err != nil {
		return nil, 0, 0, 0, 0, errors.New(errorMessage)
	}

	specularReflection, err := controller.parseFloatFromMap(lightsInterfaceListMap, "specularReflection")
	if err != nil {
		return nil, 0, 0, 0, 0, errors.New(errorMessage)
	}
	roughNess, err := controller.parseFloatFromMap(lightsInterfaceListMap, "roughNess")
	if err != nil {
		return nil, 0, 0, 0, 0, errors.New(errorMessage)
	}
	transmissionReflection, err := controller.parseFloatFromMap(lightsInterfaceListMap, "transmissionReflection")
	if err != nil {
		return nil, 0, 0, 0, 0, errors.New(errorMessage)
	}
	diffuseReflection, err := controller.parseFloatFromMap(lightsInterfaceListMap, "diffuseReflection")
	if err != nil {
		return nil, 0, 0, 0, 0, errors.New(errorMessage)
	}

	return color, specularReflection, roughNess, transmissionReflection, diffuseReflection, nil
}

// parseLightFromMap parses an object from a map.
//
// Parameters:
//  objectData - The object data.
//
// Returns:
// 	An object.
// 	An error.
//
func (controller *Controller) parseObjectFromMap(objectData map[string]interface{}) (*object.Object, error) {
	errorMessage := "unable to parse object"

	name, err := controller.parseStringFromMap(objectData, "name")
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	color, specularReflection, roughness, transmissionReflection, diffuseReflection, err :=
		controller.parseLightCharacteristicsFromMap(objectData)
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	parsedObject, err := object.Init(name, nil, nil, nil, color, specularReflection, roughness,
		transmissionReflection, diffuseReflection)
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	return parsedObject, nil
}
