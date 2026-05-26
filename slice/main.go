package main
import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5} // Creating a slice with initial values
	fmt.Println(numbers) // Output: [1 2 3 4 5]

	numbers = append(numbers, 6) // Appending a new element to the slice
	fmt.Println(numbers) // Output: [1 2 3 4 5 6]

	subSlice := numbers[1:4] // Creating a sub-slice from index 1 to 3 (4 is exclusive)
	fmt.Println(subSlice) // Output: [2 3 4]

	fmt.Printf("Length of numbers: %d\n", len(numbers)) // Output: Length of numbers: 6
	fmt.Printf("Capacity of numbers: %d\n", cap(numbers)) // Output: Capacity of numbers: 6

	// Modifying the original slice will affect the sub-slice since they share the same underlying array
	numbers[2] = 10
	fmt.Println(numbers) // Output: [1 2 10 4 5 6]
	fmt.Println(subSlice) // Output: [2 10 4] the change in the original slice is reflected in the sub-slice

	// using make function to create a slice with a specific length and capacity
	mySlice := make([]int, 3, 5)
	fmt.Println(mySlice) // Output: [0 0 0]
	fmt.Printf("Length of mySlice: %d\n", len(mySlice)) // Output: Length of mySlice: 3
	fmt.Printf("Capacity of mySlice: %d\n", cap(mySlice)) // Output: Capacity of mySlice: 5

	// appending elements to mySlice
	mySlice = append(mySlice, 1, 2)
	fmt.Println(mySlice) // Output: [0 0 0 1 2]
	fmt.Printf("Length of mySlice after appending: %d\n", len(mySlice)) // Output: Length of mySlice after appending: 5
	fmt.Printf("Capacity of mySlice after appending: %d\n", cap(mySlice)) // Output: Capacity of mySlice after appending: 5

	// appending more elements to mySlice will cause it to grow beyond its initial capacity
	mySlice = append(mySlice, 3)
	fmt.Println(mySlice) // Output: [0 0 0 1 2 3]
	fmt.Printf("Length of mySlice after appending more elements: %d\n", len(mySlice)) // Output: Length of mySlice after appending more elements: 6
	fmt.Printf("Capacity of mySlice after appending more elements: %d\n", cap(mySlice)) // Output: Capacity of mySlice after appending more elements: 10 (the capacity is doubled when the slice grows beyond its current capacity)

	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 6)
	mySlice = append(mySlice, 7)
	mySlice = append(mySlice, 8)
	mySlice = append(mySlice, 9)
	fmt.Println(mySlice) // Output: [0 0 0 1 2 3 3 6 7 8 9]
	fmt.Printf("Length of mySlice after appending more elements: %d\n", len(mySlice)) // Output: Length of mySlice after appending more elements: 11
	fmt.Printf("Capacity of mySlice after appending more elements: %d\n", cap(mySlice)) // Output: Capacity of mySlice after appending more elements: 20 (the capacity is doubled again when the slice grows beyond its current capacity)

	stock := make([]string, 0) // Creating an empty slice of strings
	fmt.Println(stock) // Output: []
	fmt.Printf("Length of stock: %d\n", len(stock)) // Output: Length of stock: 0
	fmt.Printf("Capacity of stock: %d\n", cap(stock)) // Output: Capacity of stock: 0
}