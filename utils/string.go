package utils

import (
	"bytes"
	"crypto/rand"
	"math/big"
)

func RandomString(len int) string {
	return RandomStringRange(len, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
}

func RandomStringRange(len int, str string) string {
	var res string
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		res += string(str[randomInt.Int64()])
	}
	return res
}

func ChineseLength(str string, limit int) int {
	sum := 0
	for _, r := range []rune(str) {
		switch {
		case r >= '\u0000' && r <= '\u007F':
			sum += 1
		case r >= '\u0080' && r <= '\u07FF':
			sum += 2
		case r >= '\u0800' && r <= '\uFFFF':
			sum += 3
		default:
			sum += 4
		}
		if sum >= limit {
			break
		}
	}
	return sum
}
