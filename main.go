package main

import (
	"bitwise-operations-ex02/set"
	"fmt"
)

func main() {
	s := set.Set(123)
	s2 := set.Set(0b0101)
	n := 5
	fmt.Println("1. String returns a string representation of the set.")
	fmt.Println(set.String(s))
	fmt.Println(set.String(s2))
	fmt.Println("2. IsEmpty returns true iff the set is empty.")
	fmt.Println(set.IsEmpty(s))
	fmt.Println("3. Len returns the number of elements in the set.")
	fmt.Println(set.Len(s))
	fmt.Println("4. Elements returns a slice of set elements.")
	fmt.Println(set.Elements(s))
	fmt.Println("5. Add returns a new set that contains the integer `n`.")
	fmt.Println(set.Add(s, n))
	fmt.Println("6. Contains returns true iff the element `n` exists in the set.")
	fmt.Println(set.Contains(s, n))
	fmt.Println("7. Remove returns a new set that does not contain the element `n`.")
	fmt.Println(set.String(set.Remove(s, n)))
	fmt.Println("8. Union returns a new set that is a union of two sets.")
	fmt.Println(set.String(set.Union(s, s2)))
	fmt.Println("9. Intersection returns a new set that is an intersection of two sets.")
	fmt.Println(set.String(set.Intersection(s, s2)))
	fmt.Println(set.Intersection(s, s2))
	fmt.Println("10. Difference returns a new set that is a difference of two sets.")
	fmt.Println(set.String(set.Difference(s, s2)))
	fmt.Println("11. Subtract returns a new set that contains elements that are present.")
	fmt.Println(set.String(set.Subtract(s, s2)))
}
