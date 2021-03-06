package light

import (
	"errors"
	"fmt"
)

// nonRGBColorError is the error where the color of the Light does not have 3 values.
//
// Parameters:
//	color - The color values.
//
// Returns:
//  An Error.
//
func nonRGBColorError(color []float64) error {
	errorMessage := fmt.Sprintf("There are not 3 color values: %d.", len(color))
	return errors.New(errorMessage)
}

// colorOutOfBoundsError is the error where a color coefficient is out of the bounds.
//
// Parameters:
//	color - The RGB color values.
//
// Returns:
//  An Error.
//
func colorOutOfBoundsError(color []float64) error {
	errorMessage := fmt.Sprintf("Color values out of interval [0,1]: %v.", color)
	return errors.New(errorMessage)
}
