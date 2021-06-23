package tx

import (
	"bytes"
	"encoding/hex"

	"github.com/YOwatari/programmingbitcoin/helper"
	"github.com/YOwatari/programmingbitcoin/script"
)

type Input struct {
	PrevTxID  [32]byte
	PrevIndex uint32
	ScriptSig *script.Script
	Sequence  uint32
}

func NewInput(prevTxID []byte, prevIndex uint32, scriptSig *script.Script, sequence uint32) *Input {
	var id [32]byte
	copy(id[:], prevTxID[:32])

	return &Input{
		PrevTxID:  id,
		PrevIndex: prevIndex,
		ScriptSig: scriptSig,
		Sequence:  sequence,
	}
}

func ParseInput(r *bytes.Reader) (*Input, error) {
	id := make([]byte, 32)
	if _, err := r.Read(id); err != nil {
		return nil, err
	}
	helper.ReverseBytes(id)

	buf := make([]byte, 4)
	if _, err := r.Read(buf); err != nil {
		return nil, err
	}
	index := helper.LittleEndianToInt64(buf)
	parseScript, err := script.ParseScript(r)
	if err != nil {
		return nil, err
	}
	if _, err := r.Read(buf); err != nil {
		return nil, err
	}
	seq := helper.LittleEndianToInt64(buf)
	return NewInput(id, uint32(index), parseScript, uint32(seq)), nil
}

func (in *Input) Serialize() []byte {
	buf := new(bytes.Buffer)
	buf.Write(helper.ReverseBytes(in.PrevTxID[:]))
	buf.Write(helper.Uint32ToLittleEndian(in.PrevIndex))
	buf.Write(in.ScriptSig.Serialize())
	buf.Write(helper.Uint32ToLittleEndian(in.Sequence))
	return buf.Bytes()
}

func (in *Input) fetchTransaction(testnet bool) *Transaction {
	f := getTransactionFetcher()
	return TransactionFetcherMust(f.Fetch(hex.EncodeToString(in.PrevTxID[:]), testnet, false))
}

func (in *Input) Value(testnet bool) uint64 {
	tx := in.fetchTransaction(testnet)
	return tx.Outputs[in.PrevIndex].Amount
}

func (in *Input) ScriptPubKey(testnet bool) *script.Script {
	tx := in.fetchTransaction(testnet)
	return tx.Outputs[in.PrevIndex].ScriptPubKey
}
