package gors

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP204StatusCodeWhenPluginIsEnabled(t *testing.T) {
	// Gors is enabled and should return 204 for OPTIONS request
	cfg := Config{
		Disabled: false,
	}

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
	gors, err := New(ctx, next, &cfg, "test-traefik-gors-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodOptions, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	gors.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNoContent {
		t.Fatal("Expected status code: ", http.StatusNoContent, " - Got: ", recorder.Code)
	}
}

func TestServeHTTP200StatusCodeWhenPluginIsDisabled(t *testing.T) {
	// Gors is disabled and should pass through the request to `next` no matter what the request is
	cfg := Config{
		Disabled: true,
	}

	ctx := context.Background()
	nextCalled := false
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) { nextCalled = true; rw.WriteHeader(http.StatusOK) })
	gors, err := New(ctx, next, &cfg, "test-traefik-gors-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodOptions, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	gors.ServeHTTP(recorder, req)

	if !nextCalled {
		t.Fatal("Next did not called")
	}

	if recorder.Code != http.StatusOK {
		t.Fatal("Expected status code: ", http.StatusOK, " - Got: ", recorder.Code)
	}
}
