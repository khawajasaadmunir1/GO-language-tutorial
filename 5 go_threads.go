//GO THREADS

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond) //sleep makes the go routine in which this loop is running stop execution for a while. This means that some other go routine (if present) can run.

		fmt.Println(s)
	}
}

func main() {

	// A goroutine is a lightweight  thread of execution.

	//Goroutines run in the same address space, so access to shared memory must be synchronized.

	//The goroutines’ output may be INTERLEAVED, because goroutines are being run concurrently by the Go runtime.

	//Result on console will be most probably different for each execution.
	//run the function in another go routine.
	go say("world")
	//run the function in current go routine (main go rountine)
	say("hello")

	//MORE: https://www.geeksforgeeks.org/goroutines-concurrency-in-golang/
	//https://medium.com/technofunnel/understanding-golang-and-goroutines-72ac3c9a014d

}

//rename this function as 'main' function and run (change the above function to main2 also)
func main2() {

	//counting to 10 concurrently
	for i := 0; i < 11; i++ {
		go fmt.Println(i)
	}

	//NOTE: If we comment the line below, nothing prints out. This is because if the main go routine exits, all the go routines that started off within it also exit and hence do not execute. What we want is that the main go routine should wait for all other go routines to finish executing, before it exits itself.
	time.Sleep(100 * time.Millisecond)
}
