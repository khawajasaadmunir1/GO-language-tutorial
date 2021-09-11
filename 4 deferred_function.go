package main

import (
	"fmt"
)

func main() {
	//This print will be done on to the console once the all the arguments have been evaluated of Println function i.e. AFTER the isEven(10) call returns
	fmt.Println("Result of isEven(10):", isEven(10))
}

func isEven(num int) bool {
	//the following print statement will be executed immediately after the isEven() function returns
	defer print("I have exited already")

	// normal execution
	fmt.Println("I am running")
	return num%2 == 0
}

func print(s string) {
	fmt.Println(s)
}

// HOW DOES THE ABOVE CODE WORK?
// main function start, but the Println statement does not print anything to the console until isEven function returns.

// We then move into the isEven function  call. We defer a print statement i.e. it will be called the moment isEven function returns. Moving on in the isEven function, 'I am running' is printed onto the screen (1st print). Then, isEven() returns and the deferred function is executed and 'I have exited already' (2nd print) is printed on console. Then, the 3rd print i.e. the print statement of main function runs.

//change this to main function and run
func main2() {

	tryingDEFERfunctions()
}

func tryingDEFERfunctions() {
	fmt.Println("-----------TRYING OUT DEFER FUNCTION")

	defer fmt.Println("I will run AFTER my surrounding function exits !!! :)")

	fmt.Println("-----------DONE WITH DEFER FUNCTION")

	// 	A defer statement defers the execution of a function until the surrounding function returns.

	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

}
