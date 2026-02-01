// Node 20+ recommended
const T = process.argv[2] ? parseInt(process.argv[2], 10) : 256;
const D = process.argv[3] ? parseInt(process.argv[3], 10) : 64;

const MOD = 1000000007;

// 32-bit LCG with correct overflow using Math.imul
class LCG32 {
  constructor(seed) {
    this.state = seed >>> 0;
  }
  nextU32() {
    // state = (1664525*state + 1013904223) mod 2^32
    this.state = (Math.imul(1664525, this.state) + 1013904223) >>> 0;
    return this.state;
  }
  nextF64() {
    const u = this.nextU32();
    const frac = ((u >>> 8) & 0xFFFFFF) / 16777216.0; // 2^24
    return frac * 2.0 - 1.0;
  }
}

const rng = new LCG32(123456789);

// Use typed arrays for speed/consistency
function makeMatrix(rows, cols) {
  const m = new Array(rows);
  for (let i = 0; i < rows; i++) m[i] = new Float64Array(cols);
  return m;
}

const Q = makeMatrix(T, D);
const K = makeMatrix(T, D);
const V = makeMatrix(T, D);

for (let i = 0; i < T; i++) for (let d = 0; d < D; d++) Q[i][d] = rng.nextF64();
for (let i = 0; i < T; i++) for (let d = 0; d < D; d++) K[i][d] = rng.nextF64();
for (let i = 0; i < T; i++) for (let d = 0; d < D; d++) V[i][d] = rng.nextF64();

const invSqrtD = 1.0 / Math.sqrt(D);

const scores = new Float64Array(T);
const weights = new Float64Array(T);
const out = new Float64Array(D);

let checksum = 0;

const t0 = performance.now();

for (let i = 0; i < T; i++) {
  const qi = Q[i];

  // 1) scores
  let maxS = -1e300;
  for (let j = 0; j < T; j++) {
    const kj = K[j];
    let s = 0.0;
    for (let k = 0; k < D; k++) s += qi[k] * kj[k];
    s *= invSqrtD;
    scores[j] = s;
    if (s > maxS) maxS = s;
  }

  // 2) softmax
  let denom = 0.0;
  for (let j = 0; j < T; j++) {
    const w = Math.exp(scores[j] - maxS);
    weights[j] = w;
    denom += w;
  }
  const invDenom = 1.0 / denom;

  // 3) output
  out.fill(0.0);
  for (let j = 0; j < T; j++) {
    const a = weights[j] * invDenom;
    const vj = V[j];
    for (let d = 0; d < D; d++) out[d] += a * vj[d];
  }

  // checksum
  for (let d = 0; d < D; d++) {
    const q = Math.round(out[d] * 1e6);
    checksum = (checksum + q) % MOD;
  }
}

const seconds = (performance.now() - t0) / 1000.0;
console.log(`language=node, workload=attention, t=${T}, d=${D}, seconds=${seconds.toFixed(6)}, checksum=${checksum}`);
