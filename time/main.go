package main
import(
	"fmt"
	"time"
) 
// go's time format is 2006-01-02 15:04:05 Monday
func main() {

	currentTime := time.Now()
	fmt.Println(currentTime) // 24hrs format

	// Formatting time
	formattedTime := currentTime.Format("02/01/2006 03:04:05 PM") // in 12hrs format
	fmt.Println(formattedTime)

	formattedTime1 := currentTime.Format("02/01/2006 03:04:05 PM Monday") // in 12hrs format
	fmt.Println(formattedTime1)

	// Parsing time
	layoutSring := "02/01/2006"
	dateString := "25/12/2024"

	formattedDate, _ := time.Parse(layoutSring, dateString) // parse returns a time.Time and an error, we can ignore the error here for simplicity
	fmt.Println(formattedDate) // 2024-12-25 00:00:00 +0000 UTC

	// adding 3 days to the current time

	newDate := currentTime.Add(24 * time.Hour * 3) // 24 hours in a day, so 24*3 for 3 days
	fmt.Println(newDate)
	formattedNewDate := newDate.Format("02/01/2006 03:04:05 PM Monday")
	fmt.Println(formattedNewDate)
}
