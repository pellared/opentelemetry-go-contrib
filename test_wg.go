package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Go(func() {
		fmt.Println("Inside Go")
	})
	wg.Wait()
	fmt.Println("Done")
}
