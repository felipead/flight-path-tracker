package model

import (
	"unicode"
)

type AirportCode string

func (code AirportCode) IsValid() bool {
	return len(code) == 3 &&
		unicode.IsLetter(rune(code[0])) &&
		unicode.IsLetter(rune(code[1])) &&
		unicode.IsLetter(rune(code[2]))
}
