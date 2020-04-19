package main

import (
	"fmt"
	"os"
	"strconv"
)


func grab_args() []string {
	if len(os.Args) < 2 {
		fmt.Printf("You did not pass any arguments!\n")
		os.Exit(1)
	}
	return os.Args[1:]
}


func parse_args() (base int, increment int) {
	args := grab_args()
	if len(args) != 2 {
		fmt.Print("Error: Expecting 2 integer arguments: base and increment.\n")
		os.Exit(1)
	}

	base, err1 := strconv.Atoi(args[0])
	if err1 != nil {
		fmt.Printf("You did not pass an integer for the base. You passed %v\n", args[0])
		panic(err1)
	}

	increment, err2 := strconv.Atoi(args[1])
	if err2 != nil {
		fmt.Printf("You did not pass an integer for the increment. You passed %v\n", args[1])
		panic(err2)
	}

	//fmt.Printf("%d %d\n", base, increment)
	return base, increment
}


func test_computation(base int, increment int) int {
	return base + increment
}


func main() {
	base, increment := parse_args()
	fmt.Printf("Base:\t\t%d\nIncrement:\t%d\n", base, increment)
	fmt.Printf("Result:\t\t%d\n", test_computation(base, increment))
}