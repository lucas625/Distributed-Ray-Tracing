package light

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestLight_NonRGBColorError is the error where the color of the Light does not have 3 values.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestLight_NonRGBColorError(t *testing.T) {
	color := []float64{255, 100}
	expectedErrorMessage := fmt.Sprintf("There are not 3 color values: %d.", len(color))
	err := nonRGBColorError(color)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestLight_ColorOutOfBoundsError tests the error where a color coefficient is out of the bounds.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestLight_ColorOutOfBoundsError(t *testing.T) {
	color := []float64{255, 100, 0}
	expectedErrorMessage := fmt.Sprintf("Color values out of interval [0,1]: %v.", color)
	err := colorOutOfBoundsError(color)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
