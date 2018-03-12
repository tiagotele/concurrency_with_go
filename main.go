package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	using_channel()
}

func using_channel(){
	number := flag.Int("number", 10, "Number as parameter")
	flag.Parse()
	fmt.Println("The argument is", *number)

	messages := make(chan int)

	for i := 0; i < *number; i++ {
		go task_with_channel(i, messages)
	}

	for i := 0; i < *number; i++ {
        (<-messages)
    }
	fmt.Println("Done")
}

func task_with_channel(taskNumber int, messages chan<- int){
	print_task_number(taskNumber)
	messages <- taskNumber
}

func using_sync(){
	number := flag.Int("number", 10, "Number as parameter")
	flag.Parse()
	fmt.Println("The argument is", *number)

	var wg sync.WaitGroup
	wg.Add(*number)

	for i := 0; i < *number; i++ {
		go task_with_sync(i, &wg)
	}

	wg.Wait()
	fmt.Println("Done")
}

func task_with_sync(taskNumber int, wg *sync.WaitGroup) {
	defer wg.Done()
	print_task_number(taskNumber)
}

func print_task_number(taskNumber int) {
	fmt.Println(fmt.Sprintf("Task number {%d}", taskNumber))
}
