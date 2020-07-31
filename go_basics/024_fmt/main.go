package main

import (
	"fmt"
)

func main() {
	name := "Rotem"
	fmt.Println("hello", name)

	a, b := 4.0, 5.5
	fmt.Println("a and b -> sum:", a+b, "mean:", (a+b)/2)

	figure := "Circle"
	radius := 5
	pi := 3.14159

	// formatting the print
	fmt.Printf("Figure: %s, Pi: %f, Radius is %d, circumfence: %f\n", figure, pi, radius, 2*pi*float64(radius))

	// use %v to replace with any type
	fmt.Printf("Figure: %v, Pi: %v, Radius is %v, circumfence: %v\n", figure, pi, radius, 2*pi*float64(radius))

	// quoted strings
	fmt.Printf("This is a %q figure\n", figure)

	// use %T for printing type
	fmt.Printf("figure is of type %T\n", figure)
	fmt.Printf("radius is of type %T\n", radius)
	fmt.Printf("pi is of type %T\n", pi)

	// format boolean value as "true/false" %t
	closed := true
	fmt.Printf("file closed: %t\n", closed)

	// %b converts decimal integer to a binary (base 2)
	num := 55
	fmt.Printf("%d in base2 is: %b\n", num, num)

	// in order to print a binary number using a specific number of bits
	fmt.Printf("%d in base2 with leading zeros is: %08b\n", num, num)

	// convert decimal to hexadecimal
	fmt.Printf("%d in hexademical is: %x\n", num, num)

	// format float to 3 decimal points
	x, y := 5.4, 2.6
	fmt.Printf("x * y = %.3f\n", x*y)

}
