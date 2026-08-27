package main

import (
	"fmt"
	tools "json_parser/tools"
)

func main() {
	var lexed_json []tools.Token
	var err error
	input_json := `{"str123":123, "str1234": ["another_str", "anoth", "one_more"]}`

	// Lex
	lexed_json, err = tools.Lexer(input_json)
	fmt.Println(lexed_json, err)

	if err != nil {
		fmt.Println("Error: Unable to tokenize")
		fmt.Println(err)
		return
	}

	// Parse
	parsed_json, tokens := tools.Parser(lexed_json)
	fmt.Println(parsed_json, tokens)
}