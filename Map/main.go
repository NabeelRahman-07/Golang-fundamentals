package main

import "fmt"

func main() {
	students := map[string]string{
		"Nabeel":  "A",
		"Rinshad": "B",
		"Fasil":   "A",
		"Afsal":   "C",
	}

	fmt.Println("Initial grades:", students)

	students["Fasil"] = "B"
	students["Rashid"] = "A"

	fmt.Println("Updated grades:", students)

}
