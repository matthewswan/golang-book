package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	fmt.Println("The Fahrenheit degree is:", input)
	output := ((input - 32) * 5 / 9)
	fmt.Println("The Celsius degree is:", output)
}
