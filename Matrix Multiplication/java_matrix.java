public class MatMul {
    static final int MOD = 1_000_000_007;

    public static void main(String[] args) {
        int N = 512;
        if (args.length > 0) {
            try { N = Integer.parseInt(args[0]); } catch (Exception ignored) {}
        }

        int[][] A = new int[N][N];
        int[][] B = new int[N][N];
        int[][] C = new int[N][N];

        for (int i = 0; i < N; i++) {
            for (int j = 0; j < N; j++) {
                A[i][j] = (i + j) % 100;
                B[i][j] = (i * j) % 100;
            }
        }

        long t0 = System.nanoTime();

        for (int i = 0; i < N; i++) {
            for (int k = 0; k < N; k++) {
                int aik = A[i][k];
                int[] rowB = B[k];
                int[] rowC = C[i];
                for (int j = 0; j < N; j++) {
                    rowC[j] += aik * rowB[j];
                }
            }
        }

        long checksum = 0;
        for (int i = 0; i < N; i++) {
            for (int j = 0; j < N; j++) {
                checksum = (checksum + C[i][j]) % MOD;
            }
        }

        double seconds = (System.nanoTime() - t0) / 1e9;
        System.out.printf("language=java, workload=matmul, n=%d, seconds=%.6f, checksum=%d%n",
                N, seconds, checksum);
    }
}
