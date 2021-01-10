package vector

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
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
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, dimension, vector.Dimension())
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
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, dimension, vector.Dimension())
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
	expectedErrorMessage := fmt.Sprintf("Invalid vector size %d.", dimension)

	_, err := Init(dimension)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
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

	firstVector, err := Init(dimension)
	test_helpers.AssertNilError(t, err)
	secondVector, err := Init(dimension)
	test_helpers.AssertNilError(t, err)

	isEqualDimension := firstVector.IsEqualDimension(secondVector)
	test_helpers.AssertEqual(t, true, isEqualDimension)
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
	firstVector, err := Init(1)
	test_helpers.AssertNilError(t, err)
	secondVector, err := Init(2)
	test_helpers.AssertNilError(t, err)

	isEqualDimension := firstVector.IsEqualDimension(secondVector)
	test_helpers.AssertEqual(t, false, isEqualDimension)
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

	areVectorsEqual := firstVector.IsEqual(secondVector)
	test_helpers.AssertEqual(t, true, areVectorsEqual)
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

	areVectorsEqual := firstVector.IsEqual(secondVector)
	test_helpers.AssertEqual(t, false, areVectorsEqual)
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

	areVectorsEqual := firstVector.IsEqual(secondVector)
	test_helpers.AssertEqual(t, false, areVectorsEqual)
}
