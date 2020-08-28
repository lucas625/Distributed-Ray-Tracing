package matrix

// ScalarMultiplication is a function to multiply a Matrix by a constant.
//
// Parameters:
// 	matrix - A pointer to a Matrix.
//  scalar - A constant.
//
// Returns:
// 	The resulting matrix.
//  An error.
//
func (matrix *Matrix) ScalarMultiplication(scalar float64) (*Matrix, error) {
	matrixAux, err := Init(matrix.Lines, matrix.Columns)
	if err != nil {
		return nil, err
	}
	for lineIndex := 0; lineIndex < matrix.Lines; lineIndex++ {
		for columnIndex := 0; columnIndex < matrix.Columns; columnIndex++ {
			matrixAux.Values[lineIndex][columnIndex] = matrix.Values[lineIndex][columnIndex] * scalar
		}
	}
	return matrixAux, nil
}
