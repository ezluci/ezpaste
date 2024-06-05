package main

import (
	"math/rand"
)

var az09_chars []byte

func init() {
	
	az09_chars = make([]byte, 10 + 26)
	for i := '0'; i <= '9'; i += 1 {
		az09_chars[i - '0'] = byte(i)
	}
	for i := 'a'; i <= 'z'; i += 1 {
		az09_chars[i - 'a' + 10] = byte(i)
	}
}

// random string with a-z0-9, 'count' characters
func genRandomString(count int) string {
	
	ans := ""
	for i := 0; i < count; i += 1 {
		ans += string( az09_chars[rand.Int31n(10 + 26)] );
	}

	return ans;
}