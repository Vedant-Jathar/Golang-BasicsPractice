package main

import "fmt"

func main() {
	var arr [5]float32

	arr[0] = 7.8
	arr[2] = 7
	arr[4] = 7.8

	var two_d_arr = [3][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(arr)
	fmt.Println(len(arr))

	fmt.Println(two_d_arr)
}
