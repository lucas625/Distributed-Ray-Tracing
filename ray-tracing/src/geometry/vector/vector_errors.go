package vector

import (
	"log"
)

// differentDimensionError is a function to print the error where two vectors do not have the same dimension.
//
// Parameters:
// 	vect1 - The first vector.
// 	vect2 - The second vector.
//
func differentDimensionError(vect1, vect2 *Vector) {
	log.Fatalf("Invalid dimension of vector. Expected: %d and got: %d.\n", len(vect1.Coordinates), len(vect2.Coordinates))
}

// non3DError is a function to print the error where a vector is not 3D.
//
// Parameters:
// 	vect - The first vector.
//
func non3DError(vect *Vector) {
	log.Fatalf("Invalid dimension of vector. Expected 3D and got %d.\n", len(vect.Coordinates))
}

// negativeDimensionError is a function to print the error where a vector has negative dimension.
//
// Parameters:
// 	dimension - The dimension of the vector.
//
func negativeDimensionError(dimension int) {
	log.Fatalf("Invalid vector size %d.\n", dimension)
}
