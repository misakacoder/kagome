package set

type hashSet[E comparable] map[E]struct{}

func (set hashSet[E]) Add(item E) {
	set[item] = value
}

func (set hashSet[E]) AddAll(items ...E) {
	for _, item := range items {
		set.Add(item)
	}
}

func (set hashSet[E]) Contains(item E) bool {
	_, ok := set[item]
	return ok
}

func (set hashSet[E]) ContainsAll(items ...E) bool {
	for _, item := range items {
		if !set.Contains(item) {
			return false
		}
	}
	return true
}

func (set hashSet[E]) Remove(item E) {
	delete(set, item)
}

func (set hashSet[E]) RemoveAll(items ...E) {
	for _, item := range items {
		set.Remove(item)
	}
}

func (set hashSet[E]) Clear() {
	clear(set)
}

func (set hashSet[E]) Size() int {
	return len(set)
}

func (set hashSet[E]) Range(fn func(E)) {
	for item := range set {
		fn(item)
	}
}

func (set hashSet[E]) Slice() []E {
	var slice []E
	for item := range set {
		slice = append(slice, item)
	}
	return slice
}

func New[E comparable]() Set[E] {
	return hashSet[E]{}
}
