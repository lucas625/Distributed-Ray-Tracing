package matrix

import (
	"fmt"
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
	matrix, err := BuildIdentity(size)
	if err != nil {
		t.Errorf("Identity matrix failed to be instantiated with size: %d.", size)
	}
	if matrix.Lines() != size {
		t.Errorf("Identity matrix instantiated with wrong lines: %d %d.", size, matrix.Lines())
	}
	if matrix.Columns() != size {
		t.Errorf("Identity matrix instantiated with wrong columns: %d %d.", size, matrix.Columns())
	}
	for line := 0; line < size; line++ {
		for column := 0; column < size; column++ {
			matrixValue, _ := matrix.GetValue(line, column)
			if (line != column && matrixValue != 0) ||
				(line == column && matrixValue != 1) {
				t.Errorf("Identity matrix instantiated with invalid values: %v.", matrix.values)
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
	_, err := BuildIdentity(size)

	expectedErrorMessage := fmt.Sprintf("Invalid size for matrix. lines: %d and columns: %d.", size, size)

	if err == nil {
		t.Errorf("Identity matrix instantiated with zero size: %d.", size)
	} else if err.Error() != expectedErrorMessage {
		t.Errorf(
			"Identity matrix failed to be instantiated with zero size: %d but with wrong error message: \"%s\".",
			size, err.Error())
	}
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
	matrix, _ := Init(3, 2)
	matrix.values[0] = []float64{1, 2}
	matrix.values[1] = []float64{3, 4}
	matrix.values[2] = []float64{5, 6}

	resultingMatrix := matrix.Transpose()

	expectedMatrix, _ := Init(2, 3)
	expectedMatrix.values[0] = []float64{1, 3, 5}
	expectedMatrix.values[1] = []float64{2, 4, 6}


	if !resultingMatrix.IsEqual(expectedMatrix) {
		t.Errorf("Invalid transposition of matrix %v: %v.", matrix, resultingMatrix)
	}
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
	matrix, _ := Init(3, 2)
	matrix.values[0] = []float64{1, 2}
	matrix.values[1] = []float64{3, 4}
	matrix.values[2] = []float64{5, 6}

	scalar := 3.0
	resultingMatrix := matrix.ScalarMultiplication(scalar)

	expectedMatrix, _ := Init(3, 2)
	expectedMatrix.values[0] = []float64{3, 6}
	expectedMatrix.values[1] = []float64{9, 12}
	expectedMatrix.values[2] = []float64{15, 18}

	if !resultingMatrix.IsEqual(expectedMatrix) {
		t.Errorf("Invalid multiplication of matrix %v by constant %v: %v.", matrix, scalar, resultingMatrix)
	}
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
	matrix1, _ := Init(3, 2)
	matrix1.values[0] = []float64{1, 2}
	matrix1.values[1] = []float64{3, 4}
	matrix1.values[2] = []float64{5, 6}

	matrix2, _ := Init(2, 3)
	matrix2.values[0] = []float64{7, 8, 9}
	matrix2.values[1] = []float64{10, 11, 12}

	resultingMatrix, err := matrix1.MultiplyMatrix(matrix2)
	if err != nil {
		t.Errorf("Failed to multiply matrices %v %v.", matrix1, matrix2)
	}

	expectedMatrix, _ := Init(3, 3)
	expectedMatrix.values[0] = []float64{27, 30, 33}
	expectedMatrix.values[1] = []float64{61, 68, 75}
	expectedMatrix.values[2] = []float64{95, 106, 117}

	if !resultingMatrix.IsEqual(expectedMatrix) {
		t.Errorf("Invalid multiplication of matrix %v by matrix %v: %v.", matrix1, matrix2, resultingMatrix)
	}
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
	firstMatrix, _ := Init(3, 3)
	secondMatrix, _ := Init(2, 3)

	_, err := firstMatrix.MultiplyMatrix(secondMatrix)
	if err == nil {
		t.Errorf("It shouldn't be possible to multiply matrices with incompatible size %v %v.",
			firstMatrix, secondMatrix)
	}

}
