package list

import (
	"testing"
)

func TestListMustBeEmpty(t *testing.T) {
	l := &List{}

	if l.head != nil {
		t.Fatalf("The list %v must be empty.", l)
	}

	l = NewList([]Key{})

	if l.head != nil {
		t.Fatalf("The list %v must be empty.", l)
	}

	if l.String() != "()" {
		t.Fatalf("The list %v must have empty signature ().", l)
	}
}

func TestListInsert(t *testing.T) {
	l := &List{}

	l.Insert(1)

	if l.head == nil {
		t.Fatalf("The list %v must not be empty.", l)
	}

	if l.head.Value != 1 || l.head.next != nil {
		t.Fatalf("The list %v tail is invalid.", l)
	}

	if l.String() != "(1 )" {
		t.Fatalf("The list %v must have good signature (1 ).", l)
	}

	l.Insert(2)

	if l.head == nil {
		t.Fatalf("The list %v must not be empty.", l)
	}

	if l.head.Value != 2 || l.head.next == nil {
		t.Fatalf("The list %v head is invalid.", l)
	}

	if l.head.next.Value != 1 || l.head.next.next != nil {
		t.Fatalf("The list %v tail is invalid.", l)
	}

	if l.String() != "(2 1 )" {
		t.Fatalf("The list %v must have good signature (2 1 ).", l)
	}
}

func TestListFind(t *testing.T) {
	l := &List{}

	if l.Find(0) != nil {
		t.Fatalf("The Find op on list %v should not panic.", l)
	}

	l.Insert(1)

	if l.Find(0) != nil {
		t.Fatalf("The Find(0) op on list %v should find nothing.", l)
	}

	if l.Find(1) == nil || l.Find(1).Value != 1 {
		t.Fatalf("The Find(1) op on list %v should find tail node.", l)
	}

	l.Insert(2)

	if l.Find(0) != nil {
		t.Fatalf("The Find(0) op on list %v should find nothing.", l)
	}

	if l.Find(1) == nil || l.Find(1).Value != 1 {
		t.Fatalf("The Find(1) op on list %v should find tail node.", l)
	}

	if l.Find(2) != l.head || l.Find(2).Value != 2 {
		t.Fatalf("The Find(2) op on list %v should find head node.", l)
	}
}

func TestListNew(t *testing.T) {
	l := NewList([]Key{})

	if l.String() != "()" {
		t.Fatalf("The list %v must be empty.", l)
	}

	l = NewList([]Key{1})

	if l.String() != "(1 )" {
		t.Fatalf("The list %v must contain 1: (1 ).", l)
	}

	l = NewList([]Key{1, 2, 3, 4, 5})

	if l.String() != "(1 2 3 4 5 )" {
		t.Fatalf("The list %v must contain full sequence in right order: (1 2 3 4 5 ).", l)
	}
}