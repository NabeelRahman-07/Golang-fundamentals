package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Circle struct{
	Radius float64
}
func (c Circle)Area() float64{
	return 3.14159 * c.Radius * c.Radius
}
func (c Circle)Perimeter() float64{
	return 2 * 3.14159 *c.Radius
}
type Rectangle struct{
	Length float64
	width float64
}
func (r Rectangle)Area() float64{
	return r.Length * r.width
}
func (r Rectangle)Perimeter() float64{
	return 2 * (r.Length + r.width)
}
func printshape(s Shape){
	fmt.Println("Area:",s.Area())
	fmt.Println("Perimeter:",s.Perimeter())
}

func main() {
c := Circle{Radius: 5}
r := Rectangle{Length: 4,
	 width: 3,}
	 printshape(r)
	 printshape(c)
}
