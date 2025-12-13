package main

import "fmt"

func main() {
	var age int = 43
	fmt.Println(age)
	var pointer *int
	// star has 2 usages:-
	// one is to declare a variable like used above - var pointer *int
	// second is when star is used before variable - it accesses the value stored in box that the pointer points to - dereferencing
	fmt.Println(&age) //address of operator - address of the variable
	pointer=&age
	fmt.Println(*pointer)
	*pointer = 21 //we can manipulate actual value of the variable by referencing it through * operator.
	fmt.Println(age)
}