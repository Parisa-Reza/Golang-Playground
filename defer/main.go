package main
import "fmt"

func printfunc() {
	fmt.Println("I am Parisa Reza")
}

func main() {
	// fmt.Println("1")
	// defer fmt.Println("2")
	// fmt.Println("3") // 1, 3, 2

	// fmt.Println("1")
	// defer fmt.Println("2")
	// defer fmt.Println("3")
	// fmt.Println("4") // 1, 4, 3, 2 // here defer is LIFO, so 3 will be executed before 2

	fmt.Println("Hi, there!")
	defer printfunc()
	defer fmt.Println("Let me introduce myself.")
	fmt.Println("Welcome to Go programming.")
}
