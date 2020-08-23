package matrix

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

// Matrix is a class for matrices.
//
// Members:
// 	Values  - Values of the matrix.
// 	Lines   - Number of lines of the matrix.
// 	Columns - Number of columns of the matrix.
//
type Matrix struct {
	Values  [][]float64
	Lines   int
	Columns int
}

// InitMatrix is a function to initialize a Matrix.
//
// Parameters:
// 	lin - The number of lines of the matrix.
// 	col - The number of columns of the matrix.
//
// Returns:
// 	A Matrix.
//	An error.
//
func InitMatrix(lines, columns int) (*Matrix, error) {
	if lines <= 0 || columns <= 0 {
		return nil, invalidSize(lines, columns)
	}
	matrix := Matrix{Values: make([][]float64, lin), Lines: lin, Columns: col}
	for i := range matrix.Values {
		matrix.Values[i] = make([]float64, col)
	}
	return &matrix, nil
}
