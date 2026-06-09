# ASCII Art Web - Audit Questions

## Functional

* [ ] Has the requirement for the allowed packages been respected? (only standard packages)
* [ ] Does the project contain HTML files?

## Standard Banner

Try inputting with the standard template/banner the following example:

First line:

{123}

Second line:

<Hello> (World)!

* [ ] Does it display the correct ASCII Art output?

## Standard Banner - Additional Test

Try inputting:

123??

* [ ] Does it display the correct ASCII Art output?

## Shadow Banner

Try inputting:

$% "=

* [ ] Does it display the correct ASCII Art output?

## Thinkertoy Banner

Try inputting:

123 T/fs#R

* [ ] Does it display the correct ASCII Art output?

## Generated Output

* [ ] Does it display an understandable graphical representation of the result?

## Website Navigation

* [ ] Try to navigate between all available pages.
* [ ] Are all pages working correctly?
* [ ] Does the project implement HTTP Status 404 (Not Found)?
* [ ] Does the project implement HTTP Status 400 (Bad Request)?
* [ ] Does the project implement HTTP Status 500 (Internal Server Error)?

## HTTP Communication

* [ ] Make a request to generate ASCII Art.
* [ ] Is communication between server and client properly established?
* [ ] Does the server use the correct HTTP methods?
* [ ] Does the website work without crashing?

## Server

* [ ] Is the server written in Go?
* [ ] Does the server present all required handlers and routes?
* [ ] Does the server run quickly and effectively?

## Code Quality

* [ ] Does the code follow good practices?
* [ ] Is the project structure organized?
* [ ] Are responsibilities properly separated?

## Testing

* [ ] Is there a test file for the project?
* [ ] Do all tests pass?

`go test ./...`

## Validation Package

* [ ] ValidateText tests pass.
* [ ] ValidateBanner tests pass.

## Banner Package

* [ ] Banner loading tests pass.

## Handlers Package

* [ ] 404 tests pass.
* [ ] 405 tests pass.

## Final Review

* [ ] README completed.
* [ ] LICENSE added.
* [ ] TASK.md added.
* [ ] AUDIT.md added.
* [ ] Project builds successfully.

`go run ./cmd`

* [ ] Project ready for audit.

## Social

* [ ] Did you learn something from this project?
* [ ] Can this project be open-sourced?
* [ ] Would you recommend this project as an example for future students?
