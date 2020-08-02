package vector

import (
	"fmt"
	"testing"
)


// TestDifferentDimensionError tests different dimension error.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestDifferentDimensionError(t *testing.T) {
	vect1, _ := Init(1)
	vect2, _ := Init(2)
	err := differentDimensionError(vect1, vect2)
	if err == nil {
		t.Errorf(
			"No different dimension error return for vectors dimension: %d %d.",
			len(vect1.Coordinates),
			len(vect2.Coordinates))
	} else if err.Error() != fmt.Sprintf(
		"Invalid dimension of vector. Expected: %d and got: %d.\n",
		len(vect1.Coordinates),
		len(vect2.Coordinates)) {
		t.Errorf("Wrong error message for vectors with different dimension: \"%s\".", err.Error())
	}
}

// TestNon3DError tests non 3D vector error.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestNon3DError(t *testing.T) {
	vect, _ := Init(2)
	err := non3DError(vect)
	if err == nil {
		t.Errorf(
			"No non 3D error return for vector with dimension: %d.",
			len(vect.Coordinates))
	} else if err.Error() != fmt.Sprintf(
		"Invalid dimension of vector. Expected 3D and got %d.",
		len(vect.Coordinates)) {
		t.Errorf("Wrong error message for non 3D vector: \"%s\".", err.Error())
	}
}

// TestNegativeDimensionError tests vector with negative dimension error.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestNegativeDimensionError(t *testing.T) {
	dimension := -1
	err := negativeDimensionError(dimension)
	if err == nil {
		t.Errorf("No negative dimension error return for dimension: %d.", dimension)
	} else if err.Error() != fmt.Sprintf("Invalid vector size %d.", dimension) {
		t.Errorf("Wrong error message for negative dimension: \"%s\".", err.Error())
	}
}
