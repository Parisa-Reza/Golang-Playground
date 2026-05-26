package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	person1 := Person{
		Name: "Alice",
		Age:  30,
	}

	fmt.Println("object type", person1) // {Alice 30} this is object of type Person

	//converting struct to json (string representation of the struct)

	jsonData,err:= json.Marshal(person1) //Marshal is used to convert struct to json
	if err!=nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}
	fmt.Println("JSON Data as bytes:", jsonData) // [123 34 110 97 109 101 34 58 34 65 108 105 99 101 34 44 34 97 103 101 34 58 51 48 125] this is in bytes

	fmt.Println("JSON Data as string (Marshalled):", string(jsonData)) // {"name":"Alice","age":30} this is the json string representation of the struct

	// unmarshalling json to struct
	var person2 Person
	err = json.Unmarshal(jsonData, &person2) // here
	if err!=nil {
		fmt.Println("Error converting JSON to struct:", err)
		return
	}
	fmt.Println("Unmarshalled struct:", person2) // Unmarshalled struct: {Alice 30} this is the struct representation of the json data
}