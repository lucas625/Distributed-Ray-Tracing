package screen

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestScreen_SizeError tests the error where a Screen has an invalid width or height.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestScreen_SizeError(t *testing.T) {
	width := -1
	height := -1
	expectedErrorMessage := fmt.Sprintf("Invalid size for screen: width: %d, height: %d.", width, height)
	err := sizeError(width, height)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestScreen_PixelExtraValueError tests the error where we try to add an invalid value to a pixel.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestScreen_PixelExtraValueError(t *testing.T) {
	additionalY := 1.0
	additionalX := 1.0

	expectedErrorMessage := fmt.Sprintf("Pixel extra value error. Expected [0,1), and got %v %v.",
		additionalY, additionalX)
	err := pixelExtraValueError(additionalY, additionalX)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestScreen_PixelIndexError tests the error where we try to access a pixel out of the limits of the Screen.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestScreen_PixelIndexError(t *testing.T) {
	screen, err := Init(16, 9)
	test_helpers.AssertNilError(t, err)
	pixelLineIndex := 20
	pixelColumnIndex := -1

	expectedErrorMessage := fmt.Sprintf(
		"Pixel out of limits of the screen. Expected from 0 0 to %v %v, and got %v %v.",
		screen.GetHeight(), screen.GetWidth(), pixelLineIndex, pixelColumnIndex)
	err = pixelIndexError(screen, pixelLineIndex, pixelColumnIndex)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
