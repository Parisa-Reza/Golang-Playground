package main

import (
	"fmt"
	"time"
)

func func1() {
	fmt.Println("Hello, World 1")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("End of Hello, World 1")

}

func func2() {
	fmt.Println("Hello, World 2")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("End of Hello, World 2")
}


func main () {
fmt.Println("Hello, World ")
 go func1()
  go func2()
fmt.Println("Hello, World 3")
time.Sleep(500 * time.Millisecond) // to prevent the main function from exiting before the goroutine finishes executing
}
	