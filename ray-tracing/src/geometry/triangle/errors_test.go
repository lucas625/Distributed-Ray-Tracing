package triangle

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestTriangle_indexError tests the index error of a Triangle.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestTriangle_indexError(t *testing.T) {
	index := -1
	expectedErrorMessage := fmt.Sprintf("Index out of limits of the triangle. Expected from 0 to 2 and got %v.", index)

	err := indexError(index)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestTriangle_initializationError tests the initialization error of a Triangle.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestTriangle_initializationError(t *testing.T) {
	verticesIndexes := []int{1, 2, 3, 4}
	verticesNormalsIndexes := []int{1, 2, 3, 4}
	expectedErrorMessage := fmt.Sprintf("Invalid size for triangle vertices indexes %v or vertices normals indexes %v.",
		len(verticesIndexes), len(verticesNormalsIndexes))

	err := initializationError(verticesIndexes, verticesNormalsIndexes)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
