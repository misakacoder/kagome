package ptr

var (
	True  = New(true)
	False = New(false)
)

func New[T any](v T) *T {
	return &v
}

func Value[T any](ptr *T) T {
	if ptr == nil {
		var value T
		return value
	}
	return *ptr
}

func ValueOrDefault[T any](ptr *T, defaultValue T) T {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}
