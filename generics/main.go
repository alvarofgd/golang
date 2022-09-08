package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

func main() {

	arr := []string{"uno", "dos", "tres", "cuatro", "cinco", "seis"}
	arr2 := []string{}

	r := mapping(arr, func(s string) int { return len(s) })
	fmt.Println(r)

	s := mapping(r, func(n int) string {
		str := ""
		for i := 0; i < n; i++ {
			str += "*"
		}
		return str
	})
	fmt.Println(s)

	r2 := mapping(arr2, func(s string) int { return len(s) })
	fmt.Println(r2)

	r3 := folding(r, 0, func(i int, n int) int {
		return i + n
	})
	fmt.Println(r3)

	m := folding(arr, make(map[string]int), func(m map[string]int, s string) map[string]int {
		m[s] = len(s)
		return m
	})
	fmt.Println(m)

	f := filtering(arr, func(s string) bool {
		return len(s)%2 == 0
	})
	fmt.Println(f)
	fmt.Println(arr)
}

type Predicate[T any] func(t T) bool

func filtering[T any](ts []T, f Predicate[T]) []T {

	return folding(ts, []T{}, func(ts []T, t T) []T {

		if f(t) {
			return append(ts, t)
		}

		return ts
	})
}

type Mapper[T any, R any] func(t T) R

func mapping[T any, R any](ts []T, f Mapper[T, R]) []R {

	return folding(ts, []R{}, func(rs []R, t T) []R {

		return append(rs, f(t))
	})
}

type Accumulator[T any, R any] func(r R, t T) R

func folding[T any, R any](ts []T, acc R, f Accumulator[T, R]) R {

	res := acc
	for _, v := range ts {
		res = f(res, v)
	}

	return res
}
