// Package handler_test is to test the handler package.
package handler_test

import (
	"github.com/gowizzard/mobyspulse/internal/handler"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// TestFavicon is to test the Favicon function.
func TestFavicon(t *testing.T) {

	request, err := http.NewRequest("GET", "/favicon.ico", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	handler.Favicon(recorder, request)

	if !reflect.DeepEqual(recorder.Code, http.StatusOK) {
		t.Errorf("xxpected code %d, got %d", http.StatusTemporaryRedirect, recorder.Code)
	}

	if !reflect.DeepEqual(recorder.Header().Get("Content-Type"), "image/x-icon") {
		t.Errorf("expected code %s, got %s", "image/x-icon", recorder.Header().Get("Content-Type"))
	}

}
