package main

import (
	"bitwise-operations-ex02/set"
	"fmt"
)

func main() {
	s := set.Empty
	fmt.Println(set.String(s))
	fmt.Printf("%t\n", set.IsEmpty(s))
	fmt.Printf("%v\n", set.Len(s))
	fmt.Printf("%v\n", set.Elements(s))
	sMod, err := set.Add(s, 64)
	fmt.Printf("%b, %v\n", sMod, err)
	fmt.Printf("%t\n", set.Contains(s, 5))
	fmt.Printf("%b\n", set.Remove(s, 2))
	fmt.Printf("%b\n", set.Union(s, 0b10))
	fmt.Printf("%b\n", set.Intersection(s, 0b100100))
	fmt.Printf("%b\n", set.Difference(s, 0b100111))
	fmt.Printf("%b\n", set.Subtract(s, 0b11))
}
