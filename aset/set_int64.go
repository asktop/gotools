package aset

import (
	"encoding/json"
	"github.com/asktop/gotools/async"
	"github.com/asktop/gotools/acast"
	"strings"
)

type Int64Set struct {
	mu *async.RWMutex
	m  map[int64]struct{}
}

// New create and returns a new set, which contains un-repeated items.
// The param <unsafe> used to specify whether using set in un-concurrent-safety,
// which is false in default.
func NewInt64Set(safe ...bool) *Int64Set {
	return &Int64Set{
		m:  make(map[int64]struct{}),
		mu: async.New(safe...),
	}
}

// NewInt64SetFrom returns a new set from <items>.
func NewInt64SetFrom(items []int64, safe ...bool) *Int64Set {
	m := make(map[int64]struct{})
	for _, v := range items {
		m[v] = struct{}{}
	}
	return &Int64Set{
		m:  m,
		mu: async.New(safe...),
	}
}

// Add adds one or multiple items to the set.
func (s *Int64Set) Add(item ...int64) *Int64Set {
	s.mu.Lock()
	for _, v := range item {
		s.m[v] = struct{}{}
	}
	s.mu.Unlock()
	return s
}

// Contains checks whether the set contains <item>.
func (s *Int64Set) Contains(item int64) bool {
	s.mu.RLock()
	_, exists := s.m[item]
	s.mu.RUnlock()
	return exists
}

// Iterator iterates the set with given callback function <f>,
// if <f> returns true then continue iterating; or false to stop.
func (s *Int64Set) Iterator(f func(v int64) bool) *Int64Set {
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
func (s *Int64Set) Slice() []int64 {
	s.mu.RLock()
	ret := make([]int64, len(s.m))
	i := 0
	for k, _ := range s.m {
		ret[i] = k
		i++
	}
	s.mu.RUnlock()
	return ret
}

// Remove deletes <item> from set.
func (s *Int64Set) Remove(item int64) *Int64Set {
	s.mu.Lock()
	delete(s.m, item)
	s.mu.Unlock()
	return s
}

// Clear deletes all items of the set.
func (s *Int64Set) Clear() *Int64Set {
	s.mu.Lock()
	s.m = make(map[int64]struct{})
	s.mu.Unlock()
	return s
}

// Size returns the size of the set.
func (s *Int64Set) Size() int {
	s.mu.RLock()
	l := len(s.m)
	s.mu.RUnlock()
	return l
}

func (s *Int64Set) IsEmpty() bool {
	return s.Size() == 0
}

// LockFunc locks writing with callback function <f>.
func (s *Int64Set) LockFunc(f func(m map[int64]struct{})) {
	s.mu.Lock()
	defer s.mu.Unlock()
	f(s.m)
}

// RLockFunc locks reading with callback function <f>.
func (s *Int64Set) RLockFunc(f func(m map[int64]struct{})) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	f(s.m)
}

// Equal checks whether the two sets equal.
func (s *Int64Set) Equal(other *Int64Set) bool {
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
func (s *Int64Set) IsSubsetOf(other *Int64Set) bool {
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

func (s *Int64Set) Clone() *Int64Set {
	return NewInt64SetFrom(s.Slice(), s.mu.IsSafe())
}

// Merge adds items from <others> sets into <set>.
func (s *Int64Set) Merge(others ...*Int64Set) *Int64Set {
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

// Union returns a new set which is the union of <set> and <other>.
// Which means, all the items in <newSet> are in <set> or in <other>.
func (s *Int64Set) Union(others ...*Int64Set) (newSet *Int64Set) {
	newSet = NewInt64Set(true)
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

// Diff returns a new set which is the difference set from <set> to <other>.
// Which means, all the items in <newSet> are in <set> but not in <other>.
func (s *Int64Set) Diff(others ...*Int64Set) (newSet *Int64Set) {
	newSet = NewInt64Set(true)
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

// Intersect returns a new set which is the intersection from <set> to <other>.
// Which means, all the items in <newSet> are in <set> and also in <other>.
func (s *Int64Set) Intersect(others ...*Int64Set) (newSet *Int64Set) {
	newSet = NewInt64Set(true)
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
func (s *Int64Set) Complement(full *Int64Set) (newSet *Int64Set) {
	newSet = NewInt64Set(true)
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
// Note: The items should be converted to int64 type,
// or you'd get a result that you unexpected.
func (s *Int64Set) Sum() (sum int64) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k, _ := range s.m {
		sum += k
	}
	return
}

// Join joins items with a string <sep>.
func (s *Int64Set) Join(sep string) string {
	return strings.Join(acast.ToStringSlice(s.Slice()), sep)
}

func (s *Int64Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Slice())
}

func (s *Int64Set) UnmarshalJSON(b []byte) error {
	var data []int64
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	} else {
		s.Add(data...)
		return nil
	}
}

func (s *Int64Set) String() string {
	rs, _ := s.MarshalJSON()
	return string(rs)
}
