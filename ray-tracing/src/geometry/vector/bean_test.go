package vector

import (
	"fmt"
	"testing"
)


// TestVector_Init tests the instantiation of a Vector with positive dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_Init(t *testing.T) {
	dimension := 1
	vector, err := Init(dimension)
	if err != nil {
		t.Errorf("Vector failed to be instantiated with dimension: %d.", dimension)
	}
	if vector.Dimension() != dimension {
		t.Errorf("Vector instantiated with wrong dimension: %d %d.", dimension, vector.Dimension())
	}
}

// TestVector_Init_ZeroDimension tests the instantiation of a Vector with 0 dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_Init_ZeroDimension(t *testing.T) {
	dimension := 0
	vector, err := Init(dimension)
	if err != nil {
		t.Errorf("Vector failed to be Instantiated with dimension: %d.", dimension)
	}
	if vector.Dimension() != dimension {
		t.Errorf("Vector instantiated with wrong dimension: %d %d.", dimension, vector.Dimension())
	}
}

// TestVector_Init_NegativeDimension tests the instantiation of a Vector with negative dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_Init_NegativeDimension(t *testing.T) {
	dimension := -1
	_, err := Init(dimension)
	expectedErrorMessage := fmt.Sprintf("Invalid vector size %d.", dimension)
	if err == nil {
		t.Errorf("Vector instantiated with negative dimension: %d.", dimension)
	} else if err.Error() != expectedErrorMessage {
		t.Errorf(
			"Vector failed to be instantiated with negative dimension: %d but with wrong error message: \"%s\".",
			dimension, err.Error())
	}
}

// TestVector_IsEqualDimension tests if two vectors have equal dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_IsEqualDimension(t *testing.T) {
	dimension := 1
	firstVector, _ := Init(dimension)
	secondVector, _ := Init(dimension)
	if !firstVector.IsEqualDimension(secondVector) {
		t.Errorf("Vectors with different dimensions: %d and %d.", firstVector.Dimension(), secondVector.Dimension())
	}
}

// TestVector_IsEqualDimension_DifferentDimension tests if two vectors have different dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_IsEqualDimension_DifferentDimension(t *testing.T) {
	firstVector, _ := Init(1)
	secondVector, _ := Init(2)
	if firstVector.IsEqualDimension(secondVector) {
		t.Errorf("Vectors with equal dimensions: %d and %d.", firstVector.Dimension(), secondVector.Dimension())
	}
}

// TestVector_IsEqual tests if two vectors are equal.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_IsEqual(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{10, 20, 30}}
	secondVector := &Vector{coordinates: []float64{10, 20, 30}}
	if !firstVector.IsEqual(secondVector) {
		t.Errorf("Vectors are different: %v %v", firstVector.coordinates, secondVector.coordinates)
	}
}

// TestVector_IsEqual_Different tests if two vectors are different.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_IsEqual_Different(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{10, 20, 30}}
	secondVector := &Vector{coordinates: []float64{10, 30, 20}}
	if firstVector.IsEqual(secondVector) {
		t.Errorf("Vectors are equal: %v %v", firstVector.coordinates, secondVector.coordinates)
	}
}

// TestVector_IsEqual_DifferentDimensions tests if two vectors are different by their dimension.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_IsEqual_DifferentDimensions(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{10, 20, 30, 40}}
	secondVector := &Vector{coordinates: []float64{10, 20, 30}}
	if firstVector.IsEqual(secondVector) {
		t.Errorf("Vectors are equal: %v %v", firstVector.coordinates, secondVector.coordinates)
	}
}
