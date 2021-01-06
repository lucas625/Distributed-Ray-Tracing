package vector

import (
	"testing"
	"fmt"
)


// TestVector_ScalarMultiplication tests the scalar multiplication of a Vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_ScalarMultiplication(t *testing.T) {
	vector := &Vector{coordinates: []float64{10, 20, 30}}
	multipliedVector := ScalarMultiplication(vector, 2)
	expectedVector := &Vector{coordinates: []float64{20, 40, 60}}
	if !expectedVector.IsEqual(multipliedVector) {
		t.Errorf("Vectors are different: %v %v", expectedVector.coordinates, multipliedVector.coordinates)
	}
}

// TestVector_Sum tests the sum of two vectors.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_Sum(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{10, 20, 30}}
	secondVector := &Vector{coordinates: []float64{5, 10, 20}}
	vectorSum, _ := Sum(firstVector, secondVector, 1, 1)
	expectedVector := &Vector{coordinates: []float64{15, 30, 50}}
	if !expectedVector.IsEqual(vectorSum) {
		t.Errorf("Vectors are different: %v %v", expectedVector.coordinates, vectorSum.coordinates)
	}
}

// TestVector_Sum_DifferentDimensions tests the sum of two vectors with different dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_Sum_DifferentDimensions(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{10, 20, 30}}
	secondVector := &Vector{coordinates: []float64{5, 10}}
	_, err := Sum(firstVector, secondVector, 1, 1)
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", firstVector.Dimension(), secondVector.Dimension())
	if err == nil {
		t.Errorf("Applied sum on vectors with different dimensions: %v %v", firstVector, secondVector)
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Wrong error on vector sum: %v", err.Error())
	}
}

// TestVector_Sum_Subtraction tests the subtraction of two vectors.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_Sum_Subtraction(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{10, 20, 30}}
	secondVector := &Vector{coordinates: []float64{5, 10, 20}}
	vectorSum, _ := Sum(firstVector, secondVector, 3, -1)
	expectedVector := &Vector{coordinates: []float64{25, 50, 70}}
	if !expectedVector.IsEqual(vectorSum) {
		t.Errorf("Vectors are different: %v %v", expectedVector.coordinates, vectorSum.coordinates)
	}
}

// TestVector_DotProduct tests the dot product of two vectors.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_DotProduct(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 3, 5}}
	secondVector := &Vector{coordinates: []float64{1, 2, 2}}
	dotProduct, _ := DotProduct(firstVector, secondVector)
	expectedDotProduct := float64(18)
	if expectedDotProduct != dotProduct {
		t.Errorf("Dot product is wrong: %v %v", expectedDotProduct, dotProduct)
	}
}

// TestVector_DotProduct_DifferentDimensions tests the dot product of two vectors with different dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_DotProduct_DifferentDimensions(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 3, 5}}
	secondVector := &Vector{coordinates: []float64{1, 2}}
	_, err := DotProduct(firstVector, secondVector)
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", firstVector.Dimension(),
		secondVector.Dimension())
	if err == nil {
		t.Errorf("Applied sum on vectors with different dimensions: %v %v", firstVector, secondVector)
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Wrong error on vector dot product: %v", err.Error())
	}
}

// TestVector_CrossProduct tests the cross product of two vectors.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_CrossProduct(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 3, 5}}
	secondVector := &Vector{coordinates: []float64{1, 2, 2}}
	crossProduct, _ := CrossProduct(firstVector, secondVector)
	expectedCrossProduct := &Vector{coordinates: []float64{-4, 1, 1}}
	if !expectedCrossProduct.IsEqual(crossProduct) {
		t.Errorf("Cross product is wrong: %v %v", expectedCrossProduct, crossProduct)
	}
}

// TestVector_CrossProduct_DifferentDimensions tests the cross product of two vectors, but with different dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_CrossProduct_DifferentDimensions(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 3, 5}}
	secondVector := &Vector{coordinates: []float64{1, 2}}
	_, err := CrossProduct(firstVector, secondVector)
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", firstVector.Dimension(),
		secondVector.Dimension())
	if err == nil {
		t.Errorf("Applied cross product on vectors with different dimensions: %v %v", firstVector, secondVector)
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Wrong error on vector cross product: %v", err.Error())
	}
}

// TestVector_CrossProduct_Non3D tests the cross product of two vectors with non 3D dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_CrossProduct_Non3D(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 3}}
	secondVector := &Vector{coordinates: []float64{1, 2}}
	_, err := CrossProduct(firstVector, secondVector)
	expectedErrorMessage := fmt.Sprintf("Invalid dimension of vector. Expected 3D and got %d.", firstVector.Dimension())
	if err == nil {
		t.Errorf("Applied cross product on vectors with dimensions not equal to 3: %v %v", firstVector,
			secondVector)
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Wrong error on vector cross product: %v", err.Error())
	}
}

// TestVector_Norm tests the norm of a Vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_Norm(t *testing.T) {
	vector := &Vector{coordinates: []float64{3, 4}}
	vectorNorm := Norm(vector)
	expectedVectorNorm := float64(5)
	if !(expectedVectorNorm == vectorNorm) {
		t.Errorf("Norm is wrong: %v %v", expectedVectorNorm, vectorNorm)
	}
}

// TestVector_Normalize tests the normalization of a Vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_Normalize(t *testing.T) {
	vector := &Vector{coordinates: []float64{2, 1, 2}}
	normalizedVector := Normalize(vector)
	expectedNormalizedVector := &Vector{coordinates: []float64{float64(2)/3, float64(1)/3, float64(2)/3}}
	if !expectedNormalizedVector.IsEqual(normalizedVector) {
		t.Errorf("Normalized vector is wrong: %v %v", expectedNormalizedVector, normalizedVector)
	}
}

// TestVector_Normalize_NullVector tests the normalization of a Vector with a null vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_Normalize_NullVector(t *testing.T) {
	vector := &Vector{coordinates: []float64{0, 0, 0}}
	normalizedVector := Normalize(vector)
	expectedNormalizedVector := &Vector{coordinates: []float64{0, 0, 0}}
	if !expectedNormalizedVector.IsEqual(normalizedVector) {
		t.Errorf("Normalized vector is wrong: %v %v", expectedNormalizedVector, normalizedVector)
	}
}

// TestVector_ProjectOnVector tests the projection of a Vector on another.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_ProjectOnVector(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 1, 1}}
	secondVector := &Vector{coordinates: []float64{1, -1, 0}}
	projection, _ := ProjectOnVector(firstVector, secondVector)
	expectedProjection := &Vector{coordinates: []float64{float64(1)/2, float64(-1)/2, 0}}
	if !expectedProjection.IsEqual(projection) {
		t.Errorf("Vector projection is wrong: %v %v.", expectedProjection, projection)
	}
}

// TestVector_ProjectOnVector_DifferentDimensions tests the projection of a Vector on another with different dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_ProjectOnVector_DifferentDimensions(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 1, 1}}
	secondVector := &Vector{coordinates: []float64{1, -1}}
	_, err := ProjectOnVector(firstVector, secondVector)
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", firstVector.Dimension(), secondVector.Dimension())
	if err == nil {
		t.Errorf("Applied projection on vectors with different dimensions: %v %v", firstVector, secondVector)
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Wrong error on vector projection: %v", err.Error())
	}
}

// TestVector_Orthogonalize tests the orthogonalization of a vector based on another.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_Orthogonalize(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 1, 1}}
	secondVector := &Vector{coordinates: []float64{1, -1, 0}}
	orthogonalizedVector, err := Orthogonalize(firstVector, secondVector)
	if err != nil {
		t.Errorf(err.Error())
	}
	expectedOrthogonalizedVector := &Vector{coordinates: []float64{float64(3)/2, float64(3)/2, 1}}
	if !expectedOrthogonalizedVector.IsEqual(orthogonalizedVector) {
		t.Errorf("Orthogonalized vector is wrong: %v %v.", expectedOrthogonalizedVector, orthogonalizedVector)
	}
}

// TestVector_Orthogonalize_DifferentDimensions tests the orthogonalization of a vector based on another with different
// dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_Orthogonalize_DifferentDimensions(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 1, 1}}
	secondVector := &Vector{coordinates: []float64{1, -1}}
	_, err := Orthogonalize(firstVector, secondVector)
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", firstVector.Dimension(), secondVector.Dimension())
	if err == nil {
		t.Errorf("Applied orthogonalization on vectors with different dimensions: %v %v", firstVector, secondVector)
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Wrong error on vector orthogonalization: %v", err.Error())
	}
}

// TestVector_IsOrthogonalVector tests if two vectors are orthogonal to each other.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_IsOrthogonalVector(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{1, 1, 1}}
	secondVector := &Vector{coordinates: []float64{1, -1, 0}}
	isOrthogonal, _ := IsOrthogonalVector(firstVector, secondVector)
	if !isOrthogonal {
		t.Errorf("Vectors are orthogonal to each other: %v %v", firstVector, secondVector)
	}
}

// TestVector_IsOrthogonalVector_NotOrthogonal tests if two vectors are not orthogonal to each other.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_IsOrthogonalVector_NotOrthogonal(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 1, 1}}
	secondVector := &Vector{coordinates: []float64{1, -1, 0}}
	isOrthogonal, _ := IsOrthogonalVector(firstVector, secondVector)
	if isOrthogonal {
		t.Errorf("Vectors are not orthogonal to each other: %v %v", firstVector, secondVector)
	}
}

// TestVector_IsOrthogonal_DifferentDimensions tests if two vectors are orthogonal to each other with different
// dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_IsOrthogonal_DifferentDimensions(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{1, 1, 1}}
	secondVector := &Vector{coordinates: []float64{1, -1}}
	_, err := IsOrthogonalVector(firstVector, secondVector)
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", firstVector.Dimension(), secondVector.Dimension())
	if err == nil {
		t.Errorf("Applied is orthogonal on vectors with different dimensions: %v %v", firstVector, secondVector)
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Wrong error on vector orthogonalization: %v", err.Error())
	}
}
