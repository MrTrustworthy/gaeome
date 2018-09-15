package main

import (
	"fmt"
	"math"
)


type Point struct {
	X float64
	Y float64
}

func (p Point) Abs() float64 {
	return math.Sqrt(math.Pow(p.X, 2) + math.Pow(p.Y, 2))
}

func (p *Point) Zero() {
	p.X = 0
	p.Y = 0
}

func One(p *Point){
	p.X = 1
	p.Y = 1
}

func main(){
	myobj := Point{4.2, 3.0}
	fmt.Println(myobj)
	fmt.Println("Abs:", myobj.Abs())

	myobj2 := Point{4.2, 3.0}
	fmt.Println(myobj2)
	One(&myobj2)
	fmt.Println("Zeroes:", myobj2.Abs())
}
