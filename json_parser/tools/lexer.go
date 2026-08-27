package tools

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	LEFT_BRACE    = "LEFT_BRACE"
	RIGHT_BRACE   = "RIGHT_BRACE"
	LEFT_BRACKET  = "LEFT_BRACKET"
	RIGHT_BRACKET = "RIGHT_BRACKET"
	COLON         = "COLON"
	COMMA         = "COMMA"
	STRING        = "STRING"
	// TODO: Handle +,-,e,i in numbers
	NUMBER        = "NUMBER"
	// TODO: true, false, null
)

type Token struct {
	tokenType string
	value     interface{}
}

func buildTokens(bslice []byte) ([]Token, error) {
	var tokens []Token
	var pos = 0
	var err error

	for pos < len(bslice) {
		ch := string(bslice[pos])
		switch ch {
		case "{":
			tokens = append(tokens, Token{tokenType: LEFT_BRACE, value: ch})
			pos++
		case "}":
			tokens = append(tokens, Token{tokenType: RIGHT_BRACE, value: ch})
			pos++
		case "[":
			tokens = append(tokens, Token{tokenType: LEFT_BRACKET, value: ch})
			pos++
		case "]":
			tokens = append(tokens, Token{tokenType: RIGHT_BRACKET, value: ch})
			pos++
		case ":":
			tokens = append(tokens, Token{tokenType: COLON, value: ch})
			pos++
		case ",":
			tokens = append(tokens, Token{tokenType: COMMA, value: ch})
			pos++
		case `"`:
			str := ""
			pos++
			for pos < len(bslice) && string(bslice[pos]) != `"` {
				str += string(bslice[pos])
				pos++
			}
			pos++
			tokens = append(tokens, Token{tokenType: STRING, value: str})
		default:
			// If its not in a string and its whitespace, ignore it
			if regexp.MustCompile(`\s`).MatchString(ch) {
				pos++
			// If its a number parse it as such
			} else if regexp.MustCompile(`\d`).MatchString(ch) {
				num := ch
				pos++
				for regexp.MustCompile(`\d`).MatchString(string(bslice[pos])) {
					num += string(bslice[pos])
					pos++
				}
				f, err := strconv.ParseFloat(num, 64)
				if err != nil {
					err = fmt.Errorf("Unable to convert integert %s", num)
				}
				tokens = append(tokens, Token{tokenType: NUMBER, value: f})

			} else {
				err = fmt.Errorf("Unknown character %s", ch)
				tokens = make([]Token, 0)
				pos = len(bslice)
			}
		}
	}

	return tokens, err
}

func Lexer(json string) ([]Token, error) {
	reader := strings.NewReader(json)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanBytes)
	var bslice = make([]byte, 0)
	for scanner.Scan() {
		// Skip whitespaces
		bytes := scanner.Bytes()
		for _, b := range bytes {
			bslice = append(bslice, b)
		}
	}

	tokens, err := buildTokens(bslice)

	return tokens, err
}
