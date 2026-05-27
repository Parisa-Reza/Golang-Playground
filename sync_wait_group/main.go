package main
import (
	"sync"
    "fmt"
)
// we use defer in a function to ensure that a specific action is performed when the function exits, regardless of how it exits (whether it returns normally or due to an error). In this case, we use defer to call wg.Done() at the end of func1, which signals that the task is complete and allows the main function to proceed once all tasks are done.

func func1(i int, wg *sync.WaitGroup) {
	defer wg.Done() // defer the call to Done() to ensure it is called when the function exits
	fmt.Printf("start %d\n", i)
	// some task
	fmt.Printf("end %d\n", i)
}

func main() {
	var wg sync.WaitGroup // create a wait group variable
	for i := 1; i <= 3; i++ {
   
		wg.Add(1) // increment the wait group counter for each task
		go func1(i, &wg) // pass the wait group pointer to the function
	}

	wg.Wait() // wait for all tasks to complete
	fmt.Println("All tasks completed")
		
}
