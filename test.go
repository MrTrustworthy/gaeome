package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
	Zero()
}

type Point struct {
	X float64
	Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(math.Pow(p.X, 2) + math.Pow(p.Y, 2))
}

func (p *Point) Zero() {
	p.X = 0
	p.Y = 0
}

func main() {
	myobj := Point{4.2, 3.0}
	myobj2 := Point{2.2, 4.0}
	myobjSlice := make([]*Point, 0)
	myobjSlice = append(myobjSlice, &myobj, &myobj2)

	printOneAbs(&myobj)

	myObjInterfaceSlice := make([]Abser, len(myobjSlice))
	for i, obj := range myobjSlice {
		myObjInterfaceSlice[i] = Abser(obj)
	}

	printAllAbs(myObjInterfaceSlice)
	printAllAbs(myObjInterfaceSlice)

}

func printAllAbs(absables []Abser) {
	fmt.Println("All abs:")
	for _, absable := range absables {
		fmt.Println("Abs:", absable.Abs())
		absable.Zero()
	}
}

func printOneAbs(absable Abser) {
	fmt.Println("One Abs:")
	fmt.Println("Abs:", absable.Abs())
}
