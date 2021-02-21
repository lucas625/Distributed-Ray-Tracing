package screen

import (
	"errors"
	"fmt"
)

// sizeError is the error where a Screen has an invalid width or height.
//
// Parameters:
//	width  - The width of the Screen.
//	height - The height of the Screen.
//
// Returns:
//  An Error.
//
func sizeError(width, height int) error {
	errorMessage := fmt.Sprintf("Invalid size for screen: width: %d, height: %d.", width, height)
	return errors.New(errorMessage)
}

// pixelExtraValueError is the error where we try to add an invalid value to a pixel.
//
// Parameters:
//  additionalY      - The additional value to the pixel coordinate on y [0,1).
//  additionalX      - The additional value to the pixel coordinate on x [0,1).
//
// Returns:
//  An Error.
//
func pixelExtraValueError(additionalY, additionalX float64) error {
	errorMessage := fmt.Sprintf("Pixel extra value error. Expected [0,1], and got %v %v.",
		additionalY, additionalX)
	return errors.New(errorMessage)
}

// pixelIndexError is the error where we try to access a pixel out of the limits of the Screen.
//
// Parameters:
//	screen           - The Screen that contains the pixel.
//	pixelLineIndex   - The line index of the pixel.
//	pixelColumnIndex - The column index of the pixel.
//
// Returns:
//  An Error.
//
func pixelIndexError(screen *Screen, pixelLineIndex, pixelColumnIndex int) error {
	errorMessage := fmt.Sprintf("Pixel out of limits of the screen. Expected from 0 0 to %v %v, and got %v %v.",
		screen.GetHeight(), screen.GetWidth(), pixelLineIndex, pixelColumnIndex)
	return errors.New(errorMessage)
}
