package matrix

// Controller is a class for the matrix is controller.
//
// Members:
// 	none
//
type Controller struct {}

// BuildIdentity builds an identity Matrix.
//
// Parameters:
// 	size - The number of lines and columns of the identity Matrix.
//
// Returns:
// 	A Matrix.
//  An error.
//
func (_ *Controller) BuildIdentity(size int) (*Matrix, error) {
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
//  matrix - The Matrix.
//
// Returns:
// 	The transposed Matrix.
//
func (_ *Controller) Transpose(matrix *Matrix) *Matrix {
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
//  matrix - The Matrix.
//  scalar - A number.
//
// Returns:
// 	The resulting matrix.
//
func (_ *Controller) ScalarMultiplication(matrix *Matrix, scalar float64) *Matrix {
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
//  firstMatrix  - The first Matrix.
//  secondMatrix - The second Matrix.
//
// Returns:
// 	The resulting matrix
//
func (_ *Controller) MultiplyMatrix(firstMatrix *Matrix, secondMatrix *Matrix) (*Matrix, error) {
	if firstMatrix.Columns() != secondMatrix.Lines() {
		return nil, incompatibleSize(firstMatrix, secondMatrix)
	}
	newMatrix, _ := Init(firstMatrix.Lines(), secondMatrix.Columns())
	for firstMatrixLineIndex := 0; firstMatrixLineIndex < firstMatrix.Lines(); firstMatrixLineIndex++ {
		for secondMatrixColumnIndex := 0; secondMatrixColumnIndex < secondMatrix.Columns(); secondMatrixColumnIndex++ {
			totalSlotValue := 0.0
			for firstMatrixColumnIndex := 0; firstMatrixColumnIndex < firstMatrix.Columns(); firstMatrixColumnIndex++ {
				firstValue, _ := firstMatrix.GetValue(firstMatrixLineIndex, firstMatrixColumnIndex)
				secondValue, _ := secondMatrix.GetValue(firstMatrixColumnIndex, secondMatrixColumnIndex)
				totalSlotValue += firstValue * secondValue

			}
			newMatrix.SetValue(firstMatrixLineIndex, secondMatrixColumnIndex, totalSlotValue)
		}
	}
	return newMatrix, nil
}
