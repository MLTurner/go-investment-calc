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

//Here we are importing the fmt and math packages, that are part of Go's standard library

import (
	"fmt"
	"math"
)

/*
In Go, you must use double quotes or backticks. Single quotations not allowed.
Unlike in JS, Go does not start at the top to execute functions. It begins with the one called "main."
The symbol := declares and assigns a variable. When there is no explicit type assignment, you should
use := instead of the var keyword.
*/

func main() {
	//Will not be able to change the value of a const. We do not want users to change inflation rate.
	//Therefore, constants always need an initial value.
	const inflationRate = 2.5
	//Go infers that the type will be int because the hard-coded value of 1000 is an int
	//You can add an explicit type assignment after the variable name, overriding the inferred type
	var investmentAmount float64
	//Our program should be able to accept all kinds of values, with or without decimal places

	var years float64
	var expectedReturnRate float64

	/*
		Using fmt.Print, we can provide a more user-friendly experience by prompting the user to enter
		an investment amount.
	*/
	fmt.Print("Investment amount: ")

	/*
		fmt.Scan allows us to get input from the terminal.
		You need to pass a pointer to Scan, which is notated by the &.
		The usefulness of variables is that investmentAmount was initally set at 1000,
		but the user can override it with this fmt.Scan function.
		If we just set the original investmentAmount to "var investmentAmount float64",
		the variable will be assigned its default null value (0.0 in case of float64).
	*/

	fmt.Scan(&investmentAmount)

	fmt.Print("Length of investment in years: ")
	fmt.Scan(&years)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	/*
		Want to determine the value of our money in the future
		We use math.Pow, imported from the Go standard library. math.Pow() allows us to multiply the investment amount
		to 1 + the expected return rate divided by 100, raised to the power of the number of years
	*/
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	/*
		Want to determine the value of our money in the future, accounting for inflation. Up above, we listed
		inflationRate as a constant, so the user cannot edit this.
	*/
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	/*
		Finally, we want to write the answer to the standard output
	*/
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

/*
fmt.Scan() limitations: this is a great function for easily fetching and using user input through the command line.
However, you cannot easily fetch multi-word input values with this. Fetching text that consists of more than a single
word is tricky with this function.
*/
