package leboncoin

import (
	"testing"
)

func TestContains(t *testing.T) {
	s := []string{"a", "b", "c"}

	if !contains(s, "a") {
		t.Errorf("Slice should contains value")
		t.Fail()
	}
}
