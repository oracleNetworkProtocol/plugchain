package main

import (
	"strings"
)

var (
	TokenKeyPrefix     = []byte{0x1}
	TokenNamePrefix    = "token"
	TokenNamePrefixLen = 5
	ts                 = 4
)

func main() {
}
func GetTokenKey(symbol string) []byte {
	// 10 = len([]byte(symbol))
	key := make([]byte, 10)
	key[0] = TokenKeyPrefix[0]
	copy(key[1:], []byte(symbol))
	return key
}

//返回带前缀的token-symbol ，symbol全大写
func GetSymbol(symbol string) string {

	if len(symbol) > ts && symbol[:TokenNamePrefixLen] == TokenNamePrefix {
		return strings.ToLower(symbol[:TokenNamePrefixLen]) + strings.ToTitle(symbol[TokenNamePrefixLen:])
	}
	return TokenNamePrefix + strings.ToTitle(symbol)
}
