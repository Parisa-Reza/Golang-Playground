package main

import "fmt"

func main() {
	var arr [5]int
	arr[0]=20;
	arr[1]=30;
	fmt.Println(arr) // Output: [20 30 0 0 0] index 0 and 1 are assigned values, the rest are default 0
	 
	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2) // Output: [1 2 3 4 5]
	
	var arr3 = [...]int{10, 20, 30}
	fmt.Println(arr3) // Output: [10 20 30] the size of the array is determined by the number of elements provided

	var arr4 [5]string
	fmt.Println(arr4) // Output: [     ] index 0 to 4 are default empty strings

	fmt.Printf("Length of arr4: %d\n", len(arr4)) // Output: Length of arr4: 5
	fmt.Printf("Values of arr4: %q\n", arr4)
	var arr5  [3]bool
	fmt.Println(arr5) // Output: [false false false] index 0 to 2 are default false
}

// the array has a fixed size and all elements must be of the same type. The size of the array is determined at compile time and cannot be changed at runtime. Arrays are value types, which means that when you assign an array to another variable or pass it to a function, a copy of the entire array is made. This can lead to performance issues if the array is large, so it's often recommended to use slices instead of arrays in Go for more flexible and efficient data structures.