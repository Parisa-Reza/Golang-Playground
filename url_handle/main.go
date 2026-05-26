package main
import (
	"fmt"
	"net/url"
)

func main() {
	myUrl := "https://example.com:8080/products/search?q=laptop&page=2"
    fmt.Printf("type of myUrl: %T\n", myUrl) //type of myUrl: string

	// parsing the url
	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Printf("type of parsedUrl: %T\n", parsedUrl) //type of parsedUrl: *url.URL

	// accessing different parts of the URL
	fmt.Println("Scheme:", parsedUrl.Scheme) //Scheme: https
	fmt.Println("Host:", parsedUrl.Host) //Host: example.com:8080
	fmt.Println("Path:", parsedUrl.Path) //Path: /products/search
	fmt.Println("RawQuery:", parsedUrl.RawQuery) //RawQuery: q=laptop&page=2

	// modifying the URL
	parsedUrl.Path = "/products/new"
	newUrl := parsedUrl.String() // here String() method is used to convert the URL object back to a string
	fmt.Println("Modified URL:", newUrl) //Modified URL: https://example.com:8080/products/new?q=laptop&page=2
}