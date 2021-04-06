package main

import (
	"flag"
	"fmt"

	"github.com/YOwatari/programmingbitcoin/ecc"
	"github.com/YOwatari/programmingbitcoin/helper"
)

func main() {
	pass := flag.String("pass", "default", "passphrase")
	flag.Parse()

	secret := helper.LittleEndianToBigInt(helper.Hash256([]byte(*pass)))
	key := ecc.NewPrivateKey(secret)
	fmt.Printf("address (testnet): %s\n", key.Point.Address(true, true))
}
