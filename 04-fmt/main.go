package main

import "fmt"

type ID int

var (
	id ID      = 1
	e  float64 = 1.2
)

func main() {
	fmt.Printf("The type of ID is %T\n", id)
	fmt.Printf("The value of ID is %v\n", e)
	fmt.Printf("The type of E is %T\n", e)
}
