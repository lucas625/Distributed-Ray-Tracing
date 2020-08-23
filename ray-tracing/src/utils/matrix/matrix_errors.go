package matrix

import (
"errors"
"fmt"
)

// invalidSize is a function to get the error where a matrix has a invalid size for lines or columns.
//
// Parameters:
//	lines - The number of lines of the matrix.
//	columns - The number of columns of the matrix.
//
// Returns:
//  An Error.
//
func invalidSize(lines, columns int) error {
	errorMessage := fmt.Sprintf(
		"Invalid size for matrix. Lines: %d and Columns: %d.\n",
		lines,
		columns)
	return errors.New(errorMessage)
}
