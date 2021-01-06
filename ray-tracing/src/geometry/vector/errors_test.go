package vector

import (
	"fmt"
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
	firstVector, _ := Init(1)
	secondVector, _ := Init(2)
	err := differentDimensionError(firstVector, secondVector)
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n", firstVector.Dimension(), secondVector.Dimension())
	if err == nil {
		t.Errorf(
			"No different dimension error return for vectors dimension: %d %d.", firstVector.Dimension(),
			secondVector.Dimension())
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Wrong error message for vectors with different dimension: \"%s\".", err.Error())
	}
}

// TestVector_Non3DError tests non 3D vector error.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestVector_Non3DError(t *testing.T) {
	vector, _ := Init(2)
	err := non3DError(vector)
	expectedErrorMessage := fmt.Sprintf("Invalid dimension of vector. Expected 3D and got %d.", vector.Dimension())
	if err == nil {
		t.Errorf("No non 3D error return for vector with dimension: %d.", vector.Dimension())
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Wrong error message for non 3D vector: \"%s\".", err.Error())
	}
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
	err := negativeDimensionError(dimension)
	expectedErrorMessage := fmt.Sprintf("Invalid vector size %d.", dimension)
	if err == nil {
		t.Errorf("No negative dimension error return for dimension: %d.", dimension)
	} else if err.Error() != expectedErrorMessage {
		t.Errorf("Wrong error message for negative dimension: \"%s\".", err.Error())
	}
}
