package main

import (
	"fmt"
	"regexp"
	"strings"
)

var nonword = regexp.MustCompile(`\W+`)

func match(name, file string) bool {
	name = nonword.ReplaceAllLiteralString(name, "")
	file = nonword.ReplaceAllLiteralString(file, "")
	fmt.Println(name, file)
	return strings.Contains(file, name)
}
