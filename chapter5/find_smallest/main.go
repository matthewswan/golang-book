package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	smallest := x[0]
	for _, value := range x {
		if smallest < value {
			continue
		} else {
			smallest = value
		}

		// for i := 1; i < len(x); i++ {
		// if smallest < x[i] {
		// continue
		// } else {
		// smallest = x[i]
		// }
	}
	fmt.Println("The smallest value is:", smallest)
}
