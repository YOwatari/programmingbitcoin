package exercise_test

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/YOwatari/programmingbitcoin/ecc"
)

func ExampleChapter3_one()  {
	prime := int64(223)
	a, _ := ecc.NewFieldElementFromInt64(0, prime)
	b, _ := ecc.NewFieldElementFromInt64(7, prime)

	points := []struct{
		x int64
		y int64
	} {
		{192, 105},
		{17, 56},
		{200, 119},
		{1, 193},
		{42, 99},
	}

	for _, p := range points {
		y, _ := ecc.NewFieldElementFromInt64(p.y, prime)
		x, _ := ecc.NewFieldElementFromInt64(p.x, prime)
		_, err := ecc.NewPoint(x, y, a, b)
		fmt.Printf("(%d, %d): %t\n", p.x, p.y, err == nil)
	}

	// Output:
	// (192, 105): true
	// (17, 56): true
	// (200, 119): false
	// (1, 193): true
	// (42, 99): false
}

func ExampleChapter3_two() {
	prime := int64(223)
	a, _ := ecc.NewFieldElementFromInt64(0, prime)
	b, _ := ecc.NewFieldElementFromInt64(7, prime)

	type point struct {
		x int64
		y int64
	}
	cases := []struct{
		a point
		b point
	} {
		{
			point{170, 142},
			point{60, 139},
		},
		{
			point{47, 71},
			point{17, 56},
		},
		{
			point{143, 98},
			point{76, 66},
		},
	}

	for _, c := range cases {
		y1, _ := ecc.NewFieldElementFromInt64(c.a.y, prime)
		x1, _ := ecc.NewFieldElementFromInt64(c.a.x, prime)
		y2, _ := ecc.NewFieldElementFromInt64(c.b.y, prime)
		x2, _ := ecc.NewFieldElementFromInt64(c.b.x, prime)

		p1, err := ecc.NewPoint(x1, y1, a, b)
		if err != nil {
			panic(err)
		}
		p2, err := ecc.NewPoint(x2, y2, a, b)
		if err != nil {
			panic(err)
		}

		p := new(ecc.Point)
		p, err = p.Add(p1, p2).Calc()
		if err != nil {
			panic(err)
		}

		fmt.Printf("(%d, %d) + (%d, %d) = (%d, %d)\n", c.a.x, c.a.y, c.b.x, c.b.y, p.X.(*ecc.FieldElement).Num, p.Y.(*ecc.FieldElement).Num)
	}

	// Output:
	// (170, 142) + (60, 139) = (220, 181)
	// (47, 71) + (17, 56) = (215, 68)
	// (143, 98) + (76, 66) = (47, 71)
}

func ExampleChapter3_four() {
	prime := int64(223)
	a, _ := ecc.NewFieldElementFromInt64(0, prime)
	b, _ := ecc.NewFieldElementFromInt64(7, prime)

	cases := []struct{
		n int
		x int64
		y int64
	} {
		{
			2, 192, 105,
		},
		{
			2, 143, 98,
		},
		{
			2, 47, 71,
		},
		{
			4, 47, 71,
		},
		{
			8, 47, 71,
		},
		{
			21, 47, 71,
		},
	}

	for _, c := range cases {
		ry, _ := ecc.NewFieldElementFromInt64(c.y, prime)
		rx, _ := ecc.NewFieldElementFromInt64(c.x, prime)
		result, _ := ecc.NewPoint(rx, ry, a, b)
		for i := 0; i < c.n - 1; i++ {
			y, _ := ecc.NewFieldElementFromInt64(c.y, prime)
			x, _ := ecc.NewFieldElementFromInt64(c.x, prime)
			p, _ := ecc.NewPoint(x, y, a, b)
			result.Add(result, p)
		}
		p, err := result.Calc()
		if err != nil {
			panic(err)
		}

		if p.X == nil || p.Y == nil {
			fmt.Printf("%d * (%d, %d) = %s\n", c.n, c.x, c.y, p)
		} else {
			fmt.Printf("%d * (%d, %d) = Point(%d, %d)\n", c.n, c.x, c.y, p.X.(*ecc.FieldElement).Num, p.Y.(*ecc.FieldElement).Num)
		}
	}

	// Output:
	// 2 * (192, 105) = Point(49, 71)
	// 2 * (143, 98) = Point(64, 168)
	// 2 * (47, 71) = Point(36, 111)
	// 4 * (47, 71) = Point(194, 51)
	// 8 * (47, 71) = Point(116, 55)
	// 21 * (47, 71) = Point(infinity)
}

func ExampleChapter3_five()  {
	prime := int64(223)
	a, _ := ecc.NewFieldElementFromInt64(0, prime)
	b, _ := ecc.NewFieldElementFromInt64(7, prime)
	x, _ := ecc.NewFieldElementFromInt64(15, prime)
	y, _ := ecc.NewFieldElementFromInt64(86, prime)
	p, _ := ecc.NewPoint(x, y, a, b)

	result, _ := ecc.NewPoint(x, y, a, b)
	inf := &ecc.Point{X: nil, Y: nil, A: a, B: b}

	n := 1
	for result.Ne(inf) {
		result.Add(result, p)
		n++
	}

	_, err := result.Calc()
	if err != nil {
		panic(err)
	}

	fmt.Println(n)

	// Output:
	// 7
}

func ExampleChapter3_six() {
	x, _ := new(big.Int).SetString("887387e452b8eacc4acfde10d9aaf7f6d9a0f975aabb10d006e4da568744d06c", 16)
	y, _ := new(big.Int).SetString("61de6d95231cd89026e286df3b6ae4a894a3378e393e93a0f45b666329a0ae34", 16)
	p, err := ecc.NewS256PointFromBigInt(x, y)
	if err != nil {
		panic(err)
	}

	sigs := []struct{
		z string
		r string
		s string
	} {
		{
			"ec208baa0fc1c19f708a9ca96fdeff3ac3f230bb4a7ba4aede4942ad003c0f60",
			"ac8d1c87e51d0d441be8b3dd5b05c8795b48875dffe00b7ffcfac23010d3a395",
			"68342ceff8935ededd102dd876ffd6ba72d6a427a3edb13d26eb0781cb423c4",
		},
		{
			"7c076ff316692a3d7eb3c3bb0f8b1488cf72e1afcd929e29307032997a838a3d",
			"eff69ef2b1bd93a66ed5219add4fb51e11a840f404876325a1e8ffe0529a2c",
			"c7207fee197d27c618aea621406f6bf5ef6fca38681d82b2f06fddbdce6feab6",
		},
	}

	for _, sig := range sigs {
		z, _ := new(big.Int).SetString(sig.z, 16)
		r, _ := new(big.Int).SetString(sig.r, 16)
		s, _ := new(big.Int).SetString(sig.s, 16)

		invS := new(big.Int).Exp(s, new(big.Int).Sub(ecc.N, big.NewInt(2)), ecc.N)
		u := new(big.Int)
		u = u.Mul(z, invS).Mod(u, ecc.N)
		v := new(big.Int)
		v = v.Mul(r, invS).Mod(v, ecc.N)
		n := new(ecc.S256Point)
		n.Add(new(ecc.S256Point).RMul(ecc.G, u), new(ecc.S256Point).RMul(p, v))
		fmt.Println(n.X.(*ecc.FieldElement).Num.Cmp(r) == 0)
	}

	// Output:
	// true
	// true
}

func ExampleChapter3_seven() {
	getHash := func(s string) *big.Int {
		r1 := sha256.Sum256([]byte(s))
		r2 := sha256.Sum256(r1[:])
		return new(big.Int).SetBytes(r2[:])
	}

	e := big.NewInt(12345)
	z := getHash("Programming Bitcoin!")

	k := big.NewInt(1234567890)
	r := new(ecc.S256Point).RMul(ecc.G, k).X.(*ecc.FieldElement).Num
	s := new(big.Int)
	invK := new(big.Int).Exp(k, new(big.Int).Sub(ecc.N, big.NewInt(2)), ecc.N)
	s.Add(z, new(big.Int).Mul(r, e)).Mul(s, invK).Mod(s, ecc.N)
	point := new(ecc.S256Point).RMul(ecc.G, e)
	sig := ecc.NewSignature(r, s)
	fmt.Printf("signature: %s\n", sig)
	fmt.Printf("point: %064x, %064x\n", point.X.(*ecc.FieldElement).Num, point.Y.(*ecc.FieldElement).Num)

	// Output:
	// signature: Signature(2b698a0f0a4041b77e63488ad48c23e8e8838dd1fb7520408b121697b782ef22, 1dbc63bfef4416705e602a7b564161167076d8b20990a0f26f316cff2cb0bc1a)
	// point: f01d6b9018ab421dd410404cb869072065522bf85734008f105cf385a023a80f, 0eba29d0f0c5408ed681984dc525982abefccd9f7ff01dd26da4999cf3f6a295
}
