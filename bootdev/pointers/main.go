package main

import "fmt"

//                                ---
// x := 2           => In memory | 2 | ---> has own memory location
//                                ---
// x = 5            => In memory | 5 | ---> has own memory location
//                                ---
// y = x            => In memory | 5 | ---> has own memory location
//                                ---

// in programming, we can have the approach above where every variable
// has its own memory location, and we can have multiple variables
// pointing to the same memory location

// in go, variables are passed as values, not as references

func main() {
	x := 2

	y := &x

	*y = 5

	// when I change the value on the y pointer that points to x
	// both values change because x points to x and y basically
	// points to x as well

	fmt.Println(x, " ", *y)

}
