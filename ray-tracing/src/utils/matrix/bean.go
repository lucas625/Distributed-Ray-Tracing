package matrix

import "fmt"

// Matrix is a class for matrices.
//
// Members:
// 	values  - Values of the matrix.
//
type Matrix struct {
	values  [][]float64
}

// GetValue is a getter for the Matrix's values.
//
// Parameters:
// 	lineIndex   - The index of the line.
// 	columnIndex - The index of the column.
//
// Returns:
// 	The value of the Matrix.
//  An error.
//
func (matrix *Matrix) GetValue(lineIndex, columnIndex int) (float64, error) {
	if lineIndex < 0 || lineIndex >= matrix.Lines() || columnIndex < 0 || columnIndex >= matrix.Columns() {
		return 0, indexError(matrix, lineIndex, columnIndex)
	}
	return matrix.values[lineIndex][columnIndex], nil
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
	return len(matrix.values)
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
	return len(matrix.values[0])
}

// SetValue is a setter for the Matrix's values.
//
// Parameters:
// 	lineIndex   - The index of the line.
// 	columnIndex - The index of the column.
// 	value       - The new value.
//
// Returns:
//  An error.
//
func (matrix *Matrix) SetValue(lineIndex, columnIndex int, value float64) error {
	if lineIndex < 0 || lineIndex >= matrix.Lines() || columnIndex < 0 || columnIndex >= matrix.Columns() {
		return indexError(matrix, lineIndex, columnIndex)
	}
	matrix.values[lineIndex][columnIndex] = value
	return nil
}

// IsEqual checks if two matrices are equal.
//
// Parameters:
// 	other - The second matrix.
//
// Returns:
// 	If the matrices are equal.
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

// ToString parses the Matrix to string.
//
// Parameters:
// 	none
//
// Returns:
// 	The Matrix as a string.
//
func (matrix *Matrix) ToString() string {
	return fmt.Sprintf("Lines: %v Columns: %v\n Matrix: %v\n", matrix.Lines(), matrix.Columns(), matrix.values)
}

// CopyAllValues gets all values of the Matrix as a copy.
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

// Init is the constructor.
//
// Parameters:
// 	lines   - The number of lines of the Matrix.
// 	columns - The number of columns of the Matrix.
//
// Returns:
// 	A Matrix.
//	An error.
//
func Init(lines, columns int) (*Matrix, error) {
	if lines <= 0 || columns <= 0 {
		return nil, invalidSize(lines, columns)
	}
	matrix := Matrix{values: make([][]float64, lines)}
	for lineIndex := 0; lineIndex < lines; lineIndex++ {
		matrix.values[lineIndex] = make([]float64, columns)
	}
	return &matrix, nil
}
