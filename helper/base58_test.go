package helper_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/YOwatari/programmingbitcoin/helper"
)

func TestEncodeBase58(t *testing.T) {
	cases := []struct {
		hex      string
		expected string
	}{
		{
			"7c076ff316692a3d7eb3c3bb0f8b1488cf72e1afcd929e29307032997a838a3d",
			"9MA8fRQrT4u8Zj8ZRd6MAiiyaxb2Y1CMpvVkHQu5hVM6",
		},
		{
			"eff69ef2b1bd93a66ed5219add4fb51e11a840f404876325a1e8ffe0529a2c",
			"4fE3H2E6XMp4SsxtwinF7w9a34ooUrwWe4WsW1458Pd",
		},
		{
			"c7207fee197d27c618aea621406f6bf5ef6fca38681d82b2f06fddbdce6feab6",
			"EQJsjkd6JaGwxrjEhfeqPenqHwrBmPQZjJGNSCHBkcF7",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			b, err := hex.DecodeString(c.hex)
			if err != nil {
				t.Fatal(err)
			}
			actual := helper.EncodeBase58(b)
			if actual != c.expected {
				t.Errorf("base58 diff\n got: %s\nwant: %s", actual, c.expected)
			}
		})
	}
}
