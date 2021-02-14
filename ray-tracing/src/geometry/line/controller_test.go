package line

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestController_ExtractLine tests the extraction of a Line.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_ExtractLine(t *testing.T) {
	dimension := 2

	startingPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)

	err = startingPoint.SetCoordinate(0, 3)
	test_helpers.AssertNilError(t, err)
	err = startingPoint.SetCoordinate(1, 10)
	test_helpers.AssertNilError(t, err)

	targetPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)

	err = targetPoint.SetCoordinate(0, 2)
	test_helpers.AssertNilError(t, err)
	err = targetPoint.SetCoordinate(1, 15)
	test_helpers.AssertNilError(t, err)

	expectedVectorDirector, err := vector.Init(dimension)
	test_helpers.AssertNilError(t, err)

	err = expectedVectorDirector.SetCoordinate(0, -1)
	test_helpers.AssertNilError(t, err)
	err = expectedVectorDirector.SetCoordinate(1, 5)
	test_helpers.AssertNilError(t, err)

	lineController := &Controller{}
	extractedLine, err := lineController.ExtractLine(startingPoint, targetPoint)
	test_helpers.AssertNilError(t, err)

	expectedLine := Line{startingPoint: startingPoint, vectorDirector: expectedVectorDirector}
	test_helpers.AssertEqual(t, true, expectedLine.IsEqual(extractedLine))
}

// TestController_ExtractLine_DifferentDimensions tests the extraction of a Line when the points have different
// dimensions.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_ExtractLine_DifferentDimensions(t *testing.T) {
	startingPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)
	targetPoint, err := point.Init(2)
	test_helpers.AssertNilError(t, err)
	expectedErrorMessage := fmt.Sprintf(
		"Invalid dimension of point. Expected: %d and got: %d.\n", startingPoint.Dimension(), targetPoint.Dimension())

	lineController := &Controller{}
	_, err = lineController.ExtractLine(startingPoint, targetPoint)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestController_ExtractLine tests the extraction of a Line.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestController_FindPoint(t *testing.T) {
	expectedPoint, err := point.Init(2)
	test_helpers.AssertNilError(t, err)

	err = expectedPoint.SetCoordinate(0, 9)
	test_helpers.AssertNilError(t, err)
	err = expectedPoint.SetCoordinate(1, 13)
	test_helpers.AssertNilError(t, err)

	startingPoint, err := point.Init(2)
	test_helpers.AssertNilError(t, err)

	err = startingPoint.SetCoordinate(0, 3)
	test_helpers.AssertNilError(t, err)
	err = startingPoint.SetCoordinate(1, 10)
	test_helpers.AssertNilError(t, err)

	vectorDirector, err := vector.Init(2)
	test_helpers.AssertNilError(t, err)

	err = vectorDirector.SetCoordinate(0, 2)
	test_helpers.AssertNilError(t, err)
	err = vectorDirector.SetCoordinate(1, 1)
	test_helpers.AssertNilError(t, err)

	line := &Line{startingPoint: startingPoint, vectorDirector: vectorDirector}

	lineController := &Controller{}
	linePoint, err := lineController.FindPoint(line, 3)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, true, expectedPoint.IsEqual(linePoint))
}
