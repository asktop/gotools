package aset

import (
	"encoding/json"
	"github.com/asktop/gotools/async"
	"github.com/asktop/gotools/acast"
	"strings"
)

type IfaceSet struct {
	mu *async.RWMutex
	m  map[interface{}]struct{}
}

// New create and returns a new set, which contains un-repeated items.
// The param <unsafe> used to specify whether using set in un-concurrent-safety,
// which is false in default.
func NewIfaceSet(safe ...bool) *IfaceSet {
	return &IfaceSet{
		m:  make(map[interface{}]struct{}),
		mu: async.New(safe...),
	}
}

// NewIfaceSetFrom returns a new set from <items>.
// Parameter <items> can be either a variable of any type, or a slice.
func NewIfaceSetFrom(items interface{}, unsafe ...bool) *IfaceSet {
	m := make(map[interface{}]struct{})
	for _, v := range acast.ToIfaceSlice(items) {
		m[v] = struct{}{}
	}
	return &IfaceSet{
		m:  m,
		mu: async.New(unsafe...),
	}
}

// Add adds one or multiple items to the set.
func (s *IfaceSet) Add(item ...interface{}) *IfaceSet {
	s.mu.Lock()
	for _, v := range item {
		s.m[v] = struct{}{}
	}
	s.mu.Unlock()
	return s
}

// Contains checks whether the set contains <item>.
func (s *IfaceSet) Contains(item interface{}) bool {
	s.mu.RLock()
	_, exists := s.m[item]
	s.mu.RUnlock()
	return exists
}

// Iterator iterates the set with given callback function <f>,
// if <f> returns true then continue iterating; or false to stop.
func (s *IfaceSet) Iterator(f func(v interface{}) bool) *IfaceSet {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k, _ := range s.m {
		if !f(k) {
			break
		}
	}
	return s
}

// Slice returns the a of items of the set as slice.
func (s *IfaceSet) Slice() []interface{} {
	s.mu.RLock()
	i := 0
	ret := make([]interface{}, len(s.m))
	for item := range s.m {
		ret[i] = item
		i++
	}
	s.mu.RUnlock()
	return ret
}

// Remove deletes <item> from set.
func (s *IfaceSet) Remove(item interface{}) *IfaceSet {
	s.mu.Lock()
	delete(s.m, item)
	s.mu.Unlock()
	return s
}

// Clear deletes all items of the set.
func (s *IfaceSet) Clear() *IfaceSet {
	s.mu.Lock()
	s.m = make(map[interface{}]struct{})
	s.mu.Unlock()
	return s
}

// Size returns the size of the set.
func (s *IfaceSet) Size() int {
	s.mu.RLock()
	l := len(s.m)
	s.mu.RUnlock()
	return l
}

func (s *IfaceSet) IsEmpty() bool {
	return s.Size() == 0
}

// LockFunc locks writing with callback function <f>.
func (s *IfaceSet) LockFunc(f func(m map[interface{}]struct{})) {
	s.mu.Lock()
	defer s.mu.Unlock()
	f(s.m)
}

// RLockFunc locks reading with callback function <f>.
func (s *IfaceSet) RLockFunc(f func(m map[interface{}]struct{})) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	f(s.m)
}

// Equal checks whether the two sets equal.
func (s *IfaceSet) Equal(other *IfaceSet) bool {
	if s == other {
		return true
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	if len(s.m) != len(other.m) {
		return false
	}
	for key := range s.m {
		if _, ok := other.m[key]; !ok {
			return false
		}
	}
	return true
}

// IsSubsetOf checks whether the current set is a sub-set of <other>.
func (s *IfaceSet) IsSubsetOf(other *IfaceSet) bool {
	if s == other {
		return true
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	for key := range s.m {
		if _, ok := other.m[key]; !ok {
			return false
		}
	}
	return true
}

func (s *IfaceSet) Clone() *IfaceSet {
	return NewIfaceSetFrom(s.Slice(), s.mu.IsSafe())
}

// Merge adds items from <others> sets into <set>.
func (s *IfaceSet) Merge(others ...*IfaceSet) *IfaceSet {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, other := range others {
		if s != other {
			other.mu.RLock()
		}
		for k, v := range other.m {
			s.m[k] = v
		}
		if s != other {
			other.mu.RUnlock()
		}
	}
	return s
}

// Union returns a new set which is the union of <set> and <others>.
// Which means, all the items in <newSet> are in <set> or in <others>.
func (s *IfaceSet) Union(others ...*IfaceSet) (newSet *IfaceSet) {
	newSet = NewIfaceSet(true)
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, other := range others {
		if s != other {
			other.mu.RLock()
		}
		for k, v := range s.m {
			newSet.m[k] = v
		}
		if s != other {
			for k, v := range other.m {
				newSet.m[k] = v
			}
		}
		if s != other {
			other.mu.RUnlock()
		}
	}

	return
}

// Diff returns a new set which is the difference set from <set> to <others>.
// Which means, all the items in <newSet> are in <set> but not in <others>.
func (s *IfaceSet) Diff(others ...*IfaceSet) (newSet *IfaceSet) {
	newSet = NewIfaceSet(true)
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, other := range others {
		if s == other {
			continue
		}
		other.mu.RLock()
		for k, v := range s.m {
			if _, ok := other.m[k]; !ok {
				newSet.m[k] = v
			}
		}
		other.mu.RUnlock()
	}
	return
}

// Intersect returns a new set which is the intersection from <set> to <others>.
// Which means, all the items in <newSet> are in <set> and also in <others>.
func (s *IfaceSet) Intersect(others ...*IfaceSet) (newSet *IfaceSet) {
	newSet = NewIfaceSet(true)
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, other := range others {
		if s != other {
			other.mu.RLock()
		}
		for k, v := range s.m {
			if _, ok := other.m[k]; ok {
				newSet.m[k] = v
			}
		}
		if s != other {
			other.mu.RUnlock()
		}
	}
	return
}

// Complement returns a new set which is the complement from <set> to <full>.
// Which means, all the items in <newSet> are in <full> and not in <set>.
//
// It returns the difference between <full> and <set>
// if the given set <full> is not the full set of <set>.
func (s *IfaceSet) Complement(full *IfaceSet) (newSet *IfaceSet) {
	newSet = NewIfaceSet(true)
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s != full {
		full.mu.RLock()
		defer full.mu.RUnlock()
	}
	for k, v := range full.m {
		if _, ok := s.m[k]; !ok {
			newSet.m[k] = v
		}
	}
	return
}

// Sum sums items.
// Note: The items should be converted to int type,
// or you'd get a result that you unexpected.
func (s *IfaceSet) Sum() (sum int) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k, _ := range s.m {
		sum += acast.ToInt(k)
	}
	return
}

// Join joins items with a string <sep>.
func (s *IfaceSet) Join(sep string) string {
	return strings.Join(acast.ToStringSlice(s.Slice()), sep)
}

func (s *IfaceSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Slice())
}

func (s *IfaceSet) UnmarshalJSON(b []byte) error {
	var data []interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	} else {
		s.Add(data...)
		return nil
	}
}

func (s *IfaceSet) String() string {
	rs, _ := s.MarshalJSON()
	return string(rs)
}
