package utils

func Intersection[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if containsGeneric(b, v) {
			set = append(set, v)
		}
	}

	return set
}

func containsGeneric[T comparable](b []T, e T) bool {
	for _, v := range b {
		if v == e {
			return true
		}
	}
	return false
}

func MinIntSlice(v []int) (m int) {
	if len(v) > 0 {
		m = v[0]
	}
	for i, e := range v {
		if i == 0 || e < m {
			m = e
		}
	}
	return
}
