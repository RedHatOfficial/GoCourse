package main

// Shape represents any 2D shape drawn in 2D plane
type Shape interface {
}

// OpenShape represents an "open" 2D shape, such as line, arc, spline etc.
type OpenShape interface {
	length() float64
}

// ClosedShape represents a "closed" 2D shape, such as square, circle, and ellipse
type ClosedShape interface {
	area() float64
	perimeter() float64
}

func main() {
}
