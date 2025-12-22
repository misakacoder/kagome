package maps

type Map[K comparable, V any] interface {
	Put(K, V)
	Get(K) (V, bool)
	Del(K)
	Contains(K) bool
	Len() int
	Clear()
	Range(func(K, V))
}
