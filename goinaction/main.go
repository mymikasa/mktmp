package main

import "fmt"

func main() {

	slice1 := []int{0, 1, 2, 3, 4, 5}

	slice2 := slice1[2:3:3]

	fmt.Printf("%v\n", slice2)
	fmt.Printf("%v\n", slice1)

	slice2 = append(slice2, 23)
	slice2[0] = 111

	fmt.Printf("%v\n", slice2)
	fmt.Printf("%v\n", slice1)
}
