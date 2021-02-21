package color_matrix

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/screen"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"reflect"
	"testing"
)

// TestColorMatrix_Init tests the instantiation of a ColorMatrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestColorMatrix_Init(t *testing.T) {
	targetScreen, err := screen.Init(2, 1)
	test_helpers.AssertNilError(t, err)

	colorMatrix := Init(targetScreen)
	expectedColorMatrix := &ColorMatrix{colors: [][][]int{{{0,0,0}, {0,0,0}}}}
	test_helpers.AssertEqual(t, true, expectedColorMatrix.IsEqual(colorMatrix))
}

// TestColorMatrix_SetColor tests the set color of a ColorMatrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestColorMatrix_SetColor(t *testing.T) {
	targetScreen, err := screen.Init(16, 9)
	test_helpers.AssertNilError(t, err)

	newColor := []int{255, 255, 255}

	colorMatrix := Init(targetScreen)
	err = colorMatrix.SetColor(2, 3, newColor)
	test_helpers.AssertNilError(t, err)

	receivedColor := colorMatrix.GetColors()[2][3]

	test_helpers.AssertEqual(t, true, reflect.DeepEqual(newColor, receivedColor))
}

// TestColorMatrix_SetColor_IndexError tests the set color of a ColorMatrix when there is an index error.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestColorMatrix_SetColor_IndexError(t *testing.T) {
	targetScreen, err := screen.Init(16, 9)
	test_helpers.AssertNilError(t, err)
	newColor := []int{255, 255, 255}

	colorMatrix := Init(targetScreen)
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the color matrix. Expected from 0 0 to %v %v and got %v %v.",
		colorMatrix.Lines(), colorMatrix.Columns(), -1, 3)
	err = colorMatrix.SetColor(-1, 3, newColor)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestColorMatrix_SetColor_NonRGBColorError tests the set color of a ColorMatrix when there is a non RGB color error.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestColorMatrix_SetColor_NonRGBColorError(t *testing.T) {
	targetScreen, err := screen.Init(16, 9)
	test_helpers.AssertNilError(t, err)
	colorMatrix := Init(targetScreen)

	colorWithLessThan3Values := []int{255, 255}
	expectedErrorMessage := fmt.Sprintf("Non RGB color: %v.", colorWithLessThan3Values)
	err = colorMatrix.SetColor(3, 3, colorWithLessThan3Values)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())

	outOfLimitsColor := []int{255, 255, 256}
	expectedErrorMessage = fmt.Sprintf("Non RGB color: %v.", outOfLimitsColor)
	err = colorMatrix.SetColor(3, 3, outOfLimitsColor)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
