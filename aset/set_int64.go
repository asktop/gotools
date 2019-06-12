package aset

import (
	"encoding/json"
	"github.com/asktop/gotools/async"
	"github.com/asktop/gotools/cast"
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
func (set *Int64Set) Add(item ...int64) *Int64Set {
	set.mu.Lock()
	for _, v := range item {
		set.m[v] = struct{}{}
	}
	set.mu.Unlock()
	return set
}

// Contains checks whether the set contains <item>.
func (set *Int64Set) Contains(item int64) bool {
	set.mu.RLock()
	_, exists := set.m[item]
	set.mu.RUnlock()
	return exists
}

// Iterator iterates the set with given callback function <f>,
// if <f> returns true then continue iterating; or false to stop.
func (set *Int64Set) Iterator(f func(v int64) bool) *Int64Set {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k, _ := range set.m {
		if !f(k) {
			break
		}
	}
	return set
}

// Slice returns the a of items of the set as slice.
func (set *Int64Set) Slice() []int64 {
	set.mu.RLock()
	ret := make([]int64, len(set.m))
	i := 0
	for k, _ := range set.m {
		ret[i] = k
		i++
	}
	set.mu.RUnlock()
	return ret
}

// Remove deletes <item> from set.
func (set *Int64Set) Remove(item int64) *Int64Set {
	set.mu.Lock()
	delete(set.m, item)
	set.mu.Unlock()
	return set
}

// Clear deletes all items of the set.
func (set *Int64Set) Clear() *Int64Set {
	set.mu.Lock()
	set.m = make(map[int64]struct{})
	set.mu.Unlock()
	return set
}

// Size returns the size of the set.
func (set *Int64Set) Size() int {
	set.mu.RLock()
	l := len(set.m)
	set.mu.RUnlock()
	return l
}

func (set *Int64Set) IsEmpty() bool {
	return set.Size() == 0
}

// LockFunc locks writing with callback function <f>.
func (set *Int64Set) LockFunc(f func(m map[int64]struct{})) {
	set.mu.Lock()
	defer set.mu.Unlock()
	f(set.m)
}

// RLockFunc locks reading with callback function <f>.
func (set *Int64Set) RLockFunc(f func(m map[int64]struct{})) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	f(set.m)
}

// Equal checks whether the two sets equal.
func (set *Int64Set) Equal(other *Int64Set) bool {
	if set == other {
		return true
	}
	set.mu.RLock()
	defer set.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	if len(set.m) != len(other.m) {
		return false
	}
	for key := range set.m {
		if _, ok := other.m[key]; !ok {
			return false
		}
	}
	return true
}

// IsSubsetOf checks whether the current set is a sub-set of <other>.
func (set *Int64Set) IsSubsetOf(other *Int64Set) bool {
	if set == other {
		return true
	}
	set.mu.RLock()
	defer set.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	for key := range set.m {
		if _, ok := other.m[key]; !ok {
			return false
		}
	}
	return true
}

func (set *Int64Set) Clone() *Int64Set {
	return NewInt64SetFrom(set.Slice(), set.mu.IsSafe())
}

// Merge adds items from <others> sets into <set>.
func (set *Int64Set) Merge(others ...*Int64Set) *Int64Set {
	set.mu.Lock()
	defer set.mu.Unlock()
	for _, other := range others {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range other.m {
			set.m[k] = v
		}
		if set != other {
			other.mu.RUnlock()
		}
	}
	return set
}

// Union returns a new set which is the union of <set> and <other>.
// Which means, all the items in <newSet> are in <set> or in <other>.
func (set *Int64Set) Union(others ...*Int64Set) (newSet *Int64Set) {
	newSet = NewInt64Set(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range others {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range set.m {
			newSet.m[k] = v
		}
		if set != other {
			for k, v := range other.m {
				newSet.m[k] = v
			}
		}
		if set != other {
			other.mu.RUnlock()
		}
	}

	return
}

// Diff returns a new set which is the difference set from <set> to <other>.
// Which means, all the items in <newSet> are in <set> but not in <other>.
func (set *Int64Set) Diff(others ...*Int64Set) (newSet *Int64Set) {
	newSet = NewInt64Set(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range others {
		if set == other {
			continue
		}
		other.mu.RLock()
		for k, v := range set.m {
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
func (set *Int64Set) Intersect(others ...*Int64Set) (newSet *Int64Set) {
	newSet = NewInt64Set(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range others {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range set.m {
			if _, ok := other.m[k]; ok {
				newSet.m[k] = v
			}
		}
		if set != other {
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
func (set *Int64Set) Complement(full *Int64Set) (newSet *Int64Set) {
	newSet = NewInt64Set(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	if set != full {
		full.mu.RLock()
		defer full.mu.RUnlock()
	}
	for k, v := range full.m {
		if _, ok := set.m[k]; !ok {
			newSet.m[k] = v
		}
	}
	return
}

// Sum sums items.
// Note: The items should be converted to int64 type,
// or you'd get a result that you unexpected.
func (set *Int64Set) Sum() (sum int64) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k, _ := range set.m {
		sum += k
	}
	return
}

// Join joins items with a string <sep>.
func (set *Int64Set) Join(sep string) string {
	return strings.Join(cast.ToStringSlice(set.Slice()), sep)
}

func (set *Int64Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.Slice())
}

func (set *Int64Set) UnmarshalJSON(b []byte) error {
	var data []int64
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	} else {
		set.Add(data...)
		return nil
	}
}

func (set *Int64Set) String() string {
	rs, _ := set.MarshalJSON()
	return string(rs)
}
