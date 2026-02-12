package main

import (
	"slices"
	"testing"
)

func TestSortIntegers(t *testing.T) {
	var tests = []struct {
		name     string
		numbers  []int
		expected []int
	}{
		{
			name:     "Case empty slice",
			numbers:  nil,
			expected: nil,
		},
		{
			name:     "Case numbers sorted in ascending order",
			numbers:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Case numbers sorted in descending order",
			numbers:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Case not sorted numbers",
			numbers:  []int{0, 3, 2, 5, 4, 1, 9, 7, 8, 6},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			SortIntegers(test.numbers)
			if !slices.Equal(test.numbers, test.expected) {
				t.Errorf("unexpected sorting for: got %v, expected %v\n", test.numbers, test.expected)
			}
		})
	}
}
