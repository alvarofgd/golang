package main

import (
	"fmt"
	"golang/functional"
)

type Number interface {
	int64 | float64
}

func main() {

	arr := []string{"uno", "dos", "tres", "cuatro", "cinco",
		"seis", "siete", "ocho", "nueve", "diez", "once", "doce"}
	fmt.Println("arr: ", arr)

	r := functional.Map(arr, func(s string) int { return len(s) })
	fmt.Println("arr.map(string -> len(string)): ", r)

	s := functional.Map(r, func(n int) string {
		str := ""
		for i := 0; i < n; i++ {
			str += "*"
		}
		return str
	})
	fmt.Println("arr.map(string -> '*' x len(string)): ", s)

	arr2 := []string{}
	fmt.Println("arr2: ", arr2)

	r2 := functional.Map(arr2, func(s string) int { return len(s) })
	fmt.Println("arr2.map(string -> len(string)): ", r2)

	r3 := functional.Fold(r, 0, func(i int, n int) int { return i + n })
	fmt.Println("arr.sum(string -> len(string)): ", r3)

	m := functional.Fold(arr, make(map[string]int), func(m map[string]int, s string) map[string]int {
		m[s] = len(s)
		return m
	})
	fmt.Println("arr.groupBy(string -> (string, len(string))): ", m)

	m2 := functional.Fold(arr, make(map[int][]string), func(m map[int][]string, s string) map[int][]string {
		m[len(s)] = append(m[len(s)], s)
		return m
	})
	fmt.Println("arr.groupBy(string -> (len(string), []string)): ", m2)

	f := functional.Filter(arr, func(s string) bool {
		return len(s)%2 == 0
	})
	fmt.Println("arr.even()", f)
}
