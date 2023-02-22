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
}

// Max
// square brackets for "type constraints"
func Max[T int | float64](values []T) (T, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("Max on empty slice")
	}
	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}
	return m, nil
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
