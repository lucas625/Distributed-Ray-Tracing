package object

import (
	"fmt"
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

// TestObject_Init tests the instantiation of an Object.
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

// TestObject_Init_Error tests the instantiation of an Object.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestObject_Init_Error(t *testing.T) {
	repository := buildSamplePointRepository(t)
	normals := buildNormals(t)
	firstTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)
	triangles := []*triangle.Triangle{firstTriangle}

	name := "my object"
	color := []float64{0.1, 0.25, 0.5, 10}
	specularDecay := 5.0
	specularReflection := 0.5
	roughNess := 0.0
	transmissionReflection := 0.25
	diffuseReflection := 0.25

	expectedErrorMessage := fmt.Sprintf("There are not 3 color values: %d.", len(color))

	_, err = Init(name, repository, triangles, normals, color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestObject_IsEqual_DifferentTriangleLength tests the is equal of an Object.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestObject_IsEqual_DifferentTriangleLength(t *testing.T) {
	repository := buildSamplePointRepository(t)
	normals := buildNormals(t)
	firstTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)
	secondTriangle, err := triangle.Init([]int{1, 1, 2}, []int{1, 1, 2})
	test_helpers.AssertNilError(t, err)
	firstTriangles := []*triangle.Triangle{firstTriangle, secondTriangle}
	otherTriangles := []*triangle.Triangle{firstTriangle}

	name := "my object"
	color := []float64{0.1, 0.25, 0.5}
	specularDecay := 5.0
	specularReflection := 0.5
	roughNess := 0.0
	transmissionReflection := 0.25
	diffuseReflection := 0.25

	firstObject, err := Init(name, repository, firstTriangles, normals, color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	test_helpers.AssertNilError(t, err)

	secondObject, err := Init(name, repository, otherTriangles, normals, color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	test_helpers.AssertNilError(t, err)

	test_helpers.AssertEqual(t, false, firstObject.IsEqual(secondObject))
}

// TestObject_IsEqual_DifferentNormalsLength tests the is equal of an Object.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestObject_IsEqual_DifferentNormalsLength(t *testing.T) {
	repository := buildSamplePointRepository(t)
	firstNormals := buildNormals(t)
	otherVector, err := vector.Init(3)
	test_helpers.AssertNilError(t, err)
	otherNormals := append(buildNormals(t), otherVector)
	firstTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	secondTriangle, err := triangle.Init([]int{1, 1, 2}, []int{1, 1, 2})
	test_helpers.AssertNilError(t, err)
	firstTriangles := []*triangle.Triangle{firstTriangle, secondTriangle}

	name := "my object"
	color := []float64{0.1, 0.25, 0.5}
	specularDecay := 5.0
	specularReflection := 0.5
	roughNess := 0.0
	transmissionReflection := 0.25
	diffuseReflection := 0.25

	firstObject, err := Init(name, repository, firstTriangles, firstNormals, color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	test_helpers.AssertNilError(t, err)

	secondObject, err := Init(name, repository, firstTriangles, otherNormals, color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	test_helpers.AssertNilError(t, err)

	test_helpers.AssertEqual(t, false, firstObject.IsEqual(secondObject))
}

// TestObject_IsEqual_DifferentTriangles tests the is equal of an Object.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestObject_IsEqual_DifferentTriangles(t *testing.T) {
	repository := buildSamplePointRepository(t)
	normals := buildNormals(t)
	firstTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)
	secondTriangle, err := triangle.Init([]int{1, 1, 2}, []int{1, 1, 2})
	test_helpers.AssertNilError(t, err)
	firstTriangles := []*triangle.Triangle{secondTriangle}
	otherTriangles := []*triangle.Triangle{firstTriangle}

	name := "my object"
	color := []float64{0.1, 0.25, 0.5}
	specularDecay := 5.0
	specularReflection := 0.5
	roughNess := 0.0
	transmissionReflection := 0.25
	diffuseReflection := 0.25

	firstObject, err := Init(name, repository, firstTriangles, normals, color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	test_helpers.AssertNilError(t, err)

	secondObject, err := Init(name, repository, otherTriangles, normals, color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	test_helpers.AssertNilError(t, err)

	test_helpers.AssertEqual(t, false, firstObject.IsEqual(secondObject))
}

// TestObject_IsEqual_DifferentNormals tests the is equal of an Object.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestObject_IsEqual_DifferentNormals(t *testing.T) {
	repository := buildSamplePointRepository(t)
	firstNormals := buildNormals(t)
	otherVector, err := vector.Init(4)
	test_helpers.AssertNilError(t, err)
	otherNormals := buildNormals(t)
	otherNormals[0] = otherVector
	firstTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	secondTriangle, err := triangle.Init([]int{1, 1, 2}, []int{1, 1, 2})
	test_helpers.AssertNilError(t, err)
	firstTriangles := []*triangle.Triangle{firstTriangle, secondTriangle}

	name := "my object"
	color := []float64{0.1, 0.25, 0.5}
	specularDecay := 5.0
	specularReflection := 0.5
	roughNess := 0.0
	transmissionReflection := 0.25
	diffuseReflection := 0.25

	firstObject, err := Init(name, repository, firstTriangles, firstNormals, color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	test_helpers.AssertNilError(t, err)

	secondObject, err := Init(name, repository, firstTriangles, otherNormals, color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	test_helpers.AssertNilError(t, err)

	test_helpers.AssertEqual(t, false, firstObject.IsEqual(secondObject))
}
