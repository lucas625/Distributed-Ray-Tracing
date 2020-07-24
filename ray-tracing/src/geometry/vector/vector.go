package vector

// Vector is a class for vectors.
//
// Members:
// 	Coordinates - list of coordinates.
//
type Vector struct {
	Coordinates []float64
}

// IsEqualDimension is a function to check if two vectors are of the same dimension.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	If the vectors are of the same dimension.
//
func IsEqualDimension(vect1, vect2 *Vector) bool {
	return len(vect1.Coordinates) != len(vect2.Coordinates)
}

// InitVector is a function to initialize a Vector.
//
// Parameters:
// 	size - The size of the Vector.
//
// Returns:
// 	a Vector
//
func InitVector(size int) Vector {
	if size < 0 {
		negativeDimensionError(size)
	}
	return Vector{Coordinates: make([]float64, size)}
}
