package main

import "fmt"

// Go does not provide classes but it does provide structs. Methods can be added on structs. This provides the behaviour of bundling the data and methods that operate on the data together akin to a class.

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

/*

THIS IS A METHOD.
-A method is a function with a special receiver argument.
-The receiver appears in its own argument list BETWEEN the func keyword and the method name.
	(e Employee) is the VALUE receiver (remember this. There is a difference between Value and Pointer receiver)
-LeavesRemaining method has a receiver of type 'Employee' named 'e'.

*/

// we have a struct and a method that operates on a struct bundled together like in a typical C++ class.

func (e Employee) LeavesRemaining() {

	fmt.Println("In 'LeavesRemaining' function")
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}

//Instead of calling a method of a particular struct, we call a general function that takes in as paramtre that very same struct.
func LeavesRemainingGeneral(e Employee) {

	fmt.Println("In 'LeavesRemainingGeneral' function")
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}

/*
You can declare methods with pointer receivers.
This means the receiver type has the literal syntax *T for some type T.

Methods with pointer receivers can modify the value to which the receiver points (as UpdateLeavesTaken does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Try removing the * from the declaration of the UpdateLeavesTaken function and observe how the program's behavior changes.
*/

func (e *Employee) UpdateLeavesTaken() {
	fmt.Println("In 'UpdateLeavesTaken' function")
	e.LeavesTaken = e.LeavesTaken + 3 //we are updating value of the field of structure. THIS IS ONLY POSSIBLE IF THE METHOD IS A POINTER RECEIVER METHOD INSTEAD OF VALUE RECEIVER METHOD (think of this as pass by val vs pass by ref)
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}

func main() {

	e := Employee{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}

	e.LeavesRemaining() //The LeavesRemaining() method of the Employee struct is called in main().

	//Remember: a method is just a function with a receiver argument.

	//Whatever has been done in valueReceiver() can be done via a regular function
	LeavesRemainingGeneral(e)

	//----------Now we move to Pointer receivers after having covered VALUE receivers

	e.UpdateLeavesTaken() //the structure 'e' itself is getting updated. This will not be possible using a value receiver method because in that method, a copy of structure is shared with the method.

}

// There are two reasons to use a pointer receiver.

// The first is so that the method can modify the value that its receiver points to.

// The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
