package optional

import (
	"encoding/json"
)

type Optional[T any] struct {
	Value T
	Set   bool
}

func New[T any](val T) Optional[T] {
	return Optional[T]{
		Value: val,
		Set:   true,
	}
}

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.Set {
		return []byte("null"), nil
	}
	return json.Marshal(o.Value)
}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	if data == nil {
		*o = Optional[T]{}
		return nil
	}

	var val T
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	*o = New(val)

	return nil
}
