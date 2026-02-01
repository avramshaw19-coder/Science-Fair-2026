import math
import sys
import time

def is_prime(n: int) -> bool:
    if n < 2:
        return False
    if n % 2 == 0:
        return n == 2
    r = int(math.isqrt(n))
    f = 3
    while f <= r:
        if n % f == 0:
            return False
        f += 2
    return True

def main():
    n = int(sys.argv[1]) if len(sys.argv) > 1 else 300000
    t0 = time.perf_counter()

    count = 0
    checksum = 0  # lightweight deterministic checksum
    for x in range(1, n + 1):
        if is_prime(x):
            count += 1
            checksum = (checksum + x) % 1_000_000_007

    seconds = time.perf_counter() - t0
    print(f"language=python, workload=primes, n={n}, primes={count}, seconds={seconds:.6f}, checksum={checksum}")

if __name__ == "__main__":
    main()
