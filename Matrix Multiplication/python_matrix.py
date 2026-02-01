import sys
import time

MOD = 1_000_000_007

def main():
    N = int(sys.argv[1]) if len(sys.argv) > 1 else 512

    # Allocate matrices
    A = [[(i + j) % 100 for j in range(N)] for i in range(N)]
    B = [[(i * j) % 100 for j in range(N)] for i in range(N)]
    C = [[0] * N for _ in range(N)]

    t0 = time.perf_counter()

    for i in range(N):
        for k in range(N):
            aik = A[i][k]
            rowB = B[k]
            rowC = C[i]
            for j in range(N):
                rowC[j] += aik * rowB[j]

    checksum = 0
    for i in range(N):
        for j in range(N):
            checksum = (checksum + C[i][j]) % MOD

    seconds = time.perf_counter() - t0

    print(f"language=python, workload=matmul, n={N}, seconds={seconds:.6f}, checksum={checksum}")

if __name__ == "__main__":
    main()
