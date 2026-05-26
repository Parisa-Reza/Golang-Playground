package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
  res,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
  if err!=nil {
	fmt.Println("Error fetching API:", err)
  }
  fmt.Printf("Response: %T\n", res) //Response: *http.Response
  defer res.Body.Close()

  fmt.Printf("response: %v\n", res) //
  result,_ := ioutil.ReadAll(res.Body)
   fmt.Printf("response : %v\n", result) //the response is in bytes 
  fmt.Println("readable response :", string(result)) //the response is in bytes so we need to convert it to string
}