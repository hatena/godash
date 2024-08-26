//go:build !go1.23

package godash

import (
	"github.com/samber/lo"
)

// Chunk receives the collection and chunks it into a slice of slices each of the given size.
func Chunk[T any](collection []T, size int) [][]T {
	return lo.Chunk(collection, size)
}

// Keys returns the keys of the map as a slice.
func Keys[K comparable, V any](in map[K]V) []K {
	return lo.Keys(in)
}

// Values returns the values of the map as a slice.
func Values[K comparable, V any](in map[K]V) []V {
	return lo.Values(in)
}
