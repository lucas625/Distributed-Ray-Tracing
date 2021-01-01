package matrix

import (
	"fmt"
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
	err := invalidSize(size, size)
	expectedErrorMessage := fmt.Sprintf("Invalid size for matrix. lines: %d and columns: %d.", size, size)
	if err == nil {
		t.Errorf("No invalid size error returned for size: %d.", size)
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Wrong error message for invalid size: \"%s\".", err.Error())
	}
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
	firstMatrix, _ := Init(3, 3)
	secondMatrix, _ := Init(2, 3)
	expectedErrorMessage := fmt.Sprintf(
		"Incompatible size for matrices:\nFirst matrix: lines: %d and columns: %d." +
			"\nSecond matrix: lines: %d and columns: %d.",
		firstMatrix.Lines(), firstMatrix.Columns(), secondMatrix.Lines(), secondMatrix.Columns())
	err := incompatibleSize(firstMatrix, secondMatrix)
	if err == nil {
		t.Errorf("No incompatible size error returned for matrices: %v %v.", firstMatrix, secondMatrix)
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Wrong error message for incompatible size: \"%s\".", err.Error())
	}
}

// TestMatrix_IndexError tests the index error of a Matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIndexError(t *testing.T) {
	matrix, _ := Init(1, 1)
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the matrix. Expected from 0 0 to: %v %v and got %v %v.",
		matrix.Lines(), matrix.Columns(), -1, -1)
	err := indexError(matrix, -1, -1)
	if err == nil {
		t.Errorf("IndexError not raised.")
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Wrong IndexError message: \"%s\".", err.Error())
	}
}
