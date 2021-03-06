package point

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/utils/matrix"
	"testing"
)

// TestPointController_ExtractVector tests the vector extraction between two points.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointController_ExtractVector(t *testing.T) {
	startingPoint := &Point{coordinates: []float64{10, 20, 30}}
	targetPoint := &Point{coordinates: []float64{12, 15, 45}}
	controller := &Controller{}
	expectedVector, err := vector.Init(3)
	test_helpers.AssertNilError(t, err)

	err = expectedVector.SetCoordinate(0, 2)
	test_helpers.AssertNilError(t, err)
	err = expectedVector.SetCoordinate(1, -5)
	test_helpers.AssertNilError(t, err)
	err = expectedVector.SetCoordinate(2, 15)
	test_helpers.AssertNilError(t, err)

	extractedVector, err := controller.ExtractVector(startingPoint, targetPoint)
	test_helpers.AssertNilError(t, err)

	areVectorsEqual := expectedVector.IsEqual(extractedVector)
	test_helpers.AssertEqual(t, true, areVectorsEqual)
}

// TestPointController_ExtractVector_DifferentDimensions tests the vector extraction between two points.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointController_ExtractVector_DifferentDimensions(t *testing.T) {
	startingPoint := &Point{coordinates: []float64{10, 20, 30}}
	targetPoint := &Point{coordinates: []float64{12, 15}}
	controller := &Controller{}
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of point. Expected: %d and got: %d.\n", startingPoint.Dimension(), targetPoint.Dimension())

	_, err := controller.ExtractVector(startingPoint, targetPoint)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestPointController_ToHomogeneousCoordinates tests the Point to homogeneous coordinates.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointController_ToHomogeneousCoordinates(t *testing.T) {
	point := &Point{coordinates: []float64{10, 20, 30}}
	controller := &Controller{}
	expectedMatrix, err := matrix.Init(4, 1)
	test_helpers.AssertNilError(t, err)

	err = expectedMatrix.SetValue(0, 0, 10)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(1, 0, 20)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(2, 0, 30)
	test_helpers.AssertNilError(t, err)
	err = expectedMatrix.SetValue(3, 0, 1)

	pointAsMatrix := controller.ToHomogeneousCoordinates(point)
	test_helpers.AssertNilError(t, err)

	areMatricesEqual := expectedMatrix.IsEqual(pointAsMatrix)
	test_helpers.AssertEqual(t, true, areMatricesEqual)
}

// TestPointController_SumWithVector tests the sum with vector of a Point.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointController_SumWithVector(t *testing.T) {
	expectedPoint := &Point{coordinates: []float64{3, -1, 20}}
	startingPoint := &Point{coordinates: []float64{1, 4, 8}}
	controller := &Controller{}
	sumVector, err := vector.Init(3)
	test_helpers.AssertNilError(t, err)

	err = sumVector.SetCoordinate(0, 2)
	test_helpers.AssertNilError(t, err)
	err = sumVector.SetCoordinate(1, -5)
	test_helpers.AssertNilError(t, err)
	err = sumVector.SetCoordinate(2, 12)
	test_helpers.AssertNilError(t, err)

	newPoint, err := controller.SumWithVector(startingPoint, sumVector)
	test_helpers.AssertNilError(t, err)

	test_helpers.AssertEqual(t, true, expectedPoint.IsEqual(newPoint))
}

// TestPointController_SumWithVector_PointAndVectorIncompatibleDimensionError tests the sum with vector of a Point when
// the dimension of the Point is incompatible with the dimension of the vector.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointController_SumWithVector_PointAndVectorIncompatibleDimensionError(t *testing.T) {
	startingPoint := &Point{coordinates: []float64{1, 4, 8}}
	controller := &Controller{}
	sumVector, err := vector.Init(2)
	test_helpers.AssertNilError(t, err)
	expectedErrorMessage := fmt.Sprintf(
		"Incompatible dimension for point: %d and vector: %d.", startingPoint.Dimension(), sumVector.Dimension())

	_, err = controller.SumWithVector(startingPoint, sumVector)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
