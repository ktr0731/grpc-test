package server

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestServer(t *testing.T) {
	err := New().Serve().Stop()
	if err != nil {
		t.Errorf("Stop must not return an error, but got '%s'", err)
	}
}
