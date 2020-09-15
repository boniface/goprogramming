// INTRODUCTION
//
// Golang  Site:  golang.org
// System programing Language
// Used to create Applications like Docker, Kubernetes, and used a lot in Google cloud, the best programing  language for networking professionals

// 1. Installation
// https://play.golang.org
// Get Goland Site Download https://www.jetbrains.com/go/promo/
// Free Golang Book URL https://www.golang-book.com/

// This is the beginning of a Go Program with a Main Method
// This Programing is calling a bunch of functions
// We will be doing most of the stuff in the functions and calling them into the
//  main program to test if they work
// Everything is kept in a single file to make it easier to onboard new people to th language
package main

import "fmt"

func main() {
	HelloWorld()
	flowIfElse(10)
	flowElseIf(5)
	var newSalary, reason = increaseSalary(10)
	fmt.Println(" Your New Salary is ", newSalary)
	fmt.Println("Here is the Reason why  ", reason)
	ForLoop()
}

// 1. Hello World

func HelloWorld() {
	fmt.Println(" Hello World")
}

// 2. varriables

func Variables() {
	// Ways to declare variable and assign values 3
	var name string = "Developer" //delcare with var
	fmt.Println(name)
	field := "Teaching Industry" // Not specifying var
	fmt.Println(field)

}

// 3. Primitive Data Types

func DataTypes() {

}

// 4. Data Type Conversion

// This is the end of a Go Program

//2. Basics

// 3. Arithmetic  Operators
// +, -, /, *, %
func add(number1 int, number2 int) int {
	var sum int = number1 + number2
	return sum
}

func multiply(number1 int, number2 int) int {
	var ans = number2 * number1
	return ans
}

func subtract(number1 int, number2 int) int {
	var ans = number2 - number1
	return ans
}

func divide(number1 int, number2 int) (int, int) {
	var ans = number1 / number2
	return ans, 0
}

// Logical Operators
// and (&&) or (||) not (!)
// true && false
// false && true
// false && false
// true && true

func logicalOperators() {
	println("=== PRINTING AND OPERATORS ===")
	println("THE TRUE && FALSE IS ", true && false)
	fmt.Println(" TRUE && TRUE", true && true)
	fmt.Println(" FALSE && TRUE", false && true)
	fmt.Println(" FALSE && FALSE", false && false)

	println("=== OR ===")
	fmt.Println(" FALSE || TRUE", false || true)
	fmt.Println(" TRUE || FALSE", true || false)
	fmt.Println(" TRUE || TRUE", true || true)
	fmt.Println(" FALSE || FALSE", false || false)

	println("=== Not ===")
	println(" Not on True is ", !true)
	println("Not on False ", !false)
}

// Relational Operators
// < =< <= > ==
// 3 > 8 || 7 == 7-3
// 5 < 10 && 2 > 1
func relationalOperators() {
	println("Is this True ", 1 > 2 && 3 > 4)
}

// 4. Flow Control
// if else
func flowIfElse(number int) {
	if number == 2 {
		fmt.Println(" The Number was TWO")
	} else {
		fmt.Println(" The Number was Something Else, Which is", number)
	}
}
func flowElseIf(number int) {
	// Test Several Conditions
	if number == 2 {
		fmt.Println(" The Number was TWO")
	} else if number > 5 && number < 10 {
		fmt.Println(" The Number was 5, Which is", number)
	} else if number == 6 {
		fmt.Println(" The Number was 6, Which is", number)
	} else if number == 7 {
		fmt.Println(" The Number was 7, Which is", number)
	} else {
		fmt.Println(" The Number was Something Else, Which is", number)
	}
}

// Increase Salary
// earns less 1000 multiply by 10
// earns between 1000 and 5000 multiply by 5
// earn greater 5000  and less 10000 multiply 2
// earn greater 10000 multiply by 1
// take a salary and return new Salary
//
func increaseSalary(salary int) (int, string) {
	var newSalary = 0
	var reason = ""

	if salary < 1000 {
		newSalary = salary * 10
		reason = "Your really broke, you deserve it"
	} else if salary > 1000 && salary < 5000 {
		newSalary = salary * 5
		reason = "Your have Some Money"
	} else if salary > 5000 && salary < 10000 {
		reason = "You Work Hard, you deserve it"

	} else {
		newSalary = salary * 1
		reason = "Your Have enough Money "
	}
	return newSalary, reason
}

// switch
//

func increaseSalaryWithSwitch(salary int) (int, string) {
	var newSalary = 0
	var reason = ""
	switch {
	case salary < 1000:
		newSalary = salary * 10
		reason = "Your really broke, you deserve it"
	case salary > 1000 && salary < 10000:
		newSalary = salary * 10
		reason = "Your really broke, you deserve it"

	case salary > 1000 && salary < 10000:
		newSalary = salary * 10
		reason = "Your really broke, you deserve it"
	case salary > 5000 && salary < 10000:
		newSalary = salary * 10
		reason = "Your really broke, you deserve it"
	default:
		newSalary = salary * 10
		reason = "Your really broke, you deserve it"
	}
	return newSalary, reason
}

// for Loops

func ForLoop() {
	for i := 0; i <= 10; i++ {
		if i == 0 {
			fmt.Println(" Hello World")
			break
		} else {
			println(" Am not Printing Any Number")
		}
	}
}

func InfiniteLoop() {
	// break, continue
	for {
		fmt.Println("")
	}

}

func simulate_while_loop() {

}

func simulate_do_while_loop() {

}

// 5. Data Structures

// 5.1 Array

func arrays() {

}

// 5.2 Two D Array

func TwoDarrays() {

}

// 5.4 Slice

func slices() {

}

//  5.3 Map

func Maps() {

}

// 5.6 Struct

func strusts() {

}

//6. Pointers

func pointers() {

}

// First Session Stuff
func firstSessionWork() {
	println("Hello World")
	var result = add(3, 5)
	println(" Ans", result)

	// call mult function
	var mult_result = multiply(10, 5)
	//print
	println(" the Mult result ", mult_result)

	//call sub func
	var sub_result = subtract(10, 5)
	//print
	println(" The Sub Result ", sub_result)

	//call divid
	var correctAns, err = divide(20, 10)
	//print
	println(" the Correct result is ", correctAns, "The Error is", err)

	logicalOperators()
	relationalOperators()
}
