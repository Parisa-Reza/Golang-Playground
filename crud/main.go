package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getTodo() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error occurred while making the request", err)
		return
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("request failed with status code", res.StatusCode)
		return
	}

	// read the response body using ioutil.ReadAll

	// data,errr := ioutil.ReadAll(res.Body) //returns byte slice and error
	// if errr != nil {
	// 	fmt.Println("error occurred while reading the response body",errr)
	// 	return
	// }
	// fmt.Println("response data:",string(data)) // output : response data: { "userId": 1, "id": 1, "title": "delectus aut autem", "completed": false }
	// fmt.Printf("type of data: %T\n",data) // output : type of data: []uint8
	// fmt.Printf("type of res: %T\n",res) // output : type of res: *http.Response

	// standard way to read the response body using json package using structs

	var todo Todo
	errorrs := json.NewDecoder(res.Body).Decode(&todo) // here Decode method takes a pointer to the struct variable and decodes the JSON response into the struct fields based on the json tags
	if errorrs != nil {
		fmt.Println("error occurred while decoding the response body", errorrs)
		return
	}
	fmt.Println("Response from GET request:", todo)

	// Step 1 — json.NewDecoder(res.Body)
	// Creates a decoder that reads directly from the response body stream. Instead of first reading all the bytes with io.ReadAll and then unmarshalling, it reads and decodes at the same time.
	
	// Takes → res.Body (the raw stream from the API response)
	// Returns → a Decoder object
	
	// Step 2 — .Decode(&todo)
	// Tells the decoder to convert the JSON it is reading into your struct and store the result inside todo. The & is needed because Decode has to write into your variable, so it needs the address not the value.
	
	// Takes → a pointer to your struct &todo
	// Returns → error
}

func postTodo() {

	todo := Todo{
		UserId:    45,
		Title:     "hghdfkh aut",
		Completed: true,
	}
	// convert the struct to json using json.Marshal

	jsonData, err := json.Marshal(todo) // returns byte slice and error
	if err != nil {
		fmt.Println("Error while Marshal", err)
	}

	// convering Byte slice to string

	jsonString := string(jsonData)

	// converting string to reader

	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	// making the POST request using http.Post
	response, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error while making the POST request", err)
		return
	}

	// defer closing the response body to prevent resource leaks
	defer response.Body.Close()

	// read the response body using ioutil.ReadAll
	ress, err := ioutil.ReadAll(response.Body) // returns byte slice and error
	if err != nil {
		fmt.Println("Error while reading the response body", err)
		return
	}
	fmt.Println("Response from POST request:", string(ress)) // output : Response from POST request: { "userId": 45, "id": 201, "title": "hghdfkh aut", "completed": true }

}

func updateTodo() {
	todo := Todo{
		UserId:    99999,
		Title:     "Parisa Rezaei",
		Completed: false,
	}

	// converting struct to json
	byte,err := json.Marshal(todo) // returns byte slice and error
	if err != nil {
		fmt.Println("Error while Marshal", err)
		return
	}

	// converting byte slice to string
	jsonString := string(byte)

	// converting string to reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos/34"

	//creating a new request using http.NewRequest for PUT method

	req,err := http.NewRequest(http.MethodPut,myUrl,jsonReader)

	if err != nil {
		fmt.Println("Error while creating the request", err)
		return
	}

	// setting the content type header to application/json
	req.Header.Set("Content-Type","application/json")

	// creating a new http client and sending the request using client.Do(req)

	client := http.Client{}
	responsee, errr := client.Do(req) // returns response and error
	if errr != nil {
		fmt.Println("Error while sending the request", errr)
		return
	}

	// defer closing the response body to prevent resource leaks
	defer responsee.Body.Close()

	// reading the response body using ioutil.ReadAll

	bytee,errrr := ioutil.ReadAll(responsee.Body) // returns byte slice and error
	if errrr != nil {
		fmt.Println("Error while reading the response body", err)
		return
	}
	fmt.Println("Response from PUT request:", string(bytee)) // output : Response from PUT request: { "userId": 99999, "id": 34, "title": "Parisa Rezaei", "completed": false }

fmt.Println("Status code from PUT request:", responsee.StatusCode) // output : Status code from PUT request: 200

}

func deleteTodo() {
	myUrl := "https://jsonplaceholder.typicode.com/todos/34"

	req,err := http.NewRequest(http.MethodDelete,myUrl,nil) // returns request and error
	if err != nil {
		fmt.Println("Error while creating the request", err)
		return
	}

	client := http.Client{}
	ress,errr :=  client.Do(req) // returns response and error
	if errr != nil {
		fmt.Println("Error while sending the request", errr)
		return
	}

	defer ress.Body.Close()

	fmt.Println("Status code from DELETE request:", ress.StatusCode) // output : Status code from DELETE request: 200
}


func main() {

	getTodo() // unmarshal
	postTodo() // marshal
	updateTodo() 
	deleteTodo()

}
