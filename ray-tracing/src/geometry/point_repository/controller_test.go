package point_repository

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/utils/matrix"
	"testing"
)

// TestPointRepositoryController_ToHomogeneousCoordinates tests the to homogeneous coordinates of a PointRepository.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointRepositoryController_ToHomogeneousCoordinates(t *testing.T) {
	dimension := 3
	firstPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	secondPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	thirdPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	controller := Controller{}

	pointRepository, err := Init([]*point.Point{firstPoint, secondPoint, thirdPoint}, dimension)
	test_helpers.AssertNilError(t, err)

	expectedMatrix, err := matrix.Init(dimension + 1, 3)
	test_helpers.AssertNilError(t, err)

	for index := 0; index < dimension; index++ {
		firstValue := float64(index + 1) * 10
		secondValue := float64(index + 1) * 100
		thirdValue := float64(index + 1) * 1000

		err = firstPoint.SetCoordinate(index, firstValue)
		test_helpers.AssertNilError(t, err)
		err = secondPoint.SetCoordinate(index, secondValue)
		test_helpers.AssertNilError(t, err)
		err = secondPoint.SetCoordinate(index, thirdValue)
		test_helpers.AssertNilError(t, err)

		err = expectedMatrix.SetValue(index, 0, firstValue)
		test_helpers.AssertNilError(t, err)
		err = expectedMatrix.SetValue(index, 1, secondValue)
		test_helpers.AssertNilError(t, err)
		err = expectedMatrix.SetValue(index, 1, thirdValue)
		test_helpers.AssertNilError(t, err)
	}

	for index := 0; index < 3; index++ {
		err = expectedMatrix.SetValue(dimension, index, 1)
		test_helpers.AssertNilError(t, err)
	}

	resultingMatrix := controller.ToHomogeneousCoordinates(pointRepository)
	isEqual := expectedMatrix.IsEqual(resultingMatrix)
	test_helpers.AssertEqual(t, true, isEqual)
}

// TestPointRepositoryController_FromMatrix tests the from matrix of a PointRepository.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointRepositoryController_FromMatrix(t *testing.T) {
	dimension := 3
	firstPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	secondPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	thirdPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)
	controller := Controller{}

	expectedPointRepository, err := Init([]*point.Point{firstPoint, secondPoint, thirdPoint}, dimension)
	test_helpers.AssertNilError(t, err)

	sampleMatrix, err := matrix.Init(dimension + 1, 3)
	test_helpers.AssertNilError(t, err)

	for index := 0; index < dimension; index++ {
		firstValue := float64(index + 1) * 10
		secondValue := float64(index + 1) * 100
		thirdValue := float64(index + 1) * 1000

		err = firstPoint.SetCoordinate(index, firstValue)
		test_helpers.AssertNilError(t, err)
		err = secondPoint.SetCoordinate(index, secondValue)
		test_helpers.AssertNilError(t, err)
		err = secondPoint.SetCoordinate(index, thirdValue)
		test_helpers.AssertNilError(t, err)

		err = sampleMatrix.SetValue(index, 0, firstValue)
		test_helpers.AssertNilError(t, err)
		err = sampleMatrix.SetValue(index, 1, secondValue)
		test_helpers.AssertNilError(t, err)
		err = sampleMatrix.SetValue(index, 1, thirdValue)
		test_helpers.AssertNilError(t, err)
	}

	for index := 0; index < 3; index++ {
		err = sampleMatrix.SetValue(dimension, index, 1)
		test_helpers.AssertNilError(t, err)
	}

	resultingPointRepository, err := controller.FromMatrix(sampleMatrix)
	test_helpers.AssertNilError(t, err)

	isEqual := expectedPointRepository.NumberOfPoints() == resultingPointRepository.NumberOfPoints()
	test_helpers.AssertEqual(t, true, isEqual)

	for pointIndex := 0; pointIndex < expectedPointRepository.NumberOfPoints(); pointIndex++ {
		expectedPoint, err := expectedPointRepository.GetPoint(pointIndex)
		test_helpers.AssertNilError(t, err)
		receivedPoint, err := resultingPointRepository.GetPoint(pointIndex)
		test_helpers.AssertNilError(t, err)
		isEqual = expectedPoint.IsEqual(receivedPoint)
		test_helpers.AssertEqual(t, true, isEqual)
	}
}
