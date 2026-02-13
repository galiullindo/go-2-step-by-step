package main

func Filter[T any](s []T, predicate func(T) bool) []T {
	newS := make([]T, 0)
	for _, i := range s {
		if predicate(i) {
			newS = append(newS, i)
		}
	}

	return newS
}
