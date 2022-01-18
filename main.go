package main

import (
	"fmt"
	"github.com/pkg/browser"
)

func main(){
	fmt.Printf("Hello World")
	var url string
	fmt.Printf("Enter your URL: ")
	fmt.Scan(&url)

	fmt.Printf("Opening %s ...\n", url)

	err := browser.OpenURL(url)
	if err != nil {
		fmt.Printf("Error occured")
		fmt.Print(err)
	} else {
		fmt.Printf("Successfully opened %s in the browsergo.", url)
	}

}
