package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

// TestHomeHandler ελέγχει την αρχική σελίδα (GET, 404 για λάθος paths, 405 για λάθος μεθόδους)
func TestHomeHandler(t *testing.T) {
	// Αλλάζουμε το working directory στο root ώστε να βρίσκει τα templates
	oldWd, _ := os.Getwd()
	err := os.Chdir("../..")
	if err != nil {
		t.Fatalf("Failed to change directory to root: %v", err)
	}
	defer os.Chdir(oldWd)

	// Case 1: Ένα σωστό GET request στο "/" πρέπει να επιστρέψει 200 OK
	t.Run("Valid GET Request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()

		HomeHandler(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("HomeHandler returned wrong status: got %d want %d", rr.Code, http.StatusOK)
		}
	})

	// Case 2: Ένα request σε λάθος path (π.χ. "/badpath") πρέπει να επιστρέψει 404 Not Found
	t.Run("Invalid Path 404", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/badpath", nil)
		rr := httptest.NewRecorder()

		HomeHandler(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("HomeHandler returned wrong status for bad path: got %d want %d", rr.Code, http.StatusNotFound)
		}
	})

	// Case 3: Μια λάθος μέθοδος (π.χ. POST στο "/") πρέπει να επιστρέψει 405 Method Not Allowed
	t.Run("Invalid Method 405", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		rr := httptest.NewRecorder()

		HomeHandler(rr, req)

		if rr.Code != http.StatusMethodNotAllowed {
			t.Errorf("HomeHandler returned wrong status for POST: got %d want %d", rr.Code, http.StatusMethodNotAllowed)
		}
	})
}

// TestAsciiArtHandler ελέγχει την παραγωγή του art (POST, 400 για λάθος inputs, 405 για GET)
func TestAsciiArtHandler(t *testing.T) {
	oldWd, _ := os.Getwd()
	err := os.Chdir("../..")
	if err != nil {
		t.Fatalf("Failed to change directory to root: %v", err)
	}
	defer os.Chdir(oldWd)

	// Case 1: Ένα έγκυρο POST request με σωστά Form Data πρέπει να επιστρέψει 200 OK
	t.Run("Valid POST Request", func(t *testing.T) {
		form := url.Values{}
		form.Add("text", "Hello")
		form.Add("banner", "standard")

		req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		AsciiArtHandler(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("AsciiArtHandler returned wrong status: got %d want %d", rr.Code, http.StatusOK)
		}
	})

	// Case 2: Ένα POST request με μη-ASCII χαρακτήρες πρέπει να επιστρέψει 400 Bad Request
	t.Run("Invalid Characters 400", func(t *testing.T) {
		form := url.Values{}
		form.Add("text", "Καλημέρα") // Ελληνικά (Μη-ASCII)
		form.Add("banner", "standard")

		req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		AsciiArtHandler(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("AsciiArtHandler returned wrong status for Greek text: got %d want %d", rr.Code, http.StatusBadRequest)
		}
	})

	// Case 3: Ένα GET request στο "/ascii-art" πρέπει να αποκλειστεί με 405 Method Not Allowed
	t.Run("Invalid GET Method 405", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/ascii-art", nil)
		rr := httptest.NewRecorder()

		AsciiArtHandler(rr, req)

		if rr.Code != http.StatusMethodNotAllowed {
			t.Errorf("AsciiArtHandler returned wrong status for GET: got %d want %d", rr.Code, http.StatusMethodNotAllowed)
		}
	})
}
