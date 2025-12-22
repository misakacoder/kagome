package cond

import (
	"reflect"
	"strings"
)

func Ternary[T any](condition bool, value, defaultValue T) T {
	if condition {
		return value
	}
	return defaultValue
}

func LazyTernary[T any](condition bool, valueFn, defaultValueFn func() T) T {
	if condition {
		if valueFn == nil {
			panic("valueFn cannot be nil")
		}
		return valueFn()
	}
	if defaultValueFn == nil {
		panic("defaultValueFn cannot be nil")
	}
	return defaultValueFn()
}

func RequireNonNilElse[T any](value, defaultValue T) T {
	return Ternary(!IsNil(value), value, defaultValue)
}

func LazyRequireNonNilElse[T any](valueFn, defaultValueFn func() T) T {
	if valueFn == nil {
		panic("valueFn cannot be nil")
	}
	value := valueFn()
	if !IsNil(value) {
		return value
	}
	if defaultValueFn == nil {
		panic("defaultValueFn cannot be nil")
	}
	return defaultValueFn()
}

func AnyNil(values ...any) bool {
	for _, v := range values {
		if IsNil(v) {
			return true
		}
	}
	return false
}

func AllNil(values ...any) bool {
	for _, v := range values {
		if NonNil(v) {
			return false
		}
	}
	return true
}

func NoneNil(values ...any) bool {
	for _, v := range values {
		if IsNil(v) {
			return false
		}
	}
	return true
}

func AnyEmpty(values ...any) bool {
	for _, v := range values {
		if IsEmpty(v) {
			return true
		}
	}
	return false
}

func AllEmpty(values ...any) bool {
	for _, v := range values {
		if NonEmpty(v) {
			return false
		}
	}
	return true
}

func NoneEmpty(values ...any) bool {
	for _, v := range values {
		if IsEmpty(v) {
			return false
		}
	}
	return true
}

func IsNil(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Map, reflect.Slice, reflect.Chan, reflect.Func, reflect.Ptr, reflect.UnsafePointer, reflect.Interface:
		return rv.IsNil()
	default:
	}
	return false
}

func NonNil(v any) bool {
	return !IsNil(v)
}

func IsEmpty(v any) bool {
	if IsNil(v) {
		return true
	}
	switch tp := v.(type) {
	case string:
		return strings.TrimSpace(tp) == ""
	case []byte:
		return len(tp) == 0
	default:
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Map, reflect.Array, reflect.Slice:
			return rv.Len() == 0
		default:
		}
	}
	return false
}

func NonEmpty(v any) bool {
	return !IsEmpty(v)
}

func IsZero(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	return rv.IsZero()
}

func NonZero(v any) bool {
	return !IsZero(v)
}
