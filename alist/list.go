package alist

import (
	"container/list"
	"encoding/json"
	"github.com/asktop/gotools/async"
	"github.com/asktop/gotools/cast"
	"strings"
)

type (
	List struct {
		mu   *async.RWMutex
		list *list.List
	}

	Element = list.Element
)

// New creates and returns a new empty doubly linked list.
func New(safe ...bool) *List {
	return &List{
		mu:   async.New(safe...),
		list: list.New(),
	}
}

func NewFrom(list *list.List, safe ...bool) *List {
	return &List{
		mu:   async.New(safe...),
		list: list,
	}
}

// PushFronts inserts multiple new elements with values <values> at the front of list <l>.
func (l *List) PushFront(values ...interface{}) {
	l.mu.Lock()
	for _, v := range values {
		l.list.PushFront(v)
	}
	l.mu.Unlock()
}

// PushBacks inserts multiple new elements with values <values> at the back of list <l>.
func (l *List) PushBack(values ...interface{}) {
	l.mu.Lock()
	for _, v := range values {
		l.list.PushBack(v)
	}
	l.mu.Unlock()
}

// PushFrontList inserts a copy of an other list at the front of list <l>.
// The lists <l> and <other> may be the same, but they must not be nil.
func (l *List) PushFrontList(other *List) {
	if l != other {
		other.mu.RLock()
		defer other.mu.RUnlock()
	}
	l.mu.Lock()
	l.list.PushFrontList(other.list)
	l.mu.Unlock()
}

// PushBackList inserts a copy of an other list at the back of list <l>.
// The lists <l> and <other> may be the same, but they must not be nil.
func (l *List) PushBackList(other *List) {
	if l != other {
		other.mu.RLock()
		defer other.mu.RUnlock()
	}
	l.mu.Lock()
	l.list.PushBackList(other.list)
	l.mu.Unlock()
}

// InsertBefore inserts a new element <e> with value <v> immediately before <p> and returns <e>.
// If <p> is not an element of <l>, the list is not modified.
// The <p> must not be nil.
func (l *List) InsertBefore(v interface{}, p *Element) (e *Element) {
	l.mu.Lock()
	e = l.list.InsertBefore(v, p)
	l.mu.Unlock()
	return
}

// InsertAfter inserts a new element <e> with value <v> immediately after <p> and returns <e>.
// If <p> is not an element of <l>, the list is not modified.
// The <p> must not be nil.
func (l *List) InsertAfter(v interface{}, p *Element) (e *Element) {
	l.mu.Lock()
	e = l.list.InsertAfter(v, p)
	l.mu.Unlock()
	return
}

// MoveBefore moves element <e> to its new position before <p>.
// If <e> or <p> is not an element of <l>, or <e> == <p>, the list is not modified.
// The element and <p> must not be nil.
func (l *List) MoveBefore(e, p *Element) {
	l.mu.Lock()
	l.list.MoveBefore(e, p)
	l.mu.Unlock()
}

// MoveAfter moves element <e> to its new position after <p>.
// If <e> or <p> is not an element of <l>, or <e> == <p>, the list is not modified.
// The element and <p> must not be nil.
func (l *List) MoveAfter(e, p *Element) {
	l.mu.Lock()
	l.list.MoveAfter(e, p)
	l.mu.Unlock()
}

// MoveToFront moves element <e> to the front of list <l>.
// If <e> is not an element of <l>, the list is not modified.
// The element must not be nil.
func (l *List) MoveToFront(e *Element) {
	l.mu.Lock()
	l.list.MoveToFront(e)
	l.mu.Unlock()
}

// MoveToBack moves element <e> to the back of list <l>.
// If <e> is not an element of <l>, the list is not modified.
// The element must not be nil.
func (l *List) MoveToBack(e *Element) {
	l.mu.Lock()
	l.list.MoveToBack(e)
	l.mu.Unlock()
}

// IteratorAsc iterates the list in ascending order with given callback function <f>.
// If <f> returns true, then it continues iterating; or false to stop.
func (l *List) IteratorAsc(f func(e *Element) bool) {
	l.mu.RLock()
	length := l.list.Len()
	if length > 0 {
		for i, e := 0, l.list.Front(); i < length; i, e = i+1, e.Next() {
			if !f(e) {
				break
			}
		}
	}
	l.mu.RUnlock()
}

// IteratorDesc iterates the list in descending order with given callback function <f>.
// If <f> returns true, then it continues iterating; or false to stop.
func (l *List) IteratorDesc(f func(e *Element) bool) {
	l.mu.RLock()
	length := l.list.Len()
	if length > 0 {
		for i, e := 0, l.list.Back(); i < length; i, e = i+1, e.Prev() {
			if !f(e) {
				break
			}
		}
	}
	l.mu.RUnlock()
}

// Front returns the first element of list <l> or nil if the list is empty.
func (l *List) Front() (e *Element) {
	l.mu.RLock()
	e = l.list.Front()
	l.mu.RUnlock()
	return
}

// Back returns the last element of list <l> or nil if the list is empty.
func (l *List) Back() (e *Element) {
	l.mu.RLock()
	e = l.list.Back()
	l.mu.RUnlock()
	return
}

// FrontValue returns value of the first element of <l> or nil if the list is empty.
func (l *List) FrontValue() (value interface{}) {
	l.mu.RLock()
	if e := l.list.Front(); e != nil {
		value = e.Value
	}
	l.mu.RUnlock()
	return
}

// BackValue returns value of the last element of <l> or nil if the list is empty.
func (l *List) BackValue() (value interface{}) {
	l.mu.RLock()
	if e := l.list.Back(); e != nil {
		value = e.Value
	}
	l.mu.RUnlock()
	return
}

// FrontAll copies and returns values of all elements from front of <l> as slice.
func (l *List) FrontAll() (values []interface{}) {
	l.mu.RLock()
	length := l.list.Len()
	if length > 0 {
		values = make([]interface{}, length)
		for i, e := 0, l.list.Front(); i < length; i, e = i+1, e.Next() {
			values[i] = e.Value
		}
	}
	l.mu.RUnlock()
	return
}

// BackAll copies and returns values of all elements from back of <l> as slice.
func (l *List) BackAll() (values []interface{}) {
	l.mu.RLock()
	length := l.list.Len()
	if length > 0 {
		values = make([]interface{}, length)
		for i, e := 0, l.list.Back(); i < length; i, e = i+1, e.Prev() {
			values[i] = e.Value
		}
	}
	l.mu.RUnlock()
	return
}

// Removes removes multiple elements <es> from <l> if <es> are elements of list <l>.
func (l *List) Remove(es ...*Element) {
	l.mu.Lock()
	for _, e := range es {
		l.list.Remove(e)
	}
	l.mu.Unlock()
	return
}

// RemoveFront removes <max> elements from front of <l>
// and returns values of the removed elements as slice.
func (l *List) RemoveFront(num ...int) (values []interface{}) {
	l.mu.RLock()
	if len(num) > 0 {
		max := num[0]
		length := l.list.Len()
		if length > 0 {
			if max > 0 && max < length {
				length = max
			}
			tempe := (*Element)(nil)
			values = make([]interface{}, length)
			for i := 0; i < length; i++ {
				tempe = l.list.Front()
				values[i] = l.list.Remove(tempe)
			}
		}
	} else {
		if e := l.list.Front(); e != nil {
			values = make([]interface{}, 1)
			values[0] = l.list.Remove(e)
		}
	}
	l.mu.RUnlock()
	return
}

// RemoveBack removes <max> elements from back of <l>
// and returns values of the removed elements as slice.
func (l *List) RemoveBack(num ...int) (values []interface{}) {
	l.mu.Lock()
	if len(num) > 0 {
		max := num[0]
		length := l.list.Len()
		if length > 0 {
			if max > 0 && max < length {
				length = max
			}
			tempe := (*Element)(nil)
			values = make([]interface{}, length)
			for i := 0; i < length; i++ {
				tempe = l.list.Back()
				values[i] = l.list.Remove(tempe)
			}
		}
	} else {
		if e := l.list.Back(); e != nil {
			values = make([]interface{}, 1)
			values[0] = l.list.Remove(e)
		}
	}
	l.mu.Unlock()
	return
}

// Init removes all elements from list <l>.
func (l *List) Init() {
	l.mu.Lock()
	l.list.Init()
	l.mu.Unlock()
}

// Len returns the number of elements of list <l>.
// The complexity is O(1).
func (l *List) Len() (length int) {
	l.mu.RLock()
	length = l.list.Len()
	l.mu.RUnlock()
	return
}

// LockFunc locks writing with given callback function <f> within RWMutex.Lock.
func (l *List) LockFunc(f func(list *list.List)) {
	l.mu.Lock()
	defer l.mu.Unlock()
	f(l.list)
}

// RLockFunc locks reading with given callback function <f> within RWMutex.RLock.
func (l *List) RLockFunc(f func(list *list.List)) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	f(l.list)
}

// Join joins items with a string <sep>.
func (l *List) Join(sep string) string {
	return strings.Join(cast.ToStringSlice(l.FrontAll()), sep)
}

func (l *List) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.FrontAll())
}

func (l *List) UnmarshalJSON(b []byte) error {
	var data []interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	} else {
		l.PushBack(data...)
		return nil
	}
}

func (l *List) String() string {
	rs, _ := l.MarshalJSON()
	return string(rs)
}
