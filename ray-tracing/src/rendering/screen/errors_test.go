package screen

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestScreen_ScreenSizeError tests the error where a Screen has an invalid width or height.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestScreen_ScreenSizeError(t *testing.T) {
	width := -1
	height := -1
	expectedErrorMessage := fmt.Sprintf("Invalid size for screen: width: %d, height: %d.", width, height)
	err := screenSizeError(width, height)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
