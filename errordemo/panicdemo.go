package main

import (
	"fmt"
)

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("Firsr Name can't be null")
	}
	if lastName == nil {
		panic("Last Name can't be null")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func test() {
	defer fmt.Println("deferred call in test")
	firName := "paul"
	fullName(&firName, nil)
}

func main() {
	defer fmt.Println("deferred call in main")
	test()
	fmt.Println("returned normally from main")
}
