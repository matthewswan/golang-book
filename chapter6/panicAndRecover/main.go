package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

func recover() string {
	str := "We made it to the RECOVERED function!!!!!"
	return str
}
