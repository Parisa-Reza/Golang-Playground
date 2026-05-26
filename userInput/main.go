package main
import( //import ("fmt" "os" "bufio")   // ❌
	"fmt"
	"os"
	"bufio"
)

func main(){   

	// func main()   // ❌
    // {

	// one word input
	fmt.Println("Hello, What is your name??")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("your name is",name)

	// A sentence input
	reader:= bufio.NewReader(os.Stdin)
	name,_ := reader.ReadString('\n') //"" and '' here are different, "" is for string and '' is for char
	fmt.Println("your name is",name)


}

// Line 1: reader := bufio.NewReader(os.Stdin)

// bufio.NewReader() creates a buffered reader
// os.Stdin is the standard input — i.e. the keyboard
// reader is now a reader that listens to keyboard input
// Buffered means it reads a chunk at a time instead of character by character — more efficient


// Line 2: name, _ := reader.ReadString('\n')

// ReadString('\n') reads everything typed until Enter is pressed
// '\n' is the newline character — the delimiter (stop point)
// Returns two values — the string read, and an error
// name stores the input
// _ discards the error — Go forces you to handle all return values, _ is how you intentionally ignore one