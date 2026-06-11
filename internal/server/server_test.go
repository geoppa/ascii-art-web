package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

// TestStart_RoutesRegistration ελέγχει αν όλα τα routes καταχωρούνται σωστά στο σύστημα
func TestStart_RoutesRegistration(t *testing.T) {
	// Αλλάζουμε το working directory στο root για να μην αποτύχει το FileServer setup
	oldWd, _ := os.Getwd()
	err := os.Chdir("../..")
	if err != nil {
		t.Fatalf("Failed to change working directory to root: %v", err)
	}
	defer os.Chdir(oldWd)

	// Καλούμε τη Start() σε ένα Goroutine επειδή η http.ListenAndServe μπλοκάρει.
	go func() {
		_ = Start()
	}()

	// Χρησιμοποιούμε πραγματικό χρόνο αναμονής (100 milliseconds)
	// για να βεβαιωθούμε ότι η Go πρόλαβε να καταχωρήσει όλα τα routes
	time.Sleep(100 * time.Millisecond)

	// Ελέγχουμε αν το DefaultServeMux της Go αναγνωρίζει σωστά τα routes που ορίσαμε
	tests := []struct {
		name string
		path string
	}{
		{"Home Route Registration", "/"},
		{"AsciiArt Route Registration", "/ascii-art"},
		{"Static Files Route Registration", "/static/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Η http.DefaultServeMux.Handler ελέγχει ποιος handler αντιστοιχεί σε κάθε URL
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			_, pattern := http.DefaultServeMux.Handler(req)

			// Αν το pattern είναι κενό, σημαίνει ότι το route δεν καταχωρήθηκε
			if pattern == "" {
				t.Errorf("Route pattern %s was not properly registered in DefaultServeMux", tt.path)
			}
		})
	}
}

// TestStart_StaticCSS ελέγχει αν το αρχείο style.css σερβίρεται σωστά και με το κατάλληλο Content-Type header
func TestStart_StaticCSS(t *testing.T) {
	// Αλλάζουμε το working directory στο root για να βρει τον φάκελο static/
	oldWd, _ := os.Getwd()
	err := os.Chdir("../..")
	if err != nil {
		t.Fatalf("Failed to change working directory to root: %v", err)
	}
	defer os.Chdir(oldWd)

	// Προσομοιώνουμε ένα HTTP GET request για το αρχείο style.css
	req := httptest.NewRequest(http.MethodGet, "/static/style.css", nil)
	rr := httptest.NewRecorder()

	// Χρησιμοποιούμε το DefaultServeMux για να δούμε αν το router απαντάει στο static path
	http.DefaultServeMux.ServeHTTP(rr, req)

	// 1. Ελέγχουμε αν το αρχείο σερβίρεται επιτυχώς (200 OK)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK for style.css, got %d. Make sure 'static/style.css' exists.", rr.Code)
	}

	// 2. Ελέγχουμε αν το Content-Type header είναι σωστά ρυθμισμένο σε text/css
	contentType := rr.Header().Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/css") {
		t.Errorf("Expected Content-Type to start with 'text/css', got '%s'", contentType)
	}
}
