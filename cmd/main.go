package main

import (
	"fmt"
	"json-to-db/pkg"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	//without go routine
	// pkg.LoadVotes()

	//use go routine
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		pkg.LoadVotes()
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("%v \n", time.Since(start))
}
