package screen

import (
	"fmt"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/test_helpers"
	"testing"
)

// TestScreen_Init tests the instantiation of a Screen.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestScreen_Init(t *testing.T) {
	screen, err := Init(16, 9)
	test_helpers.AssertNilError(t, err)

	expectedScreen := &Screen{width: 16, height: 9}
	test_helpers.AssertEqual(t, true, expectedScreen.IsEqual(screen))
}

// TestScreen_Init_SizeError tests the instantiation of a Screen when there is a size error.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestScreen_Init_SizeError(t *testing.T) {
	expectedErrorMessage := fmt.Sprintf("Invalid size for screen: width: %d, height: %d.", -1, 9)
	_, err := Init(-1, 9)
	test_helpers.AssertNotNilError(t, err)
	test_helpers.AssertEqual(t, expectedErrorMessage, err.Error())
}
