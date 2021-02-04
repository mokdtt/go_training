package cycle

import (
	"reflect"
	"unsafe"
)

func isCyclic(x reflect.Value, seen map[cycle]bool) bool {
	if !x.IsValid() {
		return false
	}

	if x.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		c := cycle{xptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return isCyclic(x.Elem(), seen)
	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if isCyclic(x.Index(i), seen) {
				return true
			}
		}
		return false
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if isCyclic(x.Field(i), seen) {
				return true
			}
		}
		return false
	case reflect.Map:
		for _, k := range x.MapKeys() {
			if isCyclic(x.MapIndex(k), seen) {
				return true
			}
		}
		return false
	}
	return false
}

func IsCyclic(x interface{}) bool {
	seen := make(map[cycle]bool)
	return isCyclic(reflect.ValueOf(x), seen)
}

type cycle struct {
	x unsafe.Pointer
	t reflect.Type
}
