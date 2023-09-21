package godash

import "github.com/samber/lo"

func Uniq[T comparable](collection []T) []T {
	return lo.Uniq(collection)
}

func UniqBy[T any, U comparable](collection []T, iteratee func(item T) U) []T {
	return lo.UniqBy(collection, iteratee)
}
