package amap

import (
	"encoding/json"
	"github.com/asktop/gotools/async"
	"github.com/asktop/gotools/acast"
)

type StrInt64Map struct {
	mu   *async.RWMutex
	data map[string]int64
}

// NewStrInt64Map returns an empty StrInt64Map object.
// The param <unsafe> used to specify whether using map in un-concurrent-safety,
// which is false in default, means concurrent-safe.
func NewStrInt64Map(safe ...bool) *StrInt64Map {
	return &StrInt64Map{
		data: make(map[string]int64),
		mu:   async.New(safe...),
	}
}

// NewStrInt64MapFrom returns a hash map from given map <data>.
// Note that, the param <data> map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
func NewStrInt64MapFrom(data map[string]int64, safe ...bool) *StrInt64Map {
	return &StrInt64Map{
		mu:   async.New(safe...),
		data: data,
	}
}

// Sets batch sets key-values to the hash map.
func (m *StrInt64Map) Sets(data map[string]int64) *StrInt64Map {
	m.mu.Lock()
	for k, v := range data {
		m.data[k] = v
	}
	m.mu.Unlock()
	return m
}

// Set sets key-value to the hash map.
func (m *StrInt64Map) Set(key string, val int64) *StrInt64Map {
	m.mu.Lock()
	m.data[key] = val
	m.mu.Unlock()
	return m
}

// SetOrGet returns the value by key,
// or set value with given <value> if not exist and returns this value.
func (m *StrInt64Map) SetOrGet(key string, value int64) int64 {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

// Contains checks whether a key exists.
// It returns true if the <key> exists, or else false.
func (m *StrInt64Map) Contains(key string) bool {
	m.mu.RLock()
	_, exists := m.data[key]
	m.mu.RUnlock()
	return exists
}

// Search searches the map with given <key>.
// Second return parameter <found> is true if key was found, otherwise false.
func (m *StrInt64Map) Search(key string) (value int64, found bool) {
	m.mu.RLock()
	value, found = m.data[key]
	m.mu.RUnlock()
	return
}

// Get returns the value by given <key>.
func (m *StrInt64Map) Get(key string) int64 {
	m.mu.RLock()
	val, _ := m.data[key]
	m.mu.RUnlock()
	return val
}

// Iterator iterates the hash map with custom callback function <f>.
// If <f> returns true, then it continues iterating; or false to stop.
func (m *StrInt64Map) Iterator(f func(k string, v int64) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			break
		}
	}
}

// Map returns a copy of the data of the hash map.
func (m *StrInt64Map) Map() map[string]int64 {
	m.mu.RLock()
	data := make(map[string]int64, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	m.mu.RUnlock()
	return data
}

// Keys returns all keys of the map as a slice.
func (m *StrInt64Map) Keys() []string {
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
func (m *StrInt64Map) Values() []int64 {
	m.mu.RLock()
	values := make([]int64, len(m.data))
	index := 0
	for _, value := range m.data {
		values[index] = value
		index++
	}
	m.mu.RUnlock()
	return values
}

// Removes batch deletes values of the map by keys.
func (m *StrInt64Map) Remove(keys ...string) {
	m.mu.Lock()
	for _, key := range keys {
		delete(m.data, key)
	}
	m.mu.Unlock()
}

// Clear deletes all data of the map, it will remake a new underlying data map.
func (m *StrInt64Map) Clear() {
	m.mu.Lock()
	m.data = make(map[string]int64)
	m.mu.Unlock()
}

// Size returns the size of the map.
func (m *StrInt64Map) Size() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

// IsEmpty checks whether the map is empty.
// It returns true if map is empty, or else false.
func (m *StrInt64Map) IsEmpty() bool {
	return m.Size() == 0
}

// LockFunc locks writing with given callback function <f> within RWMutex.Lock.
func (m *StrInt64Map) LockFunc(f func(data map[string]int64)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f(m.data)
}

// RLockFunc locks reading with given callback function <f> within RWMutex.RLock.
func (m *StrInt64Map) RLockFunc(f func(data map[string]int64)) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	f(m.data)
}

// Clone returns a new hash map with copy of current map data.
func (m *StrInt64Map) Clone() *StrInt64Map {
	return NewStrInt64MapFrom(m.Map(), m.mu.IsSafe())
}

// Merge merges two hash maps.
// The <other> map will be merged into the map <m>.
func (m *StrInt64Map) Merge(other *StrInt64Map) {
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
func (m *StrInt64Map) Flip() {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := make(map[string]int64, len(m.data))
	for k, v := range m.data {
		n[acast.ToString(v)] = acast.ToInt64(k)
	}
	m.data = n
}

// doSetWithLockCheck checks whether value of the key exists with mutex.Lock,
// if not exists, set value to the map with given <key>,
// or else just return the existing value.
//
// It returns value with given <key>.
func (m *StrInt64Map) doSetWithLockCheck(key string, value int64) int64 {
	m.mu.Lock()
	if v, ok := m.data[key]; ok {
		m.mu.Unlock()
		return v
	}
	m.data[key] = value
	m.mu.Unlock()
	return value
}

func (m *StrInt64Map) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Map())
}

func (m *StrInt64Map) UnmarshalJSON(b []byte) error {
	data := make(map[string]int64)
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	} else {
		m.Sets(data)
		return nil
	}
}

func (m *StrInt64Map) String() string {
	rs, _ := m.MarshalJSON()
	return string(rs)
}
