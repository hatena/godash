package godash

import "github.com/samber/lo"

func Filter[V any](collection []V, predicate func(item V, index int) bool) []V {
	return lo.Filter(collection, predicate)
}

func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	return lo.Map(collection, iteratee)
}

func Uniq[T comparable](collection []T) []T {
	return lo.Uniq(collection)
}

func UniqBy[T any, U comparable](collection []T, iteratee func(item T) U) []T {
	return lo.UniqBy(collection, iteratee)
}

func GroupBy[T any, U comparable](collection []T, iteratee func(item T) U) map[U][]T {
	return lo.GroupBy(collection, iteratee)
}

func Chunk[T any](collection []T, size int) [][]T {
	return lo.Chunk(collection, size)
}

func PartitionBy[T any, K comparable](collection []T, iteratee func(item T) K) [][]T {
	return lo.PartitionBy(collection, iteratee)
}

func Flatten[T any](collection [][]T) []T {
	return lo.Flatten(collection)
}

func KeyBy[K comparable, V any](collection []V, iteratee func(item V) K) map[K]V {
	return lo.KeyBy(collection, iteratee)
}
