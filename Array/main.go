package main

import "fmt"

func main() {
	arr := [...]int{10, 37, 86, 77, 94, 23, 8, 15}
	min := arr[0]
	max := arr[0]
	for _, v := range arr {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	fmt.Println("Minimum value:", min)
	fmt.Println("Maximum value:", max)
}
