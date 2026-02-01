All versions:

 -- Use pure loops (no NumPy / BLAS / SIMD)
 -- Multiply two N × N matrices
 -- Use the same algorithm + loop order
 -- Accept N from the command line
 -- Print a single standardized output line with a checksum
 -- This keeps the benchmark fair, CPU-only, and scientifically defensible.

Algorithm (same in every language)

We compute:

𝐶 = 𝐴 × 𝐵

Where:

A[i][j] = (i + j) % 100

B[i][j] = (i * j) % 100

Loop order (cache-friendly):

for i
  for k
    for j
      C[i][j] += A[i][k] * B[k][j]


Checksum:

sum(C[i][j]) % 1_000_000_007

Correctness check (important)

For the same N, all languages must output the same checksum.

Start with:

N = 128


Then scale to:

256 → 384 → 512 → 640 → 768


Stop when runtime hits ~30–120 seconds.