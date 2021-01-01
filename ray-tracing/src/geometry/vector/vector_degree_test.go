package vector

import (
	"fmt"
	"testing"
)


// TestProjectOnVector tests the projection of a vector on another.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestProjectOnVector(t *testing.T) {
	vect1 := &Vector{coordinates: []float64{2, 1, 1}}
	vect2 := &Vector{coordinates: []float64{1, -1, 0}}
	projection, err := ProjectOnVector(vect1, vect2)
	if err != nil {
		t.Errorf(err.Error())
	}
	expectedProjection := &Vector{coordinates: []float64{float64(1)/2, float64(-1)/2, 0}}
	if !expectedProjection.IsEqual(projection) {
		t.Errorf("Vector projection is wrong: %v %v.", expectedProjection, projection)
	}
}

// TestProjectOnVectorDifferentDimensions tests the projection of a vector on another with different dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestProjectOnVectorDifferentDimensions(t *testing.T) {
	vect1 := &Vector{coordinates: []float64{2, 1, 1}}
	vect2 := &Vector{coordinates: []float64{1, -1}}
	_, err := ProjectOnVector(vect1, vect2)
	if err == nil {
		t.Errorf("Applied projection on vectors with different dimensions: %v %v", vect1, vect2)
	} else if err.Error() != fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", vect1.Dimension(), vect2.Dimension()) {
		t.Errorf("Wrong error on vector projection: %v", err.Error())
	}
}

// TestOrthogonalize tests the orthogonalization of a vector based on another.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestOrthogonalize(t *testing.T) {
	vect1 := &Vector{coordinates: []float64{2, 1, 1}}
	vect2 := &Vector{coordinates: []float64{1, -1, 0}}
	orthogonalizedVector, err := Orthogonalize(vect1, vect2)
	if err != nil {
		t.Errorf(err.Error())
	}
	expectedOrthogonalizedVector := &Vector{coordinates: []float64{float64(3)/2, float64(3)/2, 1}}
	if !expectedOrthogonalizedVector.IsEqual(orthogonalizedVector) {
		t.Errorf("Orthogonalized vector is wrong: %v %v.", expectedOrthogonalizedVector, orthogonalizedVector)
	}
}

// TestOrthogonalizeDifferentDimensions tests the orthogonalization of a vector based on another with different
// dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestOrthogonalizeDifferentDimensions(t *testing.T) {
	vect1 := &Vector{coordinates: []float64{2, 1, 1}}
	vect2 := &Vector{coordinates: []float64{1, -1}}
	_, err := Orthogonalize(vect1, vect2)
	if err == nil {
		t.Errorf("Applied orthogonalization on vectors with different dimensions: %v %v", vect1, vect2)
	} else if err.Error() != fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", vect1.Dimension(), vect2.Dimension()) {
		t.Errorf("Wrong error on vector orthogonalization: %v", err.Error())
	}
}

// TestIsOrthogonalVector tests if two vectors are orthogonal to each other.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIsOrthogonalVector(t *testing.T) {
	vect1 := &Vector{coordinates: []float64{1, 1, 1}}
	vect2 := &Vector{coordinates: []float64{1, -1, 0}}
	isOrthogonal, err := IsOrthogonalVector(vect1, vect2)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !isOrthogonal {
		t.Errorf("Vectors are orthogonal to each other: %v %v", vect1, vect2)
	}
}

// TestIsNotOrthogonalVector tests if two vectors are not orthogonal to each other.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIsNotOrthogonalVector(t *testing.T) {
	vect1 := &Vector{coordinates: []float64{2, 1, 1}}
	vect2 := &Vector{coordinates: []float64{1, -1, 0}}
	isOrthogonal, err := IsOrthogonalVector(vect1, vect2)
	if err != nil {
		t.Errorf(err.Error())
	}
	if isOrthogonal {
		t.Errorf("Vectors are not orthogonal to each other: %v %v", vect1, vect2)
	}
}

// TestIsOrthogonalDifferentDimensions tests the projection of a vector on another with different dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestIsOrthogonalDifferentDimensions(t *testing.T) {
	vect1 := &Vector{coordinates: []float64{1, 1, 1}}
	vect2 := &Vector{coordinates: []float64{1, -1}}
	_, err := IsOrthogonalVector(vect1, vect2)
	if err == nil {
		t.Errorf("Applied is orthogonal on vectors with different dimensions: %v %v", vect1, vect2)
	} else if err.Error() != fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n",
		len(vect1.coordinates),
		len(vect2.coordinates)) {
		t.Errorf("Wrong error on vector orthogonalization: %v", err.Error())
	}
}
