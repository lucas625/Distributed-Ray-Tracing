package color_matrix

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestScreen_IndexError tests the error where we try to access an index out of the limits of the ColorMatrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestScreen_IndexError(t *testing.T) {
	colorMatrix := &ColorMatrix{colors: [][][]int{{{1, 2, 3}, {4, 5, 6}}}}
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the color matrix. Expected from 0 0 to %v %v and got %v %v.",
		colorMatrix.Lines(), colorMatrix.Columns(), -1, -1)
	err := indexError(colorMatrix, -1, -1)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestScreen_NonRGBColorError tests the error where the color is not valid for RGB.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestScreen_NonRGBColorError(t *testing.T) {
	color := []int{-1}
	expectedErrorMessage := fmt.Sprintf("Non RGB color: %v.", color)
	err := nonRGBColorError(color)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
