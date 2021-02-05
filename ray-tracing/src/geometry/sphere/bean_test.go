package sphere

import (
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestSphere_Init tests the instantiation of a Sphere.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestSphere_Init(t *testing.T) {
	centerPointIndex := 2
	radius := 5.0

	sphere := Init(centerPointIndex, radius)
	test_helpers.AssertEqual(t, centerPointIndex, sphere.GetCenterPointIndex())
	test_helpers.AssertEqual(t, radius, sphere.GetRadius())
}

// TestSphere_IsEqual tests the is equal of a Sphere.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestSphere_IsEqual(t *testing.T) {
	centerPointIndex := 2
	radius := 5.0

	sphere := Init(centerPointIndex, radius)
	otherSphere := Init(centerPointIndex, radius)
	test_helpers.AssertEqual(t, true, sphere.IsEqual(otherSphere))
}

// TestSphere_IsEqual_Different tests the is equal of a Sphere when they are different.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestSphere_IsEqual_Different(t *testing.T) {
	centerPointIndex := 2
	radius := 5.0

	sphere := Init(centerPointIndex, radius)
	otherSphere := Init(centerPointIndex, 3)
	test_helpers.AssertEqual(t, false, sphere.IsEqual(otherSphere))
}
