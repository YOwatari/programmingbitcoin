package tx

import (
	"bytes"

	"github.com/YOwatari/programmingbitcoin/helper"
	"github.com/YOwatari/programmingbitcoin/script"
)

type Output struct {
	Amount       uint64
	ScriptPubKey *script.Script
}

func NewOutput(amount uint64, scriptPubkey *script.Script) *Output {
	return &Output{
		Amount:       amount,
		ScriptPubKey: scriptPubkey,
	}
}

func ParseOutput(r *bytes.Reader) (*Output, error) {
	buf := make([]byte, 8)
	_, err := r.Read(buf)
	if err != nil {
		return nil, err
	}
	amount := helper.LittleEndianToInt64(buf)

	scriptPubkey, err := script.ParseScript(r)
	if err != nil {
		return nil, err
	}
	return NewOutput(uint64(amount), scriptPubkey), nil
}

func (out *Output) Serialize() []byte {
	buf := new(bytes.Buffer)
	buf.Write(helper.Uint64ToLittleEndian(out.Amount))
	buf.Write(out.ScriptPubKey.Serialize())
	return buf.Bytes()
}
