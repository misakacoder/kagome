package maps

import (
	"time"
)

type ExpiredMap[K comparable, V any] map[K]V

func (expiredMap ExpiredMap[K, V]) PutTimeout(key K, value V, timeout time.Duration) {
	expiredMap[key] = value
	if timeout > 0 {
		time.AfterFunc(timeout, func() {
			delete(expiredMap, key)
		})
	}
}

func (expiredMap ExpiredMap[K, V]) Put(key K, value V) {
	expiredMap.PutTimeout(key, value, -1)
}

func (expiredMap ExpiredMap[K, V]) Get(key K) (V, bool) {
	value, ok := expiredMap[key]
	return value, ok
}

func (expiredMap ExpiredMap[K, V]) Del(key K) {
	delete(expiredMap, key)
}

func (expiredMap ExpiredMap[K, V]) Contains(key K) bool {
	_, ok := expiredMap.Get(key)
	return ok
}

func (expiredMap ExpiredMap[K, V]) Len() int {
	return len(expiredMap)
}

func (expiredMap ExpiredMap[K, V]) Clear() {
	clear(expiredMap)
}

func (expiredMap ExpiredMap[K, V]) Range(fn func(K, V)) {
	for k, v := range expiredMap {
		fn(k, v)
	}
}

func NewExpiredMap[K comparable, V any]() ExpiredMap[K, V] {
	return ExpiredMap[K, V]{}
}
