package matrix

// ScalarMultiplication is a function to multiply a Matrix by a constant.
//
// Parameters:
//  matrix - A pointer to a Matrix.
//  scalar - A constant.
//
// Returns:
// 	The resulting matrix.
//  An error.
//
func ScalarMultiplication(matrix *Matrix, scalar float64) (*Matrix, error) {
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

// MultiplyMatrix is a function to multiply two Matrices.
//
// Parameters:
// 	matrix1 - A pointer to a Matrix.
//  matrix2 - A pointer to a Matrix.
//
// Returns:
// 	The resulting matrix
//
func MultiplyMatrix(matrix1, matrix2 *Matrix) (*Matrix, error) {
	if matrix1.Columns != matrix2.Lines {
		return nil, incompatibleSize(matrix1, matrix2)
	}
	multipliedMatrix, _ := Init(matrix1.Lines, matrix2.Columns)
	for firstMatrixLineIndex := 0; firstMatrixLineIndex < matrix1.Lines; firstMatrixLineIndex++ {
		for secondMatrixColumnIndex := 0; secondMatrixColumnIndex < matrix2.Columns; secondMatrixColumnIndex++ {
			for firstMatrixColumnIndex := 0; firstMatrixColumnIndex < matrix1.Columns; firstMatrixColumnIndex++ {
				firstValue := matrix1.Values[firstMatrixLineIndex][firstMatrixColumnIndex]
				secondValue := matrix2.Values[firstMatrixColumnIndex][secondMatrixColumnIndex]
				multipliedMatrix.Values[firstMatrixLineIndex][secondMatrixColumnIndex] += firstValue * secondValue
			}
		}
	}
	return multipliedMatrix, nil
}
