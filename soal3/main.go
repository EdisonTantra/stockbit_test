package main

import (
	"regexp"
)

func findFirstStringInBracket(str string) (result string) {

	if len(str) > 0 {
		regexPattern := `(?:\(\s*|\()([a-zA-Z]*)(?:\s.*\)|\))`
		re := regexp.MustCompile(regexPattern)
		m := re.FindAllStringSubmatch(str, -1)
		if m != nil {
			result = m[0][1]
		}
	}

	return result
}
