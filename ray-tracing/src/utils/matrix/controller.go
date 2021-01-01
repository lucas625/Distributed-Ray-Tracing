package matrix

// BuildIdentity builds an identity Matrix.
//
// Parameters:
// 	size - The number of lines and columns of the identity Matrix.
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

// Transpose transposes a Matrix.
//
// Parameters:
// 	none
//
// Returns:
// 	The transposed Matrix.
//
func (matrix *Matrix) Transpose() *Matrix {
	transposedMatrix, _ := Init(matrix.Columns(), matrix.Lines())
	for lineIndex := 0; lineIndex < matrix.Lines(); lineIndex++ {
		for columnIndex := 0; columnIndex < matrix.Columns(); columnIndex++ {
			matrixValue, _ := matrix.GetValue(lineIndex, columnIndex)
			transposedMatrix.SetValue(columnIndex, lineIndex, matrixValue)
		}
	}
	return transposedMatrix
}

// ScalarMultiplication multiplies a Matrix by a constant.
//
// Parameters:
//  scalar - A number.
//
// Returns:
// 	The resulting matrix.
//
func (matrix *Matrix) ScalarMultiplication(scalar float64) *Matrix {
	newMatrix, _ := Init(matrix.Lines(), matrix.Columns())
	for lineIndex := 0; lineIndex < matrix.Lines(); lineIndex++ {
		for columnIndex := 0; columnIndex < matrix.Columns(); columnIndex++ {
			matrixValue, _ := matrix.GetValue(lineIndex, columnIndex)
			newMatrix.SetValue(lineIndex, columnIndex, matrixValue * scalar)
		}
	}
	return newMatrix
}

// MultiplyMatrix multiplies two Matrices.
//
// Parameters:
//  secondMatrix - The second Matrix.
//
// Returns:
// 	The resulting matrix
//
func (matrix *Matrix) MultiplyMatrix(secondMatrix *Matrix) (*Matrix, error) {
	if matrix.Columns() != secondMatrix.Lines() {
		return nil, incompatibleSize(matrix, secondMatrix)
	}
	newMatrix, _ := Init(matrix.Lines(), secondMatrix.Columns())
	for firstMatrixLineIndex := 0; firstMatrixLineIndex < matrix.Lines(); firstMatrixLineIndex++ {
		for secondMatrixColumnIndex := 0; secondMatrixColumnIndex < secondMatrix.Columns(); secondMatrixColumnIndex++ {
			totalSlotValue := 0.0
			for firstMatrixColumnIndex := 0; firstMatrixColumnIndex < matrix.Columns(); firstMatrixColumnIndex++ {
				firstValue, _ := matrix.GetValue(firstMatrixLineIndex, firstMatrixColumnIndex)
				secondValue, _ := secondMatrix.GetValue(firstMatrixColumnIndex, secondMatrixColumnIndex)
				totalSlotValue += firstValue * secondValue

			}
			newMatrix.SetValue(firstMatrixLineIndex, secondMatrixColumnIndex, totalSlotValue)
		}
	}
	return newMatrix, nil
}
