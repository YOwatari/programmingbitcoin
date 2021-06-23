package tx

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/YOwatari/programmingbitcoin/helper"
	"io"
	"net/http"
	"sync"
)

var (
	instance *TransactionFetcher
	once sync.Once
)

type TransactionFetcher struct {
	cache map[string]*Transaction
}

func getTransactionFetcher() *TransactionFetcher {
	once.Do(func() {
		instance = &TransactionFetcher{cache: make(map[string]*Transaction)}
	})
	return instance
}

func getURL(testnet bool) string {
	if testnet {
		return "http://testnet.programmingbitcoin.com"
	}
	return "http://mainnet.programmingbitcoin.com"
}

func TransactionFetcherMust(transaction *Transaction, err error) *Transaction {
	if err != nil {
		panic(err)
	}
	return transaction
}

func (f *TransactionFetcher) Fetch(txID string, testnet bool, fresh bool) (*Transaction, error) {
	if tx, ok := f.cache[txID]; ok && !fresh {
		tx.Testnet = testnet
		return tx, nil
	}

	url := fmt.Sprintf("%s/tx/%s.hex", getURL(testnet), txID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	raw := make([]byte, hex.DecodedLen(len(b)))
	if _, err := hex.Decode(raw, b); err != nil {
		return nil, err
	}

	var tx *Transaction
	if raw[4] == 0 {
		length := len(raw)
		locktime := uint32(helper.LittleEndianToInt64(raw[length-4:]))
		copy(raw[4:], raw[6:])
		r := bytes.NewReader(raw[:length-2])
		tx, err = ParseTransaction(r, testnet)
		if err != nil {
			return nil, err
		}
		tx.Locktime = locktime
	} else {
		r := bytes.NewReader(raw)
		tx, err = ParseTransaction(r, testnet)
		if err != nil {
			return nil, err
		}
	}

	if txID != tx.ID() {
		return nil, fmt.Errorf("not the same id: %s vs %s", tx.ID(), txID)
	}

	f.cache[txID] = tx
	tx.Testnet = testnet
	return tx, nil
}
