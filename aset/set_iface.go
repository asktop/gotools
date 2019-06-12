package aset

import (
	"encoding/json"
	"github.com/asktop/gotools/async"
	"github.com/asktop/gotools/cast"
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
	for _, v := range cast.ToIfaceSlice(items) {
		m[v] = struct{}{}
	}
	return &IfaceSet{
		m:  m,
		mu: async.New(unsafe...),
	}
}

// Add adds one or multiple items to the set.
func (set *IfaceSet) Add(item ...interface{}) *IfaceSet {
	set.mu.Lock()
	for _, v := range item {
		set.m[v] = struct{}{}
	}
	set.mu.Unlock()
	return set
}

// Contains checks whether the set contains <item>.
func (set *IfaceSet) Contains(item interface{}) bool {
	set.mu.RLock()
	_, exists := set.m[item]
	set.mu.RUnlock()
	return exists
}

// Iterator iterates the set with given callback function <f>,
// if <f> returns true then continue iterating; or false to stop.
func (set *IfaceSet) Iterator(f func(v interface{}) bool) *IfaceSet {
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
func (set *IfaceSet) Slice() []interface{} {
	set.mu.RLock()
	i := 0
	ret := make([]interface{}, len(set.m))
	for item := range set.m {
		ret[i] = item
		i++
	}
	set.mu.RUnlock()
	return ret
}

// Remove deletes <item> from set.
func (set *IfaceSet) Remove(item interface{}) *IfaceSet {
	set.mu.Lock()
	delete(set.m, item)
	set.mu.Unlock()
	return set
}

// Clear deletes all items of the set.
func (set *IfaceSet) Clear() *IfaceSet {
	set.mu.Lock()
	set.m = make(map[interface{}]struct{})
	set.mu.Unlock()
	return set
}

// Size returns the size of the set.
func (set *IfaceSet) Size() int {
	set.mu.RLock()
	l := len(set.m)
	set.mu.RUnlock()
	return l
}

func (set *IfaceSet) IsEmpty() bool {
	return set.Size() == 0
}

// LockFunc locks writing with callback function <f>.
func (set *IfaceSet) LockFunc(f func(m map[interface{}]struct{})) {
	set.mu.Lock()
	defer set.mu.Unlock()
	f(set.m)
}

// RLockFunc locks reading with callback function <f>.
func (set *IfaceSet) RLockFunc(f func(m map[interface{}]struct{})) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	f(set.m)
}

// Equal checks whether the two sets equal.
func (set *IfaceSet) Equal(other *IfaceSet) bool {
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
func (set *IfaceSet) IsSubsetOf(other *IfaceSet) bool {
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

func (set *IfaceSet) Clone() *IfaceSet {
	return NewIfaceSetFrom(set.Slice(), set.mu.IsSafe())
}

// Merge adds items from <others> sets into <set>.
func (set *IfaceSet) Merge(others ...*IfaceSet) *IfaceSet {
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

// Union returns a new set which is the union of <set> and <others>.
// Which means, all the items in <newSet> are in <set> or in <others>.
func (set *IfaceSet) Union(others ...*IfaceSet) (newSet *IfaceSet) {
	newSet = NewIfaceSet(true)
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

// Diff returns a new set which is the difference set from <set> to <others>.
// Which means, all the items in <newSet> are in <set> but not in <others>.
func (set *IfaceSet) Diff(others ...*IfaceSet) (newSet *IfaceSet) {
	newSet = NewIfaceSet(true)
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

// Intersect returns a new set which is the intersection from <set> to <others>.
// Which means, all the items in <newSet> are in <set> and also in <others>.
func (set *IfaceSet) Intersect(others ...*IfaceSet) (newSet *IfaceSet) {
	newSet = NewIfaceSet(true)
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
func (set *IfaceSet) Complement(full *IfaceSet) (newSet *IfaceSet) {
	newSet = NewIfaceSet(true)
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
// Note: The items should be converted to int type,
// or you'd get a result that you unexpected.
func (set *IfaceSet) Sum() (sum int) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k, _ := range set.m {
		sum += cast.ToInt(k)
	}
	return
}

// Join joins items with a string <sep>.
func (set *IfaceSet) Join(sep string) string {
	return strings.Join(cast.ToStringSlice(set.Slice()), sep)
}

func (set *IfaceSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.Slice())
}

func (set *IfaceSet) UnmarshalJSON(b []byte) error {
	var data []interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	} else {
		set.Add(data...)
		return nil
	}
}

func (set *IfaceSet) String() string {
	rs, _ := set.MarshalJSON()
	return string(rs)
}
