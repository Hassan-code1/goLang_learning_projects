package main

import (
	"fmt"
	"math"
	"strings"
)

type cordinate struct{
	x, y float64
}

func (c cordinate) distanceFromOrigin() float64{
	sqr := c.x*c.x + c.y*c.y
	res := math.Sqrt(sqr)
	return res
}// this is a method on struct cordinate with reciever
//it is not necessary to define a method for a struct only 
//we can define it for any type

func distanceFromOriginFunction(c cordinate) float64{
	sqr := c.x*c.x + c.y*c.y
	res := math.Sqrt(sqr)
	return res
}// this is a function with same functionality but it is not specific for cordinate

type MyFloat float64
func (f MyFloat)abs() float64{
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}// it is a method defined on MyFloat
//if you define some var by MyFloat you can use this method on that
// but not on normal float64 variable 

//The above methods used Value Recievers
// now we will use pointer recievers
func (f *MyFloat) scale() {
	*f *= 10
}

func main(){
	// cor := cordinate{3, 4}
	// dist := cor.distanceFromOrigin()
	// fmt.Printf("x:%f, y:%f dist:%f", cor.x, cor.y, dist)
	var fl  (MyFloat) = -45.00
	addressOffl := &fl
	addressOffl.scale()// given address to method //  it will work

	//or we can directly do this
	fl.scale() // notice i did not gave &fl to method but it still works
	
	fmt.Println(fl.abs())

	// var fl1 (float64) = -55.00
	// fmt.Println(fl1.abs()) // wrong 

	text := "hello , **Hassan**"
	parts := strings.Split(text, "**")
	for i := 1; i < len(parts); i += 2 {
		parts[i] = "<strong>" + parts[i] + "</strong>"
	}
	text = strings.Join(parts, "")
	fmt.Println(text)
	return
}