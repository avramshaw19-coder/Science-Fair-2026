public class Attention {
    static final long MOD = 1_000_000_007L;

    static class LCG32 {
        int state;
        LCG32(int seed) { state = seed; }
        int nextU32() {
            // 32-bit overflow is intended
            state = state * 1664525 + 1013904223;
            return state;
        }
        double nextF64() {
            int u = nextU32();
            int top24 = (u >>> 8) & 0xFFFFFF;
            double frac = top24 / 16777216.0; // 2^24
            return frac * 2.0 - 1.0;
        }
    }

    public static void main(String[] args) {
        int T = 256;
        int D = 64;
        if (args.length > 0) try { T = Integer.parseInt(args[0]); } catch (Exception ignored) {}
        if (args.length > 1) try { D = Integer.parseInt(args[1]); } catch (Exception ignored) {}

        LCG32 rng = new LCG32(123456789);

        double[][] Q = new double[T][D];
        double[][] K = new double[T][D];
        double[][] V = new double[T][D];

        for (int i = 0; i < T; i++) for (int d = 0; d < D; d++) Q[i][d] = rng.nextF64();
        for (int i = 0; i < T; i++) for (int d = 0; d < D; d++) K[i][d] = rng.nextF64();
        for (int i = 0; i < T; i++) for (int d = 0; d < D; d++) V[i][d] = rng.nextF64();

        double invSqrtD = 1.0 / Math.sqrt(D);

        double[] scores = new double[T];
        double[] weights = new double[T];
        double[] out = new double[D];

        long checksum = 0;

        long t0 = System.nanoTime();

        for (int i = 0; i < T; i++) {
            double[] qi = Q[i];

            // 1) scores
            double maxS = -1e300;
            for (int j = 0; j < T; j++) {
                double[] kj = K[j];
                double s = 0.0;
                for (int k = 0; k < D; k++) s += qi[k] * kj[k];
                s *= invSqrtD;
                scores[j] = s;
                if (s > maxS) maxS = s;
            }

            // 2) softmax
            double denom = 0.0;
            for (int j = 0; j < T; j++) {
                double w = Math.exp(scores[j] - maxS);
                weights[j] = w;
                denom += w;
            }
            double invDenom = 1.0 / denom;

            // 3) output
            for (int d = 0; d < D; d++) out[d] = 0.0;
            for (int j = 0; j < T; j++) {
                double a = weights[j] * invDenom;
                double[] vj = V[j];
                for (int d = 0; d < D; d++) out[d] += a * vj[d];
            }

            // checksum
            for (int d = 0; d < D; d++) {
                long q = Math.round(out[d] * 1e6);
                checksum = (checksum + q) % MOD;
            }
        }

        double seconds = (System.nanoTime() - t0) / 1e9;
        System.out.printf("language=java, workload=attention, t=%d, d=%d, seconds=%.6f, checksum=%d%n",
                T, D, seconds, checksum);
    }
}
