package validator

import (
	"testing"
)

const (
	validJwt   = "eyJhbGciOiJFUzI1NiIsImtpZCI6IkFyU3lpVmhGbzFEVlFZUjBQNEVWR1Z1UnZTeFNZYWtsVGFQcDlNT1ZfUE09In0.eyJhdWQiOiJteS5qd3QuYXVkaWVuY2UiLCJleHAiOjE2NTM3OTU2NzMsImdyb3VwcyI6WyJhZG1pbiIsImRldmVsb3BlciJdLCJpYXQiOjE2NTM3OTUzNzMsImlzcyI6Im15Lmp3dC5pc3N1ZXIiLCJqdGkiOiJTdWJqZWN0IG9mIHRoZSBKV1QiLCJuYW1lIjoiRnVsbHkgUXVhbGlmaWVkIE5hbWUgb2YgdGhlIFN1YmplY3QiLCJuYmYiOjE2NTM3OTUzNzMsInN1YiI6IlN1YmplY3Qgb2YgdGhlIEpXVCJ9.7dlsFAH9E9mwwWVZLU7D0wBl_LEBUywIkvhfafSJOPT8MQWpMCY_ppG1kZTRk8wqUcHWmK371tXV1IpnsHGqzQ"
	invalidJwt = "eyJhbGciOiJFUzI1NiIsImtpZCI6IkFyU3lpVmhGbzFEVlFZUjBQNEVWR1Z1UnZTeFNZYWtsVGFQcDlNT1ZfUE09In0.eyJhdWQiOiJteS5qd3QuYXVkaWVuY2UiLCJleHAiOjE2NTM3OTU4ODMsImdyb3VwcyI6WyJkZXZlbG9wZXIiXSwiaWF0IjoxNjUzNzk1NTgzLCJpc3MiOiJteS5qd3QuaXNzdWVyIiwianRpIjoiU3ViamVjdCBvZiB0aGUgSldUIiwibmFtZSI6IkZ1bGx5IFF1YWxpZmllZCBOYW1lIG9mIHRoZSBTdWJqZWN0IiwibmJmIjoxNjUzNzk1NTgzLCJzdWIiOiJTdWJqZWN0IG9mIHRoZSBKV1QifQ.eJ8tyEIhbQpPCi68XwZuyWsrtWdwrr9jYS6-pXQsDktFlIVcwMLqJJ4RvvP62AwD5jJ2F896_zYxQsCrVKHfoA"
)

var testCases = []struct {
	name           string
	token          string
	expectedResult bool
}{
	{"Test Invalid JWT", invalidJwt, false},
	{"Test Valid JWT", validJwt, true},
}

func TestHttpAuthz(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := IsJWTAuthorized(&tt.token)
			if result != tt.expectedResult {
				t.Errorf("[Error] Expected %v but got %v", tt.expectedResult, result)
			}
		})
	}
}
