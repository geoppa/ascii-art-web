package validation

import (
	"errors"
	"strings"
)

func ValidateText(text string) error {
	if strings.TrimSpace(text) == "" {
		return errors.New("Please enter text to generate ASCII art.")
	}

	return nil
}

func ValidateBanner(name string) error {
	switch name {
	case "standard", "shadow", "thinkertoy":
		return nil
	default:
		return errors.New("Please select a valid banner style.")
	}
}