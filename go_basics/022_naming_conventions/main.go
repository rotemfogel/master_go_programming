package main

// package scope
var taskDone bool // OK

func main() {
	// Go names start with a letter or underscore (_)
	// case matters quickSort and QuickSort are different (variable vs function)
	// Go has only 25 keywords (https://golang.org/ref/spec#Keywords)

	var mv int        // store maximumValue
	var max_value int // python like - NOT GO IDIOMATIC

	var packetsReceived int // NOT OK, name too long
	var b int               // OK
	_, _, _, _ = mv, max_value, packetsReceived, b

	const MAX_VALUE = 100 // NOT GO IDIOMATIC
	const N = 100         // BETTER
	// note: UPPERCASE letters are exported
	// use uppercase letter variables or functions
	// if you want to use it in other functions or packages
	const n = 100 // will be local to scope

	// multi word names should be camelCase
	// applicable to variables, constants and functions

	maxValue := 1000 // recommended
	_ = maxValue

	// Naming Conventions are important for code readability and maintainability.

	// use short, concise names especially in shorter scopes
	// common names for common types:
	var s string   // string
	var i int      // index
	var num int    // number
	var msg string // message
	var v string   // value
	var err error  // error value
	var done bool  // bool, has been done?
	_, _, _, _, _, _, _ = s, i, num, msg, v, err, done

	// acronyms should be written in UPPERCASE
	writeToDB := true // OK
	writeToDb := true // NOT IDIOMATIC IN GO
	_, _ = writeToDB, writeToDb

	// packages should be given lowercase single work names
	// without underscores (_) or camelCase

	// struct Getters and Setters should be called
	// with an UPPERCASE letter
	type person struct {
		name    string
		Name    func() *string    // function to return a person's name
		SetName func(name string) // function to update a person's name
	}

	// a single method interface should be names by the method name +er suffix
	type Doer interface {
		Do()
	}

}
