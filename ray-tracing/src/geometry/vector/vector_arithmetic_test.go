package vector

import (
	"testing"
	"fmt"
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
	vect := &Vector{coordinates: []float64{10, 20, 30}}
	multipliedVector := ScalarMultiplication(vect, 2)
	expectedVector := &Vector{coordinates: []float64{20, 40, 60}}
	if !expectedVector.IsEqual(multipliedVector) {
		t.Errorf("Vectors are different: %v %v", expectedVector.coordinates, multipliedVector.coordinates)
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
	vect1 := &Vector{coordinates: []float64{10, 20, 30}}
	vect2 := &Vector{coordinates: []float64{5, 10, 20}}
	vectorSum, err := Sum(vect1, vect2, 1, 1)
	if err != nil {
		t.Errorf(err.Error())
	}
	expectedVector := &Vector{coordinates: []float64{15, 30, 50}}
	if !expectedVector.IsEqual(vectorSum) {
		t.Errorf("Vectors are different: %v %v", expectedVector.coordinates, vectorSum.coordinates)
	}
}

// TestSumDifferentDimensions tests the sum of two vectors with different dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestSumDifferentDimensions(t *testing.T) {
	vect1 := &Vector{coordinates: []float64{10, 20, 30}}
	vect2 := &Vector{coordinates: []float64{5, 10}}
	_, err := Sum(vect1, vect2, 1, 1)
	if err == nil {
		t.Errorf("Applied sum on vectors with different dimensions: %v %v", vect1, vect2)
	} else if err.Error() != fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n",
		len(vect1.coordinates),
		len(vect2.coordinates)) {
		t.Errorf("Wrong error on vector sum: %v", err.Error())
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
	vect1 := &Vector{coordinates: []float64{10, 20, 30}}
	vect2 := &Vector{coordinates: []float64{5, 10, 20}}
	vectorSum, err := Sum(vect1, vect2, 3, -1)
	if err != nil {
		t.Errorf(err.Error())
	}
	expectedVector := &Vector{coordinates: []float64{25, 50, 70}}
	if !expectedVector.IsEqual(vectorSum) {
		t.Errorf("Vectors are different: %v %v", expectedVector.coordinates, vectorSum.coordinates)
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
	vect1 := &Vector{coordinates: []float64{2, 3, 5}}
	vect2 := &Vector{coordinates: []float64{1, 2, 2}}
	dotProduct, err := DotProduct(vect1, vect2)
	if err != nil {
		t.Errorf(err.Error())
	}
	expectedDotProduct := float64(18)
	if expectedDotProduct != dotProduct {
		t.Errorf("Dot product is wrong: %v %v", expectedDotProduct, dotProduct)
	}
}

// TestDotProductDifferentDimensions tests the dot product of two vectors with different dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestDotProductDifferentDimensions(t *testing.T) {
	vect1 := &Vector{coordinates: []float64{2, 3, 5}}
	vect2 := &Vector{coordinates: []float64{1, 2}}
	_, err := DotProduct(vect1, vect2)
	if err == nil {
		t.Errorf("Applied sum on vectors with different dimensions: %v %v", vect1, vect2)
	} else if err.Error() != fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n",
		vect1.Dimension(),
		vect2.Dimension()) {
		t.Errorf("Wrong error on vector dot product: %v", err.Error())
	}
}

// TestCrossProduct tests the cross product of two vectors.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestCrossProduct(t *testing.T) {
	vect1 := &Vector{coordinates: []float64{2, 3, 5}}
	vect2 := &Vector{coordinates: []float64{1, 2, 2}}
	crossProduct, err := CrossProduct(vect1, vect2)
	if err != nil {
		t.Errorf(err.Error())
	}
	expectedCrossProduct := &Vector{coordinates: []float64{-4, 1, 1}}
	if !expectedCrossProduct.IsEqual(crossProduct) {
		t.Errorf("Cross product is wrong: %v %v", expectedCrossProduct, crossProduct)
	}
}

// TestCrossProductDifferentDimensions tests the cross product of two vectors.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestCrossProductDifferentDimensions(t *testing.T) {
	vect1 := &Vector{coordinates: []float64{2, 3, 5}}
	vect2 := &Vector{coordinates: []float64{1, 2}}
	_, err := CrossProduct(vect1, vect2)
	if err == nil {
		t.Errorf("Applied cross product on vectors with different dimensions: %v %v", vect1, vect2)
	} else if err.Error() != fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n",
		vect1.Dimension(),
		vect2.Dimension()) {
		t.Errorf("Wrong error on vector cross product: %v", err.Error())
	}
}

// TestCrossProductDifferentDimensions tests the cross product of two vectors.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestCrossProductNon3D(t *testing.T) {
	vect1 := &Vector{coordinates: []float64{2, 3}}
	vect2 := &Vector{coordinates: []float64{1, 2}}
	_, err := CrossProduct(vect1, vect2)
	if err == nil {
		t.Errorf("Applied cross product on vectors with dimensions not equal to 3: %v %v", vect1, vect2)
	} else if err.Error() !=  fmt.Sprintf(
		"Invalid dimension of vector. Expected 3D and got %d.", vect1.Dimension()) {
		t.Errorf("Wrong error on vector cross product: %v", err.Error())
	}
}

// TestNorm tests the norm of a vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestNorm(t *testing.T) {
	vect := &Vector{coordinates: []float64{3, 4}}
	norm := Norm(vect)
	expectedNorm := float64(5)
	if !(expectedNorm == norm) {
		t.Errorf("Norm is wrong: %v %v", expectedNorm, norm)
	}
}

// TestNormalize tests the normalization of a vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestNormalize(t *testing.T) {
	vect := &Vector{coordinates: []float64{2, 1, 2}}
	normalizedVector := Normalize(vect)
	expectedNormalizedVector := &Vector{coordinates: []float64{float64(2)/3, float64(1)/3, float64(2)/3}}
	if !expectedNormalizedVector.IsEqual(normalizedVector) {
		t.Errorf("Normalized vector is wrong: %v %v", expectedNormalizedVector, normalizedVector)
	}
}

// TestNormalize tests the normalization of a vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestNormalizeNullVector(t *testing.T) {
	vector := &Vector{coordinates: []float64{0, 0, 0}}
	normalizedVector := Normalize(vector)
	expectedNormalizedVector := &Vector{coordinates: []float64{0, 0, 0}}
	if !expectedNormalizedVector.IsEqual(normalizedVector) {
		t.Errorf("Normalized vector is wrong: %v %v", expectedNormalizedVector, normalizedVector)
	}
}
