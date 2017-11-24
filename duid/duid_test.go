package duid

import (
	"testing"
)

// TestGenerator checks to see that duid.Generator makes unique ids.
func TestGenerator(t *testing.T) {
	var generator = NewGenerator(0)
	var id1 = generator.NewId()
	var id2 = generator.NewId()
	if id1 == id2 {
		t.Errorf("Generator did not make unique ids: %d and %d", id1, id2)
	} else {
		t.Logf("Generator made unique ids: %d and %d.", id1, id2)
	}

}
