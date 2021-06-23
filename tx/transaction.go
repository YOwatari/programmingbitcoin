package tx

import (
	"bytes"
	"encoding/hex"
	"github.com/YOwatari/programmingbitcoin/helper"
)

type Transaction struct {
	Version  uint32
	Inputs   []*Input
	Outputs  []*Output
	Locktime uint32
	Testnet  bool
}

func NewTransaction(version uint32, inputs []*Input, outputs []*Output, locktime uint32, testnet bool) *Transaction {
	return &Transaction{
		Version: version,
		Inputs: inputs,
		Outputs: outputs,
		Locktime: locktime,
		Testnet: testnet,
	}
}

func ParseTransaction(r *bytes.Reader, testnet bool) (*Transaction, error) {
	buf := make([]byte, 4)
	if _, err := r.Read(buf); err != nil {
		return nil, err
	}
	version := helper.LittleEndianToInt64(buf)

	num := helper.ReadVarInt(r)
	inputs := make([]*Input, 0, num)
	for i := 0; i < num; i++ {
		input, err := ParseInput(r)
		if err != nil {
			return nil, err
		}
		inputs = append(inputs, input)
	}

	num = helper.ReadVarInt(r)
	outputs := make([]*Output, 0, num)
	for i := 0; i < num; i++ {
		output, err := ParseOutput(r)
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, output)
	}

	if _, err := r.Read(buf); err != nil {
		return nil, err
	}
	lt := helper.LittleEndianToInt64(buf)

	return NewTransaction(uint32(version), inputs, outputs, uint32(lt), testnet), nil
}

func (tx *Transaction) ID() string {
	return tx.String()
}

func (tx *Transaction) String() string {
	return hex.EncodeToString(tx.Hash())
}

func (tx *Transaction) Hash() []byte {
	hash := helper.Hash256(tx.Serialize())
	return helper.ReverseBytes(hash)
}

func (tx *Transaction) Serialize() []byte {
	buf := new(bytes.Buffer)
	buf.Write(helper.Uint32ToLittleEndian(tx.Version))
	buf.Write(helper.EncodeVarInt(len(tx.Inputs)))
	for _, txIn := range tx.Inputs {
		buf.Write(txIn.Serialize())
	}
	buf.Write(helper.EncodeVarInt(len(tx.Outputs)))
	for _, txOut := range tx.Outputs {
		buf.Write(txOut.Serialize())
	}
	buf.Write(helper.Uint32ToLittleEndian(tx.Locktime))
	return buf.Bytes()
}

func (tx *Transaction) Fee() uint64 {
	var result uint64
	for _, in := range tx.Inputs {
		result += in.Value(false)
	}
	for _, out := range tx.Outputs {
		result -= out.Amount
	}
	return result
}
