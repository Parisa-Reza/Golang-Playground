package main

import (
	"fmt"
	"io"
	"os"
	"io/ioutil"
)

func main() {
	// Create a new file

	file, err :=os.Create("exampleFile.txt") // this will create a new file named exampleFile.txt in the current directory. If the file already exists, it will be truncated (emptied) and overwritten.
	// this returns a file pointer and an error, we can ignore the error here for simplicity

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // this will ensure that the file is closed when the main function exits



	// write to the file

    content := "Hello world! This is a file handling example in Go."
	Byte, err :=io.WriteString(file,content + "\n") // this will return the number of bytes written and an error,
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Printf("Successfully wrote %d bytes to the file.\n", Byte)



	// read from the file using buffer

	file,errr := os.Open("exampleFile.txt") // this will open the file for reading, it returns a file pointer and an error
	if errr != nil {
		fmt.Println("Error opening file:", errr)
		return
	}
	defer file.Close() // this will ensure that the file is closed when the main function exits

	// create a buffer to read the file content
	buffer := make([]byte, 1024)

	for{
		n,err :=file.Read(buffer) // here n is the number of bytes read, and err is the error that occurred during reading
		if err != nil {
			if err == io.EOF { // this means we have reached the end of the file
				break
			}
			fmt.Println("Error reading file:", err)
			return
		}

		fmt.Println(string(buffer[:n])) // this will print the content read from the file, we use buffer[:n] to get only the bytes that were read
	}



	// read from the file using ioutil (deprecated in Go 1.16, but still commonly used)

    contentt, errr := ioutil.ReadFile("exampleFile.txt")
    if errr != nil {
	fmt.Println("Error reading file:", errr)
	return
}
	fmt.Println(string(contentt)) // this will print the content of the file as a string because ioutil.ReadFile returns a byte slice, we need to convert it to a string before printing



}