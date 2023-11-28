package godash_test

import (
	"fmt"
	"slices"

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

func ExampleGroupBy() {
	type User struct {
		id   int
		name string
	}
	users := []User{
		{1, "foo"},
		{2, "foobar"},
		{3, "foobar"},
		{4, "foo"},
		{5, "bar"},
	}
	fmt.Println(godash.GroupBy(users, func(u User) string {
		return u.name
	}))
	// Output: map[bar:[{5 bar}] foo:[{1 foo} {4 foo}] foobar:[{2 foobar} {3 foobar}]]
}

func ExampleChunk() {
	integers := []int{0, 1, 1, 3, 2, 1, 1, 0}
	fmt.Println(godash.Chunk(integers, 3))
	// Output: [[0 1 1] [3 2 1] [1 0]]
}

func ExamplePartitionBy() {
	type User struct {
		id   int
		name string
	}
	users := []User{
		{1, "foo"},
		{2, "foobar"},
		{3, "foobar"},
		{4, "foo"},
		{5, "bar"},
	}
	fmt.Println(godash.PartitionBy(users, func(u User) string {
		return u.name
	}))
	// Output: [[{1 foo} {4 foo}] [{2 foobar} {3 foobar}] [{5 bar}]]
}

func ExampleFlatten() {
	integers := [][]int{
		{1, 2, 5, 10},
		{9, 8, 1},
		{0},
		{42, 3},
	}
	fmt.Println(godash.Flatten(integers))
	// Output: [1 2 5 10 9 8 1 0 42 3]
}

func ExampleKeyBy() {
	type User struct {
		id   int
		name string
	}
	users := []User{
		{1, "foo"},
		{2, "foobar"},
		{3, "foobar"},
		{4, "foo"},
		{5, "bar"},
	}
	fmt.Println(godash.KeyBy(users, func(u User) int {
		return u.id
	}))
	// Output: map[1:{1 foo} 2:{2 foobar} 3:{3 foobar} 4:{4 foo} 5:{5 bar}]
}

func ExampleKeys() {
	m := map[int]string{
		1: "foo",
		2: "bar",
		3: "baz",
	}
	keys := godash.Keys(m)
	slices.Sort(keys)
	fmt.Println(keys)
	// Output: [1 2 3]
}

func ExampleValues() {
	m := map[int]string{
		1: "foo",
		2: "bar",
		3: "baz",
	}
	values := godash.Values(m)
	slices.Sort(values)
	fmt.Println(values)
	// Output: [bar baz foo]
}

func ExampleSum() {
	numbers := []float32{1, 5, 2, 3.2, 42}
	fmt.Println(godash.Sum(numbers))
	// Output: 53.2
}

func ExampleSumBy() {
	type User struct {
		name string
		age  int
	}
	users := []User{
		{"foo", 20},
		{"bar", 25},
		{"baz", 40},
	}
	fmt.Println(godash.SumBy(users, func(u User) int {
		return u.age
	}))
	// Output: 85
}

func ExampleEveryBy() {
	type User struct {
		name string
		age  int
	}
	users := []User{
		{"foo", 20},
		{"bar", 25},
		{"baz", 40},
	}
	fmt.Println(godash.EveryBy(users, func(u User) bool {
		return u.age >= 20
	}))
	// Output: true
}

func ExampleSomeBy() {
	type User struct {
		name string
		age  int
	}
	users := []User{
		{"foo", 20},
		{"bar", 25},
		{"baz", 40},
	}
	fmt.Println(godash.SomeBy(users, func(u User) bool {
		return u.age >= 30
	}))
	// Output: true
}

func ExampleNoneBy() {
	type User struct {
		name string
		age  int
	}
	users := []User{
		{"foo", 20},
		{"bar", 25},
		{"baz", 40},
	}
	fmt.Println(godash.NoneBy(users, func(u User) bool {
		return u.age > 40
	}))
	// Output: true
}

func ExampleFind() {
	type User struct {
		name string
		age  int
	}
	users := []User{
		{"foo", 20},
		{"bar", 25},
		{"baz", 40},
	}
	fmt.Println(godash.Find(users, func(u User) bool {
		return u.age > 24
	}))
	// Output: {bar 25} true
}

func ExampleAssociate() {
	type User struct {
		name string
		age  int
	}
	users := []User{
		{"foo", 20},
		{"bar", 25},
		{"baz", 40},
	}
	fmt.Println(godash.Associate(users, func(u User) (string, int) {
		return u.name, u.age
	}))
	// Output: map[bar:25 baz:40 foo:20]
}

func ExampleFlatMap() {
	integers := []int{0, 1, 2, 3}
	fmt.Println(godash.FlatMap(integers, func(x int, _ int) []int {
		return []int{2 * x, 2*x + 1}
	}))
	// Output: [0 1 2 3 4 5 6 7]
}
