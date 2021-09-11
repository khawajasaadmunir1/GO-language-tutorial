package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		-channel is a technique/construct which allows to let one goroutine to send data (communicate) to another goroutine.
		-Think of them as pipes through which you can connect with different concurrent goroutines.
		-By default channel is bidirectional, means the goroutines can send or receive data through the same channel
		-different types of data are not allowed to transport from the same channel.
		- In the channel, the send and receive operation block until another side is not ready by default.

	*/

	//Make a channel: make(chan [value-type]), where [value-type] is the data type of the values to send and receive, e.g., int.

	stringChan := make(chan string)
	boolChan := make(chan bool)
	intChan := make(chan int)
	sliceChan := make(chan []int)
	mapChan := make(chan map[string]int)

	// IMPORTANT: By default, sends and receives BLOCK until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

	//-----SYNCHRONIZATION OF GO ROUTINES using CHANNELS
	primes := []int{2, 3, 5, 7, 11, 13}

	// Start separate go routines. Sum of Each half calculated separately and communicated back using intChan
	go sliceSum(primes[:len(primes)/2], intChan, 1) //goroutine 1
	go sliceSum(primes[len(primes)/2:], intChan, 2) //goroutine 2

	// you can send and receive values with the channel operator, <-.
	// (The data flows in the direction of the arrow.)

	// <-intChan //this is also a valid statement, in case the value from channel is not to be used

	fmt.Println("In the main go routine, waiting for partial sum to be received :(")
	//Until we receive something here, the main go routine stalls at this point. Hence, channels can help block a go routine

	partialSum1, partialSum2 := <-intChan, <-intChan // receive from channel intChan.
	fmt.Println("In the main go routine, partial sums received FINALLY :D")

	fmt.Println("Partial Sum1:", partialSum1)
	fmt.Println("Partial Sum2:", partialSum2)
	fmt.Println("Total:", partialSum1+partialSum2)

	// ------CLOSING A CHANNEL

	// -A sender can close a channel to indicate that no more values will be sent.
	// 	-Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	close(stringChan)
	close(boolChan)
	close(intChan)
	close(sliceChan)
	close(mapChan)

	// -Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression:

	valueFromChannel, ok := <-intChan
	// ok is false if there are no more values to receive and the channel is closed.
	if !ok {
		fmt.Println("Channel has been closed already !")
	} else {
		fmt.Println("Channel open. Use 'valueFromChannel': ", valueFromChannel)
	}

	/*-------------RANGE over Channels

	-Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

	The loop for i := range c receives values from the channel repeatedly until it is closed.

	*/

	queueChan := make(chan string, 2)
	queueChan <- "first"
	queueChan <- "second"
	close(queueChan)

	// range/iterate over the queueChan and keep on taking values out of the channel until the channel is closed.
	// This range iterates over each element as it’s received from queue. Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	//If the channel has not been closed (by the sender preferably), this loop will create a deadlock situation
	for elem := range queueChan {
		fmt.Println(elem)
	}

	//------------------Buffered Channels

	/*
		- We now know how channels are helpful in transfer of data between go routines.
		-Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel. Default buffers block when 1 element is sent. The element in the channel has to be received. Until then, we are in a blocking state
		-Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

	*/

	bufferedChan := make(chan int, 100) // 100 becomes the capacity of this channel

	//sending elements to the channel without blocking. This is because we have a buffered channel. If it was a default channel, there would have been blocking scenario after sending the first element because we are only receiving the data from the channel AFTER all the data has been sent
	for s := 0; s < 50; s++ {
		bufferedChan <- s
	}
	close(bufferedChan) // dont forget to close channel if a for loop is being used over the 'range' of channel

	for val := range bufferedChan {
		fmt.Print(val, " ")
	}
	fmt.Println()

	//-------CAPACITY AND LENGTH OF CHANNELS (just like slices)

	/*
		Length of the Channel: In channel, you can find the length of the channel using len() function. Here, the length indicates the number of value queued in the channel buffer.

		Capacity of the Channel: In channel, you can find the capacity of the channel using cap() function. Here, the capacity indicates the size of the buffer.

	*/

	//-------------SELECT STATEMENT

	/*
		- Combining goroutines and channels with select is a powerful feature of Go because select lets us wait on MULTIPLE channel operations.
		- select: it is only used with channels.
		- A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready
		- In go language, select statement is just like a switch statement without any input parameter.
		-case statement refers to communication, i.e. BOTH SEND or RECEIVE operation on the channel.
		-Select statement waits until the communication(send or receive operation) is prepared for some cases to begin.

	*/

	// Creating channels
	R1 := make(chan string)
	R2 := make(chan string)

	// calling function 1 and
	// function 2 in goroutine
	go portal1(R1)
	go portal2(R2)

	select { //select statement BLOCKS until one of the cases is ready (which means until something is received/sent via a channel)

	// case 1 for portal 1
	case op1 := <-R1:
		fmt.Println(op1)

	// case 2 for portal 2
	case op2 := <-R2:
		fmt.Println(op2)

		// //Use a default case to try a send or receive without blocking:
		// default:
		// 	fmt.Println("Default select statement run")

	}

	//----------SELECT & FOR loop

	//We can iterate over select statement i.e. make the select statement be evaluated more than once using for loops. We can similarly iterate over select statement in an infinite for loop and break out of it given some condition

	//set up 2 time channels : a value is received after some time in these channels automatically
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	exitNow := false
	for {
		// fmt.Println("Waiting for a case to get selected....")
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			exitNow = true
		default:
			fmt.Println("def .")
			time.Sleep(50 * time.Millisecond) //cause the execution of program to half for some time IF default case selected
		}

		if exitNow {
			break
		}
	}

	fmt.Println("INFINITE for loop with select statements EXITED !")

	//---------GO routines + select + for loop + blocking

	fmt.Println("fibonacci TIME")
	/*
		We start a printing go routine. In that go routine, we set up an for loop in which a channel receives some data and prints it and then next iteration starts.After 10 iterations, a signal is sent via the second channel that it is time to quit. Where does the data for printing come from? It is generated in 'fibonacci' function. In this function, we have an infinite loop with a select statement based on 2 channels. Until, on one of the channels some data is received, data is sent on the other channel constantly (this data is fibonacci numbers)
	*/

	numChan := make(chan int)  //channel on which the fibonacci number will be sent for the printFibonacci function to receive it
	flagChan := make(chan int) //channel on which a 'quit' signal will be sent from printFibonacci function to fibonacci function

	go printFibonacci(numChan, flagChan)
	fibonacci(numChan, flagChan)

}

//This function takes as input a slice and calculates its sum.
//Once the sum is calculated, the sum is fed into a channel (which is also provided as argument).
//The sum value sent into this channel from slideSum go routine, will be received by the channel (same channel in this case) in another go routine (main go routine , in this case)
//NOTE: We do not return from this function. We use a channel (shared between different go routines) to transfer data / communicate
func sliceSum(thisSlice []int, myChannel chan int, goRoutineNum int) {

	fmt.Println("I Am go routine ", goRoutineNum, "Slice:", thisSlice)

	//calculate partial sum
	sum := 0
	for _, val := range thisSlice {
		sum += val
	}

	fmt.Println("I Am go routine ", goRoutineNum, "Sending sum into the channel")

	// you can send and receive values with the channel operator, <-.
	// (The data flows in the direction of the arrow.)
	myChannel <- sum

	fmt.Println("I Am go routine ", goRoutineNum, "Exiting go routine now")
}

// function 1
func portal1(channel1 chan string) {

	time.Sleep(3 * time.Second)
	channel1 <- "Welcome to channel 1"
}

// function 2
func portal2(channel2 chan string) {

	time.Sleep(9 * time.Second)
	channel2 <- "Welcome to channel 2"
}

//print 10 fibonacci nums as received on a channel
func printFibonacci(numChan, flagChan chan int) {
	//loop 10 times and receive a number on channel each time and print it
	for i := 0; i < 10; i++ {
		//NOTE: unless numChan receives some data, the channel is blocking i.e. execution halts at the point of the code
		fmt.Println(<-numChan)
	}
	//once done, send a quit singal on the channel so that the other go routine knows it is time to QUIT/stop
	flagChan <- 0
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: //sending on a channel
			//update and find fibonacci num
			x, y = y, x+y
		case <-quit: //receiving on a channel
			fmt.Println("quit")
			return

			//uncomment the default case and re-run

			// default:
			// 	fmt.Println("no case selected....")
		}

	}
}
