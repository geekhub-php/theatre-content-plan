package core

import (
	strip "github.com/grokify/html-strip-tags-go"
	"strings"
	"unicode/utf8"
)

type Description string
func (desc Description) RuneCount() int {
	striped := strip.StripTags(string(desc))
	fields := strings.Fields(striped)
	str := strings.Join(fields, "")

	return utf8.RuneCountInString(str)
}
