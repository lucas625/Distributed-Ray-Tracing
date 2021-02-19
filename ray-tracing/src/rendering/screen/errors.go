package screen

import (
	"errors"
	"fmt"
)

// screenSizeError is the error where a Screen has an invalid width or height.
//
// Parameters:
//	width  - The width of the Screen.
//	height - The height of the Screen.
//
// Returns:
//  An Error.
//
func screenSizeError(width, height int) error {
	errorMessage := fmt.Sprintf("Invalid size for screen: width: %d, height: %d.", width, height)
	return errors.New(errorMessage)
}
