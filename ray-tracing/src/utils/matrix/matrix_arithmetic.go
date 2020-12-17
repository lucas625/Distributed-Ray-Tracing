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
	auxiliar_matrix, _ := Init(matrix1.Lines, matrix2.Columns)
	for first_matrix_line_index := 0; first_matrix_line_index < matrix1.Lines; first_matrix_line_index++ {
		for second_matrix_column_index := 0; second_matrix_column_index < matrix2.Columns; second_matrix_column_index++ {
			for first_matrix_column_index := 0; first_matrix_column_index < matrix1.Columns; first_matrix_column_index++ {
				first_value := matrix1.Values[first_matrix_line_index][first_matrix_column_index]
				second_value := matrix2.Values[first_matrix_column_index][second_matrix_column_index]
				auxiliar_matrix.Values[first_matrix_line_index][first_matrix_column_index] += first_value * second_value
			}
		}
	}
	return auxiliar_matrix, nil
}
