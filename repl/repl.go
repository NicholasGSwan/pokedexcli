package repl

import "strings"

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	strArr := strings.Fields(text)
	return strArr
}
