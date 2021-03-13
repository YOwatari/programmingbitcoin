package exercise_test

import (
	"fmt"
	"sort"
	"strings"
)

func ExampleChapter1_two() {
	prime := 57
	sub := func(n int, mod int) int {
		return ((n % prime) + prime) % prime
	}

	fmt.Printf("Prime: %d\n", prime)
	fmt.Printf("44 + 33 = %d\n", (44 + 33) % prime)
	fmt.Printf("9 - 29 = %d\n", sub(9 - 29, prime))
	fmt.Printf("17 + 42 + 49 = %d\n", ((17 + 42) % prime + 49) % prime)
	fmt.Printf("52 - 30 - 38 = %d\n", sub(sub(52 - 30, prime) - 38, prime))

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
	fmt.Printf("95 * 45 * 31 = %d\n", (95 * 45 % prime) * 31 % prime)
	fmt.Printf("17 * 13 * 19 * 44 = %d\n", ((17 * 13 % prime) * 19 % prime) * 44 % prime)
	fmt.Printf("12^7 * 77^49 = %d\n", _pow(12, 7, prime) * _pow(77, 49, prime) % prime)

	// Output:
	// Prime: 97
	// 95 * 45 * 31 = 23
	// 17 * 13 * 19 * 44 = 68
	// 12^7 * 77^49 = 63
}

func ExampleChapter1_five() {
	prime := 19
	for _, k := range []int{1, 3, 7, 13, 18} {
		result := make([]int, prime)
		for i := 0; i < len(result); i++ {
			result[i] = k * i % prime
		}
		sort.Ints(result)
		s := strings.Join(strings.Fields(fmt.Sprintf("%02d", result)), ", ")
		fmt.Printf("k=%02d %s\n", k, s)
	}

	// Output:
	// k=01 [00, 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12, 13, 14, 15, 16, 17, 18]
	// k=03 [00, 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12, 13, 14, 15, 16, 17, 18]
	// k=07 [00, 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12, 13, 14, 15, 16, 17, 18]
	// k=13 [00, 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12, 13, 14, 15, 16, 17, 18]
	// k=18 [00, 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12, 13, 14, 15, 16, 17, 18]
}

func ExampleChapter1_seven() {
	for _, p := range []int{7, 11, 17, 31} {
		e := p - 1
		result := make([]int, 0)
		for i := 1; i < p; i++ {
			result = append(result, _pow(i, e, p))
		}
		s := strings.Join(strings.Fields(fmt.Sprint(result)), ", ")
		fmt.Printf("p=%02d %v\n", p, s)
	}

	// Output:
	// p=07 [1, 1, 1, 1, 1, 1]
	// p=11 [1, 1, 1, 1, 1, 1, 1, 1, 1, 1]
	// p=17 [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]
	// p=31 [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]
}

func _pow(n int, exp int, mod int) int {
	p := 1
	for exp > 0 {
		if exp & 1 == 1 {
			p = (p * n) % mod
		}

		n = (n * n) % mod
		if n == 1 {
			break
		}

		exp >>= 1
	}
	return p
}
