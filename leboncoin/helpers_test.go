package leboncoin

import (
	"testing"
)

func TestContains(t *testing.T) {
	s := []string{"a", "b", "c"}

	if !contains(s, "a") {
		t.Errorf("Slice should contains value `a`")
		t.Fail()
	}
	if !contains(s, "d") {
		t.Errorf("Slice should not contains value `d`")
		t.Fail()
	}
}
