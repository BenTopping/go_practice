package tools

import (
	// "strings"
	// "bufio"
)


func Parser(tokens []Token) (map[string]interface{}, []Token) {
	t := tokens[0]

	// if t.tokenType == LEFT_BRACKET {
	// 	return parseArray(tokens[1:])
	// }
	if t.tokenType == LEFT_BRACE {
		return parseObject(tokens[1:])
	} else {
		return map[string]interface{}{"value": t.value},  tokens[1:]
	}
}

func parseObject(tokens []Token) (map[string]interface{}, []Token) {
	parsedObj := make(map[string]interface{})

	t := tokens[0]
	if t.tokenType == RIGHT_BRACE {
		return parsedObj, tokens[1:]
	}

	for true {
		json_key := tokens[0]
		if json_key.tokenType == STRING {
			// Move to the next token
			tokens = tokens[1:]
		} else {
			// We should error here as always expect keys to be strings
		}

		if tokens[0].tokenType != COLON {
			// We should error here as always expect colon to follow a key
		}

		var json_value map[string]interface{}
		json_value, tokens = Parser(tokens[1:])

		parsedObj[json_key.value] = json_value["value"]

		t = tokens[0]
		if t.tokenType == RIGHT_BRACE {
			// We have finished this object so we can return
			return parsedObj, tokens[1:]
		} else if t.tokenType != COMMA {
			// We should error here as commas should always appear after key pairs
		}
		// Move forward to the next token
		tokens = tokens[1:]
	}

	// We should error here
	return parsedObj, nil
}

func parseArray(tokens []Token) {

}