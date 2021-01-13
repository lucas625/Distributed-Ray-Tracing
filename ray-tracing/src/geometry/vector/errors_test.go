package vector

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)


// TestVector_DifferentDimensionError tests different dimension error.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_DifferentDimensionError(t *testing.T) {
	firstVector, err := Init(1)
	test_helpers.AssertNilError(t, err)
	secondVector, err := Init(2)
	test_helpers.AssertNilError(t, err)
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", firstVector.Dimension(), secondVector.Dimension())

	err = differentDimensionError(firstVector, secondVector)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestVector_Non3DError tests non 3D Vector error.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_Non3DError(t *testing.T) {
	vector, err := Init(2)
	test_helpers.AssertNilError(t, err)
	expectedErrorMessage := fmt.Sprintf("Invalid dimension of vector. Expected 3D and got %d.", vector.Dimension())

	err = non3DError(vector)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestVector_NegativeDimensionError tests vector with negative dimension error.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_NegativeDimensionError(t *testing.T) {
	dimension := -1
	expectedErrorMessage := fmt.Sprintf("Invalid vector size %d.", dimension)

	err := negativeDimensionError(dimension)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestVector_IndexError tests Vector index error.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_IndexError(t *testing.T) {
	vector, err := Init(2)
	test_helpers.AssertNilError(t, err)
	index := 3
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the vector. Expected from 0 to %v and got %v.", vector.Dimension(), index)

	err = indexError(vector, index)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
