package ecc_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/YOwatari/programmingbitcoin/ecc"
)

func TestS256Point_Verify(t *testing.T) {
	x, _ := new(big.Int).SetString("887387e452b8eacc4acfde10d9aaf7f6d9a0f975aabb10d006e4da568744d06c", 16)
	y, _ := new(big.Int).SetString("61de6d95231cd89026e286df3b6ae4a894a3378e393e93a0f45b666329a0ae34", 16)
	p, err := ecc.NewS256PointFromBigInt(x, y)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("point: %s", p)

	cases := []struct {
		z        string
		r        string
		s        string
		expected bool
	}{
		{
			"ec208baa0fc1c19f708a9ca96fdeff3ac3f230bb4a7ba4aede4942ad003c0f60",
			"ac8d1c87e51d0d441be8b3dd5b05c8795b48875dffe00b7ffcfac23010d3a395",
			"68342ceff8935ededd102dd876ffd6ba72d6a427a3edb13d26eb0781cb423c4",
			true,
		},
		{
			"7c076ff316692a3d7eb3c3bb0f8b1488cf72e1afcd929e29307032997a838a3d",
			"eff69ef2b1bd93a66ed5219add4fb51e11a840f404876325a1e8ffe0529a2c",
			"c7207fee197d27c618aea621406f6bf5ef6fca38681d82b2f06fddbdce6feab6",
			true,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			z, _ := new(big.Int).SetString(c.z, 16)
			r, _ := new(big.Int).SetString(c.r, 16)
			s, _ := new(big.Int).SetString(c.s, 16)

			sig := ecc.NewSignature(r, s)
			if actual := p.Verify(z, sig); actual != c.expected {
				t.Errorf("\ngot: %t\nwant: %t\n\nsig: %s", actual, c.expected, sig)
			}
		})
	}
}

func TestSignature_Der(t *testing.T) {
	r, _ := new(big.Int).SetString("37206a0610995c58074999cb9767b87af4c4978db68c06e8e6e81d282047a7c6", 16)
	s, _ := new(big.Int).SetString("8ca63759c1157ebeaec0d03cecca119fc9a75bf8e6d0fa65c841c8e2738cdaec", 16)
	sec := ecc.NewSignature(r, s)
	actual := sec.Der()
	expected, err := hex.DecodeString("3045022037206a0610995c58074999cb9767b87af4c4978db68c06e8e6e81d282047a7c60221008ca63759c1157ebeaec0d03cecca119fc9a75bf8e6d0fa65c841c8e2738cdaec")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(actual, expected) {
		t.Errorf("Signiture DER format diff\n got: %x\nwant: %x", actual, expected)
	}
}
