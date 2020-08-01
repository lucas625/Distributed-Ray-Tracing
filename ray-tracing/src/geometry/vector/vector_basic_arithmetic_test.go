package vector

import (
	"testing"
)


// TestScalarMultiplication tests the scalar multiplication of a vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestScalarMultiplication(t *testing.T) {
	vect := &Vector{Coordinates: []float64{10, 20, 30}}
	multiplied_vector, err := ScalarMultiplication(vect, 2)
	if err != nil {
		t.Errorf(err.Error())
	}
	expectedVector := &Vector{Coordinates: []float64{20, 40, 60}}
	if !IsEqual(expectedVector, multiplied_vector) {
		t.Errorf("Vectors are different: %v %v", expectedVector.Coordinates, multiplied_vector.Coordinates)
	}
}

// TestSum tests the sum of two vectors.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestSum(t *testing.T) {
	vect1 := &Vector{Coordinates: []float64{10, 20, 30}}
	vect2 := &Vector{Coordinates: []float64{5, 10, 20}}
	vector_sum, err := Sum(vect1, vect2, 1, 1)
	if err != nil {
		t.Errorf(err.Error())
	}
	expectedVector := &Vector{Coordinates: []float64{15, 30, 50}}
	if !IsEqual(expectedVector, vector_sum) {
		t.Errorf("Vectors are different: %v %v", expectedVector.Coordinates, vector_sum.Coordinates)
	}
}

// TestSubtraction tests the subtraction of two vectors.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestSubtraction(t *testing.T) {
	vect1 := &Vector{Coordinates: []float64{10, 20, 30}}
	vect2 := &Vector{Coordinates: []float64{5, 10, 20}}
	vector_sum, err := Sum(vect1, vect2, 3, -1)
	if err != nil {
		t.Errorf(err.Error())
	}
	expectedVector := &Vector{Coordinates: []float64{25, 50, 70}}
	if !IsEqual(expectedVector, vector_sum) {
		t.Errorf("Vectors are different: %v %v", expectedVector.Coordinates, vector_sum.Coordinates)
	}
}

// TestDotProduct tests the dot product of two vectors.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestDotProduct(t *testing.T) {
	vect1 := &Vector{Coordinates: []float64{2, 3, 5}}
	vect2 := &Vector{Coordinates: []float64{1, 2, 2}}
	dotProduct, err := DotProduct(vect1, vect2)
	if err != nil {
		t.Errorf(err.Error())
	}
	expectedDotProduct := float64(18)
	if expectedDotProduct != dotProduct {
		t.Errorf("Dot product is wrong: %v %v", expectedDotProduct, dotProduct)
	}
}
