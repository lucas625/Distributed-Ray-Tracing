package matrix

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestMatrix_BuildIdentity tests the build identity Matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_BuildIdentity(t *testing.T) {
	size := 3
	controller := Controller{}

	matrix, err := controller.BuildIdentity(size)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, size, matrix.Lines())
	test_helpers.AssertEqual(t, size, matrix.Columns())

	for line := 0; line < size; line++ {
		for column := 0; column < size; column++ {
			matrixValue, err := matrix.GetValue(line, column)
			test_helpers.AssertNilError(t, err)
			if line != column {
				test_helpers.AssertEqual(t, 0.0, matrixValue)
			} else {
				test_helpers.AssertEqual(t, 1.0, matrixValue)
			}
		}
	}
}

// TestMatrix_BuildIdentity_ZeroSize tests the build identity matrix with zero size.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_BuildIdentity_ZeroSize(t *testing.T) {
	size := 0
	controller := Controller{}
	expectedErrorMessage := fmt.Sprintf("Invalid size for matrix. lines: %d and columns: %d.", size, size)

	_, err := controller.BuildIdentity(size)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestMatrix_Transpose tests the transposition of a matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_Transpose(t *testing.T) {
	matrix, err := Init(3, 2)
	test_helpers.AssertNilError(t, err)
	matrix.values[0] = []float64{1, 2}
	matrix.values[1] = []float64{3, 4}
	matrix.values[2] = []float64{5, 6}
	controller := Controller{}

	resultingMatrix := controller.Transpose(matrix)

	expectedMatrix, err := Init(2, 3)
	test_helpers.AssertNilError(t, err)
	expectedMatrix.values[0] = []float64{1, 3, 5}
	expectedMatrix.values[1] = []float64{2, 4, 6}

	areEqual := resultingMatrix.IsEqual(expectedMatrix)
	test_helpers.AssertEqual(t, true, areEqual)
}

// TestMatrix_ScalarMultiplication tests the multiplication of a matrix by a constant.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_ScalarMultiplication(t *testing.T) {
	matrix, err := Init(3, 2)
	test_helpers.AssertNilError(t, err)
	matrix.values[0] = []float64{1, 2}
	matrix.values[1] = []float64{3, 4}
	matrix.values[2] = []float64{5, 6}

	scalar := 3.0
	controller := Controller{}

	resultingMatrix := controller.ScalarMultiplication(matrix, scalar)

	expectedMatrix, err := Init(3, 2)
	test_helpers.AssertNilError(t, err)
	expectedMatrix.values[0] = []float64{3, 6}
	expectedMatrix.values[1] = []float64{9, 12}
	expectedMatrix.values[2] = []float64{15, 18}

	areEqual := resultingMatrix.IsEqual(expectedMatrix)
	test_helpers.AssertEqual(t, true, areEqual)
}

// TestMatrix_MultiplyMatrix tests the multiplication of a matrix by another.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_MultiplyMatrix(t *testing.T) {
	firstMatrix, err := Init(3, 2)
	test_helpers.AssertNilError(t, err)
	firstMatrix.values[0] = []float64{1, 2}
	firstMatrix.values[1] = []float64{3, 4}
	firstMatrix.values[2] = []float64{5, 6}

	secondMatrix, err := Init(2, 3)
	test_helpers.AssertNilError(t, err)
	secondMatrix.values[0] = []float64{7, 8, 9}
	secondMatrix.values[1] = []float64{10, 11, 12}

	controller := Controller{}

	resultingMatrix, err := controller.MultiplyMatrix(firstMatrix, secondMatrix)
	test_helpers.AssertNilError(t, err)

	expectedMatrix, err := Init(3, 3)
	test_helpers.AssertNilError(t, err)
	expectedMatrix.values[0] = []float64{27, 30, 33}
	expectedMatrix.values[1] = []float64{61, 68, 75}
	expectedMatrix.values[2] = []float64{95, 106, 117}

	areEqual := resultingMatrix.IsEqual(expectedMatrix)
	test_helpers.AssertEqual(t, true, areEqual)
}

// TestMatrix_MultiplyMatrix_IncompatibleSize tests the multiplication of a matrix by another.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMatrix_MultiplyMatrix_IncompatibleSize(t *testing.T) {
	firstMatrix, err := Init(3, 3)
	test_helpers.AssertNilError(t, err)

	secondMatrix, err := Init(2, 3)
	test_helpers.AssertNilError(t, err)

	controller := Controller{}

	_, err = controller.MultiplyMatrix(firstMatrix, secondMatrix)
	test_helpers.AssertNotNilError(t, err)
}
