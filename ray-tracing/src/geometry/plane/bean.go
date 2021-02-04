package plane

// Plane is a class for planes using formula AX + BY + CZ + D = 0.
//
// Members:
// 	xCoefficient - Multiplies the X of the plane is cartesian equation.
// 	yCoefficient - Multiplies the Y of the plane is cartesian equation.
// 	zCoefficient - Multiplies the Z of the plane is cartesian equation.
//  isolatedTerm - The "D" constant of the plane is cartesian equation.
//
type Plane struct {
	xCoefficient float64
	yCoefficient float64
	zCoefficient float64
	isolatedTerm float64
}

// GetXCoefficient gets the x coefficient of the plane is cartesian equation.
//
// Parameters:
// 	none
//
// Returns:
// 	The xCoefficient.
//
func (plane *Plane) GetXCoefficient() float64 {
	return plane.xCoefficient
}

// GetYCoefficient gets the y coefficient of the plane is cartesian equation.
//
// Parameters:
// 	none
//
// Returns:
// 	The yCoefficient.
//
func (plane *Plane) GetYCoefficient() float64 {
	return plane.yCoefficient
}

// GetZCoefficient gets the z coefficient of the plane is cartesian equation.
//
// Parameters:
// 	none
//
// Returns:
// 	The zCoefficient.
//
func (plane *Plane) GetZCoefficient() float64 {
	return plane.zCoefficient
}

// GetIsolatedTerm gets isolated term of the plane is cartesian equation.
//
// Parameters:
// 	none
//
// Returns:
// 	The isolatedTerm.
//
func (plane *Plane) GetIsolatedTerm() float64 {
	return plane.isolatedTerm
}

// Init initializes a Plane.
//
// Parameters:
// 	xCoefficient - Multiplies the X in the plane is cartesian equation.
// 	yCoefficient - Multiplies the Y in the plane is cartesian equation.
// 	zCoefficient - Multiplies the Z in the plane is cartesian equation.
//  isolatedTerm - The "D" constant in the plane is cartesian equation.
//
// Returns:
// 	A Plane.
//
func Init(xCoefficient, yCoefficient, zCoefficient, isolatedTerm float64) *Plane {
	return &Plane{xCoefficient: xCoefficient, yCoefficient: yCoefficient, zCoefficient: zCoefficient,
		isolatedTerm: isolatedTerm}
}
