// Package implements naive single linked list for Node[Int] type.
//
// It implements operations on native type instead of introducing
// something like:
//
// type Key interface {
//   Hash() int64
//   Equals() bool
//   String()
// }
//
// it left for the sake of future generations.
package list

import "strconv"

type Key int

type List struct {
	head *Node
}

type Node struct {
	next *Node
	Value Key
}

func NewList(a []Key) *List {
	l := &List{}
	for i := len(a) - 1; i >= 0; i -- {
		l.Insert(a[i])
	}
	return l
}

// Inserts O(1) a Key v into the beginning of the list.
func (l *List) Insert(v Key) {
	l.head = &Node{
		next: l.head,
		Value: v,
	}
}

// Lookup for first occurrence of Key v
func (l *List) Find(v Key) *Node {
	n := l.head
	for ; n != nil && n.Value != v; n = n.next {}
	return n
}

func (n *Node) String() string {
	return strconv.Itoa(int(n.Value))
}

func (l *List) String() string {
	s := "("
	for n := l.head; n != nil; n = n.next {
		s += n.String() + " "
	}
	s += ")"
	return s
}