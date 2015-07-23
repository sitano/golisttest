package set

import (
	"testing"
)

func TestSetMustBeEmpty(t *testing.T) {
	s := NewSet()

	if s.storage == nil || s.storage.String() != "()" {
		t.Fatalf("The set %v is invalid.", s)
	}
}

func TestSetAdd(t *testing.T) {
	s := NewSet()

	if ! s.Add(1) {
		t.Fatalf("The set %v must accept 1.", s)
	}

	if s.Add(1) {
		t.Fatalf("The set %v must enforce uniqueness.", s)
	}
}

func TestSetContains(t *testing.T) {
	s := NewSet()

	s.Add(1)
	s.Add(2)

	if ! s.Contains(1) {
		t.Fatalf("The set %v must contain 1.", s)
	}

	if ! s.Contains(2) {
		t.Fatalf("The set %v must contain 2.", s)
	}

	if s.Contains(3) {
		t.Fatalf("The set %v must not contain 3.", s)
	}
}