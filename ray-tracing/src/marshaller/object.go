package marshaller

import (
	"errors"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/triangle"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/object"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/point_repository"
)

// parseRepositoryFromMap parses repository from map.
//
// Parameters:
//  objectData - The object data.
//
// Returns:
// 	The point repository.
// 	An error.
//
func (controller *Controller) parseRepositoryFromMap(objectData map[string]interface{}) (
	*point_repository.PointRepository, error) {
	errorMessage := "unable to parse repository"

	repositoryInterface, found := objectData["repository"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	repositoryMap, parsed := repositoryInterface.(map[string]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}

	pointsInterface, found := repositoryMap["points"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	pointsInterfaceList, parsed := pointsInterface.([]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}

	points := make([]*point.Point, len(pointsInterfaceList))
	for pointIndex := 0; pointIndex < len(pointsInterfaceList); pointIndex++ {
		pointMap, parsed := pointsInterfaceList[pointIndex].(map[string]interface{})
		if !parsed {
			return nil, errors.New(errorMessage)
		}
		currentPoint, err := controller.parsePointFromMap(pointMap)
		if err != nil {
			return nil, errors.New(errorMessage)
		}
		points[pointIndex] = currentPoint
	}

	repository, err := point_repository.Init(points, 3)
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	return repository, nil
}

// parseTrianglesFromMap parses triangles from map.
//
// Parameters:
//  objectData - The object data.
//
// Returns:
// 	The list of triangles.
// 	An error.
//
func (controller *Controller) parseTrianglesFromMap(objectData map[string]interface{}) ([]*triangle.Triangle, error) {
	errorMessage := "unable to parse triangles"

	trianglesInterface, found := objectData["triangles"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	trianglesInterfaceList, parsed := trianglesInterface.([]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}

	triangles := make([]*triangle.Triangle, len(trianglesInterfaceList))
	for triangleIndex := 0; triangleIndex < len(trianglesInterfaceList); triangleIndex++ {
		triangleMap, parsed := trianglesInterfaceList[triangleIndex].(map[string]interface{})
		if !parsed {
			return nil, errors.New(errorMessage)
		}
		currentTriangle, err := controller.parseTriangleFromMap(triangleMap)
		if err != nil {
			return nil, errors.New(errorMessage)
		}
		triangles[triangleIndex] = currentTriangle
	}

	return triangles, nil
}

// parseNormalsFromMap parses normals from map.
//
// Parameters:
//  objectData - The object data.
//
// Returns:
// 	The list of normal vectors.
// 	An error.
//
func (controller *Controller) parseNormalsFromMap(objectData map[string]interface{}) ([]*vector.Vector, error) {
	errorMessage := "unable to parse normals"

	normalsInterface, found := objectData["normals"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	normalsInterfaceList, parsed := normalsInterface.([]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}

	normals := make([]*vector.Vector, len(normalsInterfaceList))
	for normalIndex := 0; normalIndex < len(normalsInterfaceList); normalIndex++ {
		normalMap, parsed := normalsInterfaceList[normalIndex].(map[string]interface{})
		if !parsed {
			return nil, errors.New(errorMessage)
		}
		currentNormal, err := controller.parseVectorFromMap(normalMap)
		if err != nil {
			return nil, errors.New(errorMessage)
		}
		normals[normalIndex] = currentNormal
	}

	return normals, nil
}

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

	color, err := controller.parseFloatListFromMap(lightsInterfaceListMap, "color")
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

// parseObjectFromMap parses an object from a map.
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

	normals, err := controller.parseNormalsFromMap(objectData)
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	repository, err := controller.parseRepositoryFromMap(objectData)
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	triangles, err := controller.parseTrianglesFromMap(objectData)
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	parsedObject, err := object.Init(name, repository, triangles, normals, color, specularReflection, roughness,
		transmissionReflection, diffuseReflection)
	if err != nil {
		return nil, errors.New(errorMessage)
	}
	return parsedObject, nil
}

// parseObjectsFromMap parses objects from map.
//
// Parameters:
//  pathTracingData - The path tracing data.
//
// Returns:
// 	The list of objects.
// 	An error.
//
func (controller *Controller) parseObjectsFromMap(pathTracingData map[string]interface{}) ([]*object.Object, error) {
	errorMessage := "unable to parse objects"

	objectsInterface, found := pathTracingData["objects"]
	if !found {
		return nil, errors.New(errorMessage)
	}
	objectsInterfaceList, parsed := objectsInterface.([]interface{})
	if !parsed {
		return nil, errors.New(errorMessage)
	}

	objects := make([]*object.Object, len(objectsInterfaceList))
	for objectIndex := 0; objectIndex < len(objectsInterfaceList); objectIndex++ {
		objectMap, parsed := objectsInterfaceList[objectIndex].(map[string]interface{})
		if !parsed {
			return nil, errors.New(errorMessage)
		}
		currentObject, err := controller.parseObjectFromMap(objectMap)
		if err != nil {
			return nil, errors.New(errorMessage)
		}
		objects[objectIndex] = currentObject
	}

	return objects, nil
}
