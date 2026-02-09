package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
	"time"
)

func NewTestServer(responseTime time.Duration) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/provideData", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(responseTime)
		fmt.Fprintf(w, "%s", "data")
	})
	return &http.Server{Addr: ":8081", Handler: mux}
}

func TestStartServer(t *testing.T) {
	var tests = []struct {
		name           string
		timeout        time.Duration
		responseTime   time.Duration
		expectedStatus int
		expectedData   string
	}{
		{
			name:           "Case normal",
			timeout:        10 * time.Millisecond,
			responseTime:   0,
			expectedStatus: http.StatusOK,
			expectedData:   "data",
		},
		{
			name:           "Case timeout",
			timeout:        10 * time.Millisecond,
			responseTime:   20 * time.Millisecond,
			expectedStatus: http.StatusServiceUnavailable,
			expectedData:   "timeout",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testServer := NewTestServer(test.responseTime)

			go func() {
				if err := testServer.ListenAndServe(); err != http.ErrServerClosed {
					log.Fatalf("error in the test server: %s\n", err)
				}
			}()

			defer func() {
				if err := testServer.Close(); err != nil {
					t.Errorf("error stopping the test server: %s\n", err)
				}
			}()

			defer func() {
				if err := server.Close(); err != nil {
					log.Printf("error stopping the test server: %s\n", err)
				}
			}()

			go StartServer(test.timeout)

			var client http.Client

			request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/readSource", nil)
			if err != nil {
				t.Fatalf("unexpected request error: %s\n", err)
			}

			response, err := client.Do(request)
			if err != nil {
				t.Fatalf("unexpected respone error: %s\n", err)
			}
			defer response.Body.Close()

			body, err := io.ReadAll(response.Body)
			if err != nil {
				t.Fatalf("unexpected read error: %s\n", err)
			}

			status := response.StatusCode
			if status != test.expectedStatus {
				t.Errorf("unexpected response status: got %v, expected %v\n", status, test.expectedStatus)
			}

			data := string(body)
			if data != test.expectedData {
				t.Errorf("unexpected response data: got %v, expected %v\n", data, test.expectedData)
			}
		})
	}
}
