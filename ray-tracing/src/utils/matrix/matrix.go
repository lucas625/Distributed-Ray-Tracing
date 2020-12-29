package matrix

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

// Values is a getter for the Matrix's values.
//
// Parameters:
// 	none
//
// Returns:
// 	The values of the Matrix.
//
func (matrix *Matrix) Values() [][]float64 {
	return matrix.values
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
// 	The lines of the Matrix.
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
	if matrix.lines != other.lines || matrix.columns != other.columns {
		return false
	}
	for lineIndex := 0; lineIndex < matrix.lines; lineIndex++{
		for columnIndex := 0; columnIndex < matrix.columns; columnIndex++ {
			if matrix.Values()[lineIndex][columnIndex] != other.Values()[lineIndex][columnIndex] {
				return false
			}
		}
	}
	return true
}

// BuildIdentity is a function to build an identity matrix.
//
// Parameters:
// 	size - The number of lines and columns of the matrix.
//
// Returns:
// 	A Matrix.
//  An error.
//
func BuildIdentity(size int) (*Matrix, error) {
	if size < 1 {
		return nil, invalidSize(size, size)
	}
	matrix, _ := Init(size, size)
	for lineIndex := 0; lineIndex < size; lineIndex++ {
		matrix.Values()[lineIndex][lineIndex] = 1
	}
	return matrix, nil
}

// Transpose is a function to transpose a Matrix.
//
// Parameters:
// 	matrix - the target Matrix.
//
// Returns:
// 	The transposed Matrix.
//
func Transpose(matrix *Matrix) (*Matrix, error) {
	transposedMatrix, err := Init(matrix.columns, matrix.lines)

	if err != nil {
		return nil, err
	}

	for lineIndex := 0; lineIndex < matrix.lines; lineIndex++ {
		for columnIndex := 0; columnIndex < matrix.columns; columnIndex++ {
			transposedMatrix.Values()[columnIndex][lineIndex] = matrix.Values()[lineIndex][columnIndex]
		}
	}

	return transposedMatrix, nil
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
		matrix.Values()[lineIndex] = make([]float64, columns)
	}
	return &matrix, nil
}
