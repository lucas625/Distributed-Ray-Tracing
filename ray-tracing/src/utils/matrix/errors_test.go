package matrix

import (
	"fmt"
	"testing"
)

// TestInvalidSizeError tests invalid size error for matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestInvalidSizeError(t *testing.T) {
	size := -1
	err := invalidSize(size, size)
	if err == nil {
		t.Errorf("No invalid size error returned for size: %d.", size)
	} else if err.Error() != fmt.Sprintf("Invalid size for matrix. lines: %d and columns: %d.\n", size, size) {
		t.Errorf("Wrong error message for invalid size: \"%s\".", err.Error())
	}
}

// TestIncompatibleSize tests incompatible size error for matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIncompatibleSize(t *testing.T) {
	firstMatrix, _ := Init(3, 3)
	secondMatrix, _ := Init(2, 3)
	expectedError := fmt.Sprintf(
		"Incompatible size for matrices:\nFirst matrix: lines: %d and columns: %d." +
			"\nSecond matrix: lines: %d and columns: %d.\n",
		firstMatrix.Lines(), firstMatrix.Columns(), secondMatrix.Lines(), secondMatrix.Columns())
	err := incompatibleSize(firstMatrix, secondMatrix)
	if err == nil {
		t.Errorf("No incompatible size error returned for matrices: %v %v.", firstMatrix, secondMatrix)
	} else if err.Error() != expectedError {
		t.Errorf("Wrong error message for incompatible size: \"%s\".", err.Error())
	}
}

// TestIndexError tests index error of a Matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIndexError(t *testing.T) {
	matrix, _ := Init(1, 1)
	expectedError := fmt.Sprintf(
		"Index out of limits of the matrix. Expected from 0 0 to: %v %v and got %v %v.\n",
		matrix.Lines(), matrix.Columns(), -1, -1)
	err := indexError(matrix, -1, -1)
	if err == nil {
		t.Errorf("IndexError not raised.")
	} else if err.Error() != expectedError {
		t.Errorf("Wrong IndexError message.")
	}
}
