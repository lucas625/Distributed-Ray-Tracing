package camera

import "github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"

// Controller is a class for controlling cameras.
//
// Members:
// 	none
//
type Controller struct {}

// NormalizeVectors normalizes the Camera vectors.
//
// Parameters:
//  camera - The Camera.
//
// Returns:
// 	none
//
func (*Controller) NormalizeVectors(camera *Camera) {
	vectorController := vector.Controller{}
	camera.SetLook(vectorController.Normalize(camera.look))
	camera.SetUp(vectorController.Normalize(camera.up))
	camera.SetRight(vectorController.Normalize(camera.right))
}

// CamToHomogeneousMatrix is a function to create the matrix ready(after transposition) to multiply the points.
//
// Parameters:
// 	cam - a Camera.
//
// Returns:
// 	a Matrix.
//
func CamToHomogeneousMatrix(cam *Camera) utils.Matrix {
	maux := utils.InitMatrix(3, 3)
	// placing vectors on the matrix on the right form
	maux.Values[0] = cam.Right.Coordinates
	maux.Values[1] = cam.Up.Coordinates
	maux.Values[2] = cam.Look.Coordinates
	// adding homogeneous and translation
	maux.Values = append(maux.Values, []float64{0, 0, 0, 1})
	pValues := cam.Pos.Coordinates
	for i := 0; i < 3; i++ {
		maux.Values[i] = append(maux.Values[i], pValues[i]*-1)
	}
	maux.Lines++
	maux.Columns++
	return maux
}

// CamToWorld is a function to create the matrix of camera to world.
//
// Parameters:
// 	cam - a Camera.
//
// Returns:
// 	a Matrix.
//
func CamToWorld(cam *Camera) utils.Matrix {
	maux := utils.InitMatrix(3, 3)
	// placing vectors on the matrix on the right form
	for j := 0; j < 3; j++ {
		maux.Values[j][0] = cam.Right.Coordinates[j]
		maux.Values[j][1] = cam.Up.Coordinates[j]
		maux.Values[j][2] = cam.Look.Coordinates[j]
	}
	// adding homogeneous and translation
	maux.Values = append(maux.Values, []float64{0, 0, 0, 1})
	pValues := cam.Pos.Coordinates
	for i := 0; i < 3; i++ {
		maux.Values[i] = append(maux.Values[i], pValues[i])
	}
	maux.Lines++
	maux.Columns++
	return maux
}

// InitCameraWithPoints is a function to initialize a Camera based only on its position and the target point.
//
// Parameters:
// 	pos    - the position of the camera.
// 	target - target Point.
//
// Returns:
// 	A Camera.
//
func InitCameraWithPoints(pos, target *entity.Point) Camera {
	look := entity.ExtractVector(pos, target)
	look = utils.NormalizeVector(&look)

	vectTemp := utils.Vector{Coordinates: []float64{0, 1, 0}}
	vectTemp = utils.NormalizeVector(&vectTemp)
	right := utils.VectorCrossProduct(&vectTemp, &look)
	right = utils.NormalizeVector(&right)
	up := utils.VectorCrossProduct(&look, &right)
	up = utils.NormalizeVector(&up)

	return InitCamera(*pos, look, up, right, 50.0, 1)
}
