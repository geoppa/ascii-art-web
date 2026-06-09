# ASCII Art Web

## Objective

The goal of this project is to create a web application that generates ASCII Art from user input.

The application must provide a graphical user interface accessible through a web browser and use Go's standard library to handle HTTP requests and responses.

## Requirements

* Use only Go standard packages.
* Provide an HTML user interface.
* Accept user input through a web form.
* Generate ASCII Art using banner files.
* Support the following banners:

  * standard
  * shadow
  * thinkertoy
* Display the generated ASCII Art in the browser.
* Handle HTTP requests using the correct methods.
* Implement proper error handling.

## HTTP Status Codes

The application must correctly handle:

* 400 Bad Request
* 404 Not Found
* 405 Method Not Allowed
* 500 Internal Server Error

## Features

### Text Input

Users can enter text into a form and submit it for processing.

### Banner Selection

Users can choose between multiple banner styles.

### ASCII Art Generation

The server converts the provided text into ASCII Art and returns the result to the user.

### Error Handling

The application displays custom error pages when invalid requests occur.

## Technologies

* Go
* HTML
* CSS
* Go Templates
* HTTP Server

## Project Structure

```text
ascii-art-web
├── banners
├── cmd
├── internal
├── static
├── templates
├── README.md
├── LICENSE
├── TASK.md
└── AUDIT.md
