package triangle

import (
	"errors"
	"fmt"
)

// indexError is the error where we try to access an index out of the limits of the Triangle.
//
// Parameters:
//	index - The index being accessed.
//
// Returns:
//  An Error.
//
func indexError(index int) error {
	errorMessage := fmt.Sprintf("Index out of limits of the triangle. Expected from 0 to 2 and got %v.", index)
	return errors.New(errorMessage)
}

// initializationError is the error where we try to initialize a Triangle with invalid sizes on the lists.
//
// Parameters:
// 	verticesIndexes        - The 3 indexes of the Triangle is vertices on the point list.
// 	verticesNormalsIndexes - The 3 indexes of the Triangle is vertices normals on the normals list.
//
// Returns:
//  An Error.
//
func initializationError(verticesIndexes, verticesNormalsIndexes []int) error {
	errorMessage := fmt.Sprintf("Invalid size for triangle vertices indexes %v or vertices normals indexes %v.",
		len(verticesIndexes), len(verticesNormalsIndexes))
	return errors.New(errorMessage)
}
