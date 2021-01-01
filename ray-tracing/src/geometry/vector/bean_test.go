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
	if vect.Dimension() != dimension {
		t.Errorf("Vector instantiated with wrong dimension: %d %d.", dimension, vect.Dimension())
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
	if vect.Dimension() != dimension {
		t.Errorf("Vector instantiated with wrong dimension: %d %d.", dimension, vect.Dimension())
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
	if !vector1.IsEqualDimension(vector2) {
		t.Errorf("Vectors with different dimensions: %d and %d.", vector1.Dimension(), vector2.Dimension())
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
	if vector1.IsEqualDimension(vector2) {
		t.Errorf("Vectors with equal dimensions.")
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
	vect1 := &Vector{coordinates: []float64{10, 20, 30}}
	vect2 := &Vector{coordinates: []float64{10, 20, 30}}
	if !vect1.IsEqual(vect2) {
		t.Errorf("Vectors are different: %v %v", vect1.coordinates, vect2.coordinates)
	}
}

// TestIsEqualDifferent tests if two vectors are different.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIsEqualDifferent(t *testing.T) {
	vect1 := &Vector{coordinates: []float64{10, 20, 30}}
	vect2 := &Vector{coordinates: []float64{10, 30, 20}}
	if vect1.IsEqual(vect2) {
		t.Errorf("Vectors are equal: %v %v", vect1.coordinates, vect2.coordinates)
	}
}

// TestIsEqualFailDifferentDimensions tests if two vectors are different by their dimension.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIsEqualFailDifferentDimensions(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{10, 20, 30, 40}}
	secondVector := &Vector{coordinates: []float64{10, 20, 30}}
	if firstVector.IsEqual(secondVector) {
		t.Errorf("Vectors are equal: %v %v", firstVector.coordinates, secondVector.coordinates)
	}
}
