# 📋 Project Task Checklist - ASCII Art Web

This document tracks the core requirements, constraints, bug fixes, and testing targets completed for the Zone01 `ascii-art-web` project.

---

## 🟥 1. Core Server & Routing Requirements
- [x] Write the web server entirely in native **Go** using the `net/http` standard library.
- [x] Implement standard **GET** route for the Home Dashboard (`/`).
- [x] Implement standard **POST** route for ASCII Art production (`/ascii-art`).
- [x] Create an effective routing structure that loads CSS assets via a clean static server mapping (`/static/`).
- [x] Ensure the server processes concurrent connections cleanly and executes requests effectively without lags.

---

## 🟧 2. HTTP Status Code Compliance
- [x] **`200 OK`**: Return success layout matrices when form parameters resolve perfectly.
- [x] **`400 Bad Request`**: Intercept bad requests instantly if submissions contain broken formatting or invalid data.
- [x] **`404 Not Found`**: Defend against unregistered paths (e.g., `/random`) and serve custom structural page displays.
- [x] **`405 Method Not Allowed`**: Block illegal route actions (e.g., `POST` on `/` or `GET` on `/ascii-art`) across standard traffic nodes.
- [x] **`500 Internal Server Error`**: Protect client screens and logging paths if system templates or structural banner resources are missing.

---

## 🟨 3. Algorithmic Art Generation & Security Filters
- [x] Standardize text layers into structured 8-row dynamic blocks.
- [x] Map fonts accurately inside local storage strings using the standard metric: `(char - 32) * 9 + 1`.
- [x] Convert incoming literal text inputs (`\\n`) down to active structural newline bytes (`\n`).
- [x] Strip out hidden Windows Carriage Returns (`\r`) to guarantee correct formatting when reading the `thinkertoy` asset.
- [x] Secure inputs using explicit layout string validators against malicious Directory Traversal payload injection attempts (e.g., `../`).
- [x] Verify incoming body inputs character-by-character to throw out non-ASCII symbols (e.g., Greek characters, Emojis).

---

## 🟩 4. Frontend & User Interface Architecture
- [x] Construct a responsive HTML structure with form input objects (`index.html`).
- [x] Maintain UI configuration metrics across submission actions (persisting selected radio button values and textarea logs).
- [x] Enclose algorithmic text payloads directly within standard HTML `<pre>` tags to preserve visual baseline margins.
- [x] Link layouts to custom cascading external assets (`/static/style.css`) without breaking rendering configurations.

---

## 🟦 5. Testing Framework & Verification Targets
- [x] **`main_test.go`**: Verify structural integration bounds and live network listening states over `:8080`.
- [x] **`server_test.go`**: Double-check that all application pathways register correctly on Go's `DefaultServeMux`.
- [x] **`handlers_test.go`**: Fully validate HTTP routing, 404 path handling, and 405 constraint behaviors.
- [x] **`validation_test.go`**: Confirm that `ValidateText` and `ValidateBanner` catch dangerous values or bad encoding attempts.
- [x] **`banner_test.go`**: Audit file operations, tracking missing items and trailing string sanitizations.
- [x] **`render_test.go`**: Test core formatting loops to prevent unwanted trailing newlines at the very end of output files.
- [x] **Style Assets Test**: Create dedicated testing pathways to confirm that `/static/style.css` distributes correctly with a valid `text/css` content-type header.

---

## 🛠️ Completed Optimizations & Bug Fixes
- [x] **FIXED (Bug 1 - Duplicate Headers):** Refactored the `renderHomePage` engine to explicitly capture network statuses via specific method metrics before executing layouts. This entirely resolved the `http: superfluous response.WriteHeader call` terminal diagnostic.
- [x] **FIXED (Bug 2 - Trailing Newline):** Adjusted the final output sequence using active string suffixes (`strings.TrimSuffix`). This removes unwanted extra final line returns, ensuring exact text match compliance for strict Zone01 validation checks.
