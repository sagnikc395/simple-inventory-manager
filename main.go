package main

import (
	"fmt"
	"os"
)

type Inventory struct {
	Name     string
	Quantity int
	Price    float64
}

func main() {
	for {
		fmt.Println("Welcome to simple inventory CLI!")
		fmt.Println("Enter 'exit' to quit, or anything else to continue.")

		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			fmt.Println("Exiting the application...")
			os.Exit(0)
		} else {
			fmt.Println("You have entered:", input)
		}
	}

}
