package functional

import (
	"testing"
)

func TestFilter(t *testing.T) {
	arr := []string{"uno", "dos", "tres"}
	f := Filter(arr, func(s string) bool {
		return len(s) == 3
	})

	if len(f) != 2 {
		t.Errorf("Expected 2, got %d", len(f))
	}

	f2 := Filter([]string{}, func(s string) bool {
		return true
	})

	if len(f2) != 0 {
		t.Errorf("Expected 0, got %d", len(f2))
	}
}

func TestFold(t *testing.T) {
	arr := []string{"uno", "dos", "tres", "cuatro", "cinco",
		"seis", "siete", "ocho", "nueve", "diez", "once", "doce"}

	m := Fold(arr, make(map[string]int), func(m map[string]int, s string) map[string]int {
		m[s] = len(s)
		return m
	})

	if elem, err := m["uno"]; !err || elem != 3 {
		t.Errorf("Expected 'uno' to be mapped to 3, got %d", elem)
	}

	m2 := Fold(arr, make(map[int][]string), func(m map[int][]string, s string) map[int][]string {
		m[len(s)] = append(m[len(s)], s)
		return m
	})

	if arr2, err := m2[4]; !err || len(arr2) != 6 {
		t.Errorf("Expected 4 to be mapped to [tres seis ocho diez once doce], got %v", m2[4])
	}
}

func TestMapFold(t *testing.T) {
	arr := []string{"uno", "dos", "tres", "cuatro", "cinco",
		"seis", "siete", "ocho", "nueve", "diez", "once", "doce"}

	lengths := Map(arr, func(s string) int { return len(s) })
	sum := Fold(lengths, 0, func(i int, n int) int { return i + n })

	if sum != 51 {
		t.Errorf("Expected 51, got %d", sum)
	}
}
