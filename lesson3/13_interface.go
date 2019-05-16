package main

type Shape interface {
}

type OpenShape interface {
	length() float64
}

type ClosedShape interface {
	area() float64
	perimeter() float64
}

func main() {
}
