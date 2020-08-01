package vector

import (
	"testing"
)

func TestInitPositiveDimension(t *testing.T) {
	dimension := 1
	_, err := Init(dimension)
	if err != nil {
		t.Errorf("Vector failed to be Instantiated with %dD", dimension)
	}
}
