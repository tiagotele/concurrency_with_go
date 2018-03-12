package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {

	number := flag.Int("number", 10, "Number as parameter")
	flag.Parse()
	fmt.Println("The argument is", *number)

	var wg sync.WaitGroup
	wg.Add(*number)

	for i := 0; i < *number; i++ {
		go task(i, &wg)
	}

	wg.Wait()
	fmt.Println("Done")
}

func task(taskNumber int, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println(fmt.Sprintf("Task number {%d}", taskNumber))

}
