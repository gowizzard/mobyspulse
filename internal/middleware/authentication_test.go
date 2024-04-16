// Package middleware_test is to test the middleware package.
package middleware_test

import (
	"github.com/gowizzard/mobyspulse/internal/middleware"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

// TestAuthentication is to test the Authentication function.
func TestAuthentication(t *testing.T) {

	tests := []struct {
		name        string
		environment map[string]string
		username    string
		password    string
		expected    int
	}{
		{
			name: "AUTHENTICATION_SUCCESS",
			environment: map[string]string{
				"BASIC_AUTH_USERNAME": "mobys",
				"BASIC_AUTH_PASSWORD": "pulse",
			},
			username: "mobys",
			password: "pulse",
			expected: http.StatusOK,
		},
		{
			name: "AUTHENTICATION_FAILED",
			environment: map[string]string{
				"BASIC_AUTH_USERNAME": "mobys",
				"BASIC_AUTH_PASSWORD": "pulse",
			},
			username: "mobys",
			password: "pulse1",
			expected: http.StatusUnauthorized,
		},
		{
			name: "WRONG_AUTHENTICATION_VARIABLES",
			environment: map[string]string{
				"BASIC_AUTH_USERNAME": "mobys",
			},
			username: "mobys",
			password: "pulse",
			expected: http.StatusOK,
		},
		{
			name:        "NO_AUTHENTICATION",
			environment: map[string]string{},
			username:    "",
			password:    "",
			expected:    http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for index, value := range tt.environment {
				t.Setenv(index, value)
			}

			middleware.Config.Username, middleware.Config.Ok = os.LookupEnv("BASIC_AUTH_USERNAME")
			middleware.Config.Password, middleware.Config.Ok = os.LookupEnv("BASIC_AUTH_PASSWORD")

			request, err := http.NewRequest(http.MethodGet, "/", nil)
			if err != nil {
				t.Fatal(err)
			}
			request.SetBasicAuth(tt.username, tt.password)

			recorder := httptest.NewRecorder()
			handler := middleware.Authentication(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			handler.ServeHTTP(recorder, request)

			if !reflect.DeepEqual(recorder.Code, tt.expected) {
				t.Errorf("handler returned wrong status code: got %v want %v", recorder.Code, tt.expected)
			}

		})
	}

}
