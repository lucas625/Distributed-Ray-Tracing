package screen

// Screen is a class for screens.
//
// Members:
// 	width               - The number of x pixels on the screen.
// 	height              - The number of y pixels on the screen.
//
type Screen struct {
	width  int
	height int
}

// GetWidth gets the width of the Screen.
//
// Parameters:
// 	none
//
// Returns:
// 	The width of the Screen.
//
func (screen *Screen) GetWidth() int {
	return screen.width
}

// GetHeight gets the height of the Screen.
//
// Parameters:
// 	none
//
// Returns:
// 	The height of the Screen.
//
func (screen *Screen) GetHeight() int {
	return screen.height
}

// IsEqual checks if two screens are equal.
//
// Parameters:
// 	other - The other Screen.
//
// Returns:
// 	If the screens are equal.
//
func (screen *Screen) IsEqual(other *Screen) bool {
	return screen.GetWidth() == other.GetWidth() && screen.GetHeight() == other.GetHeight() &&
}

// Init initializes a Screen.
//
// Parameters:
// 	width  - The width of the Screen.
//  height - The height of the Screen.
//
// Returns:
// 	a Screen.
// 	an error.
//
func Init(width, height int) (*Screen, error) {
	if width <= 0 || height <= 0 {
		 return nil, screenSizeError(width, height)
	}
	screen := &Screen{width: width, height: height}
	return screen, nil
}

