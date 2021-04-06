package ecc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
)

const (
	_DERMarkerStart     byte = 0x30
	_DERMarkerDelimiter byte = 0x02
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
	body.WriteByte(_DERMarkerDelimiter)
	binary.Write(&body, binary.BigEndian, uint8(len(rb)))
	body.Write(rb)

	sb := s.S.Bytes()
	if sb[0]&0x80 > 0 {
		sb = append([]byte{0x00}, sb...)
	}
	body.WriteByte(_DERMarkerDelimiter)
	binary.Write(&body, binary.BigEndian, uint8(len(sb)))
	body.Write(sb)

	var result bytes.Buffer
	result.WriteByte(_DERMarkerStart)
	binary.Write(&result, binary.BigEndian, uint8(body.Len()))
	result.Write(body.Bytes())

	return result.Bytes()
}
