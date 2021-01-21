package point

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestPoint_Init tests the instantiation of a Point.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_Init(t *testing.T) {
	dimension := 3

	point, err := Init(dimension)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, dimension, point.Dimension())

	for index := 0; index < dimension; index++ {
		pointCoordinate, err := point.GetCoordinate(index)
		test_helpers.AssertNilError(t, err)
		test_helpers.AssertEqual(t, 0.0, pointCoordinate)
	}
}

// TestPoint_Init_0D tests the instantiation of a Point with zero dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_Init_0D(t *testing.T) {
	dimension := 0

	point, err := Init(dimension)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, dimension, point.Dimension())
}

// TestPoint_Init_NegativeDimension tests the instantiation of a Point with negative dimension.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_Init_NegativeDimension(t *testing.T) {
	dimension := -1
	expectedErrorMessage := fmt.Sprintf("Invalid dimension for point: %d.", dimension)

	_, err := Init(dimension)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestPoint_GetCoordinate tests the get coordinate of a Point.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_GetCoordinate(t *testing.T) {
	point := Point{coordinates: []float64{10, 20, 30}}
	expectedCoordinate := 20.0

	receivedCoordinate, err := point.GetCoordinate(1)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, expectedCoordinate, receivedCoordinate)
}

// TestPoint_GetCoordinate_NegativeIndex tests the get coordinate of a Point with a negative index.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_GetCoordinate_NegativeIndex(t *testing.T) {
	point := Point{coordinates: []float64{10, 20, 30}}
	index := -1
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the point. Expected from 0 to %v and got %v.", point.Dimension(), index)

	_, err := point.GetCoordinate(index)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestPoint_GetCoordinate_BiggerIndex tests the get coordinate of a Point with an index out of the Point.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_GetCoordinate_BiggerIndex(t *testing.T) {
	point := Point{coordinates: []float64{10, 20, 30}}
	index := 3
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the point. Expected from 0 to %v and got %v.", point.Dimension(), index)

	_, err := point.GetCoordinate(index)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestPoint_SetCoordinate tests the set coordinate of a Point.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_SetCoordinate(t *testing.T) {
	point := Point{coordinates: []float64{10, 20, 30}}
	expectedCoordinate := 50.0
	index := 1

	err := point.SetCoordinate(index, expectedCoordinate)
	test_helpers.AssertNilError(t, err)

	receivedCoordinate, err := point.GetCoordinate(index)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, expectedCoordinate, receivedCoordinate)
}

// TestPoint_SetCoordinate_NegativeIndex tests the set coordinate of a Point with a negative index.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_SetCoordinate_NegativeIndex(t *testing.T) {
	point := Point{coordinates: []float64{10, 20, 30}}
	index := -1
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the point. Expected from 0 to %v and got %v.", point.Dimension(), index)

	err := point.SetCoordinate(index, 10)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestPoint_SetCoordinate_BiggerIndex tests the set coordinate of a Point with an index out of the Point.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_SetCoordinate_BiggerIndex(t *testing.T) {
	point := Point{coordinates: []float64{10, 20, 30}}
	index := 3
	expectedErrorMessage := fmt.Sprintf(
		"Index out of limits of the point. Expected from 0 to %v and got %v.", point.Dimension(), index)

	err := point.SetCoordinate(index, 10)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestPoint_IsEqual tests the is equal of a Point.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_IsEqual(t *testing.T) {
	firstPoint := &Point{coordinates: []float64{10, 20, 30}}
	secondPoint := &Point{coordinates: []float64{10, 20, 30}}

	isEqual := firstPoint.IsEqual(secondPoint)
	test_helpers.AssertEqual(t, true, isEqual)
}

// TestPoint_IsEqual_Different tests the is equal of a Point when the points are different.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_IsEqual_Different(t *testing.T) {
	firstPoint := &Point{coordinates: []float64{10, 20, 30}}
	secondPoint := &Point{coordinates: []float64{10, 20, 50}}

	isEqual := firstPoint.IsEqual(secondPoint)
	test_helpers.AssertEqual(t, false, isEqual)
}

// TestPoint_IsEqual_DifferentDimension tests the is equal of a Point when the points have different dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPoint_IsEqual_DifferentDimension(t *testing.T) {
	firstPoint := &Point{coordinates: []float64{10, 20, 30}}
	secondPoint := &Point{coordinates: []float64{10, 20, 30, 40}}

	isEqual := firstPoint.IsEqual(secondPoint)
	test_helpers.AssertEqual(t, false, isEqual)
}
