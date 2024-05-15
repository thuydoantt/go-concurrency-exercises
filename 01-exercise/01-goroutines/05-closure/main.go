package main

import (
	"fmt"
	"sync"
)

//TODO: run the program and check that variable i
// was pinned for access from goroutine even after
// enclosing function returns.

func main() {
	var wg sync.WaitGroup

	incr := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Printf("value of i: %v\n", i)
		}()
		fmt.Println("return from function")
		return
	}

	incr(&wg)
	wg.Wait()
	fmt.Println("done..")
}

// return from function
// value of i: 1
// done..

// bt khi function return , all the variables inside function are gone as well, but here the runtime clever engouh to know goroutine
//  still refers to this variable so the runtime copy the variable to heap , so goroutine can access this variable
