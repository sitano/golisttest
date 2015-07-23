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

type Key int

type List struct {
	start *Node
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
func (l *List)Insert(v Key) {
	l.start = &Node{
		next: l.start,
		Value: v,
	}
}

func (l *List)Find(v Key) *Node {
	n := l.start
	for ; n != nil && n.Value != v; n = n.next {}
	return n
}
