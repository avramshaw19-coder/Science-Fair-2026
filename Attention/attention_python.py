import math
import sys
import time

MOD = 1_000_000_007

class LCG32:
    # state evolves mod 2^32
    def __init__(self, seed: int):
        self.state = seed & 0xFFFFFFFF

    def next_u32(self) -> int:
        self.state = (1664525 * self.state + 1013904223) & 0xFFFFFFFF
        return self.state

    def next_f64(self) -> float:
        # Use top 24 bits to create a float in [-1, 1)
        u = self.next_u32()
        frac = ((u >> 8) & 0xFFFFFF) / 16777216.0  # 2^24
        return frac * 2.0 - 1.0

def main():
    T = int(sys.argv[1]) if len(sys.argv) > 1 else 256
    D = int(sys.argv[2]) if len(sys.argv) > 2 else 64

    rng = LCG32(123456789)

    # Q, K, V: T x D
    Q = [[rng.next_f64() for _ in range(D)] for _ in range(T)]
    K = [[rng.next_f64() for _ in range(D)] for _ in range(T)]
    V = [[rng.next_f64() for _ in range(D)] for _ in range(T)]

    inv_sqrt_d = 1.0 / math.sqrt(D)

    scores = [0.0] * T
    weights = [0.0] * T
    out = [0.0] * D

    checksum = 0

    t0 = time.perf_counter()

    for i in range(T):
        # 1) scores[j] = dot(Q[i], K[j]) / sqrt(D)
        qi = Q[i]
        max_s = -1e300
        for j in range(T):
            kj = K[j]
            s = 0.0
            for k in range(D):
                s += qi[k] * kj[k]
            s *= inv_sqrt_d
            scores[j] = s
            if s > max_s:
                max_s = s

        # 2) softmax
        denom = 0.0
        for j in range(T):
            w = math.exp(scores[j] - max_s)
            weights[j] = w
            denom += w
        inv_denom = 1.0 / denom

        # 3) out[d] = sum_j (weights[j]/denom) * V[j][d]
        for d in range(D):
            out[d] = 0.0
        for j in range(T):
            a = weights[j] * inv_denom
            vj = V[j]
            for d in range(D):
                out[d] += a * vj[d]

        # checksum fold
        for d in range(D):
            q = int(round(out[d] * 1_000_000.0))
            checksum = (checksum + q) % MOD

    seconds = time.perf_counter() - t0

    print(f"language=python, workload=attention, t={T}, d={D}, seconds={seconds:.6f}, checksum={checksum}")

if __name__ == "__main__":
    main()
