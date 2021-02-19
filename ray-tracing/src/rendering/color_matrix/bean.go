package color_matrix

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/screen"
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
func (colorScreen *ColorMatrix) Lines() int {
	return len(colorScreen.colors)
}

// Columns gets the columns of the ColorMatrix.
//
// Parameters:
// 	none
//
// Returns:
// 	The columns of the ColorMatrix.
//
func (colorScreen *ColorMatrix) Columns() int {
	return len(colorScreen.colors[0])
}

// GetColors gets the colors of the ColorMatrix.
//
// Parameters:
// 	none
//
// Returns:
// 	The colors of the ColorMatrix.
//
func (colorScreen *ColorMatrix) GetColors() [][][]int {
	return colorScreen.colors
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
func (colorScreen *ColorMatrix) SetColor(lineIndex, columnIndex int, color []int) error {
	if lineIndex >= colorScreen.Lines() || lineIndex < 0 || columnIndex >= colorScreen.Columns() || columnIndex < 0 {
		return // Add here the error
	}
	if len(color) != 3 {
		return // Add here the error
	}
	for colorIndex := 0; colorIndex < 3; colorIndex++ {
		if color[colorIndex] < 0 || color[colorIndex] > 255 {
			return // Add here the error
		}
	}
	colorScreen.colors[lineIndex][columnIndex] = color
	return nil
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
