package exercise_test

import (
	"fmt"
	"github.com/YOwatari/programmingbitcoin/ecc"
	"strings"
)

func ExampleChapter1_two() {
	prime := 57
	fmt.Printf("Prime: %d\n", prime)

	a, _ := ecc.NewFieldElement(44, prime)
	a.Calc(ecc.Add(&ecc.FieldElement{Num: 33, Prime: prime}))
	fmt.Printf("44 + 33 = %d\n", a.Num)

	b, _ := ecc.NewFieldElement(9, prime)
	b.Calc(ecc.Sub(&ecc.FieldElement{Num: 29, Prime: prime}))
	fmt.Printf("9 - 29 = %d\n", b.Num)

	c, _ := ecc.NewFieldElement(17, prime)
	c.Calc(
		ecc.Add(&ecc.FieldElement{Num: 42, Prime: prime}),
		ecc.Add(&ecc.FieldElement{Num: 49, Prime: prime}),
	)
	fmt.Printf("17 + 42 + 49 = %d\n", c.Num)

	d, _ := ecc.NewFieldElement(52, prime)
	d.Calc(
		ecc.Sub(&ecc.FieldElement{Num: 30, Prime: prime}),
		ecc.Sub(&ecc.FieldElement{Num: 38, Prime: prime}),
	)
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

	a, _ := ecc.NewFieldElement(95, prime)
	a.Calc(
		ecc.Mul(&ecc.FieldElement{Num: 45, Prime: prime}),
		ecc.Mul(&ecc.FieldElement{Num: 31, Prime: prime}),
	)
	fmt.Printf("97 * 45 * 31 = %d\n", a.Num)

	b, _ := ecc.NewFieldElement(17, prime)
	b.Calc(
		ecc.Mul(&ecc.FieldElement{Num: 13, Prime: prime}),
		ecc.Mul(&ecc.FieldElement{Num: 19, Prime: prime}),
		ecc.Mul(&ecc.FieldElement{Num: 44, Prime: prime}),
	)
	fmt.Printf("17 * 13 * 19 * 44 = %d\n", b.Num)

	c1, _ := ecc.NewFieldElement(12, prime)
	c1.Calc(ecc.Pow(7))
	c2, _ := ecc.NewFieldElement(77, prime)
	c2.Calc(ecc.Pow(49))
	c1.Calc(ecc.Mul(c2))
	fmt.Printf("12**7 * 77**49 = %d\n", c1.Num)

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
			n, _ := ecc.NewFieldElement(k, prime)
			n.Calc(ecc.Mul(&ecc.FieldElement{Num: i, Prime: prime}))
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
			n.Calc(ecc.Pow(e))
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
