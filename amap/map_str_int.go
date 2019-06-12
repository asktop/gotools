package amap

import (
	"encoding/json"
	"github.com/asktop/gotools/async"
	"github.com/asktop/gotools/acast"
)

type StrIntMap struct {
	mu   *async.RWMutex
	data map[string]int
}

// NewStrIntMap returns an empty StrIntMap object.
// The param <unsafe> used to specify whether using map in un-concurrent-safety,
// which is false in default, means concurrent-safe.
func NewStrIntMap(safe ...bool) *StrIntMap {
	return &StrIntMap{
		data: make(map[string]int),
		mu:   async.New(safe...),
	}
}

// NewStrIntMapFrom returns a hash map from given map <data>.
// Note that, the param <data> map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
func NewStrIntMapFrom(data map[string]int, safe ...bool) *StrIntMap {
	return &StrIntMap{
		mu:   async.New(safe...),
		data: data,
	}
}

// Sets batch sets key-values to the hash map.
func (m *StrIntMap) Sets(data map[string]int) *StrIntMap {
	m.mu.Lock()
	for k, v := range data {
		m.data[k] = v
	}
	m.mu.Unlock()
	return m
}

// Set sets key-value to the hash map.
func (m *StrIntMap) Set(key string, val int) *StrIntMap {
	m.mu.Lock()
	m.data[key] = val
	m.mu.Unlock()
	return m
}

// SetOrGet returns the value by key,
// or set value with given <value> if not exist and returns this value.
func (m *StrIntMap) SetOrGet(key string, value int) int {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

// Contains checks whether a key exists.
// It returns true if the <key> exists, or else false.
func (m *StrIntMap) Contains(key string) bool {
	m.mu.RLock()
	_, exists := m.data[key]
	m.mu.RUnlock()
	return exists
}

// Search searches the map with given <key>.
// Second return parameter <found> is true if key was found, otherwise false.
func (m *StrIntMap) Search(key string) (value int, found bool) {
	m.mu.RLock()
	value, found = m.data[key]
	m.mu.RUnlock()
	return
}

// Get returns the value by given <key>.
func (m *StrIntMap) Get(key string) int {
	m.mu.RLock()
	val, _ := m.data[key]
	m.mu.RUnlock()
	return val
}

// Iterator iterates the hash map with custom callback function <f>.
// If <f> returns true, then it continues iterating; or false to stop.
func (m *StrIntMap) Iterator(f func(k string, v int) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			break
		}
	}
}

// Map returns a copy of the data of the hash map.
func (m *StrIntMap) Map() map[string]int {
	m.mu.RLock()
	data := make(map[string]int, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	m.mu.RUnlock()
	return data
}

// Keys returns all keys of the map as a slice.
func (m *StrIntMap) Keys() []string {
	m.mu.RLock()
	keys := make([]string, len(m.data))
	index := 0
	for key := range m.data {
		keys[index] = key
		index++
	}
	m.mu.RUnlock()
	return keys
}

// Values returns all values of the map as a slice.
func (m *StrIntMap) Values() []int {
	m.mu.RLock()
	values := make([]int, len(m.data))
	index := 0
	for _, value := range m.data {
		values[index] = value
		index++
	}
	m.mu.RUnlock()
	return values
}

// Removes batch deletes values of the map by keys.
func (m *StrIntMap) Remove(keys ...string) {
	m.mu.Lock()
	for _, key := range keys {
		delete(m.data, key)
	}
	m.mu.Unlock()
}

// Clear deletes all data of the map, it will remake a new underlying data map.
func (m *StrIntMap) Clear() {
	m.mu.Lock()
	m.data = make(map[string]int)
	m.mu.Unlock()
}

// Size returns the size of the map.
func (m *StrIntMap) Size() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

// IsEmpty checks whether the map is empty.
// It returns true if map is empty, or else false.
func (m *StrIntMap) IsEmpty() bool {
	return m.Size() == 0
}

// LockFunc locks writing with given callback function <f> within RWMutex.Lock.
func (m *StrIntMap) LockFunc(f func(data map[string]int)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f(m.data)
}

// RLockFunc locks reading with given callback function <f> within RWMutex.RLock.
func (m *StrIntMap) RLockFunc(f func(data map[string]int)) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	f(m.data)
}

// Clone returns a new hash map with copy of current map data.
func (m *StrIntMap) Clone() *StrIntMap {
	return NewStrIntMapFrom(m.Map(), m.mu.IsSafe())
}

// Merge merges two hash maps.
// The <other> map will be merged into the map <m>.
func (m *StrIntMap) Merge(other *StrIntMap) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if other != m {
		other.mu.RLock()
		defer other.mu.RUnlock()
	}
	for k, v := range other.data {
		m.data[k] = v
	}
}

// Flip exchanges key-value of the map to value-key.
func (m *StrIntMap) Flip() {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := make(map[string]int, len(m.data))
	for k, v := range m.data {
		n[acast.ToString(v)] = acast.ToInt(k)
	}
	m.data = n
}

// doSetWithLockCheck checks whether value of the key exists with mutex.Lock,
// if not exists, set value to the map with given <key>,
// or else just return the existing value.
//
// It returns value with given <key>.
func (m *StrIntMap) doSetWithLockCheck(key string, value int) int {
	m.mu.Lock()
	if v, ok := m.data[key]; ok {
		m.mu.Unlock()
		return v
	}
	m.data[key] = value
	m.mu.Unlock()
	return value
}

func (m *StrIntMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Map())
}

func (m *StrIntMap) UnmarshalJSON(b []byte) error {
	data := make(map[string]int)
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	} else {
		m.Sets(data)
		return nil
	}
}

func (m *StrIntMap) String() string {
	rs, _ := m.MarshalJSON()
	return string(rs)
}
