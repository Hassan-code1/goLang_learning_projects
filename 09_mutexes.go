package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	_ "time"
)

//First
// func main(){
// 	counter := 0
// 	for i:= 0; i < 1000; i++{
// 		go func(){
// 			counter++
// 		}()
// 	}
// 	time.Sleep(time.Second)
// 	fmt.Println(counter)// expected 1000
// }

// output random between 1 to 1000

// WHY??
// shared meomory accessed concurrently
// ex initial coutn = 0
// it was gin to 2 routines to add 1 in it in each routine
// in concurrency it is possible both routines read when count = 0
// thus they add 1 and count become 1 in both
// out as 1

// or we can say race condn occured
//A race condition occurs when multiple goroutines access shared
//  data concurrently and at least one access is a write.

// NOTE MAP CAN'T BE ACCESSED CONCURRENTLY IF RACE OCCURS FOR MAP
// WE WILL GET A ERROR

// A critical section is:
// Code that accesses shared state.

// TO FACE THIS PROBLEM WE USE
// MUTEX OR MUTUAL EXCLUSION
// WHEN A ROUTINE USE A MEMORY EVERYOTHER HAVE TO WAIT
// TO USE THE SAME MEMORY

// PATTERN
// var mu sync.Mutex
// mu.Lock();
// Use Shared Memory
// mu.Unlock();

// func main(){
// 	var counter int
// 	var mu sync.Mutex
// 	var wg sync.WaitGroup
// 	for i := 0; i < 1000; i++{
// 		wg.Add(1)
// 		go func ()  {
// 			defer wg.Done()
// 			mu.Lock()
// 			defer mu.Unlock()
// 			counter++
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(counter)// out 1000
// }

// FORGETTING mu.Unlock() results in deadlock if other 
// routines want to use same memory

// mu.Lock()
// var1 
// var2
// mu.Unlock()

// every variable that is used in mutex are
// locked no one can access that var1 and var2


// RWMutex
// mainly when read major routines are used
// var rw sync.RWMutex
// rw.RLock() only readers allowed known as read lock
// rw.RUnlock() read unlock

// write lock rw.Lock()
// rw.Unlock()


// RACE DETECTION
// go run -race 09_mutexes.go
// or go test -race

// Atomic Operations ->  Cannot be interrupted midway.
// sync/atomic
// common atomic operations
// atomic.AddInt64() 
// atomic.LoadInt64()
// atomic.StoreInt64()
// atomic.SwapInt64()
// atomic.CompareAndSwapInt64()

func  main(){
	var counter int64
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println(counter)// out 1000
}

