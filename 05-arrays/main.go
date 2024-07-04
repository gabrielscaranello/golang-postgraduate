package main

import "fmt"

func main() {

	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	// get last item in array
	println(arr[len(arr)-1])

	// for i: index, v: value in array
	for i, v := range arr {
		fmt.Printf("The value of index %d is %d\n", i, v)
	}
}
