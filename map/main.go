package main
import "fmt"

func main() {
	grades := make (map[string]int)
	grades["Alice"] = 85
	grades["Bob"] = 90
	grades["Charlie"] = 78
	grades["David"] = 0

	fmt.Println(grades) // Output: map[Alice:85 Bob:90 Charlie:78]

	// Accessing values
	fmt.Printf("Alice's grade: %d\n", grades["Alice"]) // Output: Alice's grade: 85

	grades["Alice"] = 88 // Updating Alice's grade
	fmt.Printf("Alice's updated grade: %d\n", grades["Alice"]) // Output: Alice's updated grade: 88

	// Deleting an entry
	delete(grades, "Charlie")
	fmt.Println(grades["Charlie"]) // Output: 0 (zero value for int) // here charlie is not present in the map so it returns the zero value for int which is 0

	fmt.Println(grades["David"]) // Output: 0 (zero value for int) here david is present in the map but his value is 0

	// Checking if a key exists
	grade, exists := grades["David"]
	fmt.Println(grade, exists) // Output: 0 true
	grade, exists = grades["Charlie"]
	fmt.Println(grade, exists) // Output: 0 false

	for key, value := range grades {
		fmt.Println("Key:",key,"Value:",value)
	}

	// here we are creating a new map with some initial values and iterating over it using range to print the key-value pairs
	numbers := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	for key, value := range numbers {
		fmt.Println("Key:",key,"Value:",value)
	}

// Key: three Value: 3
// Key: one Value: 1
// Key: two Value: 2

// maps do not maintain the order of keys, so the output may vary each time you run the program

}