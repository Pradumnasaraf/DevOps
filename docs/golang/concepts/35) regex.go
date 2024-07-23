package main

import (
	"regexp"
)

func main() {

	// Test the formatAsPlainText function
	input := `# Heading 1

## Heading 2

### Heading 3

#### Heading 4

##### Heading 5

###### Heading 6

---

_Italic text_  *Italic text*


**Bold text**  __Bold text__


_Italic text_  *Italic text*

`

	output := formatAsPlainText(input)
	println(output)

}

func formatAsPlainText(input string) string {
	// Remove Markdown headings
	re := regexp.MustCompile(`(?m)^#{1,6}\s+(.*)`)
	input = re.ReplaceAllString(input, "$1")

	// Remove horizontal rules
	re = regexp.MustCompile(`(?m)^\n?-{3,}\n?`)
	input = re.ReplaceAllString(input, "\n")

	// Remove bold formatting
	re = regexp.MustCompile(`\*\*([^*]+)\*\*|__([^_]+)__`)
	input = re.ReplaceAllString(input, "$1$2")

	// Remove italic formatting
	re = regexp.MustCompile(`_([^_]+)_|\*([^*]+)\*`)
	input = re.ReplaceAllString(input, "$1$2")

	// Replace bullet points with a hyphen
	re = regexp.MustCompile(`(?m)^\* (.*)`)
	input = re.ReplaceAllString(input, "- $1")

	return input
}
