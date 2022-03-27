package set

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		want string
	}{
		{0, "{}"},
		{0b100100, "{2, 5}"},
		{0b11110101, "{0, 2, 4, 5, 6, 7}"},
	} {
		got := String(tc.s)
		if got != tc.want {
			t.Errorf("ERR: String(%b): got = %s, want = %s", tc.s, got, tc.want)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		want bool
	}{
		{0, true},
		{0b1111010101, false},
		{0b11111111111, false},
	} {
		got := IsEmpty(tc.s)
		if got != tc.want {
			t.Errorf("ERR: IsEmpty(%b): got = %t, want = %t", tc.s, got, tc.want)
		}
	}
}

func TestLen(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		want int
	}{
		{0, 0},
		{0b10, 1},
		{0b11111, 5},
		{0b10000001, 2},
	} {
		got := Len(tc.s)
		if got != tc.want {
			t.Errorf("ERR: Len(%b): got = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

func TestElements(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		want []int
	}{
		{0, nil},
		{0b1001, []int{0, 3}},
		{0b111111, []int{0, 1, 2, 3, 4, 5, 6}},
	} {
		got := Elements(tc.s)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("ERR: Elements(%b): got = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

func TestAdd(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		n    int
		want Set
	}{
		{},
		{},
	} {
		got, _ := Add(tc.s, tc.n)
		if got != tc.want {
			t.Errorf("ERR: Add(%b, %v): got = %b, want = %b", tc.s, tc.n, got, tc.want)
		}
	}
}

func TestContains(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		n    int
		want bool
	}{
		{},
		{},
		{},
		{},
	} {
		got := Contains(tc.s, tc.n)
		if got != tc.want {
			t.Errorf("ERR: Contains(%b, %v): got = %t, want = %t", tc.s, tc.n, got, tc.want)
		}
	}
}

func TestRemove(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		n    int
		want Set
	}{
		{},
		{},
		{},
	} {
		got := Remove(tc.s, tc.n)
		if got != tc.want {
			t.Errorf("ERR: Remove(%b, %v): got = %b, want = %b", tc.s, tc.n, got, tc.want)
		}
	}
}

func TestUnion(t *testing.T) {
	for _, tc := range []struct {
		s1   Set
		s2   Set
		want Set
	}{
		{},
		{},
		{},
	} {
		got := Union(tc.s1, tc.s2)
		if got != tc.want {
			t.Errorf("ERR: Union(%b, %b): got = %b, want = %b", tc.s1, tc.s2, got, tc.want)
		}
	}
}

func TestIntersection(t *testing.T) {
	for _, tc := range []struct {
		s1   Set
		s2   Set
		want Set
	}{
		{},
		{},
		{},
		{0b10100111, 0b101010, 0b100010},
	} {
		got := Intersection(tc.s1, tc.s2)
		if got != tc.want {
			t.Errorf("ERR: Intersection(%b, %b): got = %b, want = %b", tc.s1, tc.s2, got, tc.want)
		}
	}
}

func TestDifference(t *testing.T) {
	for _, tc := range []struct {
		s1   Set
		s2   Set
		want Set
	}{
		{0, 1, 1},
		{0b101, 0b100, 0b001},
		{0b101010, 0b100011, 0b001001},
	} {
		got := Difference(tc.s1, tc.s2)
		if got != tc.want {
			t.Errorf("ERR: Difference(%b, %b): got = %b, want = %b", tc.s1, tc.s2, got, tc.want)
		}
	}
}

func TestSubtract(t *testing.T) {
	for _, tc := range []struct {
		s1   Set
		s2   Set
		want Set
	}{
		{0, 1, 0},
		{0b10111101, 0b1000000, 0b00111101},
		{0b1000111, 0b111, 0b1000000},
	} {
		got := Subtract(tc.s1, tc.s2)
		if got != tc.want {
			t.Errorf("ERR: Subtract(%b, %b): got = %b, want = %b", tc.s1, tc.s2, got, tc.want)
		}
	}
}
