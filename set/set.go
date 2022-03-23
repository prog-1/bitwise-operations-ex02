// Package set implements an algebraic set and set operations.
package set

// Set represents a set of elements. The elements are integers in range [0..63].
type Set uint64

// Empty represents an empty set.
const Empty = Set(0)

// String returns a string representation of the set.
func String(s Set) string {
	return ""
}

// IsEmpty returns true iff the set is empty.
func IsEmpty(s Set) bool {
	return false
}

// Len returns the number of elements in the set.
func Len(s Set) int {
	return 0
}

// Elements returns a slice of set elements.
func Elements(s Set) []int {
	return nil
}

// Add returns a new set that contains the integer `n`.
// An error must be return if the integer is not in range [0..63].
func Add(s Set, n int) (Set, error) {
	return Empty, nil
}

// Contains returns true iff the element `n` exists in the set.
func Contains(s Set, n int) bool {
	return false
}

// Remove returns a new set that does not contain the element `n`.
func Remove(s Set, n int) Set {
	return Empty
}

// Union returns a new set that is a union of two sets.
// The union contains elements that are present in *either* set.
func Union(s1, s2 Set) Set {
	return Empty
}

// Intersection returns a new set that is an intersection of two sets.
// The intersection contains elements that are present in *both* sets.
func Intersection(s1, s2 Set) Set {
	return Empty
}

// Difference returns a new set that is a difference of two sets.
// The difference contains elements that are not present in *both* sets.
func Difference(s1, s2 Set) Set {
	return Empty
}

// Subtract returns a new set that contains elements that are present
// in the first set, but not present in the second set.
func Subtract(s1, s2 Set) Set {
	return Empty
}
