package triangle

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestTriangle_Init tests the instantiation of a Triangle.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestTriangle_Init(t *testing.T) {
	verticesIndexes := []int{1, 2, 3}
	verticesNormalsIndexes := []int{1, 2, 3}

	triangle, err := Init(verticesIndexes, verticesNormalsIndexes)
	test_helpers.AssertNilError(t, err)

	for index := 0; index < 3; index++ {
		triangleVertexIndex, err := triangle.GetVertexIndex(index)
		test_helpers.AssertNilError(t, err)
		isEqualVertex := verticesIndexes[index] == triangleVertexIndex
		test_helpers.AssertEqual(t, true, isEqualVertex)

		triangleVertexNormalIndex, err := triangle.GetVertexNormalIndex(index)
		test_helpers.AssertNilError(t, err)
		isEqualVertexNormal := verticesNormalsIndexes[index] == triangleVertexNormalIndex
		test_helpers.AssertEqual(t, true, isEqualVertexNormal)
	}
}

// TestTriangle_Init_InvalidVerticesIndexes tests the instantiation of a Triangle with invalid vertices indexes.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestTriangle_Init_InvalidVerticesIndexes(t *testing.T) {
	verticesIndexes := []int{1, 2, 3, 4}
	verticesNormalsIndexes := []int{1, 2, 3}
	expectedErrorMessage := fmt.Sprintf("Invalid size for triangle vertices indexes %v or vertices normals indexes %v.",
		len(verticesIndexes), len(verticesNormalsIndexes))

	_, err := Init(verticesIndexes, verticesNormalsIndexes)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestTriangle_Init_InvalidVerticesNormalsIndexes tests the instantiation of a Triangle with invalid vertices normals
// indexes.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestTriangle_Init_InvalidVerticesNormalsIndexes(t *testing.T) {
	verticesIndexes := []int{1, 2, 3}
	verticesNormalsIndexes := []int{1, 2, 3, 4}
	expectedErrorMessage := fmt.Sprintf("Invalid size for triangle vertices indexes %v or vertices normals indexes %v.",
		len(verticesIndexes), len(verticesNormalsIndexes))

	_, err := Init(verticesIndexes, verticesNormalsIndexes)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestTriangle_GetVertexIndex tests the get vertex index of a Triangle.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestTriangle_GetVertexIndex(t *testing.T) {
	verticesIndexes := []int{11, 15, 33}
	verticesNormalsIndexes := []int{1, 2, 3}
	expectedVertexIndex := 15

	triangle, err := Init(verticesIndexes, verticesNormalsIndexes)
	test_helpers.AssertNilError(t, err)

	vertexIndex, err := triangle.GetVertexIndex(1)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, expectedVertexIndex, vertexIndex)
}

// TestTriangle_GetVertexIndex_NegativeIndex tests the get vertex index of a Triangle with a negative index.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestTriangle_GetVertexIndex_NegativeIndex(t *testing.T) {
	verticesIndexes := []int{11, 15, 33}
	verticesNormalsIndexes := []int{1, 2, 3}
	index := -1
	expectedErrorMessage := fmt.Sprintf("Index out of limits of the triangle. Expected from 0 to 2 and got %v.", index)

	triangle, err := Init(verticesIndexes, verticesNormalsIndexes)
	test_helpers.AssertNilError(t, err)

	_, err = triangle.GetVertexIndex(index)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestTriangle_GetVertexIndex_BiggerIndex tests the get vertex index of a Triangle with a index bigger than the limit.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestTriangle_GetVertexIndex_BiggerIndex(t *testing.T) {
	verticesIndexes := []int{11, 15, 33}
	verticesNormalsIndexes := []int{1, 2, 3}
	index := 3
	expectedErrorMessage := fmt.Sprintf("Index out of limits of the triangle. Expected from 0 to 2 and got %v.", index)

	triangle, err := Init(verticesIndexes, verticesNormalsIndexes)
	test_helpers.AssertNilError(t, err)

	_, err = triangle.GetVertexIndex(index)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestTriangle_GetVertexNormalIndex tests the get vertex normal index of a Triangle.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestTriangle_GetVertexNormalIndex(t *testing.T) {
	verticesIndexes := []int{1, 2, 3}
	verticesNormalsIndexes := []int{11, 15, 33}
	expectedVertexNormalIndex := 33

	triangle, err := Init(verticesIndexes, verticesNormalsIndexes)
	test_helpers.AssertNilError(t, err)

	vertexNormalIndex, err := triangle.GetVertexNormalIndex(2)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, expectedVertexNormalIndex, vertexNormalIndex)
}

// TestTriangle_GetVertexNormalIndex_NegativeIndex tests the get vertex normal index of a Triangle with a negative index.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestTriangle_GetVertexNormalIndex_NegativeIndex(t *testing.T) {
	verticesIndexes := []int{1, 2, 3}
	verticesNormalsIndexes := []int{11, 15, 33}
	index := -1
	expectedErrorMessage := fmt.Sprintf("Index out of limits of the triangle. Expected from 0 to 2 and got %v.", index)

	triangle, err := Init(verticesIndexes, verticesNormalsIndexes)
	test_helpers.AssertNilError(t, err)

	_, err = triangle.GetVertexNormalIndex(index)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestTriangle_GetVertexNormalIndex_BiggerIndex tests the get vertex normal index of a Triangle with a index bigger than
// the limit.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestTriangle_GetVertexNormalIndex_BiggerIndex(t *testing.T) {
	verticesIndexes := []int{1, 2, 3}
	verticesNormalsIndexes := []int{11, 15, 33}
	index := 3
	expectedErrorMessage := fmt.Sprintf("Index out of limits of the triangle. Expected from 0 to 2 and got %v.", index)

	triangle, err := Init(verticesIndexes, verticesNormalsIndexes)
	test_helpers.AssertNilError(t, err)

	_, err = triangle.GetVertexNormalIndex(index)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
