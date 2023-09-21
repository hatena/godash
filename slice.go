package godash

import slo "github.com/samber/lo"

func Uniq[T comparable](collection []T) []T {
	return slo.Uniq(collection)
}

func UniqBy[T any, U comparable](collection []T, iteratee func(item T) U) []T {
	return slo.UniqBy(collection, iteratee)
}
