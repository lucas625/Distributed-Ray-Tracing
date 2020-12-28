package vector

// ProjectOnVector is a function to project one vector on another.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	The resulting vector.
//  An error.
//
func ProjectOnVector(vect1, vect2 *Vector) (*Vector, error) {
	if !IsEqualDimension(vect1, vect2) {
		return nil, differentDimensionError(vect1, vect2)
	}
	topConstant, _ := DotProduct(vect1, vect2)
	bottomConstant, _ := DotProduct(vect2, vect2)
	return ScalarMultiplication(vect2, topConstant/bottomConstant), nil
}

// Orthogonalize is a function to orthogonalize two vectors.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	The resulting vector.
//  An error.
//
func Orthogonalize(vect1, vect2 *Vector) (*Vector, error) {
	vectAux, err := ProjectOnVector(vect1, vect2)
	if err != nil {
		return nil, err
	}
	return Sum(vect1, vectAux, 1, -1)
}

// IsOrthogonalVector is a function to check if two vectors are orthogonal to each other.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	If the vectors are orthogonal to each other.
//  An error.
//
func IsOrthogonalVector(vect1, vect2 *Vector) (bool, error) {
	dotProduct, err := DotProduct(vect1, vect2)
	if err != nil {
		return false, err
	}
	return dotProduct == 0, nil
}
