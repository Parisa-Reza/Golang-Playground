package main
import "fmt"

func main() {

	type Person struct{
		name string
		age int
	}

	// first way to initialize a struct
	var person1 Person
	fmt.Println(person1) // Output: { 0} here the fields of the struct are initialized with their zero values, which is an empty string for name and 0 for age
    person1.name = "Alice"
	person1.age = 30
	fmt.Println(person1) // output: {Alice 30}


	// second way to initialize a struct
	person2 := Person{
		name: "Bob",
		age: 40,
	}
	fmt.Println(person2) // Output: {Bob 40}

    // nested struct

	type Address struct{
		city string
		country string
	}
	
	type Contact struct{	
		phone int
		email string
	}

	type Employee struct{
		employee_info Person
		employee_address Address
		employee_contact Contact
	}
   
	employee1 := Employee{
		employee_info: Person{
			name:"Parisa Reza",
			age: 25,

		},
		employee_address: Address{
			city : "Dhaka",
			country: "Bangladesh",
		},
		employee_contact: Contact{
			phone: 1234567890,
			email: "parisa.reza@example.com",
		},
	}
	
	fmt.Println(employee1) // Output: {{Parisa Reza 25} {Dhaka Bangladesh} {1234567890 parisa.reza@example.com}}
	fmt.Println(employee1.employee_address.city) // Output: Dhaka

	// creating a struct with new keyword

	type Student struct{
		name string
		age int
	}

	student1 := new(Student)
	fmt.Println(student1) // Output: &{ 0} here the fields of the struct are initialized with their zero values, which is an empty string for name and 0 for age
	student1.name = "lie"
	student1.age = 20
	fmt.Println(student1) // Output: &{Charlie 20} here we can see that student1 is a pointer to the struct, so we need to use * to access the fields of the struct
    fmt.Println(*student1) // Output: {Charlie 20} here we are dereferencing the pointer to access the fields of the struct
}
