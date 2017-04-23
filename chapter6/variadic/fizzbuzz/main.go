package main

import "fmt"

func fizzbuzz(args []int) [6]string {
	result := [6]string{}
	for i, v := range args {
		if v%5 == 0 && v%3 == 0 {
			result[i] = "FizzBuzz"
		} else if v%5 == 0 {
			result[i] = "Buzz"
		} else if v%3 == 0 {
			result[i] = "Fizz"
		} else {
			result[i] = "NONE"
		}
	}
	return result
}

func main() {
	xs := []int{6, 23, 58, 45, 15, 3}
	fmt.Println(fizzbuzz(xs))
}
