package main

import "fmt"

func printSlice[T int | string](sl []T) {
	fmt.Println(sl)
}

type stack[T comparable, S string] struct {
	elements []T
	length   S
}

func main() {
	// printSlice([]int{1, 2, 89})
	myStack := stack[string, string]{
		elements: []string{"hi", "hello"},
		length:   "ten",
	}
	fmt.Println(myStack)
}

// exception handling
// login
// apis json response handling
// gin framework
