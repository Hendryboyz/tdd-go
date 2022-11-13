package shape

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func Preimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

// func (receiverName ReceiverType) MethodName(args) returnType {}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
