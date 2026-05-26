package main
import "fmt"
import "goPlayGround/util" // here goPlayGround is the name of the module in the go.mod file
func main (){
	fmt.Println("Hello world!!!")
	util.PrintMessage("I am a util")
	//util.displayMessage("I am a second util") // this will give an error because displayMessage is not exported

	var age int=50
	fmt.Println("My age is", age)
	const pi float64=3.14
	fmt.Println("The value of pi is", pi)

	var name string = "Parisa"
	fmt.Println("My name is", name)

	country := "Bangladesh"
	fmt.Println("My address is", country)

	privateVar := "This is a private variable" // this variable can not be exported because it starts with a small letter
	fmt.Println(privateVar)

	PublicVar := "This is a public variable" // this variable can be exported because it starts with a capital letter
	fmt.Println(PublicVar)
}