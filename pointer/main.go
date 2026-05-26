package main
import "fmt"

func modifyValue(numm *int){ // here (numm *int) means 
*numm = *numm +5
}

func main() {
	num := 10
	ptr := &num
	fmt.Println("Value of num:", num) // Output: Value of num: 10
	fmt.Println("Address of num:", &num) // Output: Address of num: 0xc0000160b8 (this will vary each time you run the program)
	fmt.Println("Value of ptr:", ptr) // Output: Value of ptr: 0xc0000160b8 (this will vary each time you run the program)
	fmt.Println("Value pointed to by ptr:", *ptr) // Output: Value pointed to by ptr: 10

	modifyValue(&num)
	fmt.Println("Value of num after modification:", num) // Output: Value of num after modification: 15
}

// pointer stores nil if it is not initialized
