package helper

import (
	"bytes"
	"math/big"
	"strings"
)

const (
	_Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

func EncodeBase58(b []byte) string {
	count := 0
	for _, c := range b {
		if c == 0 {
			count++
		} else {
			break
		}
	}

	var buf bytes.Buffer
	alphabet := []byte(_Alphabet)
	num := new(big.Int).SetBytes(b)
	mod := new(big.Int)
	for num.Sign() > 0 {
		num.QuoRem(num, big.NewInt(58), mod)
		buf.WriteByte(alphabet[int(mod.Int64())])
	}

	result := strings.Repeat("1", count)
	for i := buf.Len() - 1; i >= 0; i-- {
		result = result + string(buf.Bytes()[i])
	}
	return result
}

func EncodeBase58Checksum(b []byte) string {
	return EncodeBase58(append(b, Hash256(b)[:4]...))
}
