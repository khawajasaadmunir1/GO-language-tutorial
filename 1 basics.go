/*

Go  is a general-purpose language designed with systems programming in mind.
It is strongly typed and garbage-collected and has explicit support for concurrent programming.

Keywords
The following keywords are reserved and may not be used as identifiers(variable names).

break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var


Go package inside a module

https://golang.org/doc/code : How to Write Go Code

*/

// The first statement in a Go source file must be package name. Executable commands (like this file) must always start with 'package main'

// This main package is what we are creating. It could have been named something else e.g. 'package distsys' , but as written above , if we want an executable code it has to have 'package main'

package main

//import fmt package (details on what fmt does given later). fmt stands for the Format package. This package is all about formatting input and output.

import (
	"fmt"
	"time"
)

/*

This code groups the imports into a parenthesized, "factored" import statement. Can write import statements separately as well.

*/

func main() { //main function. Code Execution starts from here

	//--------Writing your first Hello World program in Go
	fmt.Println("Hello, World") //using Println function from fmt package that we imported above

	name := "LUMS"
	batch := 2025
	fmt.Print(name, " new batch is of ", batch, " :).\n") //Here,spaces are not added automatically between arguments

	// The fmt.Printf() function in Go language formats according to a format specifier and writes to standard output
	// www.geeksforgeeks.org/fmt-printf-function-in-golang-with-examples/

	fmt.Printf("%s new batch is of %d :).\n", name, batch)

	// Moving on.....
	basics()
	operators()
	types()

	// Functions (See definations also)
	fmt.Println("Calling 'add' function: ", add(42, 13))
	fmt.Println("Calling 'add2' function: ", add(10, 13))
	val1, val2 := squareAndCube(3)
	fmt.Println("Calling 'squareAndCube' function: ", val1, val2)

	//ignore one of the return values from a function using '_':
	_, val3 := squareAndCube(10)
	fmt.Println("Calling 'squareAndCube' function: ", val3)

	loopsAndIfAndSwitch()
	arraysANDslices()
	maps()
	pointersANDstructs()

}

func basics() {

	fmt.Println("-------------BASICS-------------")

	//If  variable not initialized to any value, default value is used
	// var is a keyword written , variable name is then written, variable type is then written (AFTER the variable name unlike in C++)
	var notAssignedInt int
	var notAssignedString string
	var notAssignedBool bool

	fmt.Println(notAssignedInt)
	fmt.Println(notAssignedString) //empty string
	fmt.Println(notAssignedBool)

	// Declaring and initialization
	var assignValue int = 5
	fmt.Println(assignValue)

	//Shorter way to declare and initialize (mostly used). Using :=
	//This method works only inside functions. Cannot be used at package level (outside of functions globally)
	//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	shorterMethod := 100
	fmt.Println(shorterMethod)

	//Strings
	firstString := "Distributed Systems : CS582"
	fmt.Println("isString?", firstString) //Space will be added automatically between all arguments of fmt.Println

	//Bools
	trueBool := true
	fmt.Println("isTrue?", trueBool)
	falseBool := false
	fmt.Println("isFalse?", falseBool)

	// Declaring numerous variables all at once

	a, b, c := 1, "ThisWorks", true
	fmt.Println(a, b, c)

	//Constants
	//Constants are declared like variables, but with the const keyword.

	const iAmConst2 = true
	const iAmConst3 string = "Hey I am a const."

	fmt.Println(iAmConst2)
	fmt.Println(iAmConst3)

	// iAmConst2 =  false // This will be an error if uncommented.
	// const iAmConst4 := 431 // Constants cannot be declared using the := syntax.

	/* 	Go's basic types are

	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // alias for uint8

	rune // alias for int32
	     // represents a Unicode code point

	float32 float64

	complex64 complex128

	*/

}

//ALL THE OPERATORS WORK JUST LIKE IN C++. NO NEED TO GO OVER ALL OF THEM AND SHOULD REMOVE THE OPERATORS FUNCTION ALTOGETHER
func operators() {
	fmt.Println("-------------OPERATORS-------------")

	//Arithmatic operators

	a, b := 21, 9 //integers declared
	fmt.Println("A:", a, "B:", b)
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Division:", a/b)
	// If the desired output is a float, you have to explicitly convert the values before dividing.
	fmt.Println("Float Division:", float64(a)/float64(b))
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Modulus/Remainder:", a%b)

}

//Type conversions

func types() {

	fmt.Println("-------------TYPES and TYPE INFERENCE-------------")

	//The expression T(v) converts the value v to the type T

	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	//Inferring type of a variable declared: Use Printf instead of Println
	// %T outputs the type of variable given as parametre
	fmt.Printf("i is of type %T\n", i)

	fmt.Printf("f is of type %T\n", f)

	fmt.Printf("u is of type %T\n", u)

}

//DECLARING FUNCTIONS PROPERLY....
/*
We write the function header in the following way

func FunctionName (parametre1,parametre2,..) (returnValue1,returnValue2,....) {


}

func = keyword. start function header
function can have 0 or more parametres and return values in GO !

*/

//function name is 'add'
//2 parametres x & y , both are integers
// only 1 return value type i.e. integer
func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func add2(x, y int) int {
	return x + y
}

// returning multiple results from a function
func squareAndCube(x int) (int, int) {

	return x * x, x * x * x
}

func loopsAndIfAndSwitch() {
	fmt.Println("-------------LOOPS-------------")

	/*
		Go only features the for loop.
		-There are no parentheses surrounding the three components of the for statement and the braces { } are always required.
	*/
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i //sum = sum + i
	}
	fmt.Println(sum)

	//GO's version of WHILE loop

	// -initialization and increment statements are optional. Only the conditional statement of for loop is required
	// - In this way, the for loop can be converted to the tranditional while loop
	myCounter := 1       //initialization (as in while)
	for myCounter < 10 { //(while loops condition)
		fmt.Println(myCounter)
		myCounter = myCounter + 1
	}

	//GO's version of infinite loop

	//If you omit the loop condition it loops forever. UNCOMMENT to check

	// for {
	// 	fmt.Println("infinite. break code to exit :)")
	// }

	fmt.Println("-------------IF STATEMENTS-------------")

	cGPA := 2.4 //CHANGE this to change behavior

	if cGPA < 2.0 {
		fmt.Println("Raise it. You on probation !")
	} else if cGPA < 2.5 {
		fmt.Println("Raise it. Near probation !")

	} else {
		fmt.Println("Keep working !")
	}

	fmt.Println("-------------SWITCH STATEMENTS-------------")

	//Go only runs the selected case, not all the cases that follow.
	//The break statement that is needed at the end of languages like C++ is provided automatically in Go

	i := 2 //CHANGE THIS for variation

	fmt.Print("Write ", i, " as ")
	switch i { //switch statemtn on 'i'
	case 1: // if i == 1
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default: //if none of the above cases was true (default is optional)
		fmt.Println("IDK !!")
	}

	//Go's switch cases need not be constants, and the values involved need not be integers.
	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

	//----Switch with no condition : Switch without a condition is the same as switch true.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

}

func arraysANDslices() {
	fmt.Println("-------------ARRAYS-------------")
	//The type [n]T is an array of n values of type T.
	//An array's length is part of its type, so arrays cannot be resized.
	//indexing from 0

	var a [2]string         // 'a' is an array of strings of size 2
	a[0] = "Hello"          //first element
	a[1] = "World"          //second elements
	fmt.Println(a[0], a[1]) //printing elements separately with space between them
	fmt.Println(a)          //PRINTING THE ARRAY ITSELF

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("Primes:", primes) //PRINTING THE ARRAY ITSELF

	fmt.Println("-------------SLICES-------------")
	//The type []T is a slice with elements of type T.

	// An array has a fixed size. A slice, on the other hand, is a dynamically-size (can change size)
	//A slice is formed by specifying two indices, a low and high bound, separated by a colon:

	var s []int = primes[1:4] //index 1 to 4 (4 not included)
	x := primes[1:4]          //shortcut
	fmt.Println("Slice:", s)
	fmt.Println("Slice2:", x)

	//A slice does not store any data, it just describes a section of an underlying array. Changing the elements of a slice modifies the corresponding elements of its underlying array.

	x[0] = -1 //change the first element of the slice (note that this is the 2nd element of the original primes array)

	fmt.Println("Primes got updated?:", primes) //primes gets updated
	fmt.Println("Slice got updated?:", s)       //s gets updated
	fmt.Println("Slice2 got updated?:", x)

	x = primes[:] //complete array gets copied (just like in python)
	fmt.Println("Primes Completely copied into Slice2:", x)

	//------A slice has both a length and a capacity.
	// The length of a slice is the number of elements it contains.
	fmt.Println("Length of Slice2:", len(x))

	x = primes[1:4]
	fmt.Println("Slice2 got updated?:", x)
	fmt.Println("Length of Slice2:", len(x))

	//-------The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

	fmt.Println("Capacity of Slice2:", cap(x)) //first element in the slice is the 2nd element in the underlying array. Hence, from element 2 to 6 we have 5 elements in total and so the capacity is 5.

	//------- Nil slices

	// The zero value of a slice is nil. A nil slice has a length and capacity of 0 and has no underlying array.
	fmt.Println()
	var nilSlice []int
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice))
	if nilSlice == nil {
		fmt.Println("nil!")
	}

	// More: https://medium.com/@ishagirdhar/nil-in-golang-aaa16565a5be

	//------Creating a slice with make
	fmt.Println("Creating a slice with make function")

	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays instead of making a slice of an existing Array (as done above).
	// The make function allocates a zeroed array and returns a slice that refers to that array:

	mySlice := make([]int, 5) // len(mySlice)=5 & capacity = length , each element is 0
	fmt.Println("mySlice:", len(mySlice), cap(mySlice), mySlice)

	//To specify a capacity, pass a third argument to make:

	mySlice2 := make([]int, 0, 5) // len(b)=0, cap(b)=5 , empty slice

	fmt.Println("mySlice2:", len(mySlice2), cap(mySlice2), mySlice2)

	//string slice
	mySliceString := make([]string, 3) // len(mySlice)=5 & capacity = length , each element is 0
	fmt.Println("mySliceString:", len(mySliceString), cap(mySliceString), mySliceString)

	//Slice of slices (like array of array i.e. 2d array)
	abc := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}

	fmt.Println(abc)

	//------APPENDING TO A SLICE (length changes as a result)
	fmt.Println("Appending to a slice")
	x = primes[1:4]
	fmt.Println("Slice2 got updated?:", x)
	fmt.Println("Length of Slice2:", len(x))

	x = append(x, 10, 20)
	fmt.Println("Slice2 got updated?:", x)
	fmt.Println("Length of Slice2:", len(x))

	//------Loop over a slice using RANGE

	// When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

	fmt.Println("Loop over this Slice:", x)
	for index, val := range x {
		fmt.Printf("index: %d ,value: %d\n", index, val)

		//NOTE: You can skip the index or value by assigning to '_' instead of a variable.

	}

}

func maps() {
	fmt.Println("----------------MAPS")

	//map in GO is  like dictionary in python i.e. key value pairs
	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
	// make maps using make function

	var myMap map[int]string = make(map[int]string)
	//Alternate:
	// var myMap map[int]string
	// myMap  =  make(map[int]string)

	myMap[1] = "Khawaja"
	myMap[2] = "Saad"
	myMap[3] = "Munir"

	//map size
	fmt.Println("myMap size:", len(myMap))

	//MAP literal:
	myMap2 := map[int]string{
		1: "rock",
		2: "john",
		//NOTE: have to leave a comma at the end of a MAP LITERAL
	}

	fmt.Println("Printing map: ", myMap)
	fmt.Println("Printing map2: ", myMap2)

	//MUTATING MAPS
	myMap2[1] = "THOR" //update value
	fmt.Println("updated map:", myMap2)

	elem := myMap2[1] //retrieve value
	fmt.Println("Retrieved val: ", elem)
	delete(myMap2, 2) //delete(mapName,keyToDelete)
	fmt.Println("updated map after deletion:", myMap2)

	// 	Test that a key is present with a two-value assignment:

	// elem, ok = m[key]
	// If key is in m, ok is true. If not, ok is false.

	// If key is not in the map, then elem is the zero value for the map's element type.

	val, ok := myMap2[2] //WILL NOT BE PRESENT.
	fmt.Println("The value:", val, "Present?", ok)

}

func pointersANDstructs() {
	fmt.Println("---------POINTERS")
	//POINTERS
	// - Unlike C, Go has no pointer arithmetic.
	//The type *T is a pointer to a T value. Its zero value is nil.

	// The & operator generates a pointer to its operand.
	// The * operator denotes the pointer's underlying value.

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	//STRUCTS

	// Go’s structs are typed (variable types) collections of fields. They’re useful for grouping data together to form records.

	fmt.Println("---------STRUCTS")

	type personalInfo struct {
		age      int
		name     string
		usePhone bool
	}

	//You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
	var myInfo personalInfo = personalInfo{
		age:      20,
		name:     "Adam",
		usePhone: false,
	}

	var myInfo2 personalInfo = personalInfo{10, "Hammad", true} //have to give in all the field values in correct order if we are to initialize directly.

	fmt.Println("Printing struct details:", myInfo)
	fmt.Println("myInfo2:", myInfo2)

	var myInfoPointer *personalInfo = &myInfo
	fmt.Println("Printing struct details using pointer deref:", *myInfoPointer)

	// to access a particular field, two ways exist with the pointers
	fmt.Println("my age: ", (*myInfoPointer).age)
	fmt.Println("my age: ", myInfoPointer.age) // dont have to explicitly dereference

}
