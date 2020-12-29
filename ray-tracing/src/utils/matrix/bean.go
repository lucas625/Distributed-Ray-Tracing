package matrix

import "fmt"

// Matrix is a class for matrices.
//
// Members:
// 	values  - Values of the matrix.
// 	lines   - Number of Lines of the matrix.
// 	columns - Number of Columns of the matrix.
//
type Matrix struct {
	values  [][]float64
	lines   int
	columns int
}

// GetValue is a getter for the Matrix's values.
//
// Parameters:
// 	lineIndex - the index of the line.
// 	columnIndex - the index of the column.
//
// Returns:
// 	The value of the Matrix.
//  An error.
//
func (matrix *Matrix) GetValue(lineIndex, columnIndex int) (float64, error) {
	if lineIndex < 0 || lineIndex >= matrix.Lines() || columnIndex < 0 || columnIndex >= matrix.Columns() {
		return 0, indexOutOfLimits(matrix, lineIndex, columnIndex)
	}
	return matrix.values[lineIndex][columnIndex], nil
}

// SetValue is a setter for the Matrix's values.
//
// Parameters:
// 	lineIndex - the index of the line.
// 	columnIndex - the index of the column.
// 	value - the new value.
//
// Returns:
//  An error.
//
func (matrix *Matrix) SetValue(lineIndex, columnIndex int, value float64) error {
	if lineIndex < 0 || lineIndex >= matrix.Lines() || columnIndex < 0 || columnIndex >= matrix.Columns() {
		return indexOutOfLimits(matrix, lineIndex, columnIndex)
	}
	matrix.values[lineIndex][columnIndex] = value
	return nil
}

// Lines is a getter for the Matrix's lines.
//
// Parameters:
// 	none
//
// Returns:
// 	The lines of the Matrix.
//
func (matrix *Matrix) Lines() int {
	return matrix.lines
}

// Columns is a getter for the Matrix's columns.
//
// Parameters:
// 	none
//
// Returns:
// 	The columns of the Matrix.
//
func (matrix *Matrix) Columns() int {
	return matrix.columns
}

// IsEqual is a function to check if two matrices are equal.
//
// Parameters:
// 	other - The second matrix.
//
// Returns:
// 	A Matrix.
//
func (matrix *Matrix) IsEqual(other *Matrix) bool {
	if matrix.Lines() != other.Lines() || matrix.Columns() != other.Columns() {
		return false
	}
	for lineIndex := 0; lineIndex < matrix.Lines(); lineIndex++{
		for columnIndex := 0; columnIndex < matrix.Columns(); columnIndex++ {
			firstMatrixValue, _ := matrix.GetValue(lineIndex, columnIndex)
			secondMatrixValue, _ := other.GetValue(lineIndex, columnIndex)
			if firstMatrixValue != secondMatrixValue {
				return false
			}
		}
	}
	return true
}

// ToString is parses the matrix to string.
//
// Parameters:
// 	none
//
// Returns:
// 	The matrix as a string.
//
func (matrix *Matrix) ToString() string {
	return fmt.Sprintf("Lines: %v Columns: %v\n Matrix: %v\n", matrix.Lines(), matrix.Columns(), matrix.values)
}

// CopyAllValues gets all Values of the matrix as a deep copy.
//
// Parameters:
// 	none
//
// Returns:
// 	All the values of the matrix.
//
func (matrix *Matrix) CopyAllValues() [][]float64 {
	copiedMatrix := make([][]float64, matrix.Lines())
	for lineIndex := 0; lineIndex < matrix.Lines(); lineIndex++ {
		copiedMatrix[lineIndex] = make([]float64, matrix.Columns())
		for columnIndex := 0; columnIndex < matrix.Columns(); columnIndex++ {
			matrixValue, _ := matrix.GetValue(lineIndex, columnIndex)
			copiedMatrix[lineIndex][columnIndex] = matrixValue
		}
	}
	return copiedMatrix
}

// Init is a function to initialize a matrix.
//
// Parameters:
// 	lin - The number of lines of the matrix.
// 	col - The number of columns of the matrix.
//
// Returns:
// 	A Matrix.
//	An error.
//
func Init(lines, columns int) (*Matrix, error) {
	if lines <= 0 || columns <= 0 {
		return nil, invalidSize(lines, columns)
	}
	matrix := Matrix{values: make([][]float64, lines), lines: lines, columns: columns}
	for lineIndex := 0; lineIndex < lines; lineIndex++ {
		matrix.values[lineIndex] = make([]float64, columns)
	}
	return &matrix, nil
}
