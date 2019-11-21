// Seriál "Programovací jazyk Go"
//
// Osmnáctá část
//
// Demonstrační příklad číslo 10:
//     Testovaný balíček.

package factorial

func Factorial(n int64) int64 {
	switch {
	case n < 0:
		return 1
	case n == 0:
		return 1
	default:
		return n * Factorial(n-1)
	}
}
