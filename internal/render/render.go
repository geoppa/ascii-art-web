package render

import (
	"strings"
)

func Generate(text string, banner []string) string {
	var output strings.Builder

	parts := strings.Split(text, "\\n")

	for _, part := range parts {

		if part == "" {
			output.WriteString("\n")
			continue
		}

		for row := 0; row < 8; row++ {

			for _, char := range part {

				start := int(char-32)*9 + 1

				if start+row < len(banner) {
					output.WriteString(banner[start+row])
				}
			}

			output.WriteString("\n")
		}
	}

	return output.String()
}