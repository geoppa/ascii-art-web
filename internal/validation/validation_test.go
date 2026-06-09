package validation

import "testing"

func TestValidateText(t *testing.T) {

	err := ValidateText("hello")
	if err != nil {
	t.Errorf("expected nil, got error")
}

	err = ValidateText("")
	if err == nil {
	t.Errorf("expected error, got nil")
}

	err = ValidateText("     ")
	if err == nil {
	t.Errorf("expected error, got nil")
}
}

func TestValidateBanner(t *testing.T) {

	err := ValidateBanner("standard")
	if err != nil {
	t.Errorf("expected nil, got error")
}

	err = ValidateBanner("shadow")
	if err != nil {
	t.Errorf("expected nil, got error")
}

	err = ValidateBanner("thinkertoy")
	if err != nil {
	t.Errorf("expected nil, got error")
}

	err = ValidateBanner("banana")
	if err == nil {
	t.Errorf("expected error, got nil")
}
}
