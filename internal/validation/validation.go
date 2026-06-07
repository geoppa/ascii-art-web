package validation

func EmptyInput(text string) bool {
	return text == ""
}

func ValidBanner(name string) bool {
	switch name {
	case "standard", "shadow", "thinkertoy":
		return true
	default:
		return false
	}
}