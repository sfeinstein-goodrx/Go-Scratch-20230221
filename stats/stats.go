package main

import "fmt"

func main() {
	is := []int{2, 3, 1}
	fmt.Println(MaxInts(is))
	fmt.Println(MaxInts(nil))
	fmt.Println(Max(is))
	fmt.Println(Max[int](nil))

	fs := []float64{2, 3, 1}
	fmt.Println(MaxFloat64s(fs))
	fmt.Println(Max(fs))
	fmt.Println(Max[float64](nil))

	hs := []Height{2, 3, 1}
	fmt.Println(Max(hs))
}

// Max
// square brackets for "type constraints"
// orig:
// func Max[T int | float64](values []T) (T, error) {
func Max[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T // helper for just this situation with string, tells us the default zero value for a type
		return zero, fmt.Errorf("Max on empty slice")
		//return 0, fmt.Errorf("Max on empty slice")
	}
	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}
	return m, nil
}

type Height float64
type Ordered interface {
	//int | float64 | string
	~int | ~float64 | ~string // tilda tells us it isnt exactly that type but any type that is underneath the same
}

func MaxInts(values []int) (int, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("MaxInts on empty slice")
	}
	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}
	return m, nil
}

func MaxFloat64s(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("MaxInts on empty slice")
	}
	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}
	return m, nil
}
