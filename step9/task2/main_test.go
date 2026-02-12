package main

import "testing"

func TestContains(t *testing.T) {
	var tests = []struct {
		name     string
		numbers  []int
		target   int
		expected bool
	}{
		{
			name:     "Case contains a number",
			numbers:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			target:   0,
			expected: true,
		},
		{
			name:     "Case does not contain a number",
			numbers:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			target:   10,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := Contains(test.numbers, test.target)
			if got != test.expected {
				t.Errorf("unexpected value for %v, %v: got %v, expected %v\n", test.numbers, test.target, got, test.expected)
			}
		})
	}
}
