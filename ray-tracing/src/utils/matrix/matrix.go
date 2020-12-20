package matrix

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

// IsEqual is a function to check if two matrices are equal.
//
// Parameters:
// 	other - The second matrix.
//
// Returns:
// 	A Matrix.
//
func (matrix *Matrix) IsEqual(other *Matrix) bool {
	if matrix.Lines != other.Lines || matrix.Columns != other.Columns {
		return false
	}
	for lineIndex := 0; lineIndex < matrix.Lines; lineIndex++{
		for columnIndex := 0; columnIndex < matrix.Columns; columnIndex++ {
			if matrix.Values[lineIndex][columnIndex] != other.Values[lineIndex][columnIndex] {
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
		matrix.Values[lineIndex][lineIndex] = 1
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
	transposedMatrix, err := Init(matrix.Columns, matrix.Lines)

	if err != nil {
		return nil, err
	}

	for lineIndex := 0; lineIndex < matrix.Lines; lineIndex++ {
		for columnIndex := 0; columnIndex < matrix.Columns; columnIndex++ {
			transposedMatrix.Values[columnIndex][lineIndex] = matrix.Values[lineIndex][columnIndex]
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
	matrix := Matrix{Values: make([][]float64, lines), Lines: lines, Columns: columns}
	for lineIndex := range matrix.Values {
		matrix.Values[lineIndex] = make([]float64, columns)
	}
	return &matrix, nil
}
