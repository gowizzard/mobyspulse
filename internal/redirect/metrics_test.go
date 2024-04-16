// Package redirect_test is to test the redirect package.
package redirect_test

import (
	"github.com/gowizzard/mobyspulse/internal/redirect"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// TestMetrics is to test the Metrics function.
func TestMetrics(t *testing.T) {

	request, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(redirect.Metrics)

	handler.ServeHTTP(recorder, request)

	if !reflect.DeepEqual(recorder.Code, http.StatusMovedPermanently) {
		t.Errorf("handler returned wrong status code: got %v want %v", recorder.Code, http.StatusMovedPermanently)
	}

	if !reflect.DeepEqual(recorder.Header().Get("Location"), "/metrics") {
		t.Errorf("handler returned wrong location header: got %v want %v", recorder.Header().Get("Location"), "/metrics")
	}

}
