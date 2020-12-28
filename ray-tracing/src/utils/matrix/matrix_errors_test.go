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
	} else if err.Error() != fmt.Sprintf("Invalid size for matrix. Lines: %d and Columns: %d.\n", size, size) {
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
		"Incompatible size for matrices:\nFirst matrix: Lines: %d and Columns: %d." +
			"\nSecond matrix: Lines: %d and Columns: %d.\n",
		firstMatrix.Lines, firstMatrix.Columns, secondMatrix.Lines, secondMatrix.Columns)
	err := incompatibleSize(firstMatrix, secondMatrix)
	if err == nil {
		t.Errorf("No incompatible size error returned for matrices: %v %v.", firstMatrix, secondMatrix)
	} else if err.Error() != expectedError {
		t.Errorf("Wrong error message for incompatible size: \"%s\".", err.Error())
	}
}
