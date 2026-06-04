package main

import (
	"fmt"
	"sync"
	"time"
)

//race condition
//A race condition in Go occurs when two or more goroutines
//  access the same shared variable concurrently, and 
// at least one of those accesses is a write.
// so if we do wg.Add(1) inside goRoutine the race condition will occur
// never do like this
// go func() {
//     wg.Add(1)
// }()


// deadlock
//forgetting wg.Done()
//like this (never do this) program waits Forever
// wg.Add(1)
// go func() {
// }()

//correct way
// wg.Add(1)
// go func() {
//     defer wg.Done()
// }()


//first
// func Hello ()  {
// 	fmt.Println("Hello")
// }

// func main(){
// 	go Hello()
// 	fmt.Println("world")
// }

//second
// func hello(wg *sync.WaitGroup) {

//     defer wg.Done()

//     fmt.Println("Hello")
// }

// func main() {

//     var wg sync.WaitGroup

//     wg.Add(1)

//     go hello(&wg)

//     wg.Wait()
// }
//out Hello

//third
// func main() {
// 	var wg sync.WaitGroup
// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)
// 		go func(num int) {
// 			defer wg.Done()
// 			fmt.Println(num)
// 		}(i)
// 	}
// 	wg.Wait()
// }
//ouput 1 2 3 4 5 in random order since scheduler decides which to run first we dont
// you never assume
/*
First created
=
First executed
*/

//fourth
//Anonymous Goroutines

// go func () {
// 	fmt.Println("immediate execution")
// }()

//fifth concurrent execution
func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d started\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d finished\n", id)
}
func main(){
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
    	wg.Add(1)
    	go worker(i, &wg)
	}
	wg.Wait()// wait till all routines are executed then returns from main
}
// output
// 1 2 3 all started in random
// 1 2 3 all finished in random

/*
go keyword           -> The go keyword starts a function as a new goroutine that runs concurrently with the caller.

Goroutine lifecycle  -> A goroutine moves through creation, scheduling, execution, and termination managed by the Go runtime.

Main goroutine       -> Every Go program starts with the main goroutine, and when it exits all other goroutines are terminated.

Scheduler basics     -> The Go scheduler automatically maps goroutines onto OS threads and CPU cores for execution.

Anonymous goroutines -> Anonymous goroutines are inline functions launched concurrently using go func(){...}().

Variable capture bug -> Goroutines inside loops may capture the same loop variable, causing unexpected values unless passed as arguments.

sync.WaitGroup       -> WaitGroup synchronizes goroutines by waiting for all registered tasks to complete before proceeding.

Add()                -> Add(n) increases the WaitGroup counter by n to indicate new goroutines or tasks.

Done()               -> Done() decrements the WaitGroup counter by 1 when a goroutine finishes its work.

Wait()               -> Wait() blocks execution until the WaitGroup counter reaches zero.

Common concurrency mistakes -> Frequent mistakes include forgetting Done(), calling Add() after launching goroutines, copying WaitGroups, and assuming execution order.

Golden workflow      -> Add() → go routine → defer Done() → Wait() is the standard WaitGroup synchronization pattern.
*/