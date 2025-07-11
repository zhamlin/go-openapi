package openapi

import (
	"encoding/json"
)

// SingleOrArray holds list or single value
type SingleOrArray[T any] []T

// NewSingleOrArray creates SingleOrArray object.
func NewSingleOrArray[T any](v ...T) *SingleOrArray[T] {
	vv := SingleOrArray[T](v)
	return &vv
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (o *SingleOrArray[T]) UnmarshalJSON(data []byte) error {
	var ret []T
	if json.Unmarshal(data, &ret) != nil {
		var s T
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		ret = []T{s}
	}
	*o = ret
	return nil
}

// MarshalJSON implements json.Marshaler interface.
func (o *SingleOrArray[T]) MarshalJSON() ([]byte, error) {
	var v any = []T(*o)
	if len(*o) == 1 {
		v = (*o)[0]
	}
	return json.Marshal(&v)
}

func (o *SingleOrArray[T]) Add(v ...T) *SingleOrArray[T] {
	*o = append(*o, v...)
	return o
}
