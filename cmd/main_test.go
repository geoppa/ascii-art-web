package main

import (
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMainIntegration(t *testing.T) {
	// ΔΙΟΡΘΩΣΗ: Αναγκάζουμε το test να αλλάξει το working directory στο root
	// ώστε η template.ParseFiles() και η os.Open() να βρίσκουν τους φακέλους templates/ και banners/
	err := os.Chdir("..")
	if err != nil {
		t.Fatalf("Failed to change working directory: %v", err)
	}

	// Αν το test αρχείο σου βρίσκεται μέσα σε υποφάκελο (π.χ. cmd/),
	// άλλαξε το παραπάνω σε os.Chdir("../") για να πάει ένα επίπεδο πίσω.

	// 1. Ξεκινάμε τη main() σε ένα ξεχωριστό Goroutine
	go main()

	// 2. Δίνουμε χρόνο στον server να σηκωθεί
	time.Sleep(200 * time.Millisecond)

	// 3. Στέλνουμε GET request στην αρχική σελίδα
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatalf("Αποτυχία σύνδεσης με τον server: %v", err)
	}
	defer resp.Body.Close()

	// Επιβεβαιώνουμε το 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Αναμενόμενο Status 200 OK, αλλά πήραμε: %d", resp.StatusCode)
	}

	// 4. Test request για το 404 Not Found
	resp404, err := http.Get("http://localhost:8080/non-existent-page-xyz")
	if err != nil {
		t.Fatalf("Αποτυχία σύνδεσης με τον server στο test 404: %v", err)
	}
	defer resp404.Body.Close()

	if resp404.StatusCode != http.StatusNotFound {
		t.Errorf("Αναμενόμενο Status 404 Not Found, αλλά πήραμε: %d", resp404.StatusCode)
	}
}
