package vector

// ProjectOnVector is a function to project one vector on another.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	The resulting vector.
//
func ProjectOnVector(vect1, vect2 *Vector) Vector {
	if !IsEqualDimension(vect1, vect2) {
		log.Fatalf(differentDimensions)
	}

	topConstant := DotProduct(vect1, vect2)
	bottomConstant := DotProduct(vect2, vect2)

	return ScalarMultiplication(vect2, topConstant/bottomConstant)
}

// Orthogonalize is a function to orthogonalize two vectors.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	The resulting vector.
//
func Orthogonalize(vect1, vect2 *Vector) Vector {
	vectAux := ProjectOnVector(vect1, vect2)
	return Sum(vect1, &vectAux, 1, -1)
}

// CheckOrtogonalVector is a function to check if two vectors are orthogonal to each other.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
// Returns:
// 	If the vectors are orthogonal to each other.
//
func IsOrthogonalVector(vect1, vect2 *Vector) bool {
	return DotProduct(vect1, vect2) == 0
}
