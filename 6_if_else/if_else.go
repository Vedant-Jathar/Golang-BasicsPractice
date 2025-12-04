package main

import "fmt"

const age = 89

func main() {
	// if age > 18 {
	// 	fmt.Println(age)
	// } else if age == 89 {
	// 	fmt.Println("Age is 89")
	// }

	if age := 30; age == 20 {
		fmt.Println(age)
	} else if age == 30 {
		fmt.Println(age)
	}

	fmt.Println(age)
}
