package object

import (
	"reflect"
)

// lightCharacteristics is a class for all light characteristics data of an Object.
//
// Members:
//  color                  - RGB for the color of the object.
//  specularDecay          - Constant for how fast the specular component decays.
//  specularReflection     - The coefficient of specular reflection.
//  roughNess              - How much reflections rays get distorted.
//  transmissionReflection - The coefficient for transmission.
//  ambientReflection      - The ambient reflection.
//  diffuseReflection      - Diffuse reflection coefficient.
//
type lightCharacteristics struct {
	color              []float64
	specularDecay      float64
	specularReflection float64
	roughNess          float64
	transmissionReflection float64
	ambientReflection  float64
	diffuseReflection  float64
}

// GetColor gets the RGB color.
//
// Parameters:
// 	none
//
// Returns:
// 	The RGB color.
//
func (characteristics *lightCharacteristics) GetColor() []float64 {
	return characteristics.color
}

// GetSpecularDecay gets the specular decay.
//
// Parameters:
// 	none
//
// Returns:
// 	The specular decay.
//
func (characteristics *lightCharacteristics) GetSpecularDecay() float64 {
	return characteristics.specularDecay
}

// GetSpecularReflection gets the specular reflection.
//
// Parameters:
// 	none
//
// Returns:
// 	The specular reflection.
//
func (characteristics *lightCharacteristics) GetSpecularReflection() float64 {
	return characteristics.specularReflection
}

// GetRoughNess gets the roughness.
//
// Parameters:
// 	none
//
// Returns:
// 	The roughness.
//
func (characteristics *lightCharacteristics) GetRoughNess() float64 {
	return characteristics.roughNess
}

// GetTransmissionReflection gets the transmission reflection.
//
// Parameters:
// 	none
//
// Returns:
// 	The transmission reflection.
//
func (characteristics *lightCharacteristics) GetTransmissionReflection() float64 {
	return characteristics.transmissionReflection
}

// GetAmbientReflection gets the ambient reflection.
//
// Parameters:
// 	none
//
// Returns:
// 	The ambient reflection.
//
func (characteristics *lightCharacteristics) GetAmbientReflection() float64 {
	return characteristics.ambientReflection
}

// GetDiffuseReflection gets the diffuse reflection.
//
// Parameters:
// 	none
//
// Returns:
// 	The diffuse reflection.
//
func (characteristics *lightCharacteristics) GetDiffuseReflection() float64 {
	return characteristics.diffuseReflection
}

// IsEqual checks if a lightCharacteristics object is equal to another.
//
// Parameters:
// 	other - The other lightCharacteristics.
//
// Returns:
// 	If the lightCharacteristics are equal.
//
func (characteristics *lightCharacteristics) IsEqual(other *lightCharacteristics) bool {
	return reflect.DeepEqual(characteristics.GetColor(), other.GetColor()) &&
		characteristics.GetSpecularDecay() == other.GetSpecularDecay() &&
		characteristics.GetSpecularReflection() == other.GetSpecularReflection() &&
		characteristics.GetRoughNess() == other.GetRoughNess() &&
		characteristics.GetTransmissionReflection() == other.GetTransmissionReflection() &&
		characteristics.GetAmbientReflection() == other.GetAmbientReflection() &&
		characteristics.GetDiffuseReflection() == other.GetDiffuseReflection()
}

// initLightCharacteristics initializes the light characteristics.
//
// Parameters:
//  color                  - RGB for the color of the object.
//  specularDecay          - Constant for how fast the specular component decays.
//  specularReflection     - The coefficient of specular reflection.
//  roughNess              - How much reflections rays get distorted.
//  transmissionReflection - The coefficient for transmission.
//  ambientReflection      - The ambient reflection.
//  diffuseReflection      - Diffuse reflection coefficient.
//
// Returns:
// 	A lightCharacteristics.
// 	An error.
//
func initLightCharacteristics(color []float64, specularDecay, specularReflection, roughNess, transmissionReflection,
	ambientReflection, diffuseReflection float64) (*lightCharacteristics, error) {
	if len(color) != 3 {
		return nil, nonRGBColorError(color)
	}
	for colorIndex := 0; colorIndex < 3; colorIndex++ {
		if color[colorIndex] < 0 || color[colorIndex] > 1 {
			return nil, colorOutOfBoundsError(color)
		}
	}
	return &lightCharacteristics{color: color, specularDecay: specularDecay, specularReflection: specularReflection,
		roughNess: roughNess, transmissionReflection: transmissionReflection, ambientReflection: ambientReflection,
		diffuseReflection: diffuseReflection}, nil
}
