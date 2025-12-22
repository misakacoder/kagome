package json

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"time"
)

func ToJSONString(object any) (string, error) {
	bytes, err := json.Marshal(object)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func ParseObject[T any](text string) (T, error) {
	var result T
	if text == "" {
		return result, errors.New("empty text")
	}
	tp := reflect.TypeOf(result)
	tpKind := tp.Kind()
	if tpKind == reflect.Ptr {
		tp = tp.Elem()
		tpKind = tp.Kind()
	}
	timeType := reflect.TypeOf(time.Time{})
	if (tpKind == reflect.String || tp == timeType || tp.ConvertibleTo(timeType)) && text[0] != '"' {
		text = strconv.Quote(text)
	}
	err := json.Unmarshal([]byte(text), &result)
	return result, err
}
