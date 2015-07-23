// Naive set implementation based on single linked list and knowing internal
// structure of the Key type.
package set

import "github.com/sitano/golisttest/list"

// The Set based on single linked list impl.
// It could be even simplier:
//   type Set *list.List
// but we go with full impl.
type Set struct {
   storage *list.List
}

func NewSet() *Set {
	return &Set {
		storage: &list.List{},
	}
}

// Add an element into the Set O(n) - contains is the longest call
func (s *Set) Add(v list.Key) bool {
	if !s.Contains(v) {
		s.storage.Insert(v)
		return true
	}
	return false
}

// Check if an element contains inside the Set O(n), Say hi to HyperLogLog
func (s *Set) Contains(v list.Key) bool {
	return s.storage.Find(v) != nil
}

func (s *Set) String() string {
	return s.storage.String()
}
