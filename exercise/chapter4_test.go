package exercise

import (
	"encoding/hex"
	"fmt"
	"github.com/YOwatari/programmingbitcoin/helper"
	"math/big"

	"github.com/YOwatari/programmingbitcoin/ecc"
)

func ExampleChapter4_one() {
	secrets := []int64{
		5000,
		0x76e54a40efb620, // 2018^5
		0xdeadbeef12345,
	}
	for _, s := range secrets {
		key := ecc.NewPrivateKey(big.NewInt(s))
		fmt.Printf("%x\n", key.Point.Sec(false))
	}

	// Output:
	// 04ffe558e388852f0120e46af2d1b370f85854a8eb0841811ece0e3e03d282d57c315dc72890a4f10a1481c031b03b351b0dc79901ca18a00cf009dbdb157a1d10
	// 04027f3da1918455e03c46f659266a1bb5204e959db7364d2f473bdf8f0a13cc9dff87647fd023c13b4a4994f17691895806e1b40b57f4fd22581a4f46851f3b06
	// 04d90cd625ee87dd38656dd95cf79f65f60f7273b67d3096e68bd81e4f5342691f842efa762fd59961d0e99803c61edba8b3e3f7dc3a341836f97733aebf987121
}

func ExampleChapter4_two() {
	secrets := []int64{
		5001,
		0x7730c781f7ae53, // 2019^5
		0xdeadbeef54321,
	}
	for _, s := range secrets {
		key := ecc.NewPrivateKey(big.NewInt(s))
		fmt.Printf("%x\n", key.Point.Sec(true))
	}

	// Output:
	// 0357a4f368868a8a6d572991e484e664810ff14c05c0fa023275251151fe0e53d1
	// 02933ec2d2b111b92737ec12f1c5d20f3233a0ad21cd8b36d0bca7a0cfa5cb8701
	// 0296be5b1292f6c856b3c5654e886fc13511462059089cdf9c479623bfcbe77690
}

func ExampleChapter4_three() {
	r, _ := new(big.Int).SetString("37206a0610995c58074999cb9767b87af4c4978db68c06e8e6e81d282047a7c6", 16)
	s, _ := new(big.Int).SetString("8ca63759c1157ebeaec0d03cecca119fc9a75bf8e6d0fa65c841c8e2738cdaec", 16)
	sec := ecc.NewSignature(r, s)
	fmt.Printf("%x", sec.Der())

	// Output:
	// 3045022037206a0610995c58074999cb9767b87af4c4978db68c06e8e6e81d282047a7c60221008ca63759c1157ebeaec0d03cecca119fc9a75bf8e6d0fa65c841c8e2738cdaec
}

func ExampleChapter4_four() {
	hs := []string{
		"7c076ff316692a3d7eb3c3bb0f8b1488cf72e1afcd929e29307032997a838a3d",
		"eff69ef2b1bd93a66ed5219add4fb51e11a840f404876325a1e8ffe0529a2c",
		"c7207fee197d27c618aea621406f6bf5ef6fca38681d82b2f06fddbdce6feab6",
	}
	for _, h := range hs {
		b, err := hex.DecodeString(h)
		if err != nil {
			panic(err)
		}
		fmt.Println(helper.EncodeBase58(b))
	}

	// Output:
	// 9MA8fRQrT4u8Zj8ZRd6MAiiyaxb2Y1CMpvVkHQu5hVM6
	// 4fE3H2E6XMp4SsxtwinF7w9a34ooUrwWe4WsW1458Pd
	// EQJsjkd6JaGwxrjEhfeqPenqHwrBmPQZjJGNSCHBkcF7
}

func ExampleChapter4_five() {
	ins := []struct{
		secret int64
		compress bool
		testnet bool
	} {
		{
			5002,
			false,
			true,
		},
		{
			0x777c6b16216400,
			true,
			true,
		},
		{
			0x12345deadbeef,
			true,
			false,
		},
	}
	for _, in := range ins {
		key :=  ecc.NewPrivateKey(big.NewInt(in.secret))
		fmt.Println(key.Point.Address(in.compress, in.testnet))
	}

	// Output:
	// mmTPbXQFxboEtNRkwfh6K51jvdtHLxGeMA
	// mopVkxp8UhXqRYbCYJsbeE1h1fiF64jcoH
	// 1F1Pn2y6pDb68E5nYJJeba4TLg2U7B6KF1
}
