package helper

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

func Hash256(b []byte) []byte {
	hash1 := sha256.New()
	hash1.Write(b)
	hash2 := sha256.New()
	hash2.Write(hash1.Sum(nil))
	return hash2.Sum(nil)
}

func Hash160(b []byte) []byte {
	hash1 := sha256.New()
	hash1.Write(b)
	hash2 := ripemd160.New()
	hash2.Write(hash1.Sum(nil))
	return hash2.Sum(nil)
}
