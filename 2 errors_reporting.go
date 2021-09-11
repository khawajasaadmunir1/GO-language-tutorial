//  Nil values , Return err and logs

package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	errorsFunctions()
}

func errorsFunctions() {
	fmt.Println("------ERRORS")

	//In GO, we communicate errors via an explicit, separate return value instead of returns error values as in C
	//By convention, errors are the last return value and have type error, a built-in interface.
	var userInput float64
	fmt.Print("Input a number to find its sq root:")
	fmt.Scan(&userInput)
	ans, err := squareRoot(userInput) //recevie the returned values from function

	if err != nil { // check for error presence
		fmt.Println(("ERROR PRINTING !"))

		//a flag to disable printing the time, source file, and line number.
		// log.SetFlags(0)

		log.Fatal(err)

	} else {
		fmt.Println("Answer:", ans)
	}

	// If the returned error is not nil it usually means that there is a problem and you need to handle the error appropriately. This can mean that you use some kind of log message to warn the user, retry the function until it works or close the application entirely depending on the situation.

	// https://go.dev/blog/error-handling-and-go

	// fmt.Errorf() can also be used instead of errors.New : fmt.Errorf on the other hand also provides the ability to add formatting to your error message.

}

func squareRoot(num float64) (float64, error) {

	if num < 0 {
		//IF error, return an 'error' with any given string
		//errors.New constructs a basic error value with the given error message.

		return num, errors.New("Input was a negative number. NO SOLUTION")
	}
	// A nil value in the error position indicates that there was no error.

	return math.Sqrt(num), nil //no error is often given by a nil value returned instead of 'error'

}
