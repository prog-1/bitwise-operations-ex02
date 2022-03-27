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
		{0b1111011, "{0, 1, 3, 4, 5, 6}"},
		{0b0101, "{0, 2}"},
		{0b1000000000000000000000000000001101, "{0, 2, 3, 33}"},
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
		{0b0101, false},
		{0b1000000000000000000000000000001101, false},
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
		{0b1111011, 6},
		{0b1000000000000000000000000000001101, 4},
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
		{0b1111011, []int{0, 1, 3, 4, 5, 6}},
		{0b1000000000000000000000000000001101, []int{0, 2, 3, 33}},
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
		{0b111101, 5, 0b111101},
		{0b1000000000000000000000000000001101, 22, 0b1000000000010000000000000000001101},
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
		{0, 1, false},
		{0b111101, 2, true},
		{0b1000000000000000000000000000001101, 8, false},
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
		{0b111101, 0, 0b111100},
		{0b111101, 3, 0b110101},
		{0b1000000000000000000000000000001101, 63, 0b1000000000000000000000000000001101},
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
		{0, 0b101, 0b101},
		{0b111101, 0b11000, 0b111101},
		{0b100000000000000000000000000000110, 0b1111000000000000000000000000000110, 0b1111000000000000000000000000000110},
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
		{0b111101, 0b100, 0b100},
		{0b111101, 0b100001, 0b100001},
		{0b100000000000000000000000000000110, 0b100000000000000000000000000111110, 0b100000000000000000000000000000110},
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
		{0b111101, 0b101111, 0b010010},
		{0b100000000000000000000000000000110, 0b110000000000000000000000000000110, 0b010000000000000000000000000000000},
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
		{0b111101, 0b101111, 0b010000},
		{0b100000000000000000000000000000110, 0b110000000000000000000000000000110, 0},
	} {
		got := Subtract(tc.s1, tc.s2)
		if got != tc.want {
			t.Errorf("ERR: Subtract(%b, %b): got = %b, want = %b", tc.s1, tc.s2, got, tc.want)
		}
	}
}
