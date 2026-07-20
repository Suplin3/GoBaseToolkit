package stats

func Freq[T comparable](s []T) map[T]int {
	out := map[T]int{}
	for _, x := range s {
		out[x]++
	}
	return out
}

func UniqueCounter[T comparable](s []T) int {
	out := 0
	for _, x := range Freq(s) {
		if x == 1 {
			out++
		}
	}
	return out
}

func Top[T comparable](s []T) T {
	var topV T
	top := 0
	for v, x := range Freq(s) {
		if x > top {
			top = x
			topV = v
		}
	}
	return topV
}

func GroupBy[T any, K comparable](s []T, f func(T) K) map[K][]T {
	out := map[K][]T{}
	for _, x := range s {
		out[f(x)] = append(out[f(x)], x)
	}
	return out
}

func CountBy[T any, K comparable](s []T, f func(T) K) map[K]int {
	group := GroupBy(s, f)
	out := map[K]int{}
	for k, x := range group {
		out[k] = len(x)
	}
	return out
}
