package assert

func Equal[T any](left, right T) bool {
	switch v := (interface{})(left).(type) {
	case equaler[T]:
		return v.Equal(right)
	default:
		return (interface{})(left) == (interface{})(right)
	}
}

type equaler[T any] interface {
	Equal(u T) bool
}
