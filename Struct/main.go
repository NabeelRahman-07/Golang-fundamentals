package main

import "fmt"

type Person struct{
	Name string
	Age int
}
func (p Person)greet(){
	fmt.Println("Hello, my name is",p.Name)
	fmt.Println("I'am",p.Age,"years old.")
}

func main() {
	p1 := Person{
		Name : "Nabeel",
		Age : 20,
	}
	p1.greet()
}