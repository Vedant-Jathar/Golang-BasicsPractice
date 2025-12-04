package main

import (
	"fmt"
	"sync"
	// "time"
)

func task(i int) {
	fmt.Println(i)
}

func main() {
	var wg sync.WaitGroup
	
	for i := range 5 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
