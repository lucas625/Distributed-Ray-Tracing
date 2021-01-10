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

// TestVector_CopyAllCoordinates tests if the copy of all coordinates of a Vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_CopyAllCoordinates(t *testing.T) {
	vectorValues := []float64{10, 20, 30}
	vector := &Vector{coordinates: vectorValues}

	copiedValues := vector.CopyAllCoordinates()

	areValuesEqual := true
	for index := 0; index < len(vectorValues); index++ {
		if vectorValues[index] != copiedValues[index] {
			areValuesEqual = false
		}
	}

	test_helpers.AssertEqual(t, true, areValuesEqual)
}

// TestVector_GetCoordinate tests the get coordinate of a Vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_GetCoordinate(t *testing.T) {
	vector := &Vector{coordinates: []float64{10, 20, 30}}
	expectedValue := 30.0

	vectorValue, err := vector.GetCoordinate(2)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, expectedValue, vectorValue)
}

// TestVector_GetCoordinate_NegativeIndex tests the get coordinate of a Vector when the index is negative.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_GetCoordinate_NegativeIndex(t *testing.T) {
	vector := &Vector{coordinates: []float64{10, 20, 30}}
	index := -1
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the vector. Expected from 0 to: %v and got %v.", vector.Dimension(), index)

	_, err := vector.GetCoordinate(index)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestVector_GetCoordinate_BiggerIndex tests the get coordinate of a Vector when the index is bigger than the limit
// of the vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_GetCoordinate_BiggerIndex(t *testing.T) {
	vector := &Vector{coordinates: []float64{10, 20, 30}}
	index := 3
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the vector. Expected from 0 to: %v and got %v.", vector.Dimension(), index)

	_, err := vector.GetCoordinate(index)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestVector_SetCoordinate tests the set coordinate of a Vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_SetCoordinate(t *testing.T) {
	vector := &Vector{coordinates: []float64{10, 20, 30}}
	newValue := 50.0

	err := vector.SetCoordinate(2, newValue)
	test_helpers.AssertNilError(t, err)
	vectorValue, err := vector.GetCoordinate(2)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, newValue, vectorValue)
}

// TestVector_SetCoordinate_NegativeIndex tests the set coordinate of a Vector with a negative index.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_SetCoordinate_NegativeIndex(t *testing.T) {
	vector := &Vector{coordinates: []float64{10, 20, 30}}
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the vector. Expected from 0 to: %v and got %v.", vector.Dimension(), -1)

	err := vector.SetCoordinate(-1, 50)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestVector_SetCoordinate_BiggerIndex tests the set coordinate of a Vector when the index is bigger than the limit
// of the vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_SetCoordinate_BiggerIndex(t *testing.T) {
	vector := &Vector{coordinates: []float64{10, 20, 30}}
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the vector. Expected from 0 to: %v and got %v.", vector.Dimension(), 3)

	err := vector.SetCoordinate(3, 50)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
