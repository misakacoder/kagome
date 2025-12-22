package str

import (
	"fmt"
	"github.com/misakacoder/kagome/errs"
	"github.com/misakacoder/kagome/math"
	"math/rand"
	"strconv"
	"strings"
)

var letters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Atoi[T math.Number](str string) T {
	var result T
	switch any(result).(type) {
	case float32:
		v, err := strconv.ParseFloat(str, 32)
		errs.Panic(err)
		return T(v)
	case float64:
		v, err := strconv.ParseFloat(str, 64)
		errs.Panic(err)
		return T(v)
	default:
		v, err := strconv.Atoi(str)
		errs.Panic(err)
		return T(v)
	}
}

func Itoa[T math.Number](number T) string {
	switch v := any(number).(type) {
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", number)
	}
}

func AnyBlank(strs ...string) bool {
	for _, str := range strs {
		if IsBlank(str) {
			return true
		}
	}
	return false
}

func AllBlank(strs ...string) bool {
	for _, str := range strs {
		if NonBlank(str) {
			return false
		}
	}
	return true
}

func NoneBlank(strs ...string) bool {
	for _, str := range strs {
		if IsBlank(str) {
			return false
		}
	}
	return true
}

func IsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}

func NonBlank(str string) bool {
	return !IsBlank(str)
}

func RandString(length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
