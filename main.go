package main

import (
	"bitwise-operations-ex02/set"
	"fmt"
	"log"
)

func main() {
	// Create set
	s := set.Empty

	// Get set as string
	fmt.Println(set.String(s))

	// Check if set id Empty
	fmt.Println(set.IsEmpty(s))

	// Add number to set
	s, err := set.Add(s, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(set.String(s))

	// Length of set
	fmt.Println(set.Len(s))

	// Get set as slice of ints
	fmt.Println(set.Elements(s))

	// Check if set contains a number
	fmt.Println(set.Contains(s, 1))

	// Remove element from set
	s = set.Remove(s, 5)
	fmt.Println(set.String(s))

	s1 := set.Set(0b01111) // {0, 1, 2, 3}
	s2 := set.Set(0b10110) // {1, 2, 4}

	// Union
	fmt.Println(set.String(set.Union(s1, s2)))

	// Intersection
	fmt.Println(set.String(set.Intersection(s1, s2)))

	// Difference
	fmt.Println(set.String(set.Difference(s1, s2)))

	// Subtract
	fmt.Println(set.String(set.Subtract(s1, s2)))
}
