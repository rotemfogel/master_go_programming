package main

import "fmt"

func main() {
	//** DECLARING VARIABLES **///

	// Syntax: var var_name type
	var s1 string
	s1 = "Learning Go!"
	fmt.Println(s1) // printing string s1

	//** TYPE INFERENCE **//

	// Go deduces automatically the type of the variable by looking at the initial value (bool, int, string etc)

	var k = 6   // not necessary to say the type (int). It is inferred from the literal on the right side of =
	var i = 5   // type int
	var j = 5.6 // type float64

	// printing i, j and k
	fmt.Println("i:", i, "j:", j, "k:", k)

	// ii == jj  // -> error: cannot assign float to int (Go is a strong typed language)

	// declaring and initializing a new variable of type string (type inference)
	var s2 = "Go!"
	_ = s2 //in Go each variable must be used or there is a compile-time error
	// _ is the Blank Identifier and mutes the error of unused variables
	// _ can be only on the left hand side of the = operator

	//** Short Declaration (works only in Block Scope!) **//
	// := (colon equals syntax) used only when declaring a new variable (or at least a new variable)
	// := tells go we are going to create a new variable and go figures out what type it will be
	s := "Learning GoLang"
	fmt.Println(s)

	// multiple variable declaration
	car, cost := "Audi", 50000
	fmt.Println(
		car,
		cost,
	)
	/*
		// will cause an error since no new variable
		// is present on the left hand side
		car := "BMW"
	*/

	// multilpe declartion with reassigment
	car, year := "Audi", 2019
	_ = year

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	// another way to declare multiple variables
	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	// multiple typed variable declaration
	var a, b, c int
	fmt.Println(a, b, c)

	// multiple variable assignment
	var ii, jj int
	ii, jj = 5, 8
	fmt.Println(ii, jj)
	ii, jj = jj, ii // swapping variables
	fmt.Println(ii, jj)

	// expression in a short declaration
	sum := 5 + 2.3
	fmt.Println(sum)
}
