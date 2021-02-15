package object

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestObject_NonRGBColorError tests the error where the color of the object does not have 3 values.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestObject_NonRGBColorError(t *testing.T) {
	color := []float64{255, 100}
	expectedErrorMessage := fmt.Sprintf("There are not 3 color values: %d.", len(color))
	err := nonRGBColorError(color)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestObject_ColorOutOfBoundsError tests the error where a color coefficient is out of the bounds.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestObject_ColorOutOfBoundsError(t *testing.T) {
	color := []float64{255, 100, 0}
	expectedErrorMessage := fmt.Sprintf("Color values out of interval [0,1]: %v.", color)
	err := colorOutOfBoundsError(color)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestObject_InvalidReflectionCoefficientsError tests the error where an Object is ambient, diffuse, specular or
// transmission reflections are not valid.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestObject_InvalidReflectionCoefficientsError(t *testing.T) {
	specularReflection := 0.5
	transmissionReflection := -0.25
	diffuseReflection := 0.25
	expectedErrorMessage := fmt.Sprintf("At least one of the reflections is smaller than 0 or they do not " +
		"sum 1: diffuse %v, specular %v, transmission %v.",
		diffuseReflection, specularReflection, transmissionReflection)
	err := invalidReflectionCoefficientsError(specularReflection, transmissionReflection, diffuseReflection)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
