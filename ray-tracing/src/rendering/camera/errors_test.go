package camera

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/point"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/geometry/vector"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestCamera_Non3DCameraError tests the error where the camera is not on the third dimension.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestCamera_Non3DCameraError(t *testing.T) {
	position, err := point.Init(2)
	test_helpers.AssertNilError(t, err)
	look, err := vector.Init(2)
	test_helpers.AssertNilError(t, err)
	up, err := vector.Init(2)
	test_helpers.AssertNilError(t, err)
	right, err := vector.Init(2)
	test_helpers.AssertNilError(t, err)

	expectedErrorMessage := fmt.Sprintf("Non 3D camera. Dimensions: position %dD, look %dD, up %dD, right %dD.",
		position.Dimension(), look.Dimension(), up.Dimension(), right.Dimension())
	err = non3DCameraError(position, look, up, right)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
