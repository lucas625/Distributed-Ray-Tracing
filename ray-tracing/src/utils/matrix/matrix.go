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
	for i := 0; i < size; i++ {
		matrix.Values[i][i] = 1
	}
	return matrix, nil
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
	for i := range matrix.Values {
		matrix.Values[i] = make([]float64, columns)
	}
	return &matrix, nil
}
