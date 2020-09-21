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
	pointers()

}

// 1. Hello World

func HelloWorld() {
	fmt.Println(" Hello World")
	fmt.Println(" This is My name ")
	println("The is a Print ")
	fmt.Println("Thies ")

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
	var a int = 8
	var b int8 = 20

	fmt.Println(" the Values are ", a, b)

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
		newSalary = salary * 2
		reason = "You have some Money, you get less raise"
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
	case salary > 1000 && salary < 5000:
		newSalary = salary * 5
		reason = " You Got a raise"
	case salary > 5000:
		newSalary = salary * 2
		reason = "Your really broke, you deserve it"
	default:
		newSalary = salary * 1
		reason = "Your really broke, you deserve it"
	}
	return newSalary, reason
}

// for Loops
/**
for i := 0; i <= 10; i++

Go for loop has three sections: initialization, the condition, and effect or action. All
sections are optional.
*/

func ForLoop() {
	for i := 0; i <= 10; i++ {
		if i > 5 {
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
		fmt.Println(" Print Something here!!")
		break
	}
}

func simulate_while_loop() {
	/**
	a = 0
	b = 10
	while( a > b) {
	statement;
	 a++
	}
	*/
	var a = 0
	var b = 10
	for {
		if a > b {
			break
		}
		fmt.Println(" Loop Count", a)
		a++
	}
}

func simulate_do_while_loop() {
	/**
	do{
	statements
	}while(condition)
	*/

	/// terminate when a > 5 < 10
	//

	var i = 0
	condition := true
	for ok := true; ok; ok = condition {
		if i > 10 {
			condition = false
		}
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()
}

// 5. Data Structures

// 5.1 Array
func arrays() {
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// array with numbers 1..10
	// Print for Me Just Even numbers
	// numbers[i]%2 == 0
	// else Print Odd Numbers
	var length = len(numbers)
	for i := 0; i < length; i++ {
		if numbers[i]%2 != 0 {
			fmt.Print(" ", numbers[i])
		}

	}
}

// 5.2 Two D Array

func TwoDArrays() {
	var arr2D = [3][3]int{
		{10, 20, 30},
		{40, 60, 90},
		{20, 70, 20},
	}
	fmt.Println(arr2D[0][0])
	fmt.Println(arr2D[0][2])

	// Loop Through a 2 D Array
	var rowLength = len(arr2D)
	var colLength = len(arr2D[0])
	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {
			fmt.Printf("%d,", arr2D[i][j])
		}
		fmt.Println("")
	}
}

//  5.3 Map

func Maps() {
	var phone = map[string]int{
		"john":  234374,
		"david": 872149044,
		"mary":  9821309312,
	}

	for key, phoneNumber := range phone {
		fmt.Println("The Index", key, "The Number", phoneNumber)
	}

	fmt.Println(" The Number for John", phone["john"])
	delete(phone, "john")
	fmt.Println(" New Number For John", phone["john"])
	phone["peter"] = 111121
	fmt.Println(" The Number for Peter", phone["peter"])
}

// 5.4 Slice

func slices() {
	var numbers = []int{1, 2, 3, 4, 5}
	fmt.Println(" The Slot 0 is ", numbers[0])
	numbers = append(numbers, 100)
	numbers = append(numbers, 200)

	fmt.Println(" The Slice is ", numbers)
	fmt.Println(numbers[1:3]) // 1..3
	fmt.Println(numbers[:3])  // 0..3
	fmt.Println(numbers[3:])  // 0..3

	// Loop Through the Slice
	for index, number := range numbers {
		fmt.Println("The Index", index, "The Number", number)
	}

}

// 5.6 Struct

func structs() {
	type Person struct {
		name string
		age  int
	}

	type ComputerNode struct {
		ipaddress string
		macaddess string
		processor int
	}

	var person1 = Person{"John", 10}
	var person2 = Person{"peter", 28}
	var person3 = Person{"mary", 18}
	var person4 = Person{"jenny", 27}

	var persons = []Person{person1, person2, person3}
	persons = append(persons, person4)

	for index, person := range persons {
		fmt.Println("At Index", index, " We have ", person)
	}

}

//6. Pointers
// Each variable Memory Address &
//

func pointers() {
	var age int = 10

	//var ipaddr string
	//var name string
	// pointer is a variable that stores and adress
	var agepointer *int = &age
	fmt.Println(" the Address of age is ", &age)
	fmt.Println(" the Value of age is ", age)
	fmt.Println(" the The Pointer is ", agepointer)
	fmt.Println(" Print the Pointer Value is ", *agepointer)
	//fmt.Println(" the Address of ipaddr is ",&ipaddr)
	//fmt.Println(" the Address of aga is ",&name)
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
