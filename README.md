# ASCII Art Web

ASCII Art Web is a web-based version of the ASCII Art project written in Go.

The application allows users to enter text, select a banner style, and generate an ASCII Art representation directly from a web browser.

The project follows a modular architecture using Go packages and only standard Go libraries.

---

## Features

* Generate ASCII Art from user input.
* Support multiple banner styles:

  * standard
  * shadow
  * thinkertoy
* Simple web interface.
* Custom error pages.
* Input validation.
* HTTP status handling.
* Unit tests for core packages.

---

## Supported Banners

### Standard

The default ASCII Art banner.

### Shadow

A shadow-styled ASCII Art banner.

### Thinkertoy

A lightweight ASCII Art banner.

---

## Project Structure

```text
ascii-art-web
├── banners
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│
├── cmd
│   ├── main.go
│   └── main_test.go
│
├── internal
│   ├── banner
│   ├── handlers
│   ├── render
│   ├── server
│   └── validation
│
├── static
│   └── style.css
│
├── templates
│   ├── index.html
│   └── error.html
│
├── README.md
├── LICENSE
├── TASK.md
└── AUDIT.md
```

---

## Installation

Clone the repository:

```bash
git clone <repository-url>
cd ascii-art-web
```

---

## Running the Server

Start the application:

```bash
go run ./cmd
```

The server will start on:

```text
http://localhost:8080
```

---

## Using the Application

1. Open the website in your browser.
2. Enter text in the input field.
3. Select a banner style.
4. Click **Generate ASCII Art**.
5. View the generated result.

---

## HTTP Routes

### GET /

Displays the homepage.

### POST /ascii-art

Generates ASCII Art from the submitted form.

### /static/

Serves static assets such as CSS files.

---

## HTTP Status Codes

The application handles the following HTTP status codes:

| Status Code | Description           |
| ----------- | --------------------- |
| 400         | Bad Request           |
| 404         | Not Found             |
| 405         | Method Not Allowed    |
| 500         | Internal Server Error |

---

## Testing

Run all tests:

```bash
go test ./...
```

Current test coverage includes:

* Validation package
* Banner package
* Handlers package

---

## Technologies Used

* Go
* HTML
* CSS
* Go Templates
* HTTP Server

---

## Authors

* gpapadaki
* elgeorgiou

---

## License

This project is distributed under the MIT License.

See the LICENSE file for more information.

----------

## Σημείωση για το Merge

### Ολοκληρωμένα (elgeorgiou)

#### internal/handlers

* Homepage handler.
* ASCII Art handler.
* HTTP method validation.
* Custom error page rendering.
* 404 Not Found handling.
* 405 Method Not Allowed handling.
* Form processing and request flow.

#### internal/validation

* Text validation.
* Banner validation.
* Validation tests.

#### templates

* Homepage template.
* Error page template.

#### static

* Application styling.

#### Documentation

* README.md
* LICENSE
* TASK.md
* AUDIT.md

#### Tests

* internal/validation/validation_test.go
* internal/banner/banner_test.go
* internal/handlers/handlers_test.go

---

### Υπόλοιπο προς ολοκλήρωση

#### internal/render

* Verify rendering logic.
* Verify generated output matches expected banner output.
* Add/complete render tests.

#### Integration

* Verify complete flow after merge.
* Verify banner loading and rendering together.
* Verify all audit examples produce expected output.

#### Final Checks

* Run:

```bash
go test ./...
```

```bash
go run ./cmd
```

* Complete final audit rehearsal.
* Verify all HTTP status codes.
* Verify all supported banners.
