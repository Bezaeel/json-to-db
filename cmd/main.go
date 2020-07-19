package main

import (
	"fmt"
	"json-to-db/util"
	// "sync"
	"time"
)

func main(){
	votes := make([]map[string]interface{}, 0)
	start := time.Now()
	votes = util.ToJsonFromFile(".././example_1.json")
	//use go routine
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func ()  {
	// 	votes = util.ToJsonFromFile(".././example_1.json")
	// 	wg.Done()
	// }()
	
	// wg.Wait()
	fmt.Printf("%v \n", time.Since(start))
	for _, v := range votes{
		fmt.Println(v)
	}

}