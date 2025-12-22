package set

type Set[E comparable] interface {
	Add(item E)
	AddAll(items ...E)
	Contains(item E) bool
	ContainsAll(items ...E) bool
	Remove(item E)
	RemoveAll(items ...E)
	Clear()
	Size() int
	Range(func(item E))
	Slice() []E
}

var value = struct{}{}
