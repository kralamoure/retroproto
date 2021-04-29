package d1proto

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func EncryptPassword(pwd string, key string) string {
	hash := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

	sb := &strings.Builder{}
	for i := 0; i < len(pwd); i++ {
		loc6 := int(pwd[i])
		loc7 := int(key[i])

		loc8 := int(math.Floor(float64(loc6 / 16)))
		loc9 := loc6 % 16

		sb.WriteRune(rune(hash[(loc8+loc7%len(hash))%len(hash)]))
		sb.WriteRune(rune(hash[(loc9+loc7%len(hash))%len(hash)]))
	}
	return sb.String()
}

func SplitEncodedHostPortTicket(extra string) (host string, port int, ticket string, e error) {
	ticket = extra[11:]

	encryptedIP := extra[:8]
	encodedPort := extra[8:11]

	tmpIP := make([]string, 0)
	for i := 0; i < len(encryptedIP); i += 2 {
		tmp1 := encryptedIP[i] - 48
		tmp2 := encryptedIP[i+1] - 48
		tmpIP = append(tmpIP, fmt.Sprint((tmp1&15)<<4|tmp2&15))
	}
	host = strings.Join(tmpIP, ".")

	n1, err := Decode64(rune(encodedPort[0]))
	if err != nil {
		e = err
		return
	}
	n2, err := Decode64(rune(encodedPort[1]))
	if err != nil {
		e = err
		return
	}
	n3, err := Decode64(rune(encodedPort[2]))
	if err != nil {
		e = err
		return
	}

	port = (n1&63)<<12 | (n2&63)<<6 | n3&63

	return
}

func Encode64(n int) (rune, error) {
	zipKey := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

	if n < 0 || n > len(zipKey) {
		return 0, fmt.Errorf("%d is out of range", n)
	}

	return rune(zipKey[n]), nil
}

func Decode64(ch rune) (int, error) {
	zipKey := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

	i := strings.Index(zipKey, string(ch))
	if i < 0 {
		return i, errors.New("encoded rune is not present in zip key")
	}

	return i, nil
}
