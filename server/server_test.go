package server

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestServer(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		err := New(WithVerbose()).Serve().Stop()
		if err != nil {
			t.Errorf("Stop must not return an error, but got '%s'", err)
		}
	})

	t.Run("web", func(t *testing.T) {
		err := New(WithVerbose(), WithProtocol(ProtocolImprobableGRPCWeb)).Serve().Stop()
		if err != nil {
			t.Errorf("Stop must not return an error, but got '%s'", err)
		}
	})
}
