package main

/*
If you remove package main, you will get an error. you need the package statement first.
you split your Go code across packages; you must have at least one package per app per
Go program.
You can have multiple files that make up one package.
Idea behind packages is to help organize code.
Main is a special package name that tells Go this will be the main entry point of the app
we are building
*/
//Here we are importing the fmt package that is part of Go's standard library

import "fmt"

//in Go, you must use double quotes or backticks. Single quotations not allowed.
func main() {
	fmt.Print("Hello World!")
}
