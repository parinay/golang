package math

import "testing"

func TestAddition(t *testing.T) {
	cases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"add", 2, 2, 4},
		{"minus", 0, -2, -2},
		{"zero", 0, 0, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := Sum(tc.a, tc.b)
			if result != tc.expected {
				t.Fatalf("expected sum %v, but got %v", tc.expected, result)
			}
		})
	}

}
