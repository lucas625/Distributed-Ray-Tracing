package light

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/object"
	"reflect"
)

// Light is a class for holding a Light is data.
//
// Members:
//  LightIntensity - The intensity of the Light.
//  LightObject    - Object for the Light.
//  Color          - RGB of the Light.
//
type Light struct {
	lightIntensity   float64
	lightObject      *object.Object
	color            []float64
}

// GetLightIntensity gets the Light is intensity.
//
// Parameters:
// 	none
//
// Returns:
// 	The intensity of the Light.
//
func (light *Light) GetLightIntensity() float64 {
	return light.lightIntensity
}

// GetLightObject gets the Light is object.
//
// Parameters:
// 	none
//
// Returns:
// 	The object that makes the Light.
//
func (light *Light) GetLightObject() *object.Object {
	return light.lightObject
}

// GetColor gets the color of the Light.
//
// Parameters:
// 	none
//
// Returns:
// 	The color of the Light.
//
func (light *Light) GetColor() []float64 {
	return light.color
}

// IsEqual checks if a Light is equal to another.
//
// Parameters:
// 	other - The other Light.
//
// Returns:
// 	If the lights are equal.
//
func (light *Light) IsEqual(other *Light) bool {
	return reflect.DeepEqual(light.GetColor(), other.GetColor()) &&
		light.GetLightIntensity() == other.GetLightIntensity() &&
		light.GetLightObject().IsEqual(other.GetLightObject())
}

// Init is a function to initialize a Light.
//
// Parameters:
//  lightIntensity - The intensity of the light.
//  lightObject    - The object tha defines the light.
//  color          - The RGB of the light.
//
// Returns:
// 	A Light.
//
func Init(lightIntensity float64, lightObject *object.Object, color []float64) (*Light, error) {
	if len(color) != 3 {
		return nil, nonRGBColorError(color)
	}
	for colorIndex := 0; colorIndex < 3; colorIndex++ {
		if color[colorIndex] < 0 || color[colorIndex] > 1 {
			return nil, colorOutOfBoundsError(color)
		}
	}
	light := &Light{lightIntensity: lightIntensity, lightObject: lightObject, color: color}
	return light, nil
}

