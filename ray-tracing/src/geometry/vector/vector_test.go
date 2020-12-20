package vector

import (
	"fmt"
	"testing"
)


// TestInitPositiveDimension tests the instantiation of a vector with positive dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestInitPositiveDimension(t *testing.T) {
	dimension := 1
	vect, err := Init(dimension)
	if err != nil {
		t.Errorf("Vector failed to be instantiated with dimension: %d.", dimension)
	}
	if !(len(vect.Coordinates) == dimension) {
		t.Errorf("Vector instantiated with wrong dimension: %d %d.", dimension, len(vect.Coordinates))
	}
}

// TestInitZeroDimension tests the instantiation of a vector with 0 dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestInitZeroDimension(t *testing.T) {
	dimension := 0
	vect, err := Init(dimension)
	if err != nil {
		t.Errorf("Vector failed to be Instantiated with dimension: %d.", dimension)
	}
	if !(len(vect.Coordinates) == dimension) {
		t.Errorf("Vector instantiated with wrong dimension: %d %d.", dimension, len(vect.Coordinates))
	}
}

// TestInitNegativeDimension tests the instantiation of a vector with negative dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestInitNegativeDimension(t *testing.T) {
	dimension := -1
	_, err := Init(dimension)
	if err == nil {
		t.Errorf("Vector instantiated with negative dimension: %d.", dimension)
	} else if err.Error() != fmt.Sprintf("Invalid vector size %d.", dimension) {
		t.Errorf(
			"Vector failed to be instantiated with negative dimension: %d but with wrong error message: \"%s\".",
			dimension,
			err.Error())
	}
}

// TestIsEqualDimension tests if two vectors have equal dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIsEqualDimension(t *testing.T) {
	dimension := 1
	vector1, _ := Init(dimension)
	vector2, _ := Init(dimension)
	if !IsEqualDimension(vector1, vector2) {
		t.Errorf("Vectors with different dimensions: %d and %d.", len(vector1.Coordinates), len(vector2.Coordinates))
	}
}

// TestIsDifferentDimension tests if two vectors have different dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIsDifferentDimension(t *testing.T) {
	vector1, _ := Init(1)
	vector2, _ := Init(2)
	if IsEqualDimension(vector1, vector2) {
		t.Errorf("Vectors with equal dimensions.")
	}
}

// TestCoordinates tests the coordinates of a vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestCoordinates(t *testing.T) {
	vect := Vector{Coordinates: []float64{10, 20, 30}}
	if len(vect.Coordinates) != 3 {
		t.Errorf("Vector with unexpected coordinates: %v", vect.Coordinates)
	}
}

// TestIsEqual tests if two vectors are equal.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIsEqual(t *testing.T) {
	vect1 := &Vector{Coordinates: []float64{10, 20, 30}}
	vect2 := &Vector{Coordinates: []float64{10, 20, 30}}
	if !IsEqual(vect1, vect2) {
		t.Errorf("Vectors are different: %v %v", vect1.Coordinates, vect2.Coordinates)
	}
}

// TestIsDifferent tests if two vectors are different.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIsDifferent(t *testing.T) {
	vect1 := &Vector{Coordinates: []float64{10, 20, 30}}
	vect2 := &Vector{Coordinates: []float64{10, 30, 20}}
	if IsEqual(vect1, vect2) {
		t.Errorf("Vectors are equal: %v %v", vect1.Coordinates, vect2.Coordinates)
	}
}
