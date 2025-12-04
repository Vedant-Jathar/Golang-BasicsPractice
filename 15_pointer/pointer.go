package main

import "fmt"

// func changeNumByAddress(addr *int) {
// 	*addr = 10
// 	fmt.Println("#addr", *addr)
// }

func main() {
	// num := 3
	// fmt.Println(num)
	// changeNumByAddress(&num)
	// fmt.Println(num)

	arr1 := []int{1, 2, 3}
	arr2 := arr1

	fmt.Println(&arr1 == &arr2)
	fmt.Println(&arr1)
	fmt.Println(&arr2)
}