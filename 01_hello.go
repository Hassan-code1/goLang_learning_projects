package main

import (
	"fmt"
	"math"
	"runtime"
)

func add(x int, y int) int{
	// you can also write it like
	// func add(x, y int) int{
	return x + y
}

func swap(x, y string) (string, string){
	return y, x
}

func split(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	// return // by default it will return named  return variables x,y this method is known as naked return
	return x, y // this is normal return 
	// issue with naked return is it make code less readable
	//thus it should only be used in short functions
}

// var a, b bool, c float this is incorrect
// var a, b, c bool it is correct variables will have default zero values(0, false, "")
var a,b,c int = 1,2,3 // variable initialized while declaring

func main() {
	fmt.Printf("hello workd %g \n", math.Sqrt(4))
	fmt.Println(math.Pi)
	// a name is exported if it begins with Cap letter 
	// like here math.Pi Pi is exprted from math package
	fmt.Println(add(2,3))

	x := "hello"//  := for assignment to new var known as short assignment
	y := "world"
	x,y = swap(x,y) // = for assigning to defined var
	fmt.Println(x,y)

	fmt.Println(split(18))

	var d uint16
	var e,f,g = true, 1, "Hello"// when initializing it is not nexessary to write dtype 
	// above line can also be written as
	// e,f,g := true, 1, "Hello"
	fmt.Println(a,b,c,d, e, f, g)
	
	
	//int Printf we can use %T , %v for any variable %T will give its type
	// and %v will give its value


	//Type conversions
	i := 12
	// f := float(i) while conversion we have to be explicit
	fl := float64(i) 
	cl := 5 + 5i // complex dtype
	// fl2 := float64(cl) not possible
	in := int32(fl)
	fmt.Println(i, fl, cl, in)
	fmt.Println(true && false)
	const myCountry = "India"

	//Flow Control

	for i := 0; i < 7; i++ {
		// fmt.Println(i)
	}
	if 1 < 2 {
		fmt.Println("if condn satisfied")
	}
	if test := 9; test < 8 {//this test varible will be in this scope only except it is accessible in the next else scope
		fmt.Println("variable defined in line of if")
	}else{
		fmt.Println(test)
	}

	fmt.Println(sQrt(3))

	switch os := runtime.GOOS; os{
	case "macOs":
		fmt.Println(os)
	case "linux":
		fmt.Println(os)
	default:
		fmt.Printf("%s.\n", os)
	}
}

func sQrt(x int) float64 {
	z := 1.00
	for i:= 1; i < 11; i++ {
		fmt.Println(z)
		z -= ((z*z)-float64(x))/(2*z)//Newton's method of sqrt finding
	}
	return z
}