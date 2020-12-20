package matrix

import (
	"testing"
)

// TestScalarMultiplication tests the multiplication of a matrix by a constant.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestScalarMultiplication(t *testing.T) {
	matrix, _ := Init(3, 2)
	matrix.Values[0] = []float64{1, 2}
	matrix.Values[1] = []float64{3, 4}
	matrix.Values[2] = []float64{5, 6}

	scalar := 3.0
	resultingMatrix, err := ScalarMultiplication(matrix, scalar)
	if err != nil {
		t.Errorf("Failed to multiply matrix %v by constant %v.", matrix, scalar)
	}

	expectedMatrix, _ := Init(3, 2)
	expectedMatrix.Values[0] = []float64{3, 6}
	expectedMatrix.Values[1] = []float64{9, 12}
	expectedMatrix.Values[2] = []float64{15, 18}

	if !resultingMatrix.IsEqual(expectedMatrix) {
		t.Errorf("Invalid multiplication of matrix %v by constant %v: %v.", matrix, scalar, resultingMatrix)
	}
}

// TestMultiplyMatrix tests the multiplication of a matrix by another.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestMultiplyMatrix(t *testing.T) {
	matrix1, _ := Init(3, 2)
	matrix1.Values[0] = []float64{1, 2}
	matrix1.Values[1] = []float64{3, 4}
	matrix1.Values[2] = []float64{5, 6}

	matrix2, _ := Init(2, 3)
	matrix2.Values[0] = []float64{7, 8, 9}
	matrix2.Values[1] = []float64{10, 11, 12}

	resultingMatrix, err := MultiplyMatrix(matrix1, matrix2)
	if err != nil {
		t.Errorf("Failed to multiply matrices %v %v.", matrix1, matrix2)
	}

	expectedMatrix, _ := Init(3, 3)
	expectedMatrix.Values[0] = []float64{27, 30, 33}
	expectedMatrix.Values[1] = []float64{61, 68, 75}
	expectedMatrix.Values[2] = []float64{95, 106, 117}


	if !resultingMatrix.IsEqual(expectedMatrix) {
		t.Errorf("Invalid multiplication of matrix %v by matrix %v: %v.", matrix1, matrix2, resultingMatrix)
	}
}
