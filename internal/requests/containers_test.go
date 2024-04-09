//go:build development

// Package request_test is to test the request package.
package request_test

import (
	request "github.com/gowizzard/mobyspulse/internal/requests"
	"reflect"
	"testing"
)

// TestContainers is to test the Containers function.
func TestContainers(t *testing.T) {

	containers, err := request.Containers()
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(containers, []request.ContainerResponse{}) {
		t.Error("containers array of structs is empty")
	}

}

// TestContainer is to test the Container function.
func TestContainer(t *testing.T) {

	containers, err := request.Containers()
	if err != nil {
		t.Error(err)
	}

	for _, value := range containers {

		container, err := request.Container(value.Id)
		if err != nil {
			t.Error(err)
		}

		if reflect.DeepEqual(container, request.ContainerResponse{}) {
			t.Error("container struct is empty")
		}

	}

}
