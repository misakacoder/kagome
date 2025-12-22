package math

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type Unsigned interface {
	~int | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Number interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Abs[T Signed](number T) T {
	if number < 0 {
		return -number
	}
	return number
}

func AbsDifference[T Number](a, b T) T {
	if a > b {
		return a - b
	}
	return b - a
}

func Max[T Number](first T, rest ...T) T {
	result := first
	for _, v := range rest {
		if v > result {
			result = v
		}
	}
	return result
}

func Min[T Number](first T, rest ...T) T {
	result := first
	for _, v := range rest {
		if v < result {
			result = v
		}
	}
	return result
}

func Sum[T Number](numbers ...T) T {
	var result T
	for _, v := range numbers {
		result += v
	}
	return result
}

func Avg[T Number](numbers ...T) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := Sum(numbers...)
	return float64(sum) / float64(len(numbers))
}
