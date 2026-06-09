package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHomeHandler404(t *testing.T) {
	err := os.MkdirAll("templates", 0755)
	if err != nil {
		t.Fatalf("failed to create templates directory: %v", err)
	}

	err = os.WriteFile(
		"templates/error.html",
		[]byte("{{ .Code }} {{ .Message }}"),
		0644,
	)
	if err != nil {
		t.Fatalf("failed to create error template: %v", err)
	}

	defer os.RemoveAll("templates")

	req := httptest.NewRequest(
		http.MethodGet,
		"/banana",
		nil,
	)

	rr := httptest.NewRecorder()

	HomeHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf(
			"expected %d, got %d",
			http.StatusNotFound,
			rr.Code,
		)
	}
}

func TestHomeHandler405(t *testing.T) {
	err := os.MkdirAll("templates", 0755)
	if err != nil {
		t.Fatalf("failed to create templates directory: %v", err)
	}

	err = os.WriteFile(
		"templates/error.html",
		[]byte("{{ .Code }} {{ .Message }}"),
		0644,
	)
	if err != nil {
		t.Fatalf("failed to create error template: %v", err)
	}

	defer os.RemoveAll("templates")

	req := httptest.NewRequest(
		http.MethodPost,
		"/",
		nil,
	)

	rr := httptest.NewRecorder()

	HomeHandler(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf(
			"expected %d, got %d",
			http.StatusMethodNotAllowed,
			rr.Code,
		)
	}
}

func TestAsciiArtHandler405(t *testing.T) {
	err := os.MkdirAll("templates", 0755)
	if err != nil {
		t.Fatalf("failed to create templates directory: %v", err)
	}

	err = os.WriteFile(
		"templates/error.html",
		[]byte("{{ .Code }} {{ .Message }}"),
		0644,
	)
	if err != nil {
		t.Fatalf("failed to create error template: %v", err)
	}

	defer os.RemoveAll("templates")

	req := httptest.NewRequest(
		http.MethodGet,
		"/ascii-art",
		nil,
	)

	rr := httptest.NewRecorder()

	AsciiArtHandler(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf(
			"expected %d, got %d",
			http.StatusMethodNotAllowed,
			rr.Code,
		)
	}
}
