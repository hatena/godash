package godash_test

import (
	"fmt"

	"github.com/hatena/godash"
)

func ExampleFilter() {
	integers := []int{-2, 0, 19, 5, 42}
	fmt.Println(godash.Filter(integers, func(x int, _ int) bool {
		return x%2 == 0
	}))
	// Output: [-2 0 42]
}

func ExampleMap() {
	integers := []int{-2, 0, 19, 5, 42}
	fmt.Println(godash.Map(integers, func(x int, _ int) int {
		return x * 2
	}))
	// Output: [-4 0 38 10 84]
}

func ExampleReduce() {
	integers := []int{-2, 0, 19, 5, 42}
	fmt.Println(godash.Reduce(integers, func(agg int, x int, _ int) int {
		return agg + x
	}, 0))
	// Output: 64
}

func ExampleUniq() {
	integers := []int{0, 1, 1, 3, 2, 1, 1, 0}
	fmt.Println(godash.Uniq(integers))
	// Output: [0 1 3 2]
}

func ExampleUniqBy() {
	type User struct {
		id   int
		name string
	}
	users := []User{
		{1, "foo"},
		{2, "foobar"},
		{2, "foobar"},
		{3, "foo"},
	}
	fmt.Println(godash.UniqBy(users, func(u User) int {
		return u.id
	}))
	// Output: [{1 foo} {2 foobar} {3 foo}]
}
