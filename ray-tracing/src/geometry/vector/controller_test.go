package vector

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)


// TestVectorController_ScalarMultiplication tests the scalar multiplication of a Vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_ScalarMultiplication(t *testing.T) {
	vector := &Vector{coordinates: []float64{10, 20, 30}}
	expectedVector := &Vector{coordinates: []float64{20, 40, 60}}
	controller := Controller{}

	multipliedVector := controller.ScalarMultiplication(vector, 2)

	areVectorsEqual := expectedVector.IsEqual(multipliedVector)
	test_helpers.AssertEqual(t, true, areVectorsEqual)
}

// TestVectorController_Sum tests the sum of two vectors.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_Sum(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{10, 20, 30}}
	secondVector := &Vector{coordinates: []float64{5, 10, 20}}
	expectedVector := &Vector{coordinates: []float64{15, 30, 50}}
	controller := Controller{}

	vectorSum, err := controller.Sum(firstVector, secondVector, 1, 1)
	test_helpers.AssertNilError(t, err)

	areVectorsEqual := expectedVector.IsEqual(vectorSum)
	test_helpers.AssertEqual(t, true, areVectorsEqual)
}

// TestVectorController_Sum_DifferentDimensions tests the sum of two vectors with different dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_Sum_DifferentDimensions(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{10, 20, 30}}
	secondVector := &Vector{coordinates: []float64{5, 10}}
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", firstVector.Dimension(), secondVector.Dimension())
	controller := Controller{}

	_, err := controller.Sum(firstVector, secondVector, 1, 1)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestVectorController_Sum_Subtraction tests the subtraction of two vectors.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_Sum_Subtraction(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{10, 20, 30}}
	secondVector := &Vector{coordinates: []float64{5, 10, 20}}
	expectedVector := &Vector{coordinates: []float64{25, 50, 70}}
	controller := Controller{}

	vectorSum, err := controller.Sum(firstVector, secondVector, 3, -1)
	test_helpers.AssertNilError(t, err)

	areVectorsEqual := expectedVector.IsEqual(vectorSum)
	test_helpers.AssertEqual(t, true, areVectorsEqual)
}

// TestVectorController_DotProduct tests the dot product of two vectors.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_DotProduct(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 3, 5}}
	secondVector := &Vector{coordinates: []float64{1, 2, 2}}
	expectedDotProduct := float64(18)
	controller := Controller{}

	dotProduct, err := controller.DotProduct(firstVector, secondVector)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, expectedDotProduct, dotProduct)
}

// TestVectorController_DotProduct_DifferentDimensions tests the dot product of two vectors with different dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_DotProduct_DifferentDimensions(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 3, 5}}
	secondVector := &Vector{coordinates: []float64{1, 2}}
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", firstVector.Dimension(),
		secondVector.Dimension())
	controller := Controller{}

	_, err := controller.DotProduct(firstVector, secondVector)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestVectorController_CrossProduct tests the cross product of two vectors.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_CrossProduct(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 3, 5}}
	secondVector := &Vector{coordinates: []float64{1, 2, 2}}
	expectedCrossProduct := &Vector{coordinates: []float64{-4, 1, 1}}
	controller := Controller{}

	crossProduct, err := controller.CrossProduct(firstVector, secondVector)

	areVectorsEqual := expectedCrossProduct.IsEqual(crossProduct)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, true, areVectorsEqual)
}

// TestVectorController_CrossProduct_DifferentDimensions tests the cross product of two vectors, but with different dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_CrossProduct_DifferentDimensions(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 3, 5}}
	secondVector := &Vector{coordinates: []float64{1, 2}}
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", firstVector.Dimension(),
		secondVector.Dimension())
	controller := Controller{}

	_, err := controller.CrossProduct(firstVector, secondVector)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestVectorController_CrossProduct_Non3D tests the cross product of two vectors with non 3D dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_CrossProduct_Non3D(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 3}}
	secondVector := &Vector{coordinates: []float64{1, 2}}
	expectedErrorMessage := fmt.Sprintf("Invalid dimension of vector. Expected 3D and got %d.", firstVector.Dimension())
	controller := Controller{}

	_, err := controller.CrossProduct(firstVector, secondVector)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestVectorController_Norm tests the norm of a Vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_Norm(t *testing.T) {
	vector := &Vector{coordinates: []float64{3, 4}}
	expectedVectorNorm := float64(5)
	controller := Controller{}

	vectorNorm := controller.Norm(vector)
	test_helpers.AssertEqual(t, expectedVectorNorm, vectorNorm)
}

// TestVectorController_Normalize tests the normalization of a Vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_Normalize(t *testing.T) {
	vector := &Vector{coordinates: []float64{2, 1, 2}}
	expectedNormalizedVector := &Vector{coordinates: []float64{float64(2)/3, float64(1)/3, float64(2)/3}}
	controller := Controller{}

	normalizedVector := controller.Normalize(vector)
	areVectorsEqual := expectedNormalizedVector.IsEqual(normalizedVector)
	test_helpers.AssertEqual(t, true, areVectorsEqual)
}

// TestVectorController_Normalize_NullVector tests the normalization of a Vector with a null vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_Normalize_NullVector(t *testing.T) {
	vector := &Vector{coordinates: []float64{0, 0, 0}}
	expectedNormalizedVector := &Vector{coordinates: []float64{0, 0, 0}}
	controller := Controller{}

	normalizedVector := controller.Normalize(vector)
	areVectorsEqual := expectedNormalizedVector.IsEqual(normalizedVector)
	test_helpers.AssertEqual(t, true, areVectorsEqual)
}

// TestVectorController_ProjectOnVector tests the projection of a Vector on another.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_ProjectOnVector(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 1, 1}}
	secondVector := &Vector{coordinates: []float64{1, -1, 0}}
	expectedProjection := &Vector{coordinates: []float64{float64(1)/2, float64(-1)/2, 0}}
	controller := Controller{}

	projection, err := controller.ProjectOnVector(firstVector, secondVector)
	test_helpers.AssertNilError(t, err)

	areVectorsEqual := expectedProjection.IsEqual(projection)
	test_helpers.AssertEqual(t, true, areVectorsEqual)
}

// TestVectorController_ProjectOnVector_DifferentDimensions tests the projection of a Vector on another with different dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_ProjectOnVector_DifferentDimensions(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 1, 1}}
	secondVector := &Vector{coordinates: []float64{1, -1}}
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", firstVector.Dimension(), secondVector.Dimension())
	controller := Controller{}

	_, err := controller.ProjectOnVector(firstVector, secondVector)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestVectorController_Orthogonalize tests the orthogonalization of a vector based on another.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_Orthogonalize(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 1, 1}}
	secondVector := &Vector{coordinates: []float64{1, -1, 0}}
	expectedOrthogonalizedVector := &Vector{coordinates: []float64{float64(3)/2, float64(3)/2, 1}}
	controller := Controller{}

	orthogonalizedVector, err := controller.Orthogonalize(firstVector, secondVector)
	test_helpers.AssertNilError(t, err)

	areVectorsEqual := expectedOrthogonalizedVector.IsEqual(orthogonalizedVector)
	test_helpers.AssertEqual(t, true, areVectorsEqual)
}

// TestVectorController_Orthogonalize_DifferentDimensions tests the orthogonalization of a vector based on another with different
// dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_Orthogonalize_DifferentDimensions(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 1, 1}}
	secondVector := &Vector{coordinates: []float64{1, -1}}
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", firstVector.Dimension(), secondVector.Dimension())
	controller := Controller{}

	_, err := controller.Orthogonalize(firstVector, secondVector)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestVectorController_IsOrthogonalVector tests if two vectors are orthogonal to each other.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_IsOrthogonalVector(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{1, 1, 1}}
	secondVector := &Vector{coordinates: []float64{1, -1, 0}}
	controller := Controller{}

	isOrthogonal, err := controller.IsOrthogonalVector(firstVector, secondVector)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, true, isOrthogonal)
}

// TestVectorController_IsOrthogonalVector_NotOrthogonal tests if two vectors are not orthogonal to each other.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_IsOrthogonalVector_NotOrthogonal(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{2, 1, 1}}
	secondVector := &Vector{coordinates: []float64{1, -1, 0}}
	controller := Controller{}

	isOrthogonal, err := controller.IsOrthogonalVector(firstVector, secondVector)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, false, isOrthogonal)
}

// TestVectorController_IsOrthogonal_DifferentDimensions tests if two vectors are orthogonal to each other with different
// dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVectorController_IsOrthogonal_DifferentDimensions(t *testing.T) {
	firstVector := &Vector{coordinates: []float64{1, 1, 1}}
	secondVector := &Vector{coordinates: []float64{1, -1}}
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", firstVector.Dimension(), secondVector.Dimension())
	controller := Controller{}

	_, err := controller.IsOrthogonalVector(firstVector, secondVector)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
