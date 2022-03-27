package set

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	for _, tc := range []struct {
		input Set
		want  string
	}{
		{0b111, "{0, 1, 2}"},
		{Empty, "{}"},
	} {
		t.Run("", func(t *testing.T) {
			if got := String(tc.input); got != tc.want {
				t.Errorf("String(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	for _, tc := range []struct {
		input Set
		want  bool
	}{
		{0b0, true},
		{0b1, false},
	} {
		t.Run("", func(t *testing.T) {
			if got := IsEmpty(tc.input); got != tc.want {
				t.Errorf("IsEmpty(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestLen(t *testing.T) {
	for _, tc := range []struct {
		input Set
		want  int
	}{
		{0b0, 0},
		{0b1, 1},
		{0b101, 2},
	} {
		t.Run("", func(t *testing.T) {
			if got := Len(tc.input); got != tc.want {
				t.Errorf("Len(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestElements(t *testing.T) {
	for _, tc := range []struct {
		input Set
		want  []int
	}{
		{0b0, nil},
		{0b1, []int{0}},
		{0b101, []int{0, 2}},
	} {
		t.Run("", func(t *testing.T) {
			if got := Elements(tc.input); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Elements(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	for _, tc := range []struct {
		input  Set
		input2 int
		want   Set
		want2  error
	}{
		{0b1, 5, 0b100001, nil},
		{Empty, 64, 0b0, OutOfRange},
		{Empty, 63, 1 << 63, nil},
	} {
		t.Run("", func(t *testing.T) {
			if got, err := Add(tc.input, tc.input2); got != tc.want || err != tc.want2 {
				t.Errorf("Add(%v, %v) = %v, %v, want %v, %v", tc.input, tc.input2, got, err, tc.want, tc.want2)
			}
		})
	}
}

func TestContains(t *testing.T) {
	for _, tc := range []struct {
		input  Set
		input2 int
		want   bool
	}{
		{0b0, 0, false},
		{0b1, 0, true},
		{0b100, 0, false},
		{0b100, 2, true},
	} {
		t.Run("", func(t *testing.T) {
			if got := Contains(tc.input, tc.input2); got != tc.want {
				t.Errorf("Contains(%v, %v) = %v, want %v", tc.input, tc.input2, got, tc.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	for _, tc := range []struct {
		input  Set
		input2 int
		want   Set
	}{
		{0b1, 0, 0b0},
		{0b0, 0, 0b0},
		{0b10, 0, 0b10},
		{0b10, 1, 0b0},
	} {
		t.Run("", func(t *testing.T) {
			if got := Remove(tc.input, tc.input2); got != tc.want {
				t.Errorf("Remove(%v, %v) = %v, want %v", tc.input, tc.input2, got, tc.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	for _, tc := range []struct {
		input  Set
		input2 Set
		want   Set
	}{
		{0b1, 0, 0b1},
		{0b10, 0b01, 0b11},
		{0b11, 0b01, 0b11},
	} {
		t.Run("", func(t *testing.T) {
			if got := Union(tc.input, tc.input2); got != tc.want {
				t.Errorf("Union(%v, %v) = %v, want %v", tc.input, tc.input2, got, tc.want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	for _, tc := range []struct {
		input  Set
		input2 Set
		want   Set
	}{
		{0b1, 0, 0b0},
		{0b10, 0b01, 0b0},
		{0b11, 0b01, 0b01},
	} {
		t.Run("", func(t *testing.T) {
			if got := Intersection(tc.input, tc.input2); got != tc.want {
				t.Errorf("Intersection(%v, %v) = %v, want %v", tc.input, tc.input2, got, tc.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	for _, tc := range []struct {
		input  Set
		input2 Set
		want   Set
	}{
		{0b1, 0, 0b1},
		{0b10, 0b01, 0b11},
		{0b11, 0b01, 0b10},
	} {
		t.Run("", func(t *testing.T) {
			if got := Difference(tc.input, tc.input2); got != tc.want {
				t.Errorf("Difference(%v, %v) = %v, want %v", tc.input, tc.input2, got, tc.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	for _, tc := range []struct {
		input  Set
		input2 Set
		want   Set
	}{
		{0b1, 0, 0b1},
		{0b10, 0b01, 0b10},
		{0b11, 0b01, 0b10},
	} {
		t.Run("", func(t *testing.T) {
			if got := Subtract(tc.input, tc.input2); got != tc.want {
				t.Errorf("Subtract(%v, %v) = %v, want %v", tc.input, tc.input2, got, tc.want)
			}
		})
	}
}
