package main

import "fmt"

func counter() func() int {

	count := 20

	return func() int {
		count += 1
		return count
	}
}

func main() {
	fn := counter()
	fmt.Println(fn())
	fmt.Println(fn())
}
