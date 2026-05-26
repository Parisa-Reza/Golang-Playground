package main
import "fmt"

func main() {
	x:=15
	if x>10{
		fmt.Println("x is greater than 10")
	}else if x==10{
		fmt.Println("x is equal to 10")
	}else{
		fmt.Println("x is less than 10")
	}

	day := 1
	switch day {
	case 1:
		fmt.Println("It's Monday")
	case 2:
		fmt.Println("It's Tuesday")
	case 3:
		fmt.Println("It's Wednesday")
	default:
		fmt.Println("It's some other day")
	}

	// multiple cases
	grade := "B"
	switch grade {
	case "A", "A+":
		fmt.Println("Excellent")
	case "B", "B+":
		fmt.Println("Good")
	case "C", "C+":
		fmt.Println("Average")
	default:
		fmt.Println("Needs Improvement")
	}

	// expression in switch
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	// for loop with range
	numbers:= []int{10, 20, 30, 40, 50}
	for index,value := range numbers {
		fmt.Println("Index:",index,"Value:",value)
	}

	// infinite loop

	count := 0
	for {
		fmt.Println("Infinite loop iteration:")
		count++
		if count >= 5 {
			break
		}
	}
}