package main

import (
	"testing"
	"time"
)

func TestTimeoutFibonacci(t *testing.T) {
	var tests = []struct {
		name          string
		n             int
		timeout       time.Duration
		expected      int
		isExpectedErr bool
		expectedErr   error
	}{
		{
			name:          "Case negative n error",
			n:             -1,
			timeout:       10 * time.Millisecond,
			expected:      0,
			isExpectedErr: true,
			expectedErr:   ErrNegative,
		},
		{
			name:     "Case 0th Fibonacci number",
			n:        0,
			timeout:  10 * time.Millisecond,
			expected: 0,
		},
		{
			name:     "Case 1th Fibonacci number",
			n:        1,
			timeout:  10 * time.Millisecond,
			expected: 1,
		},
		{
			name:     "Case 2th Fibonacci number",
			n:        2,
			timeout:  10 * time.Millisecond,
			expected: 1,
		},
		{
			name:     "Case 40th Fibonacci number",
			n:        40,
			timeout:  10 * time.Millisecond,
			expected: 102334155,
		},
		{
			name:     "Case 92th Fibonacci number",
			n:        92,
			timeout:  10 * time.Millisecond,
			expected: 7540113804746346429,
		},
		{
			name:          "Case timeout zero",
			n:             92,
			timeout:       0,
			expected:      0,
			isExpectedErr: true,
			expectedErr:   ErrTimeout,
		},
		{
			name:          "Case timeout 1mcs",
			n:             1000,
			timeout:       1 * time.Microsecond,
			expected:      0,
			isExpectedErr: true,
			expectedErr:   ErrTimeout,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := time.Now()
			got, err := TimeoutFibonacci(test.n, test.timeout)
			duration := time.Since(start)

			if (err != nil) != test.isExpectedErr {
				t.Errorf("unexpected error: got %v, error is expected %v\n", err, test.isExpectedErr)
			}
			if test.expectedErr != nil && err != test.expectedErr {
				t.Errorf("unexpected error: got %v, expected %v\n", err, test.expectedErr)
			}

			if got != test.expected {
				t.Errorf("unexpected value for %v, %v: got %v, expected %v\n", test.n, test.timeout, got, test.expected)
			}

			if duration > test.timeout+2*time.Millisecond {
				t.Errorf("unexpected timeout: got %v, expected near %v\n", duration, test.timeout)

			}
		})
	}
}
