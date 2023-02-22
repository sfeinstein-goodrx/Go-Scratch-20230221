package main

import "fmt"

func main() {
	var i any // using "any" generally means doing something wrong

	// any is basically the empty interface

	i = 7
	fmt.Println(i)

	// i += 1 // won't compile, since type is any, must convert to number first

	i = "Hi"
	fmt.Println(i)

	s := i.(string) // type assertion
	fmt.Println("s:", s)
	fmt.Printf("T:%T\n", s)

	// s = i.(int) // panic
	// solve with:
	n, ok := i.(int)
	if ok {
		fmt.Println("type assertion succeeded:", n)
	} else {
		fmt.Println("type assertion failed")
	}
}
