//go:build development

// Package request_test is to test the request package.
package request_test

import (
	"github.com/gowizzard/mobyspulse/internal/request"
	"reflect"
	"testing"
)

// TestPing is to test the Ping function.
func TestPing(t *testing.T) {

	status, err := request.Ping()
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(status, 200) {
		t.Error("status is not 200")
	}

	t.Logf("status: %d", status)

}
