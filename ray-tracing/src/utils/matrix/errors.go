package matrix

import (
"errors"
"fmt"
)

// invalidDimensionForHomogeneousCoordinates is the error where a Matrix has a invalid dimension for homogeneous
// coordinates.
//
// Parameters:
//	dimension - The dimension of the Matrix.
//
// Returns:
//  An Error.
//
func invalidDimensionForHomogeneousCoordinates(dimension int) error {
	errorMessage := fmt.Sprintf("Invalid dimension for homogeneous coodinates matrix: %d.", dimension)
	return errors.New(errorMessage)
}

// invalidSize is the error where a Matrix has a invalid size for lines or columns.
//
// Parameters:
//	lines   - The number of lines of the Matrix.
//	columns - The number of columns of the Matrix.
//
// Returns:
//  An Error.
//
func invalidSize(lines, columns int) error {
	errorMessage := fmt.Sprintf(
		"Invalid size for matrix. lines: %d and columns: %d.",
		lines,
		columns)
	return errors.New(errorMessage)
}

// incompatibleSize is the error where two matrices don't have the size compatible.
//
// Parameters:
//	firstMatrix  - The first Matrix.
//	secondMatrix - The second Matrix.
//
// Returns:
//  An Error.
//
func incompatibleSize(firstMatrix, secondMatrix *Matrix) error {
	errorMessage := fmt.Sprintf(
		"Incompatible size for matrices:\nFirst matrix: lines: %d and columns: %d." +
			"\nSecond matrix: lines: %d and columns: %d.",
		firstMatrix.Lines(), firstMatrix.Columns(), secondMatrix.Lines(), secondMatrix.Columns())
	return errors.New(errorMessage)
}

// indexError is the error where we try to access an index out of the limits of the Matrix.
//
// Parameters:
//	matrix      - The Matrix.
//	lineIndex   - The index of the line.
//	columnIndex - The index of the column.
//
// Returns:
//  An Error.
//
func indexError(matrix *Matrix, lineIndex, columnIndex int) error {
	errorMessage := fmt.Sprintf(
		"Index out of limits of the matrix. Expected from 0 0 to %v %v and got %v %v.",
		matrix.Lines(), matrix.Columns(), lineIndex, columnIndex)
	return errors.New(errorMessage)
}
