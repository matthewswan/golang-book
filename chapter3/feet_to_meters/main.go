package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	fmt.Println("The amount in feet:", input)
	output := input * 0.3048
	fmt.Println("The amount in meters:", output)
}
