package main

import "fmt"

func main() {
	fmt.Println("welcome to class on pointers")
	//var ptr *int

	myNumber := 12

	var ptr = &myNumber

	fmt.Println("value of pointer is ", ptr)
	fmt.Println("value of pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is : ", myNumber)

}
