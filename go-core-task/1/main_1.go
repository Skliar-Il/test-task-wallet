package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
)

func GenerateHashedString(args []interface{}, salt string) string {
	var resultStr string

	for _, value := range args {
		switch v := value.(type) {
		case int:
			resultStr += strconv.Itoa(v)
		case float64:
			resultStr += fmt.Sprintf("%f", v)
		case string:
			resultStr += v
		case bool:
			resultStr += fmt.Sprintf("%t", v)
		case complex64:
			resultStr += fmt.Sprintf("%f+%fi", real(v), imag(v))
		default:
			resultStr += fmt.Sprintf("<unsupported:%s>", reflect.TypeOf(value).String())
		}
	}

	runes := []rune(resultStr)
	runesSalt := []rune(salt)
	mid := len(runes) / 2
	runesWithSalt := append(runes[:mid], append(runesSalt, runes[mid:]...)...)

	hash := sha256.Sum256([]byte(string(runesWithSalt)))
	return hex.EncodeToString(hash[:])
}

func main() {
	fmt.Println(GenerateHashedString([]interface{}{1, 052, 0x52, 2.5, "hello", true, complex(1, 2)}, "go-2024"))
}
