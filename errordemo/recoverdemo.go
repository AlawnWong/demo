package main

import (
	"fmt"
	"runtime/debug"
)

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
		debug.PrintStack()
	}
}

func fullNameRecover(firstName *string, lastName *string) {
	defer recoverName()
	if firstName == nil {
		panic("Firsr Name can't be null")
	}
	if lastName == nil {
		panic("Last Name can't be null")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func testRecover() {
	defer fmt.Println("deferred call in test")
	firName := "paul"
	fullNameRecover(&firName, nil)
}

func main() {
	defer fmt.Println("deferred call in main")
	testRecover()
	fmt.Println("returned normally from main")
}
