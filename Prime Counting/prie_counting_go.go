package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return n == 2
	}
	r := int(math.Sqrt(float64(n)))
	for f := 3; f <= r; f += 2 {
		if n%f == 0 {
			return false
		}
	}
	return true
}

func main() {
	N := 300000
	if len(os.Args) > 1 {
		if v, err := strconv.Atoi(os.Args[1]); err == nil {
			N = v
		}
	}

	start := time.Now()

	count := 0
	checksum := 0
	const MOD = 1000000007

	for x := 1; x <= N; x++ {
		if isPrime(x) {
			count++
			checksum = (checksum + x) % MOD
		}
	}

	seconds := time.Since(start).Seconds()
	fmt.Printf("language=go, workload=primes, n=%d, primes=%d, seconds=%.6f, checksum=%d\n",
		N, count, seconds, checksum)
}
