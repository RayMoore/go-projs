package main

import "fmt"

// each go file needs a main func to exec
// each folder is a package
// each package should have only one main func
func main() {

	// this is single line comment
	/*
	 this is multi-line comments
	 */

	fmt.Println("hello go lang")
	fmt.Println("You can seperate long string by using \",\".",
		"Like + in Java and other langs")

	// variables
	var i int
	i = 10
	fmt.Println(i)
	// use "," to concat string, a blank space will be added
	fmt.Println("i =", i)
	// other methods to declare variables
	var j = 10.11 // dynamic binding
	fmt.Println("j =", j)
	k := "this is a string" // omit "var" to declare a variable
	fmt.Println("k =", k)
}
