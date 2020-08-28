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
	resultingMatrix, err := matrix.ScalarMultiplication(scalar)
	if err != nil {
		t.Errorf("Failed to multiply matrix %v by constant %v.", matrix, scalar)
	}

	expectedMatrix, _ := Init(3, 2)
	for lineIndex := 0; lineIndex < expectedMatrix.Lines; lineIndex++ {
		for columnIndex := 0; columnIndex < expectedMatrix.Columns; columnIndex++ {
			expectedMatrix.Values[lineIndex][columnIndex] = scalar * matrix.Values[lineIndex][columnIndex]
		}
	}

	if !resultingMatrix.IsEqual(matrix) {
		t.Errorf("Invalid multiplication of matrix %v by constant %v: %v.", matrix, scalar, expectedMatrix)
	}
}
