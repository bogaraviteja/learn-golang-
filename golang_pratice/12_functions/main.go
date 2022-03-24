package main

import "fmt"

func plus(a int, b int) int { // function with 2 parameters

	return a + b
}

func plusPlus(a, b, c int) int { // function with 3 parameters
	return a + b + c
}

func main() {

	res := plus(1, 2) // function call
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3) // function call
	fmt.Println("1+2+3 =", res)
}
