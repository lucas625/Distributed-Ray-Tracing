package matrix

import (
	"fmt"
	"testing"
)

// TestInitPositiveSize tests the instantiation of a matrix with positive lines and columns.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestInitPositiveSize(t *testing.T) {
	size := 1
	matrix, err := Init(size, size)
	if err != nil {
		t.Errorf("Matrix failed to be instantiated with size: %d.", size)
	}
	if matrix.lines != size {
		t.Errorf("Matrix instantiated with wrong lines: %d %d.", size, matrix.lines)
	}
	if matrix.columns != size {
		t.Errorf("Matrix instantiated with wrong columns: %d %d.", size, matrix.columns)
	}
	for line := 0; line < size; line++ {
		for column := 0; column < size; column++ {
			if matrix.Values()[line][column] != 0 {
				t.Errorf("Matrix instantiated with invalid values: %v.", matrix.Values())
			}
		}
	}
}

// TestInitZeroSize tests the instantiation of a matrix with zero size.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestInitZeroSize(t *testing.T) {
	size := 0
	_, err := Init(size, size)
	if err == nil {
		t.Errorf("Matrix instantiated with zero size: %d.", size)
	} else if err.Error() != fmt.Sprintf("Invalid size for matrix. lines: %d and columns: %d.\n", size, size) {
		t.Errorf(
			"Matrix failed to be instantiated with zero size: %d but with wrong error message: \"%s\".",
			size,
			err.Error())
	}
}

// TestInitNegativeSize tests the instantiation of a matrix with negative size.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestInitNegativeSize(t *testing.T) {
	size := -1
	_, err := Init(size, size)
	if err == nil {
		t.Errorf("Matrix instantiated with negative size: %d.", size)
	} else if err.Error() != fmt.Sprintf("Invalid size for matrix. lines: %d and columns: %d.\n", size, size) {
		t.Errorf(
			"Matrix failed to be instantiated with negative size: %d but with wrong error message: \"%s\".",
			size,
			err.Error())
	}
}

// TestBuildIdentityZeroSize tests the build identity matrix with zero size.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestBuildIdentityZeroSize(t *testing.T) {
	size := 0
	_, err := BuildIdentity(size)
	if err == nil {
		t.Errorf("Identity matrix instantiated with zero size: %d.", size)
	} else if err.Error() != fmt.Sprintf("Invalid size for matrix. lines: %d and columns: %d.\n", size, size) {
		t.Errorf(
			"Identity matrix failed to be instantiated with zero size: %d but with wrong error message: \"%s\".",
			size,
			err.Error())
	}
}


// TestBuildIdentityPositiveSize tests the build identity matrix with positive size.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestBuildIdentityPositiveSize(t *testing.T) {
	size := 3
	matrix, err := BuildIdentity(size)
	if err != nil {
		t.Errorf("Identity matrix failed to be instantiated with size: %d.", size)
	}
	if matrix.lines != size {
		t.Errorf("Identity matrix instantiated with wrong lines: %d %d.", size, matrix.lines)
	}
	if matrix.columns != size {
		t.Errorf("Identity matrix instantiated with wrong columns: %d %d.", size, matrix.columns)
	}
	for line := 0; line < size; line++ {
		for column := 0; column < size; column++ {
			if (line != column && matrix.Values()[line][column] != 0) ||
				(line == column && matrix.Values()[line][column] != 1) {
				t.Errorf("Identity matrix instantiated with invalid values: %v.", matrix.Values())
			}
		}
	}
}

// TestTranspose tests the transposition of a matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestTranspose(t *testing.T) {
	matrix, _ := Init(3, 2)
	matrix.Values()[0] = []float64{1, 2}
	matrix.Values()[1] = []float64{3, 4}
	matrix.Values()[2] = []float64{5, 6}

	resultingMatrix, err := Transpose(matrix)
	if err != nil {
		t.Errorf("Failed to transpose matrix %v.", matrix)
	}

	expectedMatrix, _ := Init(2, 3)
	expectedMatrix.Values()[0] = []float64{1, 3, 5}
	expectedMatrix.Values()[1] = []float64{2, 4, 6}


	if !resultingMatrix.IsEqual(expectedMatrix) {
		t.Errorf("Invalid transposition of matrix %v: %v.", matrix, resultingMatrix)
	}
}

// TestIsEqual tests the is equal of a matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIsEqual(t *testing.T) {
	firstMatrix, _ := Init(3, 2)
	firstMatrix.Values()[0] = []float64{1, 2}
	firstMatrix.Values()[1] = []float64{3, 4}
	firstMatrix.Values()[2] = []float64{5, 6}

	secondMatrix, _ := Init(3, 2)
	secondMatrix.Values()[0] = []float64{1, 2}
	secondMatrix.Values()[1] = []float64{3, 4}
	secondMatrix.Values()[2] = []float64{5, 6}

	if !firstMatrix.IsEqual(secondMatrix) {
		t.Errorf("Matrix are different of matrix %v: %v.", firstMatrix, secondMatrix)
	}
}

// TestIsEqualDifferent tests the is equal of a matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIsEqualDifferent(t *testing.T) {
	firstMatrix, _ := Init(3, 2)
	firstMatrix.Values()[0] = []float64{1, 2}
	firstMatrix.Values()[1] = []float64{3, 4}
	firstMatrix.Values()[2] = []float64{5, 6}

	secondMatrix, _ := Init(3, 2)
	secondMatrix.Values()[0] = []float64{10, 20}
	secondMatrix.Values()[1] = []float64{30, 40}
	secondMatrix.Values()[2] = []float64{50, 60}

	if firstMatrix.IsEqual(secondMatrix) {
		t.Errorf("Matrix are equal of matrix %v: %v.", firstMatrix, secondMatrix)
	}
}

// TestIsEqualDifferentLines tests the is equal of a matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIsEqualDifferentLines(t *testing.T) {
	firstMatrix, _ := Init(2, 2)
	secondMatrix, _ := Init(3, 2)

	if firstMatrix.IsEqual(secondMatrix) {
		t.Errorf("Matrix are equal of matrix %v: %v.", firstMatrix, secondMatrix)
	}
}

// TestIsEqualDifferentColumns tests the is equal of a matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIsEqualDifferentColumns(t *testing.T) {
	firstMatrix, _ := Init(3, 3)
	secondMatrix, _ := Init(3, 2)

	if firstMatrix.IsEqual(secondMatrix) {
		t.Errorf("Matrix are equal of matrix %v: %v.", firstMatrix, secondMatrix)
	}
}
