package main

import "fmt"

func main() {
	// constants must be initialized when declared
	const days int = 7

	/*
	 ERR: missing value in const declaration
	*/
	// const pi float64

	// as opposed to variables which get assigned default values
	var i int
	fmt.Println(i)

	const secondsInHour int = 3600
	duration := 234
	// use of constants
	fmt.Println("Duration in seconds:", duration*secondsInHour)

	// constants are handled in compile time, whereas variables are handled in runtime
	x, y := 5, 0
	_, _ = x, y
	/*
		this yields:
		panic: runtime error: integer divide by zero
	*/
	// fmt.Println(x / y)

	const a = 5
	const b = 0
	/*
	  $ go build main.go
	  # command-line-arguments
	  ./main.go:34:16: division by zero
	*/
	// fmt.Println(a / b)

	// declare multiple constants
	const (
		boilingPoint = 100      // °C
		density      = 997      // kg/m³
		molarMass    = 18.01528 // g/mol
		meltingPoint = 0        // °C
		formula      = "H2O"
	)

	fmt.Println("Water boiling point: ", boilingPoint,
		" density: ", density,
		" molar mass: ", molarMass,
		" melting point: ", meltingPoint,
		" formula: ", formula)

	// consequtive constants get values from previous constant
	const (
		c1 = 100
		c2
		c3
	)
	fmt.Println(c1, c2, c3)

}
