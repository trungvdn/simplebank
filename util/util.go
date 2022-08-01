package util

func ContainFunc[T any](arr []T, fn func(T) bool) []T {
	res := []T{}
	for _, v := range arr {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}
