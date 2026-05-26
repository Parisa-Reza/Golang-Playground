package main
import "fmt"

func main(){
	fmt.Println("Hello world!!!")
	 age := 50
	 bankBalance := 1000.50
	 name := "Parisa"
	 fmt.Printf("My age is %d\n", age)
	 fmt.Printf("My bank balance is %.2f\n", bankBalance)
	 fmt.Printf("My name is %s\n", name)
	 fmt.Printf("My age is %d and my bank balance is %.2f\n", age, bankBalance)
	 // here without /n the next line will be in the same line
	 fmt.Printf("Type of age is %T\n", age)
}

// for running this file in the terminal, use the command: go run print/main.go