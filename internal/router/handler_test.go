// Package router_test is to test the router package.
package router_test

import (
	"github.com/gowizzard/mobyspulse/internal/handler"
	"github.com/gowizzard/mobyspulse/internal/redirect"
	"github.com/gowizzard/mobyspulse/internal/router"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// TestHandler is to test the Handler function.
func TestHandler(t *testing.T) {

	tests := []struct {
		name     string
		method   string
		path     string
		handler  http.HandlerFunc
		expected int
	}{
		{
			name:     "METRICS_REDIRECT",
			method:   http.MethodGet,
			path:     "/",
			handler:  redirect.Metrics,
			expected: http.StatusMovedPermanently,
		},
		{
			name:     "FAVICON_HANDLER",
			method:   http.MethodGet,
			path:     "/favicon.ico",
			handler:  handler.Favicon,
			expected: http.StatusOK,
		},
		{
			name:     "METRICS_HANDLER",
			method:   http.MethodGet,
			path:     "/metrics",
			handler:  handler.Metrics,
			expected: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			request, err := http.NewRequest(tt.method, tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()

			mux := http.NewServeMux()
			router.Handler(mux)
			mux.ServeHTTP(recorder, request)

			if !reflect.DeepEqual(recorder.Code, tt.expected) {
				t.Errorf("handler returned wrong status code: got %v want %v", recorder.Code, tt.expected)
			}

		})
	}

}
