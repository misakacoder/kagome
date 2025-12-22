package set

import (
	"cmp"
	"github.com/misakacoder/kagome/maps"
)

type linkedSet[E cmp.Ordered] struct {
	data maps.Map[E, struct{}]
}

func (set *linkedSet[E]) Add(item E) {
	set.data.Put(item, value)
}

func (set *linkedSet[E]) AddAll(items ...E) {
	for _, item := range items {
		set.Add(item)
	}
}

func (set *linkedSet[E]) Contains(item E) bool {
	_, ok := set.data.Get(item)
	return ok
}

func (set *linkedSet[E]) ContainsAll(items ...E) bool {
	for _, item := range items {
		if !set.Contains(item) {
			return false
		}
	}
	return true
}

func (set *linkedSet[E]) Remove(item E) {
	set.data.Del(item)
}

func (set *linkedSet[E]) RemoveAll(items ...E) {
	for _, item := range items {
		set.Remove(item)
	}
}

func (set *linkedSet[E]) Clear() {
	set.data.Clear()
}

func (set *linkedSet[E]) Size() int {
	return set.data.Len()
}

func (set *linkedSet[E]) Range(fn func(item E)) {
	set.data.Range(func(k E, v struct{}) {
		fn(k)
	})
}

func (set *linkedSet[E]) Slice() []E {
	var slice []E
	set.data.Range(func(k E, v struct{}) {
		slice = append(slice, k)
	})
	return slice
}

func NewLinkedSet[E cmp.Ordered]() Set[E] {
	return &linkedSet[E]{
		data: maps.NewLinkedMap[E, struct{}](),
	}
}
