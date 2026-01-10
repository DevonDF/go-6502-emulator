package assembler

import (
	"strings"
)

type TokenType int

const (
	allowedStringStartCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_@"
	hexCharacters                = "0123456789abcdefABCDEF"
	allowedStringCharacters      = allowedStringStartCharacters + "0123456789"
	specialCharacters            = "#$<>():-+.\","

	TokenTypeString TokenType = iota
	TokenTypeSymbol
	TokenTypeNumber
)

type Token struct {
	Token string    // the string token that represents this token
	Type  TokenType // the type of the Token
}

// Tokenise takes a string and returns a list of Token contained within said string
func Tokenise(line string) []Token {

	tokens := make([]Token, 0)

	for lineIndex := 0; lineIndex < len(line); lineIndex++ {
		character := string(line[lineIndex])
		if strings.Contains(" \t", character) {
			// whitespace, continue
			continue
		} else if character == ";" {
			// comment, we can stop here
			break
		} else if strings.Contains(specialCharacters, character) {
			// a special character, make a symbol token
			tokens = append(tokens, Token{Token: character, Type: TokenTypeSymbol})
		} else if strings.Contains(allowedStringStartCharacters, character) {
			// start reading a string
			readStr := ""
			for lineIndex < len(line) {
				strChar := string(line[lineIndex])
				if !strings.Contains(allowedStringCharacters, strChar) {
					lineIndex--
					break
				}
				readStr += strChar
				lineIndex++
			}
			tokens = append(tokens, Token{Token: readStr, Type: TokenTypeString})
		} else if strings.Contains(hexCharacters, character) {
			// start reading a number
			readNumStr := ""
			for lineIndex < len(line) {
				strChar := string(line[lineIndex])
				if !strings.Contains(hexCharacters, strChar) {
					lineIndex--
					break
				}
				readNumStr += strChar
				lineIndex++
			}
			tokens = append(tokens, Token{Token: readNumStr, Type: TokenTypeNumber})
		}

	}

	return tokens
}
