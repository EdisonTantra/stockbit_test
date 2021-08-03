package main

import (
	"regexp"
)

func findFirstStringInBracket(str string) (result string) {

	if len(str) > 0 {
		regexPattern := `\((.*?)\)`
		re := regexp.MustCompile(regexPattern)
		m := re.FindAllStringSubmatch(str, -1)
		if m != nil {
			result = m[0][1]
		}
	}

	// if you really want same return with the old function
	// if you want to whole word just remove 3 lines below
	if len(result) > 1 {
		// slice last char
		result = result[:len(result)-1]
	}

	return result
}
