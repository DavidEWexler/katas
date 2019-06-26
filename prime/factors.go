package prime

func Factors(n int) (primes []int) {
	for c := 2; c*c <= n; c++ {
		for n%c == 0 {
			primes = append(primes, c)
			n /= c
		}
	}
	if n > 1 {
		primes = append(primes, n)
	}
	return primes
}
