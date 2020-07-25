package vector

import (
	"testing"
)

func TestInitPositiveDimension(t *testing.T) {
	vect, err := Init(1)
	if err == nil {
		t.Errorf("Vector Initiated with %dD", len(vect.Coordinates))
	}
}
