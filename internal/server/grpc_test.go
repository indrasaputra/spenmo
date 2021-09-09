package server_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/internal/config"
	"github.com/indrasaputra/spenmo/internal/server"
)

var (
	testPort = "8080"
)

func TestNewGrpc(t *testing.T) {
	t.Run("successfully create a development gRPC server", func(t *testing.T) {
		rate := &config.RateLimit{}
		srv := server.NewGrpc(testPort, rate)
		defer srv.Stop()
		assert.NotNil(t, srv)
	})
}

func TestGrpc_Run(t *testing.T) {
	t.Run("listener fails", func(t *testing.T) {
		rate := &config.RateLimit{}
		srv := server.NewGrpc("abc", rate)

		err := srv.Run()
		defer srv.Stop()

		assert.NotNil(t, err)
	})

	t.Run("success run", func(t *testing.T) {
		rate := &config.RateLimit{}
		srv := server.NewGrpc("8018", rate)

		err := srv.Run()
		defer srv.Stop()
		time.Sleep(1 * time.Second)

		assert.Nil(t, err)
	})
}
