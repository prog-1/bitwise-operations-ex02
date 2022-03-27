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
		{0, ""},
		{0b100101, "0 2 5 "},
		{0b10111, "0 1 2 4 "},
		{0b0100000000000000000000000000000000000000000000000000000000000011, "0 1 62 "},
		{0b1000000000000000000000000000000000000000000000000000000000000011, "0 1 63 "},
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
		{0b100101, false},
		{0b1000000000000000000000000000000000000000000000000000000000000011, false},
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
		{0b100101, 3},
		{0b10111, 4},
		{0b1000000000000000000000000000000000000000000000000000000000000011, 3},
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
		{0b100101, []int{0, 2, 5}},
		{0b10111, []int{0, 1, 2, 4}},
		{0b0100000000000000000000000000000000000000000000000000000000000011, []int{0, 1, 62}},
		{0b1000000000000000000000000000000000000000000000000000000000000011, []int{0, 1, 63}},
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
		{0, 1, 0b10},
		{0b100101, 4, 0b110101},
		{0b10111, 62, 0b0100000000000000000000000000000000000000000000000000000000010111},
		{0b1000000000000000000000000000000000000000000000000000000000000011, 62, 0b1100000000000000000000000000000000000000000000000000000000000011},
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
		{0, 0, false},
		{0, 1, false},
		{0b100101, 2, true},
		{0b10111, 3, false},
		{0b1000000000000000000000000000000000000000000000000000000000000011, 62, false},
		{0b1000000000000000000000000000000000000000000000000000000000000011, 63, true},
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
		{0b100101, 2, 0b100001},
		{0b10111, 3, 0b10111},
		{0b10111, 0, 0b10110},
		{0b1000000000000000000000000000000000000000000000000000000000000011, 63, 0b11},
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
		{0, 0b10, 0b10},
		{0b100101, 0b10, 0b100111},
		{0b100101, 0b11000, 0b111101},
		{0b1000000000000000000000000000000000000000000000000000000000000011, 0b0100000000000000000000000000000000000000000000000000000000000000, 0b1100000000000000000000000000000000000000000000000000000000000011},
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
		{0, 1, 0},
		{0b100101, 1, 1},
		{0b100101, 0b100, 0b100},
		{0b10100111, 0b101010, 0b100010},
		{0b1100000000000000000000000000000000000000000000000000000000000011, 0b0100000001100000000111100000000010000001100000000010000000000010, 0b100000000000000000000000000000000000000000000000000000000000010},
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
		{0b101, 0b111, 0b10},
		{0b100101, 0b1011, 0b101110},
		{0b1000000000000000000000000000000000000000000000000000000000000011, 0b0100000000000000000000000000000000000000000000000000000000000010, 0b1100000000000000000000000000000000000000000000000000000000000001},
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
		{0b101, 1, 0b100},
		{0b100101, 0b11, 0b100100},
		{0b1000000000000000000000000000000000000000000000000000000000000011, 0b1000000000000000000000000000000000000000000000000000000000000000, 0b11},
	} {
		got := Subtract(tc.s1, tc.s2)
		if got != tc.want {
			t.Errorf("ERR: Subtract(%b, %b): got = %b, want = %b", tc.s1, tc.s2, got, tc.want)
		}
	}
}
