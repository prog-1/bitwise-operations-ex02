package main

import (
	"bitwise-operations-ex02/set"
	"fmt"
)

func main() {
	s1 := set.Empty // empty
	fmt.Println("Empty set.String:", set.String(s1))
	fmt.Println("Empty set.IsEmpty:", set.IsEmpty(s1))
	fmt.Println("Empty set.Len:", set.Len(s1))
	fmt.Println("Empty set.Elements:", set.Elements(s1))

	s2 := set.Set(0b0101)   // {0, 2}
	s2another := set.Set(5) // also {0, 2}
	fmt.Println("set.String:", set.String(s2))
	fmt.Println("set.String (another):", set.String(s2another))
	fmt.Println("set.IsEmpty", set.IsEmpty(s2))
	fmt.Println("set.Len:", set.Len(s2))
	fmt.Println("set.Elements:", set.Elements(s2))
	setadd, _ := set.Add(s2, 4)
	fmt.Println("set.Add:", set.String(setadd))
	fmt.Println("set.Contains #1:", set.Contains(s2, 2))
	fmt.Println("set.Contains #2:", set.Contains(s2, 3))
	fmt.Println("set.Remove #1:", set.String(set.Remove(s2, 0)))
	fmt.Println("set.Remove #2:", set.String(set.Remove(s2, 1)))
	fmt.Println("set.Remove #3:", set.String(set.Remove(s2, 2)))
	s3, s4 := set.Set(0b0111), set.Set(0b11010) // {0, 1, 2}, {1, 3, 4}
	fmt.Println("set.Union:", set.String(set.Union(s3, s4)))
	s5 := set.Set(0b11100) // {2, 3, 5}
	fmt.Println("set.Intersection:", set.String(set.Intersection(s3, s5)))
	s6 := set.Set(0b100010) // {1, 5}
	fmt.Println("set.Difference:", set.String(set.Difference(s3, s6)))
	s7 := set.Set(0b10010) // {1, 4}
	fmt.Println("set.Subtract:", set.String(set.Subtract(s3, s7)))
}
