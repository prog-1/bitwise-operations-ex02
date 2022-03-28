// Package set implements an algebraic set and operations for sets.
package set

import (
	"fmt"
	"strconv"
)

// Set represents a set of elements. The elements are integers in range [0..63].
//
// Examples:
// {0, 1, 2, 63} is a valid set.
// {63, 64} is an invalid set.
type Set uint64

// Empty represents an empty set: {}.
const Empty = Set(0)

// String returns a string representation of the set, e.g. "{0, 1, 2}".
func String(s Set) (st string) {
	st += "{"
	for i := 0; i < 64; i, s = i+1, s>>1 {
		if s&1 == 1 {
			if st != "{" {
				st += ", "
			}
			st += strconv.Itoa(i)
		}
	}
	st += "}"
	return st
}

// IsEmpty returns true iff the set is empty.
//
// Examples:
// IsEmpty({}) -> true
// IsEmpty({5}) -> false.
func IsEmpty(s Set) bool {
	return s == 0 
}

// Len returns the number of elements in the set.
//
// Examples:
// Len({}) -> 0
// Len({0, 1, 2}) -> 3
func Len(s Set) (len int) {
	for i := 0; i < 64; i, s = i+1, s>>1 {
		if s&1 == 1 {
			len++
		}
	}
	return len
}

// Elements returns a slice of set elements.
//
// Examples:
// Elements({0, 1, 2}) -> []int{0, 1, 2}
func Elements(s Set) (result []int) {
	for i := 0; i < 64; i, s = i+1, s>>1 {
		if s&1 == 1 {
			result = append(result, i)
		}
	}
	return result
}

// Add returns a new set that contains the integer `n`.
// An error must be return if the integer is not in range [0..63].
//
// Examples:
// Add({0, 1, 2}, 5) -> {0, 1, 2, 5}, nil
// Add({0, 1, 2}, 64) -> {}, error
func Add(s Set, n int) (Set, error) {
	var err error
	if n >= 64 {
		fmt.Println("Can't add")
	}
	if n < 0 {
		fmt.Println("Can't add")
	}
	return (1 << n) | s, err
}

// Contains returns true iff the element `n` exists in the set.
//
// Examples:
// Contains({0, 1, 2}, 2) -> true
// Contains({0, 1, 2}, 3) -> false
func Contains(s Set, n int) bool {
	if s&(1<<n) > 0 {
		return true
	}
	return false
}

// Remove returns a new set that does not contain the element `n`.
//
// Examples:
// Remove({0, 1, 2}, 1) -> {0, 2}
// Remove({0, 1, 2}, 4) -> {0, 1, 2}
func Remove(s Set, n int) Set {
	return s &^ (1 << n)
}

// Union returns a new set that is a union of two sets.
// The union contains elements that are present in *either* set.
//
// Examples:
// Union({0, 1, 2}, {1, 3, 4}) -> {0, 1, 2, 3, 4}
// Union({0, 1, 2}, {}) -> {0, 1, 2}
func Union(s1, s2 Set) Set {
	return s1 | s2
}

// Intersection returns a new set that is an intersection of two sets.
// The intersection contains elements that are present in *both* sets.
//
// Examples:
// Intersection({0, 1, 2}, {2, 3, 4}) -> {2}
// Intersection({0, 1, 2}, {}) -> {}
func Intersection(s1, s2 Set) Set {
	return s1 & s2
}

// Difference returns a new set that is a difference of two sets.
// The difference contains elements that are not present in *both* sets.
//
// Examples:
// Difference({0, 1, 2}, {1, 5}) -> {0, 2, 5}
func Difference(s1, s2 Set) Set {
	return s1 ^ s2
}

// Subtract returns a new set that contains elements that are present
// in the first set, but not present in the second set.
//
// Examples:
// Subtract({0, 1, 2}, {1, 4}) -> {0, 2}
func Subtract(s1, s2 Set) Set {
	return s1 &^ s2
}
