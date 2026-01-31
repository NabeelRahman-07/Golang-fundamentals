package main

import "fmt"

func removeDuplicate(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, v := range arr {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func main() {
	sl := []int{10, 20, 30, 40, 30, 60, 10, 40, 90, 10, 80, 70, 10, 110, 120, 50, 60, 30, 100}

	unique := removeDuplicate(sl)

	fmt.Println(unique)
}
