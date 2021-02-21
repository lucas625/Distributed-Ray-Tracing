package color_matrix

import (
	"errors"
	"fmt"
)

// indexError is the error where we try to access an index out of the limits of the ColorMatrix.
//
// Parameters:
//	colorMatrix - The ColorMatrix.
//	lineIndex   - The index of the line.
//	columnIndex - The index of the column.
//
// Returns:
//  An Error.
//
func indexError(colorMatrix *ColorMatrix, lineIndex, columnIndex int) error {
	errorMessage := fmt.Sprintf(
		"Index out of limits of the color matrix. Expected from 0 0 to %v %v and got %v %v.",
		colorMatrix.Lines(), colorMatrix.Columns(), lineIndex, columnIndex)
	return errors.New(errorMessage)
}

// nonRGBColorError is the error where the color is not valid for RGB.
//
// Parameters:
//	color - The non RGB color.
//
// Returns:
//  An Error.
//
func nonRGBColorError(color []int) error {
	errorMessage := fmt.Sprintf("Non RGB color: %v.", color)
	return errors.New(errorMessage)
}