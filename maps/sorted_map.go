package maps

import (
	"cmp"
	"slices"
)

type sortedMap[K cmp.Ordered, V any] struct {
	keys []K
	data map[K]V
	sort bool
}

func (sortedMap *sortedMap[K, V]) Put(key K, value V) {
	if _, ok := sortedMap.data[key]; !ok {
		sortedMap.keys = append(sortedMap.keys, key)
		if sortedMap.sort {
			slices.Sort(sortedMap.keys)
		}
	}
	sortedMap.data[key] = value
}

func (sortedMap *sortedMap[K, V]) Get(key K) (V, bool) {
	value, ok := sortedMap.data[key]
	return value, ok
}

func (sortedMap *sortedMap[K, V]) Del(key K) {
	delete(sortedMap.data, key)
	i, found := slices.BinarySearch(sortedMap.keys, key)
	if found {
		sortedMap.keys = slices.Delete(sortedMap.keys, i, i+1)
	}
}

func (sortedMap *sortedMap[K, V]) Contains(key K) bool {
	_, ok := sortedMap.Get(key)
	return ok
}

func (sortedMap *sortedMap[K, V]) Len() int {
	return len(sortedMap.data)
}

func (sortedMap *sortedMap[K, V]) Clear() {
	clear(sortedMap.keys)
	clear(sortedMap.data)
}

func (sortedMap *sortedMap[K, V]) Range(fn func(key K, value V)) {
	for _, key := range sortedMap.keys {
		value := sortedMap.data[key]
		fn(key, value)
	}
}

func NewSortedMap[K cmp.Ordered, V any]() Map[K, V] {
	return &sortedMap[K, V]{
		keys: make([]K, 0, 8),
		data: make(map[K]V, 8),
		sort: true,
	}
}

func NewLinkedMap[K cmp.Ordered, V any]() Map[K, V] {
	return &sortedMap[K, V]{
		keys: make([]K, 0, 8),
		data: make(map[K]V, 8),
		sort: false,
	}
}
