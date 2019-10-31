package acast

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/asktop/gotools/ascan"
	"reflect"
	"strings"
)

func MapsToStructs(src []map[string]interface{}, value interface{}) error {
	var column []string
	if len(src) == 0 {
		return errors.New("page: src can't be nil")
	}
	for k, _ := range src[0] {
		column = append(column, k)
	}
	ptr := make([]interface{}, len(column))

	var v reflect.Value
	var elemType reflect.Type

	if il, ok := value.(interfaceLoader); ok {
		v = reflect.ValueOf(il.v)
		elemType = il.typ
	} else {
		v = reflect.ValueOf(value)
	}

	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("page: attempt to load into an invalid pointer")
	}
	v = v.Elem()
	isScanner := v.Addr().Type().Implements(typeScanner)
	isSlice := v.Kind() == reflect.Slice && v.Type().Elem().Kind() != reflect.Uint8 && !isScanner
	isMap := v.Kind() == reflect.Map && !isScanner
	isMapOfSlices := isMap && v.Type().Elem().Kind() == reflect.Slice && v.Type().Elem().Elem().Kind() != reflect.Uint8
	if isMap {
		v.Set(reflect.MakeMap(v.Type()))
	}

	s := newTagStore()
	count := 0
	for _, data := range src {
		var elem, keyElem reflect.Value

		if elemType != nil {
			elem = reflectAlloc(elemType)
		} else if isMapOfSlices {
			elem = reflectAlloc(v.Type().Elem().Elem())
		} else if isSlice || isMap {
			elem = reflectAlloc(v.Type().Elem())
		} else {
			elem = v
		}

		if isMap {
			err := s.findPtr(elem, column[1:], ptr[1:])
			if err != nil {
				return err
			}
			keyElem = reflectAlloc(v.Type().Key())
			err = s.findPtr(keyElem, column[:1], ptr[:1])
			if err != nil {
				return err
			}
		} else {
			err := s.findPtr(elem, column, ptr)
			if err != nil {
				return err
			}
		}

		// Before scanning, set nil pointer to dummy dest.
		// After that, reset pointers to nil for the next batch.
		for i := range ptr {
			if ptr[i] == nil {
				ptr[i] = dummyDest
			}
		}

		for i, col := range column {
			err := ascan.ConvertAssign(ptr[i], data[col])
			if err != nil {
				return fmt.Errorf(`page: Scan error on column name %q: %v`, col, err)
			}
		}

		for i := range ptr {
			ptr[i] = nil
		}

		count++

		if isSlice {
			v.Set(reflect.Append(v, elem))
		} else if isMapOfSlices {
			s := v.MapIndex(keyElem)
			if !s.IsValid() {
				s = reflect.Zero(v.Type().Elem())
			}
			v.SetMapIndex(keyElem, reflect.Append(s, elem))
		} else if isMap {
			v.SetMapIndex(keyElem, elem)
		} else {
			break
		}
	}
	return nil
}

func MapsStrToStructs(src []map[string]string, value interface{}) error {
	var column []string
	if len(src) == 0 {
		return errors.New("page: src can't be nil")
	}
	for k, _ := range src[0] {
		column = append(column, k)
	}
	ptr := make([]interface{}, len(column))

	var v reflect.Value
	var elemType reflect.Type

	if il, ok := value.(interfaceLoader); ok {
		v = reflect.ValueOf(il.v)
		elemType = il.typ
	} else {
		v = reflect.ValueOf(value)
	}

	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("page: attempt to load into an invalid pointer")
	}
	v = v.Elem()
	isScanner := v.Addr().Type().Implements(typeScanner)
	isSlice := v.Kind() == reflect.Slice && v.Type().Elem().Kind() != reflect.Uint8 && !isScanner
	isMap := v.Kind() == reflect.Map && !isScanner
	isMapOfSlices := isMap && v.Type().Elem().Kind() == reflect.Slice && v.Type().Elem().Elem().Kind() != reflect.Uint8
	if isMap {
		v.Set(reflect.MakeMap(v.Type()))
	}

	s := newTagStore()
	count := 0
	for _, data := range src {
		var elem, keyElem reflect.Value

		if elemType != nil {
			elem = reflectAlloc(elemType)
		} else if isMapOfSlices {
			elem = reflectAlloc(v.Type().Elem().Elem())
		} else if isSlice || isMap {
			elem = reflectAlloc(v.Type().Elem())
		} else {
			elem = v
		}

		if isMap {
			err := s.findPtr(elem, column[1:], ptr[1:])
			if err != nil {
				return err
			}
			keyElem = reflectAlloc(v.Type().Key())
			err = s.findPtr(keyElem, column[:1], ptr[:1])
			if err != nil {
				return err
			}
		} else {
			err := s.findPtr(elem, column, ptr)
			if err != nil {
				return err
			}
		}

		// Before scanning, set nil pointer to dummy dest.
		// After that, reset pointers to nil for the next batch.
		for i := range ptr {
			if ptr[i] == nil {
				ptr[i] = dummyDest
			}
		}

		for i, col := range column {
			err := ascan.ConvertAssign(ptr[i], data[col])
			if err != nil {
				return fmt.Errorf(`page: Scan error on column name %q: %v`, col, err)
			}
		}

		for i := range ptr {
			ptr[i] = nil
		}

		count++

		if isSlice {
			v.Set(reflect.Append(v, elem))
		} else if isMapOfSlices {
			s := v.MapIndex(keyElem)
			if !s.IsValid() {
				s = reflect.Zero(v.Type().Elem())
			}
			v.SetMapIndex(keyElem, reflect.Append(s, elem))
		} else if isMap {
			v.SetMapIndex(keyElem, elem)
		} else {
			break
		}
	}
	return nil
}

var (
	dummyDest   sql.Scanner = dummyScanner{}
	typeScanner             = reflect.TypeOf((*sql.Scanner)(nil)).Elem()
	typeValuer              = reflect.TypeOf((*driver.Valuer)(nil)).Elem()
	NameMapping             = camelCaseToSnakeCase
)

type dummyScanner struct{}

func (dummyScanner) Scan(interface{}) error {
	return nil
}

type interfaceLoader struct {
	v   interface{}
	typ reflect.Type
}

func camelCaseToSnakeCase(name string) string {
	var buf strings.Builder
	buf.Grow(len(name) * 2)

	for i := 0; i < len(name); i++ {
		buf.WriteByte(toLower(name[i]))
		if i != len(name)-1 && isUpper(name[i+1]) &&
			(isLower(name[i]) || isDigit(name[i]) ||
				(i != len(name)-2 && isLower(name[i+2]))) {
			buf.WriteByte('_')
		}
	}

	return buf.String()
}

func isUpper(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

func isLower(b byte) bool {
	return 'a' <= b && b <= 'z'
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func toLower(b byte) byte {
	if isUpper(b) {
		return b - 'A' + 'a'
	}
	return b
}

func reflectAlloc(typ reflect.Type) reflect.Value {
	if typ.Kind() == reflect.Ptr {
		return reflect.New(typ.Elem())
	}
	return reflect.New(typ).Elem()
}

type tagStore struct {
	m map[reflect.Type][]string
}

func newTagStore() *tagStore {
	return &tagStore{
		m: make(map[reflect.Type][]string),
	}
}

func (s *tagStore) get(t reflect.Type) []string {
	if t.Kind() != reflect.Struct {
		return nil
	}
	if _, ok := s.m[t]; !ok {
		l := make([]string, t.NumField())
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.PkgPath != "" && !field.Anonymous {
				// unexported
				continue
			}
			tag := field.Tag.Get("db")
			if tag == "-" {
				// ignore
				continue
			}
			if tag == "" {
				tag = field.Tag.Get("json")
				if tag == "-" {
					// ignore
					continue
				}
			}
			if tag == "" {
				// no tag, but we can record the field name
				tag = NameMapping(field.Name)
			}
			l[i] = tag
		}
		s.m[t] = l
	}
	return s.m[t]
}

func (s *tagStore) findPtr(value reflect.Value, name []string, ptr []interface{}) error {
	if value.CanAddr() && value.Addr().Type().Implements(typeScanner) {
		ptr[0] = value.Addr().Interface()
		return nil
	}
	switch value.Kind() {
	case reflect.Struct:
		s.findValueByName(value, name, ptr, true)
		return nil
	case reflect.Ptr:
		if value.IsNil() {
			value.Set(reflect.New(value.Type().Elem()))
		}
		return s.findPtr(value.Elem(), name, ptr)
	default:
		ptr[0] = value.Addr().Interface()
		return nil
	}
}

func (s *tagStore) findValueByName(value reflect.Value, name []string, ret []interface{}, retPtr bool) {
	if value.Type().Implements(typeValuer) {
		return
	}
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			return
		}
		s.findValueByName(value.Elem(), name, ret, retPtr)
	case reflect.Struct:
		l := s.get(value.Type())
		for i := 0; i < value.NumField(); i++ {
			tag := l[i]
			if tag == "" {
				continue
			}
			fieldValue := value.Field(i)
			for i, want := range name {
				if want != tag {
					continue
				}
				if ret[i] == nil {
					if retPtr {
						ret[i] = fieldValue.Addr().Interface()
					} else {
						ret[i] = fieldValue
					}
				}
			}
			s.findValueByName(fieldValue, name, ret, retPtr)
		}
	}
}
