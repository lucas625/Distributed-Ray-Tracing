package matrix

import (
"errors"
"fmt"
)

// invalidSize is a function to get the error where a matrix has a invalid size for lines or columns.
//
// Parameters:
//	lines - The number of lines of the matrix.
//	columns - The number of columns of the matrix.
//
// Returns:
//  An Error.
//
func invalidSize(lines, columns int) error {
	errorMessage := fmt.Sprintf(
		"Invalid size for matrix. lines: %d and columns: %d.\n",
		lines,
		columns)
	return errors.New(errorMessage)
}

// incompatibleSize is a function to get the error where two matrices don't have a compatible size.
//
// Parameters:
//	matrix1 - The first matrix.
//	matrix2 - The second matrix.
//
// Returns:
//  An Error.
//
func incompatibleSize(matrix1, matrix2 *Matrix) error {
	errorMessage := fmt.Sprintf(
		"Incompatible size for matrices:\nFirst matrix: lines: %d and columns: %d." +
			"\nSecond matrix: lines: %d and columns: %d.\n",
		matrix1.Lines(), matrix1.Columns(), matrix2.Lines(), matrix2.Columns())
	return errors.New(errorMessage)
}

// indexError is a function to get the error where we try to access an index out of the matrix.
//
// Parameters:
//	matrix - The Matrix.
//	lineIndex - The index of the line.
//	columnIndex - The index of column.
//
// Returns:
//  An Error.
//
func indexError(matrix *Matrix, lineIndex, columnIndex int) error {
	errorMessage := fmt.Sprintf(
		"Index out of limits of the matrix. Expected from 0 0 to: %v %v and got %v %v.\n",
		matrix.Lines(), matrix.Columns(), lineIndex, columnIndex)
	return errors.New(errorMessage)
}
