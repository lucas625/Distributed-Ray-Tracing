package plane

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestPlane_Init tests the instantiation of a Plane.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPlane_Init(t *testing.T) {
	xCoefficient := 3.0
	yCoefficient := 5.0
	zCoefficient := 1.0
	isolatedTerm := 10.0

	plane := Init(xCoefficient, yCoefficient, zCoefficient, isolatedTerm)
	test_helpers.AssertEqual(t, xCoefficient, plane.GetXCoefficient())
	test_helpers.AssertEqual(t, yCoefficient, plane.GetYCoefficient())
	test_helpers.AssertEqual(t, zCoefficient, plane.GetZCoefficient())
	test_helpers.AssertEqual(t, isolatedTerm, plane.GetIsolatedTerm())
}

// TestPlane_IsEqual tests the is equal of a Plane.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPlane_IsEqual(t *testing.T) {
	xCoefficient := 3.0
	yCoefficient := 5.0
	zCoefficient := 1.0
	isolatedTerm := 10.0

	plane := Init(xCoefficient, yCoefficient, zCoefficient, isolatedTerm)
	otherPlane := Init(xCoefficient, yCoefficient, zCoefficient, isolatedTerm)
	test_helpers.AssertEqual(t, true, plane.IsEqual(otherPlane))
}

// TestPlane_IsEqual_DifferentPlanes tests the is equal of a Plane when the planes are different.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestPlane_IsEqual_DifferentPlanes(t *testing.T) {
	xCoefficient := 3.0
	yCoefficient := 5.0
	zCoefficient := 1.0
	isolatedTerm := 10.0

	plane := Init(xCoefficient, -1, zCoefficient, isolatedTerm)
	otherPlane := Init(xCoefficient, yCoefficient, zCoefficient, isolatedTerm)
	test_helpers.AssertEqual(t, false, plane.IsEqual(otherPlane))
}
