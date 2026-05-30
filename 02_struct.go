package main

import(
	"fmt"
	"strings"
	// "math"
)
type coordinate struct {
	x int
	y int
}

func main() {
	//pointers

	i, j := 12, 13.09
	p := &i//p will give address of i and *p will give value at address pointed by pointer
	q := &j
	fmt.Println(*p, *q, p, q)

	//structures
	fmt.Println(coordinate{1,2})
	a := coordinate{4,5}
	fmt.Println(a.x)

	//arrays
	var arr [2]string
	arr[0] = "Hello"; arr[1] = "World!"
	primes := [5] int {2,3,5,7,9}
	fmt.Println(arr, primes)

	//slices
	var slice []int = primes[1:4]//index 1 included but not 4
	fmt.Println(slice)
	//changing anything in slice changes back in array
	slice[0] = 4 // it will change primes[1]
	fmt.Println(primes, slice)

	names := []struct{
		name string
		age int
	}{
		{"Hassan", 21},
		{"Anwar", 13},
		{"Amjad", 17},
	}
	fmt.Println(names)
	fmt.Println(primes[1:4], primes[:2], primes[2:], primes[:], len(primes[:2]), cap(primes[1:3]))
	// cap gives max elements in arr starting from  index i if slice is s[i,j]

	//we can create a slice with make function
	// make([]dtype, len, cap)
	mk := make([]int, 2, 5)
	fmt.Println(mk)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	fmt.Println(len(board), len(board[1]), board[0])
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// append to a slice
	var slicee []int
	fmt.Println(slicee)
	slicee = append(slicee, 0, 1)
	fmt.Println(slicee)

	//using rnage to iterate over loop or map
	for idx, val := range primes{
		fmt.Println(idx, ":", val)
	}
	// if you only want one variable you can omit other
	// for idx, _ := range
	// for idx := range
	// for _, val := range all these are correct

	//Maps
	mp := make(map[string]int32)//mp anything to anything even struct
	mp["age"] = 21
	fmt.Println(mp["age"])

	type vertex struct {
		lat, long float32
	}
	mp1 := map[string]vertex{
		"jodhpur":vertex{21, 12},
	}
	fmt.Println(mp1)
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}


func WordCount(s string) map[string]int {
	wordToCount := make(map[string]int)
	var word []string
	for i := 0; i < len(s); i++{
		if s[i:i+1] != " " {
			word = append(word, s[i:i+1])
		}else{
			wordToCount[strings.Join(word, "")]++
			word = word[:0]
		}
	}
	return wordToCount
}
