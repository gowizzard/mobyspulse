//go:build development

// Package request_test is to test the request package.
package request_test

import (
	"github.com/gowizzard/mobyspulse/internal/request"
	"testing"
)

// TestVersion is to test the Version function.
func TestVersion(t *testing.T) {

	response, err := request.Version()
	if err != nil {
		t.Error(err)
	}

	t.Logf("response: %v", response)

}
