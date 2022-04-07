package set

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    Set
		want string
	}{
		{"1", 0, "{}"},
		{"2", 0b100101, "{0, 2, 5}"},
		{"3", 37 /* the same as 0b100101 */, "{0, 2, 5}"},
		{"4", 0b10111, "{0, 1, 2, 4}"},
		{"5", 0b0100000000000000000000000000000000000000000000000000000000000011, "{0, 1, 62}"},
		{"6", 0b1000000000000000000000000000000000000000000000000000000000000011, "{0, 1, 63}"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := String(tc.s); got != tc.want {
				t.Errorf("String(%b): got = %s, want = %s", tc.s, got, tc.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    Set
		want bool
	}{
		{"1", 0, true},
		{"2", 0b100101, false},
		{"3", 0b1000000000000000000000000000000000000000000000000000000000000011, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsEmpty(tc.s); got != tc.want {
				t.Errorf("IsEmpty(%b): got = %t, want = %t", tc.s, got, tc.want)
			}
		})
	}
}

func TestLen(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    Set
		want int
	}{
		{"1", 0, 0},
		{"2", 0b100101, 3},
		{"3", 0b10111, 4},
		{"4", 0b1000000000000000000000000000000000000000000000000000000000000011, 3},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Len(tc.s); got != tc.want {
				t.Errorf("Len(%b): got = %v, want = %v", tc.s, got, tc.want)
			}
		})
	}
}

func TestElements(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    Set
		want []int
	}{
		{"1", 0, nil},
		{"2", 0b100101, []int{0, 2, 5}},
		{"3", 0b10111, []int{0, 1, 2, 4}},
		{"4", 0b0100000000000000000000000000000000000000000000000000000000000011, []int{0, 1, 62}},
		{"5", 0b1000000000000000000000000000000000000000000000000000000000000011, []int{0, 1, 63}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Elements(tc.s); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Elements(%b): got = %v, want = %v", tc.s, got, tc.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    Set
		n    int
		want Set
	}{
		{"1", 0, 1, 0b10},
		{"2", 0b100101, 4, 0b110101},
		{"3", 0b10111, 62, 0b0100000000000000000000000000000000000000000000000000000000010111},
		{"4", 0b1000000000000000000000000000000000000000000000000000000000000011, 62, 0b1100000000000000000000000000000000000000000000000000000000000011},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, _ := Add(tc.s, tc.n); got != tc.want {
				t.Errorf("Add(%b, %v): got = %b, want = %b", tc.s, tc.n, got, tc.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    Set
		n    int
		want bool
	}{
		{"1", 0, 0, false},
		{"2", 0, 1, false},
		{"3", 0b100101, 2, true},
		{"4", 0b10111, 3, false},
		{"5", 0b1000000000000000000000000000000000000000000000000000000000000011, 62, false},
		{"6", 0b1000000000000000000000000000000000000000000000000000000000000011, 63, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Contains(tc.s, tc.n); got != tc.want {
				t.Errorf("Contains(%b, %v): got = %t, want = %t", tc.s, tc.n, got, tc.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    Set
		n    int
		want Set
	}{
		{"1", 0b100101, 2, 0b100001},
		{"2", 0b10111, 3, 0b10111},
		{"3", 0b10111, 0, 0b10110},
		{"4", 0b1000000000000000000000000000000000000000000000000000000000000011, 63, 0b11},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Remove(tc.s, tc.n); got != tc.want {
				t.Errorf("Remove(%b, %v): got = %b, want = %b", tc.s, tc.n, got, tc.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	for _, tc := range []struct {
		name   string
		s1, s2 Set
		want   Set
	}{
		{"1", 0, 0b10, 0b10},
		{"2", 0b100101, 0b10, 0b100111},
		{"3", 0b100101, 0b11000, 0b111101},
		{"4", 0b1000000000000000000000000000000000000000000000000000000000000011, 0b0100000000000000000000000000000000000000000000000000000000000000, 0b1100000000000000000000000000000000000000000000000000000000000011},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Union(tc.s1, tc.s2); got != tc.want {
				t.Errorf("Union(%b, %b): got = %b, want = %b", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	for _, tc := range []struct {
		name   string
		s1, s2 Set
		want   Set
	}{
		{"1", 0, 1, 0},
		{"2", 0b100101, 1, 1},
		{"3", 0b100101, 0b100, 0b100},
		{"4", 0b10100111, 0b101010, 0b100010},
		{"5", 0b1100000000000000000000000000000000000000000000000000000000000011, 0b0100000001100000000111100000000010000001100000000010000000000010, 0b100000000000000000000000000000000000000000000000000000000000010},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Intersection(tc.s1, tc.s2); got != tc.want {
				t.Errorf("Intersection(%b, %b): got = %b, want = %b", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	for _, tc := range []struct {
		name   string
		s1, s2 Set
		want   Set
	}{
		{"1", 0, 1, 1},
		{"2", 0b101, 0b111, 0b10},
		{"3", 0b100101, 0b1011, 0b101110},
		{"4", 0b1000000000000000000000000000000000000000000000000000000000000011, 0b0100000000000000000000000000000000000000000000000000000000000010, 0b1100000000000000000000000000000000000000000000000000000000000001},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Difference(tc.s1, tc.s2); got != tc.want {
				t.Errorf("Difference(%b, %b): got = %b, want = %b", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	for _, tc := range []struct {
		name   string
		s1, s2 Set
		want   Set
	}{
		{"1", 0, 1, 0},
		{"2", 0b101, 1, 0b100},
		{"3", 0b100101, 0b11, 0b100100},
		{"4", 0b1000000000000000000000000000000000000000000000000000000000000011, 0b1000000000000000000000000000000000000000000000000000000000000000, 0b11},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Subtract(tc.s1, tc.s2); got != tc.want {
				t.Errorf("Subtract(%b, %b): got = %b, want = %b", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}
