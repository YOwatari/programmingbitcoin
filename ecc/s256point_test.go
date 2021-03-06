package ecc_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/YOwatari/programmingbitcoin/ecc"
)

func TestS256Point_Sec(t *testing.T) {
	cases := []struct {
		secret   int64
		compress bool
		expected string
	}{
		{
			5000,
			false,
			"04ffe558e388852f0120e46af2d1b370f85854a8eb0841811ece0e3e03d282d57c315dc72890a4f10a1481c031b03b351b0dc79901ca18a00cf009dbdb157a1d10",
		},
		{
			0x76e54a40efb620,
			false,
			"04027f3da1918455e03c46f659266a1bb5204e959db7364d2f473bdf8f0a13cc9dff87647fd023c13b4a4994f17691895806e1b40b57f4fd22581a4f46851f3b06",
		},
		{
			0xdeadbeef12345,
			false,
			"04d90cd625ee87dd38656dd95cf79f65f60f7273b67d3096e68bd81e4f5342691f842efa762fd59961d0e99803c61edba8b3e3f7dc3a341836f97733aebf987121",
		},
		{
			5001,
			true,
			"0357a4f368868a8a6d572991e484e664810ff14c05c0fa023275251151fe0e53d1",
		},
		{
			0x7730c781f7ae53,
			true,
			"02933ec2d2b111b92737ec12f1c5d20f3233a0ad21cd8b36d0bca7a0cfa5cb8701",
		},
		{
			0xdeadbeef54321,
			true,
			"0296be5b1292f6c856b3c5654e886fc13511462059089cdf9c479623bfcbe77690",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			key := ecc.NewPrivateKey(big.NewInt(c.secret))
			actual := key.Point.Sec(c.compress)
			expected, err := hex.DecodeString(c.expected)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(actual, expected) {
				t.Errorf("Public key's SEC format diff\n got: %x\nwant: %x", actual, expected)
			}
		})
	}
}

func TestS256Point_Address(t *testing.T) {
	cases := []struct{
		secret int64
		compress bool
		testnet bool
		expected string
	} {
		{
			5002,
			false,
			true,
			"mmTPbXQFxboEtNRkwfh6K51jvdtHLxGeMA",
		},
		{
			0x777c6b16216400,
			true,
			true,
			"mopVkxp8UhXqRYbCYJsbeE1h1fiF64jcoH",
		},
		{
			0x12345deadbeef,
			true,
			false,
			"1F1Pn2y6pDb68E5nYJJeba4TLg2U7B6KF1",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			key := ecc.NewPrivateKey(big.NewInt(c.secret))
			actual := key.Point.Address(c.compress, c.testnet)
			if actual != c.expected {
				t.Errorf("adress diff:\n got: %s\nwant: %s", actual, c.expected)
			}
		})
	}
}
