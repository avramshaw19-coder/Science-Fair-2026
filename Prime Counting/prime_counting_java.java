public class Prime {
    static boolean isPrime(int n) {
        if (n < 2) return false;
        if (n % 2 == 0) return n == 2;
        int r = (int)Math.sqrt(n);
        for (int f = 3; f <= r; f += 2) {
            if (n % f == 0) return false;
        }
        return true;
    }

    public static void main(String[] args) {
        int N = 300000;
        if (args.length > 0) {
            try { N = Integer.parseInt(args[0]); } catch (Exception ignored) {}
        }

        long t0 = System.nanoTime();

        int count = 0;
        int checksum = 0;
        final int MOD = 1_000_000_007;

        for (int x = 1; x <= N; x++) {
            if (isPrime(x)) {
                count++;
                checksum = (checksum + x) % MOD;
            }
        }

        double seconds = (System.nanoTime() - t0) / 1e9;
        System.out.printf("language=java, workload=primes, n=%d, primes=%d, seconds=%.6f, checksum=%d%n",
                N, count, seconds, checksum);
    }
}
