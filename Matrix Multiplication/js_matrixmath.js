const N = process.argv[2] ? parseInt(process.argv[2], 10) : 512;
const MOD = 1000000007;

// Allocate
const A = Array.from({ length: N }, (_, i) =>
  Array.from({ length: N }, (_, j) => (i + j) % 100)
);
const B = Array.from({ length: N }, (_, i) =>
  Array.from({ length: N }, (_, j) => (i * j) % 100)
);
const C = Array.from({ length: N }, () => Array(N).fill(0));

const t0 = performance.now();

for (let i = 0; i < N; i++) {
  for (let k = 0; k < N; k++) {
    const aik = A[i][k];
    const rowB = B[k];
    const rowC = C[i];
    for (let j = 0; j < N; j++) {
      rowC[j] += aik * rowB[j];
    }
  }
}

let checksum = 0;
for (let i = 0; i < N; i++) {
  for (let j = 0; j < N; j++) {
    checksum = (checksum + C[i][j]) % MOD;
  }
}

const seconds = (performance.now() - t0) / 1000;
console.log(
  `language=node, workload=matmul, n=${N}, seconds=${seconds.toFixed(6)}, checksum=${checksum}`
);
