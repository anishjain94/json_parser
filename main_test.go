package main

import (
	"fmt"
	"testing"
)

func TestTokens(t *testing.T) {

	str := `{"hello": 12345}`
	tokens, _ := lexure(str)

	fmt.Println(tokens)
}
