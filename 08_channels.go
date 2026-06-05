package main

import  "fmt"

// what is a channel -> A pipe connecting goroutines.

// channel can carry specific type values
/*
make(chan string)
make(chan bool)
make(chan float64)
*/

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		ch <- 1
// 	}()

// 	value := <-ch
// 	fmt.Println(value)
// }//out 1

//unbuffered channel 
// make(chan int)
/*
Send blocks until receive exists.
Receive blocks until send exists.
*/

// func  main(){
// 	ch := make(chan int)
// 	ch <- 10
// }//out dead lock 
/*
Nobody receives.
Sender waits forever.
*/

// synchronization using channels
// func worker(done chan bool) {
//     fmt.Println("Working")
//     done <- true
// }

// func main(){
// 	done := make(chan bool)
// 	go worker(done)
// 	<-done // works likewaitGroup.Wait()
// }//out Working


//buffered channels
// func main(){
// 	ch := make(chan int, 3)// channel with capacity 3
// 	ch <- 1
// 	fmt.Printf("Capacity of channel :%d current length of channel : %d\n",cap(ch), len(ch)) 
// 	ch <- 2
// 	fmt.Printf("Capacity of channel :%d current length of channel : %d\n",cap(ch), len(ch))
// 	ch <- 3
// 	fmt.Printf("Capacity of channel :%d current length of channel : %d\n",cap(ch), len(ch))
// 	// ch <- 4//out error buffer blocked send
// 	//ch = [1 2 3]
// 	<-ch//ch = [2 3]
// 	<-ch//ch = [3]
// 	<-ch// ch empty
// 	// <-ch// error ch empty recieving blocked
// }

//directional channels
// ch chan<- sendOnly
// ch <-chan recieveOnly


//Closing of channels
//close(ch)
//No more values will be sent. but can recieve till empty
// closing of channel doesn't delete data currently in channel
// func main(){
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	close(ch)
// 	fmt.Println(<-ch, <-ch)// works fine
// 	//out 1 2
// }

// checking if channel is closed
// val, ok := <-ch
// ok == true -> channel open
// ok == false -> channel closed and empty
// doubt if channel -> empty will ok give false
// func main(){
// 	ch := make(chan int, 1)
// 	ch <- 1
// 	<-ch
// 	val, ok := <-ch// error if ch not closed but empty
// 	fmt.Println(val, ok)
// }

// closed a channel 2 times -> panic: close of closed channel

//range over channel

func main(){
	ch := make(chan int, 3)
	// ch<-1;ch<-2;ch<-3;
	// for val := range ch {
	// 	fmt.Println(val, len(ch))
	// }
	//<-ch;<-ch;<-ch; no need
	// close(ch); no  need
	// while using range iteration stops when ch gets closed so
	// after len(ch) == 0 channel automatically closes in range loop
	// but if we use an other iterator we have to close channel
	ch<-1;ch<-2;ch<-3;
	for i := 1; i <= 3; i++ {
		fmt.Println(<-ch)
	}
	close(ch)

}
