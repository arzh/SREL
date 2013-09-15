package srel

import (
	"strings"
	"regexp"
)

// *builder class and *builder functions
// Heavily influenced by GoVisualExpressions
// github.com/VerbalExpressions/GoVerbalExpressions

// Used to build a Regexp
type builder struct {
	exp string
	pre string
	post string
	mod string
}

func (b *builder) String() string {
	mod := ""
	if len(b.mod) > 0 {
		mod += "(?" + b.mod + ")"
	}
	
	return mod + b.pre + b.exp + b.post
}

// Finally puts the regex together in string
func (b *builder) compile() (*regexp.Regexp, error) {
	return regexp.Compile(b.String())
}

// Wrapper for QuoteMeta
func quote(s string) string {
	return regexp.QuoteMeta(s)
}

// functions that take one argument
// ----------------------------------------------------
type builderFuncVar func(*builder, string)


// Match anything except given string
func anythingBut(b *builder, s string) {
	b.exp += `(?:[^` + quote(s) + `]*)`
}

// Search string, doesn't fail if not found
func maybe(b *builder, s string) {
	b.exp += `(?:` + quote(s) + `)?`
}

// Search for string, string must MUST be there, unlike 'maybe'
func find(b *builder, s string) {
	b.exp += `(?:` + quote(s) + `)`
}

// Matches any character given
func any(b *builder, s string) {
	b.exp += `(?:[` + quote(s) + `])`
}


// functions that take no arguments
// -----------------------------------------------------
type builderFunc func(*builder)


// Search at the beginning of a line
func startOfLine(b *builder) {
	b.pre += `^`
}

// Search at the end of a line
func endOfLine(b *builder) {
	b.post += `$`
}

// Match all characters
func anything(b *builder) {
	b.exp += `(?:.*)`
}

// Match a line-break (\n or \r\n)
func lineBreak(b *builder) {
	b.exp += `(?:(?:\n)|(?:\r\n))`
}

// Match for 'tab' 
func tab(b *builder) {
	b.exp += `\t+`
}

// Match any word
func word(b *builder) {
	b.exp += `\w+`
}

// Binding function, to allow for simple chaining an alt expression
func or(b *builder) {
	if strings.Index(b.pre, "(") == -1 {
		b.pre += "(?:"
	}
	if strings.Index(b.post, ")") == -1 {
		b.post = ")" + b.post
	}

	b.exp += ")|(?:"
}