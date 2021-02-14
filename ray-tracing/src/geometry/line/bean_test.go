package line

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestLine_Init tests the instantiation of a Line.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestLine_Init(t *testing.T) {
	dimension := 3

	startingPoint, err := point.Init(dimension)
	test_helpers.AssertNilError(t, err)

	vectorDirector, err := vector.Init(dimension)
	test_helpers.AssertNilError(t, err)

	line, err := Init(startingPoint, vectorDirector)
	test_helpers.AssertNilError(t, err)

	expectedLine := Line{startingPoint: startingPoint, vectorDirector: vectorDirector}
	test_helpers.AssertEqual(t, true, expectedLine.IsEqual(line))
}

// TestLine_Init_PointAndVectorIncompatibleDimensionError tests the instantiation of a Line when the starting point and
// the vector are not of the same dimension.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestLine_Init_PointAndVectorIncompatibleDimensionError(t *testing.T) {
	startingPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)

	vectorDirector, err := vector.Init(2)
	test_helpers.AssertNilError(t, err)

	_, err = Init(startingPoint, vectorDirector)
	test_helpers.AssertNotNilError(t, err)
	expectedErrorMessage := fmt.Sprintf(
		"Incompatible dimension for point: %d and vector: %d.", startingPoint.Dimension(), vectorDirector.Dimension())
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}

// TestLine_IsEqual_DifferentDimensions tests the instantiation of a Line.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestLine_IsEqual_DifferentDimensions(t *testing.T) {
	firstLineStartingPoint, err := point.Init(3)
	test_helpers.AssertNilError(t, err)

	firstLineVectorDirector, err := vector.Init(3)
	test_helpers.AssertNilError(t, err)

	firstLine, err := Init(firstLineStartingPoint, firstLineVectorDirector)
	test_helpers.AssertNilError(t, err)

	secondLineStartingPoint, err := point.Init(2)
	test_helpers.AssertNilError(t, err)

	secondLineVectorDirector, err := vector.Init(2)
	test_helpers.AssertNilError(t, err)

	secondLine, err := Init(secondLineStartingPoint, secondLineVectorDirector)
	test_helpers.AssertNilError(t, err)
	test_helpers.AssertEqual(t, false, firstLine.IsEqual(secondLine))
}
