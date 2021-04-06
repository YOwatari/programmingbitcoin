package ecc

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/YOwatari/programmingbitcoin/helper"
)

const (
	_WIFPrefixMainnet  byte = 0x80
	_WIFPrefixTestnet  byte = 0xef
	_WIFSuffixCompress byte = 0x01
)

type PrivateKey struct {
	Secret *big.Int
	Point  *S256Point
}

func NewPrivateKey(secret *big.Int) *PrivateKey {
	return &PrivateKey{Secret: secret, Point: new(S256Point).RMul(G, secret)}
}

func (key *PrivateKey) Hex() string {
	return fmt.Sprintf("%064x", key.Secret.Bytes())
}

func (key *PrivateKey) Sign(z *big.Int) *Signature {
	k := key.deterministicK(z)
	r := new(S256Point).RMul(G, k).X.(*FieldElement).Num
	invK := new(big.Int).Exp(k, new(big.Int).Sub(N, big.NewInt(2)), N)
	s := new(big.Int)
	s.Add(z, new(big.Int).Mul(r, key.Secret)).Mul(s, invK).Mod(s, N)
	if s.Cmp(new(big.Int).Div(r, big.NewInt(2))) > 0 {
		s.Mul(N, s)
	}
	return NewSignature(r, s)
}

func (key *PrivateKey) deterministicK(z *big.Int) *big.Int {
	k := bytes.Repeat([]byte{0x00}, 32)
	v := bytes.Repeat([]byte{0x01}, 32)

	if z.Cmp(N) > 0 {
		z.Sub(z, N)
	}

	zb := int2bytes(z)
	sb := int2bytes(key.Secret)

	_hmac := func(k []byte, values ...[]byte) []byte {
		h := hmac.New(sha256.New, k)
		for _, v := range values {
			h.Write(v)
		}
		return h.Sum(nil)
	}
	k = _hmac(k, v, []byte{0x00}, sb, zb)
	v = _hmac(k, v)
	k = _hmac(k, v, []byte{0x01}, sb, zb)
	v = _hmac(k, v)

	candidate := new(big.Int)
	for {
		v = _hmac(k, v)
		candidate.SetBytes(v)
		if candidate.Sign() > 0 && candidate.Cmp(N) < 0 {
			break
		}

		k = _hmac(k, v, []byte{0x00})
		v = _hmac(k, v)
	}
	return candidate
}

func (key *PrivateKey) Wif(compress, testnet bool) string {
	var b bytes.Buffer
	if testnet {
		b.WriteByte(_WIFPrefixTestnet)
	} else {
		b.WriteByte(_WIFPrefixMainnet)
	}
	b.Write(int2bytes(key.Secret))
	if compress {
		b.WriteByte(_WIFSuffixCompress)
	}
	return helper.EncodeBase58Checksum(b.Bytes())
}

func int2bytes(n *big.Int) []byte {
	size := 32
	raw := n.Bytes()
	result := make([]byte, size)
	copy(result[size-len(raw):], raw)
	return result
}
