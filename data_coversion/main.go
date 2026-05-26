package main
import (
	"fmt"
	"strconv"
)

func main() {

	num := 10
	strNum := strconv.Itoa(num) // converting int to string
	fmt.Printf("num var is of type %T and value is %d\n", num, num) // Output: num var is of type int and value is 10
	fmt.Printf("strNum var is of type %T and value is %s\n", strNum, strNum) // Output: strNum var is of type string and value is 10

	str := "20"
    num2,_ := strconv.Atoi(str) // converting string to int here Atoi returns two values, the converted integer and an error. We are ignoring the error using the blank identifier _ since we are sure that the conversion will succeed.
	fmt.Printf("str var is of type %T and value is %s\n", str, str) // Output: str var is of type string and value is 20
	fmt.Printf("num2 var is of type %T and value is %d\n", num2, num2) // Output: num2 var is of type int and value is 20

	pi := "3.14"
	piFloat, _ := strconv.ParseFloat(pi, 64) // converting string to float64 here ParseFloat returns two values, the converted float and an error. We are ignoring the error using the blank identifier _ since we are sure that the conversion will succeed.
	fmt.Printf("pi var is of type %T and value is %s\n", pi, pi) // Output: pi var is of type string and value is 3.14
	fmt.Printf("piFloat var is of type %T and value is %f\n", piFloat, piFloat) // Output: piFloat var is of type float64 and value is 3.140000
}
