package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"testing"
)

type Params struct {
	Name string
	ISE  bool
}

func ParseParams(r *http.Request) (Params, error) {
	query := r.URL.Query()
	params := Params{}

	params.Name = query.Get("name")
	if params.Name == "" {
		return params, fmt.Errorf("missing name")
	}

	iseStr := query.Get("ise")
	if iseStr == "" {
		params.ISE = false
	} else {
		ise, err := strconv.ParseBool(iseStr)
		if err != nil {
			params.ISE = false
		} else {
			params.ISE = ise
		}
	}

	return params, nil
}

func NewTestServer(addr string) *http.Server {
	studentMap := map[string]int{
		"Jack50": 50,
		"John40": 40,
		"Bob50":  50,
		"Sara60": 60,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/mark", func(w http.ResponseWriter, r *http.Request) {
		params, err := ParseParams(r)
		if err != nil {
			http.Error(w, "", http.StatusNotFound)
			return
		}

		if params.ISE {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		mark, found := studentMap[params.Name]
		if !found {
			http.Error(w, "", http.StatusNotFound)
			return
		}

		fmt.Fprintf(w, "%d", mark)
	})

	return &http.Server{Addr: addr, Handler: mux}
}
func TestCompare(t *testing.T) {
	var tests = []struct {
		name            string
		studentName1    string
		studentName2    string
		expected        string
		isExpectedError bool
		expectedError   error
	}{
		{
			name:         "Case firs greater than second",
			studentName1: "Jack50",
			studentName2: "John40",
			expected:     ">",
		},
		{
			name:         "Case firs equals second",
			studentName1: "Jack50",
			studentName2: "Bob50",
			expected:     "=",
		},
		{
			name:         "Case firs less than second",
			studentName1: "Jack50",
			studentName2: "Sara60",
			expected:     "<",
		},
		{
			name:         "Case identical names",
			studentName1: "Jack50",
			studentName2: "Jack50",
			expected:     "=",
		},
		{
			name:            "Case student not found",
			studentName1:    "Jack50",
			studentName2:    "",
			expected:        "",
			isExpectedError: true,
			expectedError:   StudentNotFoundError,
		},
		{
			name:            "Case internal server error",
			studentName1:    "Jack50",
			studentName2:    "Barbara25&ise=true",
			expected:        "",
			isExpectedError: true,
			expectedError:   InternalServerError,
		},
	}

	testServer := NewTestServer(":8082")

	go func() {
		if err := testServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("error in the test server: %s\n", err)
		}
	}()

	defer func() {
		if err := testServer.Close(); err != nil {
			log.Printf("error stopping the test server: %s\n", err)
		}
	}()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Compare(test.studentName1, test.studentName2)

			if (err != nil) != test.isExpectedError {
				t.Errorf("unexpected error: got %v, error is expected %v\n", err, test.isExpectedError)
			}
			if err != test.expectedError && test.expectedError != nil {
				t.Errorf("unexpected error: got %v, expected %v\n", err, test.expectedError)
			}

			if got != test.expected {
				t.Errorf("unexpected value: got %v, expected %v\n", got, test.expected)
			}
		})
	}
}
