package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const MOD = 1000000007

func main() {
	N := 512
	if len(os.Args) > 1 {
		if v, err := strconv.Atoi(os.Args[1]); err == nil {
			N = v
		}
	}

	A := make([][]int, N)
	B := make([][]int, N)
	C := make([][]int, N)

	for i := 0; i < N; i++ {
		A[i] = make([]int, N)
		B[i] = make([]int, N)
		C[i] = make([]int, N)
		for j := 0; j < N; j++ {
			A[i][j] = (i + j) % 100
			B[i][j] = (i * j) % 100
		}
	}

	start := time.Now()

	for i := 0; i < N; i++ {
		for k := 0; k < N; k++ {
			aik := A[i][k]
			rowB := B[k]
			rowC := C[i]
			for j := 0; j < N; j++ {
				rowC[j] += aik * rowB[j]
			}
		}
	}

	checksum := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			checksum = (checksum + C[i][j]) % MOD
		}
	}

	seconds := time.Since(start).Seconds()
	fmt.Printf("language=go, workload=matmul, n=%d, seconds=%.6f, checksum=%d\n",
		N, seconds, checksum)
}
