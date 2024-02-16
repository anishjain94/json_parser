package main

type EnumTokenType string

const (
	TokenObjectStart EnumTokenType = "tokenObjectStart"
	TokenObjectEnd   EnumTokenType = "tokenObjectEnd"
	TokenArrayStart  EnumTokenType = "tokenArrayStart"
	TokenArrayEnd    EnumTokenType = "tokenArrayEnd"
	TokenComma       EnumTokenType = "tokenComma"
	TokenColon       EnumTokenType = "tokenColon"
	TokenBoolean     EnumTokenType = "tokenBoolean"
	TokenDigit       EnumTokenType = "tokenDigit"
	TokenNull        EnumTokenType = "tokenNull"
	TokenString      EnumTokenType = "tokenString"
)

var SpecialSymbolsMap = map[string]EnumTokenType{
	"{": TokenObjectStart,
	"}": TokenObjectEnd,
	"[": TokenArrayStart,
	"]": TokenArrayEnd,
	":": TokenColon,
}
