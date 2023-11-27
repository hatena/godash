// Package godash implements slice/map related methods and is a wrapper of the [github.com/samber/lo] package.
//
// This package focuses on only presenting a subset of functions from [github.com/samber/lo],
// in order to make sure that the user takes advantage of Go's standard methods from the slice package etc.
// If any of the functions presented here gets included in the standard library, that function will be
// marked deprecated (the function will not get removed, but will point the user to the standard library).
package godash

import (
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

// Filter iterates through the collection and returns a slice with only the values that match the predicate.
func Filter[V any](collection []V, predicate func(item V, index int) bool) []V {
	return lo.Filter(collection, predicate)
}

// Map iterates through the collection and returns a slice with the values converted though the iteratee.
func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	return lo.Map(collection, iteratee)
}

// Reduce iterates through the collection and reduces it to one value, using the accumulator.
func Reduce[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	return lo.Reduce(collection, accumulator, initial)
}

// Uniq returns a slice with only unique values.
func Uniq[T comparable](collection []T) []T {
	return lo.Uniq(collection)
}

// UniqBy returns a slice with only unique values. The uniqueness is determined by the iteratee.
func UniqBy[T any, U comparable](collection []T, iteratee func(item T) U) []T {
	return lo.UniqBy(collection, iteratee)
}

// GroupBy groups the values of the collection using the iteratee, and returns it as a map.
func GroupBy[T any, U comparable](collection []T, iteratee func(item T) U) map[U][]T {
	return lo.GroupBy(collection, iteratee)
}

// Chunk receives the collection and chunks it into a slice of slices each of the given size.
func Chunk[T any](collection []T, size int) [][]T {
	return lo.Chunk(collection, size)
}

// PartitionBy partitions the collection by the iteratee. This is similar to Chunk, but instead of a given size, a function can be applied.
// Use GroupBy if you want a map of slices.
func PartitionBy[T any, K comparable](collection []T, iteratee func(item T) K) [][]T {
	return lo.PartitionBy(collection, iteratee)
}

// Flatten flattens a slice of slices to a single slice.
func Flatten[T any](collection [][]T) []T {
	return lo.Flatten(collection)
}

// KeyBy iterates through the collection and returns a map with the key generated via the iteratee.
func KeyBy[K comparable, V any](collection []V, iteratee func(item V) K) map[K]V {
	return lo.KeyBy(collection, iteratee)
}

// Keys returns the keys of the map as a slice.
func Keys[K comparable, V any](in map[K]V) []K {
	return lo.Keys(in)
}

// Values returns the values of the map as a slice.
func Values[K comparable, V any](in map[K]V) []V {
	return lo.Values(in)
}

// Sum returns a sum of the values within the slice. The values of the slice need to be numbers.
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T {
	return lo.Sum(collection)
}

// SumBy returns a sum of the values within the slice, added using the specified iteratee function. The values of the slice need to be numbers.
func SumBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection []T, iteratee func(item T) R) R {
	return lo.SumBy(collection, iteratee)
}

// EveryBy returns whether or not all the values within the collection meet the predicate.
func EveryBy[T any](collection []T, predicate func(item T) bool) bool {
	return lo.EveryBy(collection, predicate)
}

// NoneBy returns whether or not all the values within the collection do not meet the predicate.
func NoneBy[T any](collection []T, predicate func(item T) bool) bool {
	return lo.NoneBy(collection, predicate)
}

// Find returns the first value that meets the predicate, as well as whether or not a match was found.
func Find[T any](collection []T, predicate func(item T) bool) (T, bool) {
	return lo.Find(collection, predicate)
}

func Associate[T any, K comparable, V any](collection []T, transform func(item T) (K, V)) map[K]V {
	return lo.Associate(collection, transform)
}

// FlatMap iterates through the collection and returns a slice with the values converted through the iteratee and flattened.
func FlatMap[T any, R any](collection []T, iteratee func(item T, index int) []R) []R {
	return lo.FlatMap(collection, iteratee)
}
