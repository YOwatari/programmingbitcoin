package ecc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func NewSignature(r, s *big.Int) *Signature {
	return &Signature{R: r, S: s}
}

func (s *Signature) String() string {
	return fmt.Sprintf("Signature(%064x, %064x)", s.R, s.S)
}

func (s *Signature) Der() []byte {
	var body bytes.Buffer

	rb := s.R.Bytes()
	if rb[0]&0x80 > 0 {
		rb = append([]byte{0x00}, rb...)
	}
	body.Write([]byte{0x02})
	binary.Write(&body, binary.BigEndian, uint8(len(rb)))
	body.Write(rb)

	sb := s.S.Bytes()
	if sb[0]&0x80 > 0 {
		sb = append([]byte{0x00}, sb...)
	}
	body.Write([]byte{0x02})
	binary.Write(&body, binary.BigEndian, uint8(len(sb)))
	body.Write(sb)

	var h bytes.Buffer
	h.Write([]byte{0x30})
	binary.Write(&h, binary.BigEndian, uint8(body.Len()))
	h.Write(body.Bytes())

	return h.Bytes()
}
