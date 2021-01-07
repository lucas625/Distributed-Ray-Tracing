package matrix

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
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
	size := 3

	matrix, err := Init(size, size)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, size, matrix.Lines())
	test_helpers.AssertEqual(t, size, matrix.Columns())

	for lineIndex := 0; lineIndex < size; lineIndex++ {
		for columnIndex := 0; columnIndex < size; columnIndex++ {
			matrixValue, err := matrix.GetValue(lineIndex, columnIndex)
			test_helpers.AssertNilError(t, err)
			test_helpers.AssertEqual(t, 0.0, matrixValue)
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
	expectedErrorMessage := fmt.Sprintf("Invalid size for matrix. lines: %d and columns: %d.", size, size)

	_, err := Init(size, size)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
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
	expectedErrorMessage := fmt.Sprintf("Invalid size for matrix. lines: %d and columns: %d.", size, size)

	_, err := Init(size, size)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
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

	matrixValue, err := matrix.GetValue(1, 1)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, 5.0, matrixValue)
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
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the matrix. Expected from 0 0 to: %v %v and got %v %v.", matrix.Lines(),
		matrix.Columns(), -1, -1)

	_, err := matrix.GetValue(-1, -1)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
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
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the matrix. Expected from 0 0 to: %v %v and got %v %v.", matrix.Lines(),
		matrix.Columns(), 3, 4)

	_, err := matrix.GetValue(3, 4)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
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
	test_helpers.AssertNilError(t, err)

	matrixValue, err := matrix.GetValue(1, 1)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, newValue, matrixValue)
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
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the matrix. Expected from 0 0 to: %v %v and got %v %v.", matrix.Lines(),
		matrix.Columns(), 1, -1)

	err := matrix.SetValue(1, -1, 10)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
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

	areMatricesEqual := firstMatrix.IsEqual(secondMatrix)
	test_helpers.AssertEqual(t, true, areMatricesEqual)
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

	areMatricesEqual := firstMatrix.IsEqual(secondMatrix)
	test_helpers.AssertEqual(t, false, areMatricesEqual)
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

	areMatricesEqual := firstMatrix.IsEqual(secondMatrix)
	test_helpers.AssertEqual(t, false, areMatricesEqual)
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

	areMatricesEqual := firstMatrix.IsEqual(secondMatrix)
	test_helpers.AssertEqual(t, false, areMatricesEqual)
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

	test_helpers.AssertEqual(t, true, areEqual)
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

	test_helpers.AssertEqual(t, expectedMatrixAsString, matrix.ToString())
}
