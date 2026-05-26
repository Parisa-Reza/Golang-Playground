package main
import (
  "fmt"
  "strings"
) 

func main() {
	fruits := "apple,banana,orange"
	fruit := strings.Split(fruits,",")
	fmt.Println(fruit) // [apple banana orange]

	// Join
	joinedFruits := strings.Join(fruit, " - ")
	fmt.Println(joinedFruits) // apple - banana - orange

	// Contains
	fmt.Println(strings.Contains(fruits, "banana")) // true
	fmt.Println(strings.Contains(fruits, "grape"))  // false

	message := "               hello world!               "
	fmt.Println(strings.TrimSpace(message)) // hello world!

    message1 := "               hello      world!               "
	fmt.Println(strings.TrimSpace(message1)) // hello      world!

	fullName := []string{"Parisa", "Reza"}
	fmt.Println(strings.Join(fullName, " ")) // Parisa Reza

}
