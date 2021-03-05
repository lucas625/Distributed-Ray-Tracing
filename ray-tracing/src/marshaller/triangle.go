package marshaller

import (
	"errors"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/triangle"
)

// parseTriangleFromMap parses a triangle from a map.
//
// Parameters:
//  triangleAsMap - The triangle as a map.
//
// Returns:
// 	The triangle.
// 	An error.
//
func (controller *Controller) parseTriangleFromMap(triangleAsMap map[string]interface{}) (*triangle.Triangle, error) {
	errorMessage := "invalid triangle"

	verticesIndices, err := controller.parseIntListFromMap(triangleAsMap, "verticesIndices")
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	verticesNormalsIndices, err := controller.parseIntListFromMap(triangleAsMap, "verticesNormalsIndices")
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	parsedTriangle, err := triangle.Init(verticesIndices, verticesNormalsIndices)
	if err != nil {
		return nil, errors.New(errorMessage)
	}

	return parsedTriangle, nil
}
