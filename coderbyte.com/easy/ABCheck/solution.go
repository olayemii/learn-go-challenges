package main

import (
	"strings"
)

/*
Have the function ABCheck(str) take the str parameter being passed
and return the string true if the characters a and b are separated
by exactly 3 places anywhere in the string at least once
(ie. "lane borrowed" would result in true because there is exactly three characters between a and b).
Otherwise return the string false.
*/

func ABCheck(str string) bool {
	for i, letter := range str {
		if strings.ToLower(string(letter)) == "a" &&
			len(str) > (i+4) &&
			strings.ToLower(string(str[i+4])) == "b" {
			return true
		}
		if strings.ToLower(string(letter)) == "b" &&
			len(str) > (i+4) &&
			strings.ToLower(string(str[i+4])) == "a" {
			return true
		}
	}
	return false
}
