package object

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/triangle"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"reflect"
	"testing"
)

// TestController_GetBoundingBox tests the bounding box of an Object.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_GetBoundingBox(t *testing.T) {
	repository := buildSamplePointRepository(t)
	firstNormals := buildNormals(t)
	otherVector, err := vector.Init(4)
	test_helpers.AssertNilError(t, err)
	otherNormals := buildNormals(t)
	otherNormals[0] = otherVector
	firstTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)
	secondTriangle, err := triangle.Init([]int{1, 1, 2}, []int{1, 1, 2})
	test_helpers.AssertNilError(t, err)
	triangles := []*triangle.Triangle{firstTriangle, secondTriangle}

	name := "my object"
	color := []float64{0.1, 0.25, 0.5}
	specularDecay := 5.0
	specularReflection := 0.5
	roughNess := 0.0
	transmissionReflection := 0.25
	diffuseReflection := 0.25

	object, err := Init(name, repository, triangles, firstNormals, color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	test_helpers.AssertNilError(t, err)

	objectController := Controller{}
	boundingBox := objectController.GetBoundingBox(object)

	expectedBoundingBox := []float64{0, 0, 0, 2, 2, 2}

	test_helpers.AssertEqual(t, true, reflect.DeepEqual(expectedBoundingBox, boundingBox))
}

// TestController_GetCenter tests the get center of an Object.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_GetCenter(t *testing.T) {
	repository := buildSamplePointRepository(t)
	firstNormals := buildNormals(t)
	otherVector, err := vector.Init(4)
	test_helpers.AssertNilError(t, err)
	otherNormals := buildNormals(t)
	otherNormals[0] = otherVector
	firstTriangle, err := triangle.Init([]int{0, 1, 2}, []int{0, 1, 2})
	test_helpers.AssertNilError(t, err)
	secondTriangle, err := triangle.Init([]int{1, 1, 2}, []int{1, 1, 2})
	test_helpers.AssertNilError(t, err)
	triangles := []*triangle.Triangle{firstTriangle, secondTriangle}

	name := "my object"
	color := []float64{0.1, 0.25, 0.5}
	specularDecay := 5.0
	specularReflection := 0.5
	roughNess := 0.0
	transmissionReflection := 0.25
	diffuseReflection := 0.25

	object, err := Init(name, repository, triangles, firstNormals, color, specularDecay, specularReflection, roughNess,
		transmissionReflection, diffuseReflection)
	test_helpers.AssertNilError(t, err)

	objectController := Controller{}
	centerPoint := objectController.GetCenter(object)

	expectedCenterPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)

	err = expectedCenterPoint.SetCoordinate(0, 1)
	test_helpers.AssertNilError(t, err)
	err = expectedCenterPoint.SetCoordinate(1, 1)
	test_helpers.AssertNilError(t, err)
	err = expectedCenterPoint.SetCoordinate(2, 1)
	test_helpers.AssertNilError(t, err)

	test_helpers.AssertEqual(t, true, expectedCenterPoint.IsEqual(centerPoint))
}
