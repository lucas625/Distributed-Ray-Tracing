package point_repository

import (
	"fmt"
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

	isEqual := expectedPointRepository.IsEqual(resultingPointRepository)
	test_helpers.AssertEqual(t, true, isEqual)
}

// TestPointRepositoryController_MultiplyByMatrix tests the multiply by matrix of a PointRepository.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointRepositoryController_MultiplyByMatrix(t *testing.T) {
	controller := Controller{}
	pointRepository := setUpPointRepository(t)
	multiplyingMatrix := setUpMatrix(t, 3)

	expectedPointRepository := setUpPointRepository(t)

	expectedValues := [][]float64{{8, 9, 4}, {8, 7, 5}, {4, 7, 2}}

	for pointIndex := 0; pointIndex < expectedPointRepository.NumberOfPoints(); pointIndex++ {
		currentPoint, err := expectedPointRepository.GetPoint(pointIndex)
		test_helpers.AssertNilError(t, err)
		for coordinateIndex := 0; coordinateIndex < currentPoint.Dimension(); coordinateIndex++ {
			err = currentPoint.SetCoordinate(coordinateIndex, expectedValues[pointIndex][coordinateIndex])
			test_helpers.AssertNilError(t, err)
		}
	}

	resultingPointRepository, err := controller.MultiplyByMatrix(pointRepository, multiplyingMatrix)
	test_helpers.AssertNilError(t, err)

	isEqual := expectedPointRepository.IsEqual(resultingPointRepository)
	test_helpers.AssertEqual(t, true, isEqual)
}

// TestPointRepositoryController_MultiplyByMatrix_InvalidMultiplication tests the multiply by matrix of a
// PointRepository with an invalid multiplication.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointRepositoryController_MultiplyByMatrix_InvalidMultiplication(t *testing.T) {
	controller := Controller{}
	pointRepository := setUpPointRepository(t)
	multiplyingMatrix := setUpMatrix(t, 2)
	expectedErrorMessage := fmt.Sprintf(
		"Incompatible size for matrices:\nFirst matrix: lines: %d and columns: %d." +
			"\nSecond matrix: lines: %d and columns: %d.",
		multiplyingMatrix.Lines(), multiplyingMatrix.Columns(), 4, 3)

	_, err := controller.MultiplyByMatrix(pointRepository, multiplyingMatrix)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// setUpMatrix builds a sample matrix.
//
// Parameters:
// 	t         - The testing instance.
//  dimension - The dimension of the sample matrix (2 or 3).
//
// Returns:
//  A sample matrix.
//
func setUpMatrix(t *testing.T, dimension int) *matrix.Matrix {
	matrixController := matrix.Controller{}

	sampleMatrix, err := matrixController.BuildHomogeneousCoordinates(dimension)
	test_helpers.AssertNilError(t, err)

	err = sampleMatrix.SetValue(0, 0, 1)
	test_helpers.AssertNilError(t, err)
	err = sampleMatrix.SetValue(0, 1, 2)
	test_helpers.AssertNilError(t, err)
	err = sampleMatrix.SetValue(0, 2, 3)
	test_helpers.AssertNilError(t, err)

	err = sampleMatrix.SetValue(1, 0, 4)
	test_helpers.AssertNilError(t, err)
	err = sampleMatrix.SetValue(1, 1, 1)
	test_helpers.AssertNilError(t, err)
	err = sampleMatrix.SetValue(1, 2, 3)
	test_helpers.AssertNilError(t, err)

	err = sampleMatrix.SetValue(2, 0, 0)
	test_helpers.AssertNilError(t, err)
	err = sampleMatrix.SetValue(2, 1, 1)
	test_helpers.AssertNilError(t, err)
	err = sampleMatrix.SetValue(2, 2, 2)
	test_helpers.AssertNilError(t, err)
	return sampleMatrix
}

