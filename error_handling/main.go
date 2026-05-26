package main

import "fmt"



func divide(a,b float64)(result float64,err error) {
if (b==0){
	err=fmt.Errorf("cannot divide by zero")
	return 0,err
}
return a/b,nil
}

func main() {
	
	res,err:= divide(10,0)
	if (err!=nil){
	fmt.Println("Error:",err)
	}
	fmt.Println("the result is",res)
}