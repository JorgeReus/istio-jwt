package istio

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

const (
	validJwt   = "eyJhbGciOiJFUzI1NiIsImtpZCI6IkFyU3lpVmhGbzFEVlFZUjBQNEVWR1Z1UnZTeFNZYWtsVGFQcDlNT1ZfUE09In0.eyJhdWQiOiJteS5qd3QuYXVkaWVuY2UiLCJleHAiOjE2NTM3OTU2NzMsImdyb3VwcyI6WyJhZG1pbiIsImRldmVsb3BlciJdLCJpYXQiOjE2NTM3OTUzNzMsImlzcyI6Im15Lmp3dC5pc3N1ZXIiLCJqdGkiOiJTdWJqZWN0IG9mIHRoZSBKV1QiLCJuYW1lIjoiRnVsbHkgUXVhbGlmaWVkIE5hbWUgb2YgdGhlIFN1YmplY3QiLCJuYmYiOjE2NTM3OTUzNzMsInN1YiI6IlN1YmplY3Qgb2YgdGhlIEpXVCJ9.7dlsFAH9E9mwwWVZLU7D0wBl_LEBUywIkvhfafSJOPT8MQWpMCY_ppG1kZTRk8wqUcHWmK371tXV1IpnsHGqzQ"
	invalidJwt = "eyJhbGciOiJFUzI1NiIsImtpZCI6IkFyU3lpVmhGbzFEVlFZUjBQNEVWR1Z1UnZTeFNZYWtsVGFQcDlNT1ZfUE09In0.eyJhdWQiOiJteS5qd3QuYXVkaWVuY2UiLCJleHAiOjE2NTM3OTU4ODMsImdyb3VwcyI6WyJkZXZlbG9wZXIiXSwiaWF0IjoxNjUzNzk1NTgzLCJpc3MiOiJteS5qd3QuaXNzdWVyIiwianRpIjoiU3ViamVjdCBvZiB0aGUgSldUIiwibmFtZSI6IkZ1bGx5IFF1YWxpZmllZCBOYW1lIG9mIHRoZSBTdWJqZWN0IiwibmJmIjoxNjUzNzk1NTgzLCJzdWIiOiJTdWJqZWN0IG9mIHRoZSBKV1QifQ.eJ8tyEIhbQpPCi68XwZuyWsrtWdwrr9jYS6-pXQsDktFlIVcwMLqJJ4RvvP62AwD5jJ2F896_zYxQsCrVKHfoA"
)

type HttpReqBuilder func() *http.Request

var httpTests = []struct {
	name                string
	expectedBody        string
	excpectedStatusCode int
	reqBuilder          HttpReqBuilder
}{
	{"Test Invalid JWT",
		denyBody, http.StatusForbidden, func() *http.Request {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", invalidJwt))
			return req
		}},

	{"Test Valid JWT",
		"", http.StatusOK, func() *http.Request {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", validJwt))
			return req
		}},
}

func TestHttpAuthz(t *testing.T) {
	// Start the service in a system defined port
	httpAuthz := NewHttpAuthorizer(0)
	var wg sync.WaitGroup
	wg.Add(1)
	go httpAuthz.Start(&wg)
	time.Sleep(500 * time.Millisecond)
	for _, tt := range httpTests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			httpAuthz.ServeHTTP(rec, tt.reqBuilder())
			resp := rec.Result()
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			// Verify body
			if string(data) != tt.expectedBody {
				t.Errorf("expected %v got %v", tt.expectedBody, string(data))
			}
			// Verify status code
			if resp.StatusCode != tt.excpectedStatusCode {
				t.Errorf("expected status code %v got %v", tt.excpectedStatusCode, resp.StatusCode)
			}
		})
	}
	httpAuthz.Stop()
	wg.Wait()
}
