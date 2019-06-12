package amap

import (
	"encoding/json"
	"github.com/asktop/gotools/async"
	"github.com/asktop/gotools/cast"
)

type StrIfaceMap struct {
	mu   *async.RWMutex
	data map[string]interface{}
}

// NewStrIfaceMap returns an empty StrIfaceMap object.
// The param <unsafe> used to specify whether using map in un-concurrent-safety,
// which is false in default, means concurrent-safe.
func NewStrIfaceMap(safe ...bool) *StrIfaceMap {
	return &StrIfaceMap{
		mu:   async.New(safe...),
		data: make(map[string]interface{}),
	}
}

// NewStrIfaceMapFrom returns a hash map from given map <data>.
// Note that, the param <data> map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
func NewStrIfaceMapFrom(data map[string]interface{}, safe ...bool) *StrIfaceMap {
	return &StrIfaceMap{
		mu:   async.New(safe...),
		data: data,
	}
}

// Sets batch sets key-values to the hash map.
func (m *StrIfaceMap) Sets(data map[string]interface{}) *StrIfaceMap {
	m.mu.Lock()
	for k, v := range data {
		m.data[k] = v
	}
	m.mu.Unlock()
	return m
}

// Set sets key-value to the hash map.
func (m *StrIfaceMap) Set(key string, val interface{}) *StrIfaceMap {
	m.mu.Lock()
	m.data[key] = val
	m.mu.Unlock()
	return m
}

// SetOrGet returns the value by key,
// or set value with given <value> if not exist and returns this value.
func (m *StrIfaceMap) SetOrGet(key string, value interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

// Contains checks whether a key exists.
// It returns true if the <key> exists, or else false.
func (m *StrIfaceMap) Contains(key string) bool {
	m.mu.RLock()
	_, exists := m.data[key]
	m.mu.RUnlock()
	return exists
}

// Search searches the map with given <key>.
// Second return parameter <found> is true if key was found, otherwise false.
func (m *StrIfaceMap) Search(key string) (value interface{}, found bool) {
	m.mu.RLock()
	value, found = m.data[key]
	m.mu.RUnlock()
	return
}

// Get returns the value by given <key>.
func (m *StrIfaceMap) Get(key string) interface{} {
	m.mu.RLock()
	val, _ := m.data[key]
	m.mu.RUnlock()
	return val
}

// Iterator iterates the hash map with custom callback function <f>.
// If <f> returns true, then it continues iterating; or false to stop.
func (m *StrIfaceMap) Iterator(f func(k string, v interface{}) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			break
		}
	}
}

// Map returns a copy of the data of the hash map.
func (m *StrIfaceMap) Map() map[string]interface{} {
	m.mu.RLock()
	data := make(map[string]interface{}, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	m.mu.RUnlock()
	return data
}

// Keys returns all keys of the map as a slice.
func (m *StrIfaceMap) Keys() []string {
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
func (m *StrIfaceMap) Values() []interface{} {
	m.mu.RLock()
	values := make([]interface{}, len(m.data))
	index := 0
	for _, value := range m.data {
		values[index] = value
		index++
	}
	m.mu.RUnlock()
	return values
}

// Removes batch deletes values of the map by keys.
func (m *StrIfaceMap) Remove(keys ...string) {
	m.mu.Lock()
	for _, key := range keys {
		delete(m.data, key)
	}
	m.mu.Unlock()
}

// Clear deletes all data of the map, it will remake a new underlying data map.
func (m *StrIfaceMap) Clear() {
	m.mu.Lock()
	m.data = make(map[string]interface{})
	m.mu.Unlock()
}

// IsEmpty checks whether the map is empty.
// It returns true if map is empty, or else false.
func (m *StrIfaceMap) IsEmpty() bool {
	m.mu.RLock()
	empty := len(m.data) == 0
	m.mu.RUnlock()
	return empty
}

// Size returns the size of the map.
func (m *StrIfaceMap) Size() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

// Clone returns a new hash map with copy of current map data.
func (m *StrIfaceMap) Clone() *StrIfaceMap {
	return NewStrIfaceMapFrom(m.Map(), m.mu.IsSafe())
}

// Merge merges two hash maps.
// The <other> map will be merged into the map <m>.
func (m *StrIfaceMap) Merge(other *StrIfaceMap) {
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
func (m *StrIfaceMap) Flip() {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := make(map[string]interface{}, len(m.data))
	for k, v := range m.data {
		n[cast.ToString(v)] = k
	}
	m.data = n
}

// LockFunc locks writing with given callback function <f> within RWMutex.Lock.
func (m *StrIfaceMap) LockFunc(f func(m map[string]interface{})) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f(m.data)
}

// RLockFunc locks reading with given callback function <f> within RWMutex.RLock.
func (m *StrIfaceMap) RLockFunc(f func(m map[string]interface{})) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	f(m.data)
}

// doSetWithLockCheck checks whether value of the key exists with mutex.Lock,
// if not exists, set value to the map with given <key>,
// or else just return the existing value.
//
// When setting value, if <value> is type of <func() interface {}>,
// it will be executed with mutex.Lock of the hash map,
// and its return value will be set to the map with <key>.
//
// It returns value with given <key>.
func (m *StrIfaceMap) doSetWithLockCheck(key string, value interface{}) interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	if v, ok := m.data[key]; ok {
		return v
	}
	if f, ok := value.(func() interface{}); ok {
		value = f()
	}
	if value != nil {
		m.data[key] = value
	}
	return value
}

func (m *StrIfaceMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Map())
}

func (m *StrIfaceMap) UnmarshalJSON(b []byte) error {
	data := make(map[string]interface{})
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	} else {
		m.Sets(data)
		return nil
	}
}
