package main

import "fmt"

func main() {
	flowIfElse(10)
	flowElseIf(5)
	var newSalary, reason= increaseSalary(10)
	fmt.Println(" Your New Salary is ", newSalary)
	fmt.Println("Here is the Reason why  ", reason)
	ForLoop()
}

// ``` ```
// Golang  Site:  golang.org
// System programing Language
// Docker, Kubernetes, and used in Google cloud, the best programing  networking professionals

// 1. Installation
// https://play.golang.org
// Get Goland Site Download https://www.jetbrains.com/go/promo/
// Free Golang Book URL https://www.golang-book.com/

//2. Basics



// 3. Operators
// +, -, /, *, %
func add(number1 int, number2 int) int {
	var sum int  = number1 + number2
	return sum
}

func multiply(number1 int, number2 int) int {
   var ans = number2 * number1
	return ans
}

func subtract(number1 int, number2 int) int {
	var ans = number2 - number1
	return  ans
}

func divide(number1 int, number2 int) (int,int) {
	var ans = number1/number2
	return ans,0
}


// Logical Operators
// and (&&) or (||) not (!)
// true && false
// false && true
// false && false
// true && true

func logicalOperators()  {
	println("=== PRINTING AND OPERATORS ===")
	println( "THE TRUE && FALSE IS ", true && false )
	fmt.Println(" TRUE && TRUE",true && true)
	fmt.Println(" FALSE && TRUE",false && true)
	fmt.Println(" FALSE && FALSE",false && false)

	println("=== OR ===")
	fmt.Println(" FALSE || TRUE",false || true)
	fmt.Println(" TRUE || FALSE",true || false)
	fmt.Println(" TRUE || TRUE",true || true)
	fmt.Println(" FALSE || FALSE",false || false)


	println("=== Not ===")
	println(" Not on True is ", !true)
	println("Not on False ", !false)
}

// Relational Operators
// < =< <= > ==
// 3 > 8 || 7 == 7-3
// 5 < 10 && 2 > 1
func relationalOperators(){
	println("Is this True ", 1 > 2 && 3 > 4)
}


// 4. Flow Control
// if else
func flowIfElse(number int){
	if number == 2 {
		fmt.Println(" The Number was TWO")
	}else {
		fmt.Println(" The Number was Something Else, Which is",number)
	}
}
func flowElseIf(number int)  {
// Test Several Conditions
	if number == 2 {
		fmt.Println(" The Number was TWO")
	}else if number > 5 && number < 10 {
		fmt.Println(" The Number was 5, Which is",number)
	}else if number == 6 {
		fmt.Println(" The Number was 6, Which is",number)
	}else if number == 7 {
		fmt.Println(" The Number was 7, Which is",number)
	} else {
		fmt.Println(" The Number was Something Else, Which is",number)
	}
}

// Increase Salary
// earns less 1000 multiply by 10
// earns between 1000 and 5000 multiply by 5
// earn greater 5000  and less 10000 multiply 2
// earn greater 10000 multiply by 1
// take a salary and return new Salary
//
func increaseSalary( salary int) (int,string){
	var newSalary = 0
	var reason = ""

	if salary < 1000 {
		newSalary = salary*10
		reason = "Your really broke, you deserve it"
	}else if salary > 1000 && salary < 5000 {
		newSalary = salary *5
		reason = "Your have Some Money"
	}else if salary > 5000 && salary < 10000 {
		reason = "You Work Hard, you deserve it"

	} else {
		newSalary = salary * 1
		reason = "Your Have enough Money "
	}
	return newSalary,reason
}

// switch

func increaseSalaryWithSwitch(salary int) (int,string){
	var newSalary = 0
	var reason =""
	switch salary{
	case 5000:
		newSalary = salary*10
		reason = "Your really broke, you deserve it"
	case 2000:
		newSalary = salary*10
		reason = "Your really broke, you deserve it"
	default:
		newSalary = salary*10
		reason = "Your really broke, you deserve it"

	}
	return newSalary, reason
}

// for

func ForLoop()  {
	for i:=0; i<=10; i++{
		if(i==0){
			fmt.Println(" Hello World")
			break
		} else{
			println(" Am not Printing Any Number")
		}
	}
}

func EnhanceForLoop()  {
	// break, continue
	for{
		fmt.Println("")
	}

}






// 5. Data Structures

// 5.1 Array

// 5.2 Two D Array

//  5.3 Map

// 5.4 Slice

// 5.6 Struct

//6. Pointers

// First Session Stuff
func firstSessionWork()  {
	println("Hello World")
	var result = add(3,5)
	println(" Ans", result)

	// call mult function
	var mult_result = multiply(10,5)
	//print
	println(" the Mult result ", mult_result)

	//call sub func
	var sub_result = subtract(10, 5)
	//print
	println(" The Sub Result ", sub_result)

	//call divid
	var correctAns,err = divide(20, 10)
	//print
	println(" the Correct result is ", correctAns, "The Error is",err)

	logicalOperators()
	relationalOperators()
}
