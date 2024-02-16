package main

func lexString(input string, pos int) (Token, int) {
	endPos := pos

	for i := pos; i < len(input); i++ {
		if string(input[i]) == "\"" {
			endPos = i
			break
		}
	}

	return Token{Value: input[pos:endPos], TokenType: TokenString}, endPos - pos + 1
}

func lexDigits(input string, pos int) (Token, int) {

	var endPos int

	for endPos = pos; endPos < len(input); endPos++ {
		if string(input[endPos]) == "\"" || string(input[endPos]) == "}" {
			break
		}
	}

	return Token{Value: input[pos:endPos], TokenType: TokenDigit}, endPos - pos + 1
}

func isDigit(ch string) bool {

	if ch > "0" && ch <= "9" {
		return true
	}

	return false
}
