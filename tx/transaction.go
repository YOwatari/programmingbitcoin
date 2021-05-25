package tx

import (
	"io"

	"github.com/YOwatari/programmingbitcoin/helper"
)

type Transaction struct {
	Version uint32
}

func NewTransaction(version uint32) *Transaction {
	return &Transaction{
		Version: version,
	}
}

func ParseTransaction(r io.Reader) *Transaction {
	buf := make([]byte, 4)
	r.Read(buf)
	version := helper.LittleEndianToInt64(buf)
	return NewTransaction(uint32(version))
}
