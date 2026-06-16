package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/paerser/types"
	"github.com/traefik/traefik/v2/pkg/config/dynamic"
)

func TestForwardAuthTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	nextCalled := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
		w.WriteHeader(http.StatusOK)
	})

	auth, err := New(context.Background(), next, dynamic.ForwardAuth{
		Address: server.URL,
		Timeout: types.Duration(50 * time.Millisecond),
	}, "auth")
	require.NoError(t, err)

	ts := httptest.NewServer(auth)
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	require.NoError(t, err)
	assert.Equal(t, http.StatusGatewayTimeout, resp.StatusCode)
	assert.False(t, nextCalled)
}
