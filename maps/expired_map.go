package maps

import (
	"time"
)

type expiredMap[K comparable, V any] map[K]V

func (expiredMap expiredMap[K, V]) PutTimeout(key K, value V, timeout time.Duration) {
	expiredMap[key] = value
	if timeout > 0 {
		time.AfterFunc(timeout, func() {
			delete(expiredMap, key)
		})
	}
}

func (expiredMap expiredMap[K, V]) Put(key K, value V) {
	expiredMap.PutTimeout(key, value, -1)
}

func (expiredMap expiredMap[K, V]) Get(key K) (V, bool) {
	value, ok := expiredMap[key]
	return value, ok
}

func (expiredMap expiredMap[K, V]) Del(key K) {
	delete(expiredMap, key)
}

func (expiredMap expiredMap[K, V]) Contains(key K) bool {
	_, ok := expiredMap.Get(key)
	return ok
}

func (expiredMap expiredMap[K, V]) Len() int {
	return len(expiredMap)
}

func (expiredMap expiredMap[K, V]) Clear() {
	clear(expiredMap)
}

func (expiredMap expiredMap[K, V]) Range(fn func(K, V)) {
	for k, v := range expiredMap {
		fn(k, v)
	}
}

func NewExpiredMap[K comparable, V any]() expiredMap[K, V] {
	return expiredMap[K, V]{}
}
