package screen

import (
	"errors"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/utils/matrix"
	"strconv"
)

package screen

import (
"errors"
"math"
"strconv"

"github.com/lucas625/Projeto-CG/src/utils"
)

// Screen is a class for screens.
//
// Members:
// 	width               - The number of x pixels on the screen.
// 	height              - The number of y pixels on the screen.
//  cameraToWorldMatrix - The matrix from camera coordinates to world coordinates.
//
type Screen struct {
	width         int
	height        int
	cameraToWorldMatrix *matrix.Matrix
}

// Init initializes a Screen.
//
// Parameters:
// 	width               - The width of the screen.
//  height              - The height of the screen.
//  cameraToWorldMatrix - The matrix from camera coordinates to world coordinates.
//
// Returns:
// 	a Screen.
//
func InitScreen(width, height int, cameraToWorldMatrix *matrix.Matrix) (*Screen, error) {
	if width < 0 || height < 0 {
		utils.ShowError(errors.New("Invalid Screen"), "width("+strconv.Itoa(width)+") or height("+strconv.Itoa(height)+") invalid for screen.")
	}
	screen := &Screen{width: width, height: height, cameraToWorldMatrix: cameraToWorldMatrix}
	return screen, nil
}

