package util

import "fmt"

func PrintMessage(message string){ // this function is exported because it starts with a capital letter
	fmt.Println(message)
}
func displayMessage(message string){ // this function is not exported because it starts with a small letter
	fmt.Println(message)
}