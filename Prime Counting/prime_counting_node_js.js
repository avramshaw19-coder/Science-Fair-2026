// Node 20+ recommended

const N = process.argv[2] ? parseInt(process.argv[2], 10) : 300000;

function isPrime(n) {
  if (n < 2) return false;
  if (n % 2 === 0) return n === 2;
  const r = Math.floor(Math.sqrt(n));
  for (let f = 3; f <= r; f += 2) {
    if (n % f === 0) return false;
  }
  return true;
}

const t0 = performance.now();

let count = 0;
let checksum = 0;
const MOD = 1000000007;

for (let x = 1; x <= N; x++) {
  if (isPrime(x)) {
    count++;
    checksum = (checksum + x) % MOD;
  }
}

const seconds = (performance.now() - t0) / 1000.0;
console.log(
  `language=node, workload=primes, n=${N}, primes=${count}, seconds=${seconds.toFixed(6)}, checksum=${checksum}`
);
