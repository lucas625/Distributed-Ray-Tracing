package screen

import (
	"errors"
	"math"
	"strconv"
)

// PixelToWorld is a function to get the position of a pixel in world coordinates.
//
// Parameters:
// 	x        - position of the pixel.
//  y        - position of the pixel.
//  d        - distance viewport to cam.
//  camWorld - the matrix camera to world.
//  px       - the additional on x (0->1)
//  py       - the additional on y (0->1)
//  fov      - field of view in degrees.
//
// Returns:
// 	a Vector.
//
func (sc *Screen) PixelToWorld(x, y int, d float64, px, py, fov float64) utils.Vector {
	if x >= sc.Height || y >= sc.Width {
		utils.ShowError(errors.New("Invalid Pixel"), "X("+strconv.Itoa(x)+") or Y("+strconv.Itoa(y)+") invalid for screen("+strconv.Itoa(sc.Height)+", "+strconv.Itoa(sc.Width)+").")
	}
	camWorld := sc.CamToWorld

	aspectRatio := float64(sc.Width) / float64(sc.Height)
	alpha := (fov / 2) * math.Pi / 180.0
	z := d

	camerax := (2*(float64(x)+px)/float64(sc.Width) - 1) * aspectRatio * math.Tan(alpha)
	cameray := (1 - 2*(float64(y)+py)/float64(sc.Height)) * math.Tan(alpha)

	v := utils.InitVector(3)

	v.Coordinates[0] = camerax
	v.Coordinates[1] = cameray
	v.Coordinates[2] = z

	vMat := utils.VectorToHomogeneousCoord(&v)

	vMatPos := utils.MultMatrix(camWorld, &vMat)
	for i := 0; i < 3; i++ {
		v.Coordinates[i] = vMatPos.Values[i][0]
	}
	vNormalized := utils.NormalizeVector(&v)

	return vNormalized
}
