package helper

import (
	"bytes"
)

func ReadVarInt(r *bytes.Reader) int {
	b, err := r.ReadByte()
	if err != nil {
		panic(err)
	}

	varint := func(n int) int {
		buf := make([]byte, n)
		r.Read(buf)
		return int(LittleEndianToInt64(buf))
	}

	switch b {
	case 0xfd:
		return varint(2)
	case 0xfe:
		return varint(4)
	case 0xff:
		return varint(8)
	default:
		return int(b)
	}
}

func EncodeVarInt(n int) []byte {
	switch {
	case n < 0xfd:
		return []byte{byte(n)}
	case n < 0x10000:
		result := make([]byte, 3)
		copy(result[1:], Uint16ToLittleEndian(uint16(n)))
		result[0] = 0xfd
		return result
	case n < 0x100000000:
		result := make([]byte, 5)
		copy(result[1:], Uint32ToLittleEndian(uint32(n)))
		result[0] = 0xfe
		return result
	default:
		result := make([]byte, 9)
		copy(result[1:], Uint64ToLittleEndian(uint64(n)))
		result[0] = 0xff
		return result
	}
}
