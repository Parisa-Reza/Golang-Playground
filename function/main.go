package main

import "fmt"

func textDisplay() {
	fmt.Println("Hello, I am a function")
}


func add(a,b,c,d int)(result int) {
 result = a+b+c+d
 return
}

func main() {
	fmt.Println("Hello, World!")
	textDisplay()
	res:=add(1,2,3,4)
	fmt.Println("the result is",res)
}