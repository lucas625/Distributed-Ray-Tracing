package matrix

import (
	"fmt"
	"testing"
)

// TestMatrix_Init tests the instantiation of a Matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_Init(t *testing.T) {
	size := 1
	matrix, err := Init(size, size)
	if err != nil {
		t.Errorf("Matrix failed to be instantiated with size: %d.", size)
	}
	if matrix.Lines() != size {
		t.Errorf("Matrix instantiated with wrong lines: Expected: %d and Got: %d.", size, matrix.Lines())
	}
	if matrix.Columns() != size {
		t.Errorf("Matrix instantiated with wrong columns: Expected: %d and Got: %d.", size, matrix.Columns())
	}
	for lineIndex := 0; lineIndex < size; lineIndex++ {
		for columnIndex := 0; columnIndex < size; columnIndex++ {
			matrixValue, _ := matrix.GetValue(lineIndex, columnIndex)
			if matrixValue != 0 {
				t.Errorf("Matrix instantiated with non zero values: %v.", matrix.values)
			}
		}
	}
}

// TestMatrix_Init_ZeroSize tests the instantiation of a Matrix with zero size.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_Init_ZeroSize(t *testing.T) {
	size := 0
	_, err := Init(size, size)
	expectedErrorMessage := fmt.Sprintf("Invalid size for matrix. lines: %d and columns: %d.", size, size)
	if err == nil {
		t.Errorf("Matrix instantiated with zero size: %d.", size)
	} else if err.Error() != expectedErrorMessage {
		t.Errorf(
			"Matrix failed to be instantiated with zero size: %d but with wrong error message: \"%s\".",
			size, err.Error())
	}
}

// TestMatrix_Init_NegativeSize tests the instantiation of a Matrix with negative size.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_Init_NegativeSize(t *testing.T) {
	size := -1
	_, err := Init(size, size)
	if err == nil {
		t.Errorf("Matrix instantiated with negative size: %d.", size)
	} else if err.Error() != fmt.Sprintf("Invalid size for matrix. lines: %d and columns: %d.", size, size) {
		t.Errorf(
			"Matrix failed to be instantiated with negative size, but with wrong error message: \"%s\".",
			err.Error())
	}
}

// TestMatrix_GetValue tests the GetValue.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_GetValue(t *testing.T) {
	matrix, _ := Init(2, 3)
	matrix.values[0] = []float64{1, 2, 3}
	matrix.values[1] = []float64{4, 5, 6}

	matrixValue, _ := matrix.GetValue(1, 1)

	if matrixValue != 5 {
		t.Errorf("Got the wrong value. Expected: %v and got: %v.", 5, matrixValue)
	}
}

// TestMatrix_GetValue_NegativeIndex tests the GetValue with a negative index.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_GetValue_NegativeIndex(t *testing.T) {
	matrix, _ := Init(2, 3)
	matrix.values[0] = []float64{1, 2, 3}
	matrix.values[1] = []float64{4, 5, 6}

	_, err := matrix.GetValue(-1, -1)

	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the matrix. Expected from 0 0 to: %v %v and got %v %v.", matrix.Lines(),
		matrix.Columns(), -1, -1)

	if err == nil {
		t.Errorf("IndexError not raised.")
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Wrong IndexError message: %s.", err.Error())
	}
}

// TestMatrix_GetValue_IndexError tests the GetValue with a an index error.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_GetValue_IndexError(t *testing.T) {
	matrix, _ := Init(2, 3)
	matrix.values[0] = []float64{1, 2, 3}
	matrix.values[1] = []float64{4, 5, 6}

	_, err := matrix.GetValue(3, 4)

	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the matrix. Expected from 0 0 to: %v %v and got %v %v.", matrix.Lines(),
		matrix.Columns(), 3, 4)

	if err == nil {
		t.Errorf("IndexError not raised.")
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Wrong IndexError message.")
	}
}

// TestMatrix_SetValue tests the SetValue.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_SetValue(t *testing.T) {
	matrix, _ := Init(2, 3)
	matrix.values[0] = []float64{1, 2, 3}
	matrix.values[1] = []float64{4, 5, 6}

	newValue := 10.0
	err := matrix.SetValue(1, 1, newValue)

	if err != nil {
		t.Errorf("Unexpected error while setting a matrix value: %v.", err.Error())
	}

	matrixValue, _ := matrix.GetValue(1, 1)

	if matrixValue != newValue {
		t.Errorf("Set the wrong value. Expected: %v and got: %v.", newValue, matrixValue)
	}
}

// TestMatrix_SetValue_IndexError tests the SetValue with an index error.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_SetValue_IndexError(t *testing.T) {
	matrix, _ := Init(2, 3)
	matrix.values[0] = []float64{1, 2, 3}
	matrix.values[1] = []float64{4, 5, 6}

	err := matrix.SetValue(1, -1, 10)

	if err == nil {
		t.Errorf("Error not raised when trying to set a matrix value on an invalid index.")
	} else if err.Error() !=
		fmt.Sprintf("Index out of limits of the matrix. Expected from 0 0 to: %v %v and got %v %v.",
			matrix.Lines(), matrix.Columns(), 1, -1) {
		t.Errorf("Wrong IndexError message.")
	}
}

// TestMatrix_IsEqual tests the is equal of a Matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_IsEqual(t *testing.T) {
	firstMatrix, _ := Init(3, 2)
	firstMatrix.values[0] = []float64{1, 2}
	firstMatrix.values[1] = []float64{3, 4}
	firstMatrix.values[2] = []float64{5, 6}

	secondMatrix, _ := Init(3, 2)
	secondMatrix.values[0] = []float64{1, 2}
	secondMatrix.values[1] = []float64{3, 4}
	secondMatrix.values[2] = []float64{5, 6}

	if !firstMatrix.IsEqual(secondMatrix) {
		t.Errorf("Matrices are different: %v %v.", firstMatrix, secondMatrix)
	}
}

// TestMatrix_IsEqual_Different tests the is equal of a Matrix when they are different.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_IsEqual_Different(t *testing.T) {
	firstMatrix, _ := Init(3, 2)
	firstMatrix.values[0] = []float64{1, 2}
	firstMatrix.values[1] = []float64{3, 4}
	firstMatrix.values[2] = []float64{5, 6}

	secondMatrix, _ := Init(3, 2)
	secondMatrix.values[0] = []float64{10, 20}
	secondMatrix.values[1] = []float64{30, 40}
	secondMatrix.values[2] = []float64{50, 60}

	if firstMatrix.IsEqual(secondMatrix) {
		t.Errorf("Matrices are equal: %v %v.", firstMatrix, secondMatrix)
	}
}

// TestMatrix_IsEqual_DifferentLines tests the is equal of a Matrix with different lines.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_IsEqual_DifferentLines(t *testing.T) {
	firstMatrix, _ := Init(2, 2)
	secondMatrix, _ := Init(3, 2)

	if firstMatrix.IsEqual(secondMatrix) {
		t.Errorf("Matrices are equal: %v %v.", firstMatrix, secondMatrix)
	}
}

// TestMatrix_IsEqual_DifferentColumns tests the is equal of a Matrix with different columns.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_IsEqual_DifferentColumns(t *testing.T) {
	firstMatrix, _ := Init(3, 3)
	secondMatrix, _ := Init(3, 2)

	if firstMatrix.IsEqual(secondMatrix) {
		t.Errorf("Matrices are equal: %v %v.", firstMatrix, secondMatrix)
	}
}

// TestMatrix_CopyAllValues tests the CopyAllValues.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_CopyAllValues(t *testing.T) {
	matrix, _ := Init(2, 3)
	matrix.values[0] = []float64{1, 2, 3}
	matrix.values[1] = []float64{4, 5, 6}

	copiedValues := matrix.CopyAllValues()

	areEqual := true
	for lineIndex := 0; lineIndex < matrix.Lines(); lineIndex++ {
		for columnIndex := 0; columnIndex < matrix.Columns(); columnIndex++ {
			expectedValue, _ := matrix.GetValue(lineIndex, columnIndex)
			if copiedValues[lineIndex][columnIndex] != expectedValue {
				areEqual = false
			}
		}
	}

	if !areEqual {
		t.Errorf("Copied values are different from expected %v: %v.", matrix.values, copiedValues)
	}
}

// TestMatrix_ToString tests the ToString.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_ToString(t *testing.T) {
	matrix, _ := Init(2, 3)
	matrix.values[0] = []float64{1, 2, 3}
	matrix.values[1] = []float64{4, 5, 6}
	expectedMatrixAsString := fmt.Sprintf(
		"Lines: %v Columns: %v\n Matrix: %v\n", matrix.Lines(), matrix.Columns(), matrix.values)
	if matrix.ToString() != expectedMatrixAsString {
		t.Errorf("Invalid to matrix to string: %v.", matrix)
	}
}
