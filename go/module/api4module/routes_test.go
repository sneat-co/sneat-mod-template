package api4module

import (
	"github.com/sneat-co/sneat-go-core/extension"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterHttpRoutes(t *testing.T) {
	var gotMethod, gotPath string
	var gotHandler http.HandlerFunc

	var handle extension.HTTPHandleFunc = func(method, path string, handler http.HandlerFunc) {
		gotMethod, gotPath, gotHandler = method, path, handler
	}

	RegisterHttpRoutes(handle)

	if gotMethod != "POST" {
		t.Errorf("expected method POST, got %q", gotMethod)
	}
	if gotPath != "/api4module/about" {
		t.Errorf("expected path /api4module/about, got %q", gotPath)
	}
	if gotHandler == nil {
		t.Fatal("expected a handler to be registered")
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/api4module/about", nil)
	gotHandler(recorder, request)

	if body := recorder.Body.String(); body != "api4module" {
		t.Errorf("expected body %q, got %q", "api4module", body)
	}
}
