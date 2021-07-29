package main

var (
	TokenKeyPrefix = []byte{0x1}
)

func main() {
	// sym := "token2222"
	// sym1 := "token2a"
	// log.Println(len(TokenKeyPrefix))
	// log.Println(GetTokenKey(sym1))
	// log.Println(GetTokenKey(sym))
}
func GetTokenKey(symbol string) []byte {
	// 10 = len([]byte(symbol))
	key := make([]byte, 10)
	key[0] = TokenKeyPrefix[0]
	copy(key[1:], []byte(symbol))
	return key
}
