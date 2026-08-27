package tools

func Parser(tokens []Token) (interface{}, []Token) {
	t := tokens[0]

	if t.tokenType == LEFT_BRACKET {
		return parseArray(tokens[1:])
	}
	if t.tokenType == LEFT_BRACE {
		return parseObject(tokens[1:])
	} else {
		return t.value,  tokens[1:]
	}
}

func parseObject(tokens []Token) (interface{}, []Token) {
	parsedObj := make(map[string]interface{})

	t := tokens[0]
	if t.tokenType == RIGHT_BRACE {
		return parsedObj, tokens[1:]
	}

	for true {
		jsonKey := tokens[0]
		if jsonKey.tokenType == STRING {
			// Move to the next token
			tokens = tokens[1:]
		} else {
			// We should error here as always expect keys to be strings
		}

		if tokens[0].tokenType != COLON {
			// We should error here as always expect colon to follow a key
		}

		jsonValue, remainingTokens := Parser(tokens[1:])
		parsedObj[jsonKey.value] = jsonValue
		tokens = remainingTokens

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

func parseArray(tokens []Token) (interface{}, []Token) {
	parsedArray := make([]interface{}, 0)

	t := tokens[0]
	if t.tokenType == RIGHT_BRACKET { 
		return parsedArray, tokens[1:]
	}

	for true {
		var jsonValue interface{}
		jsonValue, remainingTokens := Parser(tokens)
		parsedArray = append(parsedArray, jsonValue)
		tokens = remainingTokens

		t = tokens[0]
		if t.tokenType == RIGHT_BRACKET {
			// We have finished this Array so we can return
			return parsedArray, tokens[1:]
		} else if t.tokenType != COMMA {
			// We should error here as commas should always appear after key pairs
		}
		// Move forward to the next token
		tokens = tokens[1:]
	}

	return parsedArray, nil
}