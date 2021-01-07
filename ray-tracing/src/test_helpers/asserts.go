//!test

package test_helpers

import (
	"testing"
)

// AssertEqual throws an testing error when the values are not the expected, only usable on values that can be
// checked directly by "!=".
//
// Parameters:
//	t             - The test_helpers instance.
//	expectedValue - The expected value.
//	receivedValue - The value received.
//
// Returns:
//  none
//
func AssertEqual(t *testing.T, expectedValue, receivedValue interface{}) {
	if expectedValue != receivedValue {
		t.Errorf("%v is not %v", receivedValue, expectedValue)
	}
}

// AssertNilError throws an testing error when the error is not nil.
//
// Parameters:
//	t   - The test_helpers instance.
//	err - The error.
//
// Returns:
//  none
//
func AssertNilError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

// AssertNotNilError throws an testing error when the error is nil.
//
// Parameters:
//	t   - The test_helpers instance.
//	err - The error.
//
// Returns:
//  none
//
func AssertNotNilError(t *testing.T, err error) {
	if err == nil {
		t.Error("Error should not be nil.")
	}
}
