package point_repository

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestPointRepository_Init tests the instantiation of a PointRepository.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointRepository_Init(t *testing.T) {
	firstPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	pointRepository, err := Init([]*point.Point{firstPoint}, 3)
	test_helpers.AssertNilError(t, err)

	isEqual := firstPoint == pointRepository.points[0]
	test_helpers.AssertEqual(t, true, isEqual)
}

// TestPointRepository_Init_IncompatibleDimensionError tests the instantiation of a PointRepository with incompatible
// dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointRepository_Init_IncompatibleDimensionError(t *testing.T) {
	dimension := 3
	firstPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	secondPoint, err := point.Init(2)
	test_helpers.AssertNilError(t, err)
	points := []*point.Point{firstPoint, secondPoint}
	expectedErrorMessage := fmt.Sprintf("Not all points have %v dimensions. Points: %v.", dimension, points)

	_, err = Init(points, dimension)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestPointRepository_Init_InvalidSizeError tests the instantiation of a PointRepository with 0 points.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointRepository_Init_InvalidSizeError(t *testing.T) {
	var points []*point.Point
	expectedErrorMessage := fmt.Sprintf("Invalid points list: %v. There must be at least one point.", points)

	_, err := Init(points, 3)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestPointRepository_GetPoint tests the get point of a PointRepository.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointRepository_GetPoint(t *testing.T) {
	firstPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	secondPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	pointRepository, err := Init([]*point.Point{firstPoint, secondPoint}, 3)
	test_helpers.AssertNilError(t, err)

	receivedPoint, err := pointRepository.GetPoint(1)
	test_helpers.AssertNilError(t, err)

	isEqual := secondPoint == receivedPoint
	test_helpers.AssertEqual(t, true, isEqual)
}

// TestPointRepository_GetPoint_NegativeIndex tests the get point of a PointRepository with negative index.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointRepository_GetPoint_NegativeIndex(t *testing.T) {
	firstPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	secondPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	pointRepository, err := Init([]*point.Point{firstPoint, secondPoint}, 3)
	test_helpers.AssertNilError(t, err)
	index := -1

	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the point repository. Expected from 0 to %v and got %v.",
		len(pointRepository.points), index)

	_, err = pointRepository.GetPoint(index)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestPointRepository_GetPoint_NegativeIndex tests the get point of a PointRepository with an index out of the
// PointRepository.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointRepository_GetPoint_BiggerIndex(t *testing.T) {
	firstPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	secondPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	pointRepository, err := Init([]*point.Point{firstPoint, secondPoint}, 3)
	test_helpers.AssertNilError(t, err)
	index := 2

	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the point repository. Expected from 0 to %v and got %v.",
		len(pointRepository.points), index)

	_, err = pointRepository.GetPoint(index)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestPointRepository_NumberOfPoints tests the number of points of a PointRepository.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointRepository_NumberOfPoints(t *testing.T) {
	firstPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	secondPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	pointRepository, err := Init([]*point.Point{firstPoint, secondPoint}, 3)
	test_helpers.AssertNilError(t, err)

	test_helpers.AssertEqual(t, 2, pointRepository.NumberOfPoints())
}

// TestPointRepository_PointsDimension tests the points dimension of a PointRepository.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointRepository_PointsDimension(t *testing.T) {
	firstPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	secondPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	pointRepository, err := Init([]*point.Point{firstPoint, secondPoint}, 3)
	test_helpers.AssertNilError(t, err)

	test_helpers.AssertEqual(t, 3, pointRepository.PointsDimension())
}

// TestPointRepository_IsEqual tests the is equal of a PointRepository.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointRepository_IsEqual(t *testing.T) {
	firstPointRepository := setUpPointRepository(t)
	secondPointRepository := setUpPointRepository(t)
	isEqual := firstPointRepository.IsEqual(secondPointRepository)
	test_helpers.AssertEqual(t, true, isEqual)
}

// TestPointRepository_IsEqual_Different tests the is equal of a PointRepository when the repositories are not equal.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointRepository_IsEqual_Different(t *testing.T) {
	firstPointRepository := setUpPointRepository(t)
	secondPointRepository := setUpPointRepository(t)
	changingPoint, err := secondPointRepository.GetPoint(0)
	test_helpers.AssertNilError(t, err)
	err = changingPoint.SetCoordinate(0, -1)
	test_helpers.AssertNilError(t, err)
	isEqual := firstPointRepository.IsEqual(secondPointRepository)
	test_helpers.AssertEqual(t, false, isEqual)
}

// TestPointRepository_IsEqual_DifferentNumberOfPoints tests the is equal of a PointRepository when the repositories are
//  not equal by their number of points.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPointRepository_IsEqual_DifferentNumberOfPoints(t *testing.T) {
	firstPointRepository := setUpPointRepository(t)
	firstPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	secondPointRepository, err := Init([]*point.Point{firstPoint}, 3)
	test_helpers.AssertNilError(t, err)
	isEqual := firstPointRepository.IsEqual(secondPointRepository)
	test_helpers.AssertEqual(t, false, isEqual)
}

// setUpPointRepository builds a sample PointRepository.
//
// Parameters:
// 	t - The testing instance.
//
// Returns:
//  A sample PointRepository.
//
func setUpPointRepository(t *testing.T) *PointRepository {
	dimension := 3

	firstPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)

	err = firstPoint.SetCoordinate(0, 1)
	test_helpers.AssertNilError(t, err)
	err = firstPoint.SetCoordinate(1, 2)
	test_helpers.AssertNilError(t, err)
	err = firstPoint.SetCoordinate(2, 1)
	test_helpers.AssertNilError(t, err)

	secondPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)

	err = secondPoint.SetCoordinate(0, 0)
	test_helpers.AssertNilError(t, err)
	err = secondPoint.SetCoordinate(1, 1)
	test_helpers.AssertNilError(t, err)
	err = secondPoint.SetCoordinate(2, 2)
	test_helpers.AssertNilError(t, err)

	thirdPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)

	err = thirdPoint.SetCoordinate(0, 1)
	test_helpers.AssertNilError(t, err)
	err = thirdPoint.SetCoordinate(1, 0)
	test_helpers.AssertNilError(t, err)
	err = thirdPoint.SetCoordinate(2, 1)
	test_helpers.AssertNilError(t, err)

	pointRepository, err := Init([]*point.Point{firstPoint, secondPoint, thirdPoint}, dimension)
	test_helpers.AssertNilError(t, err)
	return pointRepository
}
