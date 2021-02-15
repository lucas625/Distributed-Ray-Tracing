package object

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point_repository"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/triangle"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
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

// TestObject_Init tests the instantiation of a Object.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestObject_Init(t *testing.T) {
	repository := buildSamplePointRepository(t)
	normals := buildNormals(t)
	firstTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)
	triangles := []*triangle.Triangle{firstTriangle}

	name := "my object"
	color := []float64{0.1, 0.25, 0.5}
	specularDecay := 5.0
	specularReflection := 0.5
	roughNess := 0.0
	transmissionReflection := 0.25
	diffuseReflection := 0.25

	object, err := Init(name, repository, triangles, normals, color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	test_helpers.AssertNilError(t, err)

	expectedLightCharacteristics := &lightCharacteristics{color: color, specularDecay: specularDecay,
		specularReflection: specularReflection, roughNess: roughNess, transmissionReflection: transmissionReflection,
		diffuseReflection: diffuseReflection}

	expectedObject := &Object{
		name: name, repository: repository, triangles: triangles, normals: normals,
		lightCharacteristics: expectedLightCharacteristics}
	test_helpers.AssertEqual(t, true, expectedObject.IsEqual(object))
}
