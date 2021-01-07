package matrix

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestMatrix_InvalidSizeError tests the invalid size error for matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_InvalidSizeError(t *testing.T) {
	size := -1
	expectedErrorMessage := fmt.Sprintf("Invalid size for matrix. lines: %d and columns: %d.", size, size)

	err := invalidSize(size, size)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestMatrix_IncompatibleSize tests the incompatible size error for matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_IncompatibleSize(t *testing.T) {
	firstMatrix, err := Init(3, 3)
	test_helpers.AssertNilError(t, err)
	secondMatrix, err := Init(2, 3)
	test_helpers.AssertNilError(t, err)
	expectedErrorMessage := fmt.Sprintf(
		"Incompatible size for matrices:\nFirst matrix: lines: %d and columns: %d." +
			"\nSecond matrix: lines: %d and columns: %d.",
		firstMatrix.Lines(), firstMatrix.Columns(), secondMatrix.Lines(), secondMatrix.Columns())

	err = incompatibleSize(firstMatrix, secondMatrix)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestMatrix_IndexError tests the index error of a Matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_IndexError(t *testing.T) {
	matrix, err := Init(1, 1)
	test_helpers.AssertNilError(t, err)
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the matrix. Expected from 0 0 to: %v %v and got %v %v.",
		matrix.Lines(), matrix.Columns(), -1, -1)

	err = indexError(matrix, -1, -1)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
