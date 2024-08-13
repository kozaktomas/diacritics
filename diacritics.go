package diacritics

import (
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type mns struct{}

func (a mns) Contains(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: non-spacing marks
}

func Remove(s string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(mns{}), norm.NFC)
	ns, _, err := transform.String(t, s)

	return ns, err
}
