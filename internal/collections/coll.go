package collections

func IndexOf[T comparable](s []T, value T) int {
	for i, v := range s {
		if v == value {
			return i
		}
	}
	return -1
}

func RemoveDuplicates[T comparable](s []T) []T {
	out := make([]T, 0, len(s))
	for _, v := range s {
		if IndexOf(out, v) < 0 {
			out = append(out, v)
		}
	}
	return out
}

func Reverse[T any](s []T) []T {
	out := make([]T, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return out
}

func Filter[T any](s []T, f func(T) bool) []T {
	out := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

func Chunk[T any](s []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(s); i += chunkSize {
		chunks = append(chunks, s[i:min(i+chunkSize, len(s))])
	}
	return chunks
}
