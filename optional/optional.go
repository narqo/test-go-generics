package optional

import (
	"encoding/json"
)

type Value[T any] struct {
	Value T
}

func New[T any](val T) *Value[T] {
	return &Value[T]{
		Value: val,
	}
}

func (o *Value[T]) MarshalJSON() ([]byte, error) {
	if o == nil {
		return []byte("null"), nil
	}
	return json.Marshal(o.Value)
}

func (o *Value[T]) UnmarshalJSON(data []byte) error {
	if data == nil {
		*o = Value[T]{}
		return nil
	}

	var val T
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	*o = Value[T]{
		Value: val,
	}

	return nil
}
