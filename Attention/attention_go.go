package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

const MOD int64 = 1000000007

type LCG32 struct {
	state uint32
}

func NewLCG32(seed uint32) *LCG32 {
	return &LCG32{state: seed}
}

func (r *LCG32) nextU32() uint32 {
	r.state = 1664525*r.state + 1013904223
	return r.state
}

func (r *LCG32) nextF64() float64 {
	u := r.nextU32()
	frac := float64((u >> 8) & 0xFFFFFF) / 16777216.0 // 2^24
	return frac*2.0 - 1.0
}

func main() {
	T := 256
	D := 64
	if len(os.Args) > 1 {
		if v, err := strconv.Atoi(os.Args[1]); err == nil {
			T = v
		}
	}
	if len(os.Args) > 2 {
		if v, err := strconv.Atoi(os.Args[2]); err == nil {
			D = v
		}
	}

	rng := NewLCG32(123456789)

	Q := make([][]float64, T)
	K := make([][]float64, T)
	V := make([][]float64, T)
	for i := 0; i < T; i++ {
		Q[i] = make([]float64, D)
		K[i] = make([]float64, D)
		V[i] = make([]float64, D)
		for d := 0; d < D; d++ {
			Q[i][d] = rng.nextF64()
		}
	}
	for i := 0; i < T; i++ {
		for d := 0; d < D; d++ {
			K[i][d] = rng.nextF64()
		}
	}
	for i := 0; i < T; i++ {
		for d := 0; d < D; d++ {
			V[i][d] = rng.nextF64()
		}
	}

	invSqrtD := 1.0 / math.Sqrt(float64(D))

	scores := make([]float64, T)
	weights := make([]float64, T)
	out := make([]float64, D)

	var checksum int64 = 0

	start := time.Now()

	for i := 0; i < T; i++ {
		qi := Q[i]

		// 1) scores
		maxS := -1e300
		for j := 0; j < T; j++ {
			kj := K[j]
			s := 0.0
			for k := 0; k < D; k++ {
				s += qi[k] * kj[k]
			}
			s *= invSqrtD
			scores[j] = s
			if s > maxS {
				maxS = s
			}
		}

		// 2) softmax
		denom := 0.0
		for j := 0; j < T; j++ {
			w := math.Exp(scores[j] - maxS)
			weights[j] = w
			denom += w
		}
		invDenom := 1.0 / denom

		// 3) output
		for d := 0; d < D; d++ {
			out[d] = 0.0
		}
		for j := 0; j < T; j++ {
			a := weights[j] * invDenom
			vj := V[j]
			for d := 0; d < D; d++ {
				out[d] += a * vj[d]
			}
		}

		// checksum
		for d := 0; d < D; d++ {
			q := int64(math.Round(out[d] * 1e6))
			checksum = (checksum + q) % MOD
		}
	}

	seconds := time.Since(start).Seconds()
	fmt.Printf("language=go, workload=attention, t=%d, d=%d, seconds=%.6f, checksum=%d\n",
		T, D, seconds, checksum)
}
