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
		{0b11101, "{0, 2, 3, 4}"},
	} {
		t.Run("", func(t *testing.T) {
			if got := String(tc.s); got != tc.want {
				t.Errorf("String(%v) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		want bool
	}{
		{0, true},
		{0b100101, false},
		{0b1110, false},
	} {
		t.Run("", func(t *testing.T) {
			if got := IsEmpty(tc.s); got != tc.want {
				t.Errorf("IsEmpty(%v) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
}

func TestLen(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		want int
	}{
		{0, 0},
		{0b100111, 4},
		{0b11111, 5},
	} {
		t.Run("", func(t *testing.T) {
			if got := Len(tc.s); got != tc.want {
				t.Errorf("Len(%v) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
}

func TestElements(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		want []int
	}{
		{0, nil},
		{0b101110, []int{1, 2, 3, 5}},
	} {
		got := Elements(tc.s)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Elements(%b): got = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

/*func TestAdd(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		n    int
		want Set
	}{
		{},
	} {
		t.Run("", func(t *testing.T) {
			if got := Add(tc.s, tc.n); got != tc.want {
				t.Errorf("Add(%v, %v) = %v, want %v", tc.s, tc.n, got, tc.want)
			}
		})
	}
}
*/
// This test isn't correct

func TestContains(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		n    int
		want bool
	}{
		{0, 0, false},
		{0, 1, false},
		{0b110111, 2, true},
	} {
		t.Run("", func(t *testing.T) {
			if got := Contains(tc.s, tc.n); got != tc.want {
				t.Errorf("Contains(%v, %v) = %v, want %v", tc.s, tc.n, got, tc.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	for _, tc := range []struct {
		s    Set
		n    int
		want Set
	}{
		{0b100101, 2, 0b100001},
		{0b11001, 3, 0b10001},
	} {
		t.Run("", func(t *testing.T) {
			if got := Remove(tc.s, tc.n); got != tc.want {
				t.Errorf("Remove(%v, %v) = %v, want %v", tc.s, tc.n, got, tc.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	for _, tc := range []struct {
		s1   Set
		s2   Set
		want Set
	}{
		{0, 0b10, 0b10},
		{0b1100, 0b0111, 0b1111},
	} {
		t.Run("", func(t *testing.T) {
			if got := Union(tc.s1, tc.s2); got != tc.want {
				t.Errorf("Union(%v, %v) = %v, want %v", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	for _, tc := range []struct {
		s1   Set
		s2   Set
		want Set
	}{
		{0b111, 0b11100, 0b100},
		{0b1100, 0b11010, 0b1000},
	} {
		t.Run("", func(t *testing.T) {
			if got := Intersection(tc.s1, tc.s2); got != tc.want {
				t.Errorf("Intersection(%v, %v) = %v, want %v", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	for _, tc := range []struct {
		s1   Set
		s2   Set
		want Set
	}{
		{0b111, 0b100010, 0b100101},
		{0b1010, 0b111, 0b1101},
	} {
		t.Run("", func(t *testing.T) {
			if got := Difference(tc.s1, tc.s2); got != tc.want {
				t.Errorf("Difference(%v, %v) = %v, want %v", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	for _, tc := range []struct {
		s1   Set
		s2   Set
		want Set
	}{
		{0b10011, 0b001, 0b10010},
		{0b111, 0b10010, 0b101},
	} {
		t.Run("", func(t *testing.T) {
			if got := Subtract(tc.s1, tc.s2); got != tc.want {
				t.Errorf("Substract(%v, %v) = %v, want %v", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}
