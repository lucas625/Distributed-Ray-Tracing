package color_matrix

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/screen"
	"reflect"
)

// ColorMatrix is a class for a ColorMatrix.
//
// Members:
// 	colors - The RGB colors matrix.
//
type ColorMatrix struct {
	colors [][][]int
}

// Lines gets the lines of the ColorMatrix.
//
// Parameters:
// 	none
//
// Returns:
// 	The lines of the ColorMatrix.
//
func (colorMatrix *ColorMatrix) Lines() int {
	return len(colorMatrix.colors)
}

// Columns gets the columns of the ColorMatrix.
//
// Parameters:
// 	none
//
// Returns:
// 	The columns of the ColorMatrix.
//
func (colorMatrix *ColorMatrix) Columns() int {
	return len(colorMatrix.colors[0])
}

// GetColors gets the colors of the ColorMatrix.
//
// Parameters:
// 	none
//
// Returns:
// 	The colors of the ColorMatrix.
//
func (colorMatrix *ColorMatrix) GetColors() [][][]int {
	return colorMatrix.colors
}

// SetColor sets the a color of the ColorMatrix.
//
// Parameters:
// 	lineIndex   - The line index of the color.
// 	columnIndex - The column index of the color.
// 	color       - The color.
//
// Returns:
// 	none
//
func (colorMatrix *ColorMatrix) SetColor(lineIndex, columnIndex int, color []int) error {
	if lineIndex >= colorMatrix.Lines() || lineIndex < 0 || columnIndex >= colorMatrix.Columns() || columnIndex < 0 {
		return indexError(colorMatrix, lineIndex, columnIndex)
	}
	if len(color) != 3 {
		return nonRGBColorError(color)
	}
	for colorIndex := 0; colorIndex < 3; colorIndex++ {
		if color[colorIndex] < 0 || color[colorIndex] > 255 {
			return nonRGBColorError(color)
		}
	}
	colorMatrix.colors[lineIndex][columnIndex] = color
	return nil
}

// IsEqual checks if two color matrices are equal.
//
// Parameters:
// 	other - the other matrix.
//
// Returns:
// 	If the two color matrices are equal.
//
func (colorMatrix *ColorMatrix) IsEqual(other *ColorMatrix) bool {
	return reflect.DeepEqual(colorMatrix.GetColors(), other.GetColors())
}

// Init initializes a ColorMatrix.
//
// Parameters:
// 	targetScreen - The screen of the ColorMatrix.
//
// Returns:
// 	A ColorMatrix.
//
func Init(targetScreen *screen.Screen) *ColorMatrix {
	colors := make([][][]int, targetScreen.GetHeight())
	for lineIndex := 0; lineIndex < targetScreen.GetHeight(); lineIndex++ {
		colors[lineIndex] = make([][]int, targetScreen.GetWidth())
		for columnIndex := 0; columnIndex < targetScreen.GetWidth(); columnIndex++ {
			colors[lineIndex][columnIndex] = make([]int, 3)
		}
	}
	return &ColorMatrix{colors: colors}
}
