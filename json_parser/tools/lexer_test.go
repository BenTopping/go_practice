package tools

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestJsonTokens(t *testing.T) {
	result, errors := Lexer(`
	{
		"full":"test",
		"multi":"line"
	}`)
	expected := []Token{
		{
			tokenType: LEFT_BRACE,
			value: "{",
		},
		{
			tokenType: STRING,
			value: "full",
		},
		{
			tokenType: COLON,
			value: ":",
		},
		{
			tokenType: STRING,
			value: "test",
		},
		{
			tokenType: COMMA,
			value: ",",
		},
		{
			tokenType: STRING,
			value: "multi",
		},
		{
			tokenType: COLON,
			value: ":",
		},
		{
			tokenType: STRING,
			value: "line",
		},
		{
			tokenType: RIGHT_BRACE,
			value: "}",
		},
	}

	assert.Equal(t, expected, result)
	assert.Equal(t, nil, errors)
}

func TestInvalidCharacter(t *testing.T) {
	result, errors := Lexer(`
	{
		"full":"test",
	} &`)
	expected := make([]Token,0)

	assert.Equal(t, expected, result)
	assert.Equal(t, fmt.Errorf("Unknown character &"), errors)
}
