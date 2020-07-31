package main

import (
	"fmt"
)

func main() {
	// you must provide a type for each variable you declare or Go should be able to infer it
	var a = 10
	var b = 15.5      // type inference (deduction)
	c := "Gopher"     // short declaration, type inference
	_, _, _ = a, b, c // Blank Identifier (_) to get rid of unused variable error

	// Go is a Statically and Strong Typed Programming Language
	// a = 3.14 -> error. A variable cannot change it's type
	// a = b    -> error. It's not allowed to assign a type to another type

	var d int
	var e float64

	// Type Casting is allowed
	d = int(b)
	e = float64(a)
	_, _ = d, e

	//** ZERO VALUES **//

	// go has default variable values:
	// - numeric -> 0
	// - boolean -> false
	// - string  -> "" (empty string)
	// - pointer -> Nil
	var value int
	var name string
	var price float64
	var done bool
	// should print 0 "" 0 false
	fmt.Println(value, name, price, done)
}
