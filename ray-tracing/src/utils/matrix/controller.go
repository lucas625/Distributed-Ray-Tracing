package matrix

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
		matrix.SetValue(lineIndex, lineIndex, 1)
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
func Transpose(matrix *Matrix) *Matrix {
	transposedMatrix, _ := Init(matrix.Columns(), matrix.Lines())

	for lineIndex := 0; lineIndex < matrix.Lines(); lineIndex++ {
		for columnIndex := 0; columnIndex < matrix.Columns(); columnIndex++ {
			matrixValue, _ := matrix.GetValue(lineIndex, columnIndex)
			transposedMatrix.SetValue(columnIndex, lineIndex, matrixValue)
		}
	}

	return transposedMatrix
}

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
	matrixAux, err := Init(matrix.Lines(), matrix.Columns())
	if err != nil {
		return nil, err
	}
	for lineIndex := 0; lineIndex < matrix.Lines(); lineIndex++ {
		for columnIndex := 0; columnIndex < matrix.Columns(); columnIndex++ {
			matrixValue, _ := matrix.GetValue(lineIndex, columnIndex)
			matrixAux.SetValue(lineIndex, columnIndex, matrixValue * scalar)
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
	if matrix1.Columns() != matrix2.Lines() {
		return nil, incompatibleSize(matrix1, matrix2)
	}
	multipliedMatrix, _ := Init(matrix1.Lines(), matrix2.Columns())
	for firstMatrixLineIndex := 0; firstMatrixLineIndex < matrix1.Lines(); firstMatrixLineIndex++ {
		for secondMatrixColumnIndex := 0; secondMatrixColumnIndex < matrix2.Columns(); secondMatrixColumnIndex++ {
			totalSlotValue := 0.0
			for firstMatrixColumnIndex := 0; firstMatrixColumnIndex < matrix1.Columns(); firstMatrixColumnIndex++ {
				firstValue, _ := matrix1.GetValue(firstMatrixLineIndex, firstMatrixColumnIndex)
				secondValue, _ := matrix2.GetValue(firstMatrixColumnIndex, secondMatrixColumnIndex)
				totalSlotValue += firstValue * secondValue

			}
			multipliedMatrix.SetValue(firstMatrixLineIndex, secondMatrixColumnIndex, totalSlotValue)
		}
	}
	return multipliedMatrix, nil
}
