package screen

// ColoredScreen is a class for image screen with colors.
//
// Members:
// 	Colors - the color matrix.
//
type ColoredScreen struct {
	Colors [][][]int
	Screen
}

// InitColoredScreen is a function to initialize a colored screen.
//
// Parameters:
// 	width  - the screen width.
//  height - the screen height.
//
// Returns:
// 	a colored Screen.
//
func InitColoredScreen(width, height int) ColoredScreen {
	colors := make([][][]int, height)
	for i := 0; i < height; i++ {
		colors[i] = make([][]int, width)
		for j := 0; j < width; j++ {
			colors[i][j] = make([]int, 3)
		}
	}
	return ColoredScreen{Screen: Screen{Width: width, Height: height}, Colors: colors}
}
