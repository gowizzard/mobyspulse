//go:build development

package request_test

import (
	request "github.com/gowizzard/mobyspulse/internal/requests"
	"reflect"
	"testing"
)

// TestImage is to test the Image function.
func TestImage(t *testing.T) {

	containers, err := request.Containers()
	if err != nil {
		t.Error(err)
	}

	for _, value := range containers {

		container, err := request.Container(value.Id)
		if err != nil {
			t.Error(err)
		}

		image, err := request.Image(container.Image)
		if err != nil {
			t.Error(err)
		}

		if reflect.DeepEqual(image, request.ImageResponse{}) {
			t.Error("image struct is empty")
		}

	}

}
