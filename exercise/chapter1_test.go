package exercise_test

import (
	"fmt"
	"github.com/YOwatari/programmingbitcoin/ecc"
	"strings"
)

func ExampleChapter1_two() {
	prime := 57
	fmt.Printf("Prime: %d\n", prime)

	a1, _ := ecc.NewFieldElement(44, prime)
	a2, _ := ecc.NewFieldElement(33, prime)
	a, _ := a1.Add(a2).Calc()
	fmt.Printf("44 + 33 = %d\n", a.Num)

	b1, _ := ecc.NewFieldElement(9, prime)
	b2, _ := ecc.NewFieldElement(29, prime)
	b, _ := b1.Sub(b2).Calc()
	fmt.Printf("9 - 29 = %d\n", b.Num)

	c1, _ := ecc.NewFieldElement(17, prime)
	c2, _ := ecc.NewFieldElement(42, prime)
	c3, _ := ecc.NewFieldElement(49, prime)
	c, _ := c1.Add(c2).Add(c3).Calc()
	fmt.Printf("17 + 42 + 49 = %d\n", c.Num)

	d1, _ := ecc.NewFieldElement(52, prime)
	d2, _ := ecc.NewFieldElement(30, prime)
	d3, _ := ecc.NewFieldElement(38, prime)
	d, _ := d1.Sub(d2).Sub(d3).Calc()
	fmt.Printf("52 - 30 - 38 = %d\n", d.Num)

	// Output:
	// Prime: 57
	// 44 + 33 = 20
	// 9 - 29 = 37
	// 17 + 42 + 49 = 51
	// 52 - 30 - 38 = 41
}

func ExampleChapter1_four() {
	prime := 97
	fmt.Printf("Prime: %d\n", prime)

	a1, _ := ecc.NewFieldElement(95, prime)
	a2, _ := ecc.NewFieldElement(45, prime)
	a3, _ := ecc.NewFieldElement(31, prime)
	a, _ := a1.Mul(a2).Mul(a3).Calc()
	fmt.Printf("97 * 45 * 31 = %d\n", a.Num)

	b1, _ := ecc.NewFieldElement(17, prime)
	b2, _ := ecc.NewFieldElement(13, prime)
	b3, _ := ecc.NewFieldElement(19, prime)
	b4, _ := ecc.NewFieldElement(44, prime)
	b, _ := b1.Mul(b2).Mul(b3).Mul(b4).Calc()
	fmt.Printf("17 * 13 * 19 * 44 = %d\n", b.Num)

	c1, _ := ecc.NewFieldElement(12, prime)
	c2, _ := ecc.NewFieldElement(77, prime)
	c, _ := c1.Pow(7).Mul(c2.Pow(49)).Calc()
	fmt.Printf("12**7 * 77**49 = %d\n", c.Num)

	// Output:
	// Prime: 97
	// 97 * 45 * 31 = 23
	// 17 * 13 * 19 * 44 = 68
	// 12**7 * 77**49 = 63
}

func ExampleChapter1_five() {
	prime := 19
	for _, k := range []int{1, 3, 7, 13, 18} {
		result := make([]string, prime)
		for i := 0; i < len(result); i++ {
			n1, _ := ecc.NewFieldElement(k, prime)
			n2, _ := ecc.NewFieldElement(i, prime)
			n, _ := n1.Mul(n2).Calc()
			result[i] = fmt.Sprintf("%02d", n.Num)
		}
		fmt.Printf("k=%02d [%v]\n", k, strings.Join(result, ", "))
	}

	// Output:
	// k=01 [00, 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12, 13, 14, 15, 16, 17, 18]
	// k=03 [00, 03, 06, 09, 12, 15, 18, 02, 05, 08, 11, 14, 17, 01, 04, 07, 10, 13, 16]
	// k=07 [00, 07, 14, 02, 09, 16, 04, 11, 18, 06, 13, 01, 08, 15, 03, 10, 17, 05, 12]
	// k=13 [00, 13, 07, 01, 14, 08, 02, 15, 09, 03, 16, 10, 04, 17, 11, 05, 18, 12, 06]
	// k=18 [00, 18, 17, 16, 15, 14, 13, 12, 11, 10, 09, 08, 07, 06, 05, 04, 03, 02, 01]
}

func ExampleChapter1_seven() {
	for _, p := range []int{7, 11, 17, 31} {
		e := p - 1
		result := make([]string, 0)
		for i := 1; i < p; i++ {
			n, _ := ecc.NewFieldElement(i, p)
			n, _ = n.Pow(e).Calc()
			result = append(result, fmt.Sprintf("%d", n.Num))
		}
		fmt.Printf("p=%02d [%v]\n", p, strings.Join(result, ", "))
	}

	// Output:
	// p=07 [1, 1, 1, 1, 1, 1]
	// p=11 [1, 1, 1, 1, 1, 1, 1, 1, 1, 1]
	// p=17 [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]
	// p=31 [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]
}
