package max

type Comparable interface {
	int | int32 | int64 | string | float64 | float32
}

func Max[T Comparable](values []T) T {
	if len(values) == 0 {
		panic("Empty string")
	}
	_max := values[0]
	for _, value := range values[1:] {
		if value > _max {
			_max = value
		}
	}
	return _max
}
