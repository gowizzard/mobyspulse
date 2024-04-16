//go:build development

package handler_test

import (
	"github.com/gowizzard/mobyspulse/internal/handler"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// TestMetrics is to test the Metrics function.
func TestMetrics(t *testing.T) {

	request, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	handler.Metrics(recorder, request)

	if !reflect.DeepEqual(recorder.Code, http.StatusOK) {
		t.Errorf("expected code %d, got %d", http.StatusOK, recorder.Code)
	}

}
