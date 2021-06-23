package script

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/YOwatari/programmingbitcoin/helper"
	"strings"
)

type Script struct {
	cmds [][]byte
}

func NewScript(cmds [][]byte) *Script {
	return &Script{cmds: cmds}
}

func ParseScript(r *bytes.Reader) (*Script, error) {
	length := helper.ReadVarInt(r)
	var cmds [][]byte
	count := 0
	for count < length {
		current, err := r.ReadByte() // opcode or element
		if err != nil {
			return nil, err
		}
		count++

		switch {
		case current >= 1 && current <= 75:
			n := int(current)
			buf := make([]byte, n)
			r.Read(buf)
			cmds = append(cmds, buf)
			count += n
		case current == 76:
			data, err := r.ReadByte()
			if err != nil {
				return nil, err
			}
			n := int(data)

			buf := make([]byte, n)
			r.Read(buf)
			cmds = append(cmds, buf)

			count += n + 1
		case current == 77:
			data := make([]byte, 2)
			r.Read(data)
			n := int(helper.LittleEndianToInt64(data))

			buf := make([]byte, n)
			r.Read(buf)
			cmds = append(cmds, buf)

			count += n + 2
		default:
			opcode := current
			cmds = append(cmds, []byte{opcode})
		}
	}

	if count != length {
		return nil, fmt.Errorf("cannot parse script. length: %d, count: %d", length, count)
	}

	return NewScript(cmds), nil
}

func (s *Script) Serialize() []byte {
	raw := new(bytes.Buffer)
	for _, cmd := range s.cmds {
		length := len(cmd)
		switch {
		case length == 1:
			raw.WriteByte(cmd[0])
		case length < 76:
			raw.WriteByte(byte(length))
		case length < 0x100:
			raw.WriteByte(byte(76))
			raw.WriteByte(byte(length))
		case length <= 520:
			raw.WriteByte(byte(77))
			raw.WriteByte(byte(length))
		default:
			raw.Write(cmd)
		}
	}
	buf := new(bytes.Buffer)
	buf.Write(helper.EncodeVarInt(raw.Len()))
	buf.Write(raw.Bytes())
	return buf.Bytes()
}

func (s *Script) String() string {
	result := make([]string, len(s.cmds))
	for i, cmd := range s.cmds {
		if len(cmd) == 1 {
			opcode := int(cmd[0])
			if name, ok := opcodeName[opcode]; ok {
				result[i] = name
			} else {
				result[i] = fmt.Sprintf("OP_[%d]", opcode)
			}
		} else {
			result[i] = hex.EncodeToString(cmd)
		}
	}
	return strings.Join(result, " ")
}
