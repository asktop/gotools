package aset

import (
	"encoding/json"
	"github.com/asktop/gotools/async"
	"github.com/asktop/gotools/cast"
	"strings"
)

type StringSet struct {
	mu *async.RWMutex
	m  map[string]struct{}
}

// New create and returns a new set, which contains un-repeated items.
// The param <unsafe> used to specify whether using set in un-concurrent-safety,
// which is false in default.
func NewStringSet(safe ...bool) *StringSet {
	return &StringSet{
		m:  make(map[string]struct{}),
		mu: async.New(safe...),
	}
}

// NewStringSetFrom returns a new set from <items>.
func NewStringSetFrom(items []string, safe ...bool) *StringSet {
	m := make(map[string]struct{})
	for _, v := range items {
		m[v] = struct{}{}
	}
	return &StringSet{
		m:  m,
		mu: async.New(safe...),
	}
}

// Add adds one or multiple items to the set.
func (s *StringSet) Add(item ...string) *StringSet {
	s.mu.Lock()
	for _, v := range item {
		s.m[v] = struct{}{}
	}
	s.mu.Unlock()
	return s
}

// Contains checks whether the set contains <item>.
func (s *StringSet) Contains(item string) bool {
	s.mu.RLock()
	_, exists := s.m[item]
	s.mu.RUnlock()
	return exists
}

// Iterator iterates the set with given callback function <f>,
// if <f> returns true then continue iterating; or false to stop.
func (s *StringSet) Iterator(f func(v string) bool) *StringSet {
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
func (s *StringSet) Slice() []string {
	s.mu.RLock()
	ret := make([]string, len(s.m))
	i := 0
	for item := range s.m {
		ret[i] = item
		i++
	}
	s.mu.RUnlock()
	return ret
}

// Remove deletes <item> from set.
func (s *StringSet) Remove(item ...string) *StringSet {
	s.mu.Lock()
	for _, v := range item {
		delete(s.m, v)
	}
	s.mu.Unlock()
	return s
}

// Clear deletes all items of the set.
func (s *StringSet) Clear() *StringSet {
	s.mu.Lock()
	s.m = make(map[string]struct{})
	s.mu.Unlock()
	return s
}

// Size returns the size of the set.
func (s *StringSet) Size() int {
	s.mu.RLock()
	l := len(s.m)
	s.mu.RUnlock()
	return l
}

func (s *StringSet) IsEmpty() bool {
	return s.Size() == 0
}

// LockFunc locks writing with callback function <f>.
func (s *StringSet) LockFunc(f func(m map[string]struct{})) {
	s.mu.Lock()
	defer s.mu.Unlock()
	f(s.m)
}

// RLockFunc locks reading with callback function <f>.
func (s *StringSet) RLockFunc(f func(m map[string]struct{})) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	f(s.m)
}

// Equal checks whether the two sets equal.
func (s *StringSet) Equal(other *StringSet) bool {
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
func (s *StringSet) IsSubsetOf(other *StringSet) bool {
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

func (s *StringSet) Clone() *StringSet {
	return NewStringSetFrom(s.Slice(), s.mu.IsSafe())
}

// Merge adds items from <others> sets into <set>.
func (s *StringSet) Merge(others ...*StringSet) *StringSet {
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
func (s *StringSet) Union(others ...*StringSet) (newSet *StringSet) {
	newSet = NewStringSet(true)
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
func (s *StringSet) Diff(others ...*StringSet) (newSet *StringSet) {
	newSet = NewStringSet(true)
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
func (s *StringSet) Intersect(others ...*StringSet) (newSet *StringSet) {
	newSet = NewStringSet(true)
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
func (s *StringSet) Complement(full *StringSet) (newSet *StringSet) {
	newSet = NewStringSet(true)
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
func (s *StringSet) Sum() (sum int) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k, _ := range s.m {
		sum += cast.ToInt(k)
	}
	return
}

// Join joins items with a string <sep>.
func (s *StringSet) Join(sep string) string {
	return strings.Join(s.Slice(), sep)
}

func (s *StringSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Slice())
}

func (s *StringSet) UnmarshalJSON(b []byte) error {
	var data []string
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	} else {
		s.Add(data...)
		return nil
	}
}

func (s *StringSet) String() string {
	rs, _ := s.MarshalJSON()
	return string(rs)
}
