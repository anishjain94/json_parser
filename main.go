package main

type Token struct {
	Value     string
	TokenType EnumTokenType
}

func lexure(input string) ([]Token, error) {
	var tokens []Token

	for i := 0; i < len(input); i++ {

		ch := string(input[i])

		if _, ok := SpecialSymbolsMap[ch]; ok {
			tokens = append(tokens, Token{
				Value:     ch,
				TokenType: SpecialSymbolsMap[ch],
			})
		} else if ch == " " {

			for temp := i + 1; temp <= len(input); temp++ {
				if string(input[temp]) == " " {
					i++
				} else {
					break
				}
			}

		} else if ch == "\"" {
			token, offset := lexString(input, i+1)
			tokens = append(tokens, token)

			i += offset
		} else if ch == "t" {
			if input[i:i+4] == "true" {
				tokens = append(tokens, Token{
					Value:     "true",
					TokenType: TokenBoolean,
				})
				i += 4
			}

		} else if ch == "f" {
			if input[i:i+5] == "false" {
				tokens = append(tokens, Token{
					Value:     "false",
					TokenType: TokenBoolean,
				})
				i += 5

			}
		} else if ch == "n" {
			if input[i:i+4] == "null" {
				tokens = append(tokens, Token{
					Value:     "true",
					TokenType: TokenBoolean,
				})
				i += 4
			}
		} else if isDigit(ch) {
			token, offset := lexDigits(input, i)
			tokens = append(tokens, token)

			i += offset
		}

	}

	return tokens, nil
}

func main() {

}
