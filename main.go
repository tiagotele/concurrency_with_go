package main

import (
	"fmt"
	"flag"
)

func main(){
	
	number := flag.Int("number", 10, "Number as parameter")

	flag.Parse()
	fmt.Println("The argument is", *number)
	
	for i := 0 ; i < *number ; i++{
		go task(i)
	}

	fmt.Scanln()
    fmt.Println("Done")
}

func task(taskNumber int){
	fmt.Println(fmt.Sprintf("Task number {%d}", taskNumber))
}