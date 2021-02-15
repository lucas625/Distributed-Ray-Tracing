package object

import (
	"errors"
	"fmt"
)

// nonRGBColorError is the error where the color of the object does not have 3 values.
//
// Parameters:
//	color - The color values.
//
// Returns:
//  An Error.
//
func nonRGBColorError(color []float64) error {
	errorMessage := fmt.Sprintf("There are not 3 color values: %d.", len(color))
	return errors.New(errorMessage)
}

// colorOutOfBoundsError is the error where a color coefficient is out of the bounds.
//
// Parameters:
//	color - The RGB color values.
//
// Returns:
//  An Error.
//
func colorOutOfBoundsError(color []float64) error {
	errorMessage := fmt.Sprintf("Color values out of interval [0,1]: %v.", color)
	return errors.New(errorMessage)
}

// invalidReflectionCoefficientsError is the error where an Object is ambient, diffuse, specular or transmission
//reflections are not valid.
//
// Parameters:
//  specularReflection     - Percentage of specular rays. .
//  transmissionReflection - Percentage of transmission rays.
//  ambientReflection      - Percentage of ambient rays.
//  diffuseReflection      - Percentage of diffuse rays.
//
// Returns:
//  An Error.
//
func invalidReflectionCoefficientsError(specularReflection, transmissionReflection, ambientReflection,
	diffuseReflection float64) error {
	errorMessage := fmt.Sprintf("At least one of the reflections is smaller than 0 or they do not sum 1: " +
		"ambient %v, diffuse %v, specular %v, transmission %v.",
		ambientReflection, diffuseReflection, specularReflection, transmissionReflection)
	return errors.New(errorMessage)
}
