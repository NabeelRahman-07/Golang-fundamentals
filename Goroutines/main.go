package main

import "fmt"

func odd(limit int){
	for i := 1 ; i<=limit ; i++ {
		if (i%2)!=0{
			fmt.Println(i)
		}
	}
}
func even(limit int){
	for i := 1 ; i<=limit ; i++ {
		if (i%2)==0{
			fmt.Println(i)
		}
	}
}
func main(){
	var x int
	fmt.Println("Enter limit:")
	fmt.Scan(&x)
	go odd(x)
	go even(x)
}