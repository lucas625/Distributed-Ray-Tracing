package object

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestLightCharacteristics_Init tests the instantiation of a lightCharacteristics.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestLightCharacteristics_Init(t *testing.T) {
	color := []float64{0.1, 0.25, 0.5}
	specularDecay := 5.0
	specularReflection := 0.5
	roughNess := 0.0
	transmissionReflection := 0.25
	diffuseReflection := 0.25

	receivedLightCharacteristics, err := initLightCharacteristics(
		color, specularDecay, specularReflection, roughNess, transmissionReflection, diffuseReflection)
	test_helpers.AssertNilError(t, err)

	expectedLightCharacteristics := &lightCharacteristics{color: color, specularDecay: specularDecay,
		specularReflection: specularReflection, roughNess: roughNess, transmissionReflection: transmissionReflection,
		diffuseReflection: diffuseReflection}
	test_helpers.AssertEqual(t, true, expectedLightCharacteristics.IsEqual(receivedLightCharacteristics))
}

// TestLightCharacteristics_Init_NonRGBColorError tests the instantiation of a lightCharacteristics.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestLightCharacteristics_Init_NonRGBColorError(t *testing.T) {
	color := []float64{0.1, 0.25}
	specularDecay := 5.0
	specularReflection := 0.5
	roughNess := 0.0
	transmissionReflection := 0.25
	diffuseReflection := 0.25

	expectedErrorMessage := fmt.Sprintf("There are not 3 color values: %d.", len(color))

	_, err := initLightCharacteristics(
		color, specularDecay, specularReflection, roughNess, transmissionReflection, diffuseReflection)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestLightCharacteristics_Init_ColorOutOfBoundsError tests the instantiation of a lightCharacteristics.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestLightCharacteristics_Init_ColorOutOfBoundsError(t *testing.T) {
	color := []float64{1.5, 0.25, 0.5}
	specularDecay := 5.0
	specularReflection := 0.5
	roughNess := 0.0
	transmissionReflection := 0.25
	diffuseReflection := 0.25

	expectedErrorMessage := fmt.Sprintf("Color values out of interval [0,1]: %v.", color)

	_, err := initLightCharacteristics(
		color, specularDecay, specularReflection, roughNess, transmissionReflection, diffuseReflection)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestLightCharacteristics_Init_InvalidReflectionCoefficientsError tests the instantiation of a lightCharacteristics.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestLightCharacteristics_Init_InvalidReflectionCoefficientsError(t *testing.T) {
	color := []float64{0.5, 0.25, 0.5}
	specularDecay := 5.0
	specularReflection := -1.0
	roughNess := 0.0
	transmissionReflection := 0.25
	diffuseReflection := 0.25

	expectedErrorMessage := fmt.Sprintf("At least one of the reflections is smaller than 0 or they do not " +
		"sum 1: diffuse %v, specular %v, transmission %v.",
		diffuseReflection, specularReflection, transmissionReflection)

	_, err := initLightCharacteristics(
		color, specularDecay, specularReflection, roughNess, transmissionReflection, diffuseReflection)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
