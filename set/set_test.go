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
		{3, "0, 1"},
		{5, "0, 2"},
		{15, "0, 1, 2, 3"},
	} {
		if got := String(tc.s); got != tc.want {
			t.Errorf("String(%v): got = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
func TestIsEmpty(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		want bool
	}{
		{0, true},
		{1, false},
		{128, false},
	} {
		if got := IsEmpty(tc.s); got != tc.want {
			t.Errorf("IsEmpty(%v): got = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

func TestLen(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		want int
	}{
		{0, 0},
		{1, 1},
		{5, 2},
		{7, 3},
	} {
		if got := Len(tc.s); got != tc.want {
			t.Errorf("Len(%v): got = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

func TestElements(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		want []int
	}{
		{0, []int{}},
		{1, []int{0}},
		{3, []int{0, 1}},
		{7, []int{0, 1, 2}},
		{15, []int{0, 1, 2, 3}},
	} {
		if got := Elements(tc.s); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Elements(%v): got = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

func TestAdd(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		n    int
		want Set
	}{
		{7, 5, 39},
		{7, 64, 0},
		{0, 0, 1},
		{7, -1, 0},
	} {
		if got, err := Add(tc.s, tc.n); got != tc.want {
			t.Errorf("Add(%v, %v): got = %v, %v, want = %v", tc.s, tc.n, got, err, tc.want)
		}
	}
}

func TestContains(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		n    int
		want bool
	}{
		{7, 2, true},
		{7, 3, false},
		{0, 3, false},
		{1, 0, true},
	} {
		if got := Contains(tc.s, tc.n); got != tc.want {
			t.Errorf("Contains(%v, %v): got = %v, want = %v", tc.s, tc.n, got, tc.want)
		}
	}
}

func TestRemove(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		n    int
		want Set
	}{
		{7, 1, 5},
		{7, 4, 7},
		{0, 0, 0},
		{1, 0, 0},
	} {
		if got := Remove(tc.s, tc.n); got != tc.want {
			t.Errorf("Remove(%v, %v): got = %v, want = %v", tc.s, tc.n, got, tc.want)
		}
	}
}

func TestUnion(t *testing.T) {
	for _, tc := range []struct {
		s1   Set
		s2   Set
		want Set
	}{
		{0, 0, 0},
		{1, 2, 3},
		{7, 26, 31},
		{7, 0, 7},
	} {
		if got := Union(tc.s1, tc.s2); got != tc.want {
			t.Errorf("Union(%v, %v): got = %v, want = %v", tc.s1, tc.s2, got, tc.want)
		}
	}
}

func TestIntersection(t *testing.T) {
	for _, tc := range []struct {
		s1   Set
		s2   Set
		want Set
	}{
		{7, 28, 4},
		{7, 0, 0},
		{0, 0, 0},
		{1, 1, 1},
	} {
		if got := Intersection(tc.s1, tc.s2); got != tc.want {
			t.Errorf("Intersection(%v, %v): got = %v, want = %v", tc.s1, tc.s2, got, tc.want)
		}
	}
}

func TestDifference(t *testing.T) {
	for _, tc := range []struct {
		s1   Set
		s2   Set
		want Set
	}{
		{0, 0, 0},
		{1, 2, 3},
		{7, 34, 37},
		{7, 7, 0},
	} {
		if got := Difference(tc.s1, tc.s2); got != tc.want {
			t.Errorf("Difference(%v, %v): got = %v, want = %v", tc.s1, tc.s2, got, tc.want)
		}
	}
}

func TestSubtract(t *testing.T) {
	for _, tc := range []struct {
		s1   Set
		s2   Set
		want Set
	}{
		{0, 0, 0},
		{1, 0, 1},
		{7, 18, 5},
		{7, 7, 0},
	} {
		if got := Subtract(tc.s1, tc.s2); got != tc.want {
			t.Errorf("Subtract(%v, %v): got = %v, want = %v", tc.s1, tc.s2, got, tc.want)
		}
	}
}
