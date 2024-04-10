package request_test

import (
	"github.com/gowizzard/mobyspulse/internal/request"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
	request.Client.CloseIdleConnections()
}
