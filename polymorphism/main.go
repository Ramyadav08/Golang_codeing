package main

import (
   "fmt"
   "math"
)

type Shape interface {
   Area() float64
}

type Circle struct {
   Radius float64
}

type Square struct {
   Length float64
}

func (c Circle) Area() float64 {
   return math.Pi * c.Radius * c.Radius
}

func (s Square) Area() float64 {
   return s.Length * s.Length
}

func PrintArea(s Shape) {
   fmt.Println("Area of shape is:", s.Area())
}

func main() {
   c := Circle{Radius: 5}
   s := Square{Length: 4}
   PrintArea(c)
   PrintArea(s)
}