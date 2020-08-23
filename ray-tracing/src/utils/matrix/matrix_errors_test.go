package matrix

import (
	"fmt"
	"testing"
)

// TestInvalidSizeError tests invalid size error for matrix.
//
// Parameters:
//  t - Test instance.
//
// Returns:
//  none
//
func TestInvalidSizeError(t *testing.T) {
	size := -1
	err := invalidSize(size, size)
	if err == nil {
		t.Errorf("No invalid size error return for size: %d.", size)
	} else if err.Error() != fmt.Sprintf("Invalid size for matrix. Lines: %d and Columns: %d.\n", size, size) {
		t.Errorf("Wrong error message for invalid size: \"%s\".", err.Error())
	}
}
