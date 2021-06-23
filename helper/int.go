package helper

import (
	"bytes"
	"encoding/binary"
	"math/big"
)

func LittleEndianToBigInt(b []byte) *big.Int {
	var buf bytes.Buffer
	for i := len(b) - 1; i >= 0; i-- {
		buf.WriteByte(b[i])
	}
	return new(big.Int).SetBytes(buf.Bytes())
}

func LittleEndianToInt64(b []byte) int64 {
	if len(b) > 8 {
		panic("bytes is too large")
	}
	if len(b) < 8 {
		b = append(b, make([]byte, 8-len(b))...)
	}

	var result int64
	buf := bytes.NewReader(b)
	if err := binary.Read(buf, binary.LittleEndian, &result); err != nil {
		panic(err)
	}
	return result
}

func Uint64ToLittleEndian(n uint64) []byte {
	b := new(bytes.Buffer)
	if err := binary.Write(b, binary.LittleEndian, &n); err != nil {
		panic(err)
	}
	return b.Bytes()
}

func Uint32ToLittleEndian(n uint32) []byte {
	b := new(bytes.Buffer)
	if err := binary.Write(b, binary.LittleEndian, &n); err != nil {
		panic(err)
	}
	return b.Bytes()
}

func Uint16ToLittleEndian(n uint16) []byte {
	b := new(bytes.Buffer)
	if err := binary.Write(b, binary.LittleEndian, &n); err != nil {
		panic(err)
	}
	return b.Bytes()
}
