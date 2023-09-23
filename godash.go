package godash

import (
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

func Filter[V any](collection []V, predicate func(item V, index int) bool) []V {
	return lo.Filter(collection, predicate)
}

func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	return lo.Map(collection, iteratee)
}

func Reduce[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	return lo.Reduce(collection, accumulator, initial)
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

func Keys[K comparable, V any](in map[K]V) []K {
	return lo.Keys(in)
}

func Values[K comparable, V any](in map[K]V) []V {
	return lo.Values(in)
}

func Sum[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T {
	return lo.Sum(collection)
}

func SumBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection []T, iteratee func(item T) R) R {
	return lo.SumBy(collection, iteratee)
}

func EveryBy[T any](collection []T, predicate func(item T) bool) bool {
	return lo.EveryBy(collection, predicate)
}

func NoneBy[T any](collection []T, predicate func(item T) bool) bool {
	return lo.NoneBy(collection, predicate)
}

func Find[T any](collection []T, predicate func(item T) bool) (T, bool) {
	return lo.Find(collection, predicate)
}
