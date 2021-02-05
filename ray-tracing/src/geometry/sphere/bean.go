package sphere

// Sphere is a class for spheres.
//
// Members:
// 	centerPointIndex - The index of the center point of the Sphere on the points list.
//  radius           - The radius of the Sphere.
//
type Sphere struct {
	centerPointIndex int
	radius float64
}

// GetCenterPointIndex gets the center point index of the Sphere.
//
// Parameters:
// 	none
//
// Returns:
// 	The index of the center point of the Sphere on the points list.
//
func (sphere *Sphere) GetCenterPointIndex() int {
	return sphere.centerPointIndex
}

// GetRadius gets the center point index of the Sphere.
//
// Parameters:
// 	none
//
// Returns:
// 	The radius of the Sphere.
//
func (sphere *Sphere) GetRadius() float64 {
	return sphere.radius
}

// IsEqual checks if two spheres are equal.
//
// Parameters:
// 	other - The other Sphere
//
// Returns:
// 	If the spheres are equal.
//
func (sphere *Sphere) IsEqual(other *Sphere) bool {
	return sphere.GetCenterPointIndex() == other.GetCenterPointIndex() && sphere.GetRadius() == other.GetRadius()
}

// Init initializes a Sphere.
//
// Parameters:
// 	center - The index of the center point of the Sphere on the points list.
//  radius - The radius of the Sphere.
//
// Returns:
// 	A Sphere.
//
func Init(centerPointIndex int, radius float64) *Sphere {
	return &Sphere{centerPointIndex: centerPointIndex, radius: radius}
}
