package main

import (
	"fmt"
	"sort"
)

func main() {
	var s1 []int
	fmt.Println("s1:", s1)
	fmt.Printf("s1: %#v of %T\n", s1, s1)
	fmt.Println("len:", len(s1))

	cart := []string{
		"apples",
		"oranges",
		"bananas",
	}
	for i := range cart { // indices
		fmt.Println(i)
	}
	for i, v := range cart { // index + value
		fmt.Println(i, v)
	}
	for _, v := range cart { // values
		fmt.Println(v)
	}

	s2 := []int{1, 2, 3}
	fmt.Println("len:", len(s2), "cap:", cap(s2))
	s2 = append(s2, 4)
	fmt.Println("len:", len(s2), "cap:", cap(s2))

	var s3 []int
	for i := 0; i < 1000; i++ {
		s3 = appendInt(s3, i)
	}
	fmt.Println(s3[:10], len(s3))
	s3[100] = -1

	// will panic with index out of range, since Go does bounds checking
	// doesn't matter that underlying cap is greater than 1000
	// s3[1001] = -1

	vals := []float64{3, 1, 2}
	fmt.Println(median(vals))
	vals = []float64{3, 1, 2, 4}
	fmt.Println(median(vals))
	fmt.Println(vals)
}

func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}

	// slices are always passed by value but contain a pointer to the data (e.g. in slice impl you'll see "array" field is "unsafe.Pointer"
	// this is a reference to the input and so is sorting the input in place;  will be changed in the caller's scope

	// so we updated to copy instead
	vals := make([]float64, len(values))
	copy(vals, values)

	sort.Float64s(vals)
	i := len(vals) / 2
	if len(vals)%2 == 1 {
		return vals[i], nil
	}
	v := (vals[i-1] + vals[i]) / 2
	return v, nil
}

func appendInt(vals []int, v int) []int {
	i := len(vals)
	if len(vals) < cap(vals) { // we have enough space in the underlying array
		vals = vals[:len(vals)+1]
	} else { // need to re-allocate & copy values
		size := 2 * (len(vals) + 1)
		fmt.Println(len(vals), "->", size)
		s := make([]int, size)
		copy(s, vals) // never fails, copies as much as it can
		vals = s[:len(vals)+1]
	}
	vals[i] = v
	return vals
}
