package light

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point_repository"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/triangle"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rendering/object"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// buildSamplePointRepository builds a sample point repository for testing.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  A sample point repository.
//
func buildSamplePointRepository(t *testing.T) *point_repository.PointRepository {
	dimension := 3

	firstPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = firstPoint.SetCoordinate(0, 2)
	test_helpers.AssertNilError(t, err)

	secondPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = secondPoint.SetCoordinate(1, 2)
	test_helpers.AssertNilError(t, err)

	thirdPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	err = thirdPoint.SetCoordinate(2, 2)
	test_helpers.AssertNilError(t, err)

	repository, err := point_repository.Init([]*point.Point{firstPoint, secondPoint, thirdPoint}, dimension)
	test_helpers.AssertNilError(t, err)
	return repository
}

// buildNormals builds the normals.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  A list of vectors.
//
func buildNormals(t *testing.T) []*vector.Vector {
	dimension := 3

	firstVector, err := vector.Init(dimension)
	test_helpers.AssertNilError(t, err)
	secondVector, err := vector.Init(dimension)
	test_helpers.AssertNilError(t, err)
	ThirdVector, err := vector.Init(dimension)
	test_helpers.AssertNilError(t, err)

	return []*vector.Vector{firstVector, secondVector, ThirdVector}
}

// TestLight_Init tests the instantiation of a Light.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestLight_Init(t *testing.T) {
	repository := buildSamplePointRepository(t)
	normals := buildNormals(t)
	firstTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)
	triangles := []*triangle.Triangle{firstTriangle}

	name := "my light"
	color := []float64{0, 0, 0}
	specularDecay := 0.0
	specularReflection := 1.0
	roughNess := 0.0
	transmissionReflection := 0.0
	diffuseReflection := 0.0

	lightObject, err := object.Init(name, repository, triangles, normals, color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	test_helpers.AssertNilError(t, err)

	lightIntensity := 5.0
	lightColor := []float64{1, 0.5, 0.25}

	expectedLight := &Light{lightIntensity: lightIntensity, lightObject: lightObject, color: lightColor}

	receivedLight, err := Init(lightIntensity, lightObject, lightColor)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, true, expectedLight.IsEqual(receivedLight))
}

// TestLight_Init_NonRGBColorError tests the instantiation of a Light.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestLight_Init_NonRGBColorError(t *testing.T) {
	repository := buildSamplePointRepository(t)
	normals := buildNormals(t)
	firstTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)
	triangles := []*triangle.Triangle{firstTriangle}

	name := "my light"
	color := []float64{0, 0, 0}
	specularDecay := 0.0
	specularReflection := 1.0
	roughNess := 0.0
	transmissionReflection := 0.0
	diffuseReflection := 0.0

	lightObject, err := object.Init(name, repository, triangles, normals, color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	test_helpers.AssertNilError(t, err)

	lightIntensity := 5.0
	lightColor := []float64{1, 0.5}

	expectedErrorMessage := fmt.Sprintf("There are not 3 color values: %d.", len(lightColor))

	_, err = Init(lightIntensity, lightObject, lightColor)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestLight_Init_ColorOutOfBoundsError tests the instantiation of a Light.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestLight_Init_ColorOutOfBoundsError(t *testing.T) {
	repository := buildSamplePointRepository(t)
	normals := buildNormals(t)
	firstTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)
	triangles := []*triangle.Triangle{firstTriangle}

	name := "my light"
	color := []float64{0, 0, 0}
	specularDecay := 0.0
	specularReflection := 1.0
	roughNess := 0.0
	transmissionReflection := 0.0
	diffuseReflection := 0.0

	lightObject, err := object.Init(name, repository, triangles, normals, color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	test_helpers.AssertNilError(t, err)

	lightIntensity := 5.0
	lightColor := []float64{255, 0.5, 1}

	expectedErrorMessage := fmt.Sprintf("Color values out of interval [0,1]: %v.", lightColor)

	_, err = Init(lightIntensity, lightObject, lightColor)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
