using System;
using System.Diagnostics;

class Program
{
    const long MOD = 1_000_000_007;

    struct Lcg32
    {
        public uint State;
        public Lcg32(uint seed) { State = seed; }

        public uint NextU32()
        {
            unchecked
            {
                State = 1664525u * State + 1013904223u;
                return State;
            }
        }

        public double NextF64()
        {
            uint u = NextU32();
            uint top24 = (u >> 8) & 0xFFFFFFu;
            double frac = top24 / 16777216.0; // 2^24
            return frac * 2.0 - 1.0;
        }
    }

    static void Main(string[] args)
    {
        int T = (args.Length > 0 && int.TryParse(args[0], out int t)) ? t : 256;
        int D = (args.Length > 1 && int.TryParse(args[1], out int d)) ? d : 64;

        var rng = new Lcg32(123456789u);

        double[][] Q = new double[T][];
        double[][] K = new double[T][];
        double[][] V = new double[T][];

        for (int i = 0; i < T; i++)
        {
            Q[i] = new double[D];
            K[i] = new double[D];
            V[i] = new double[D];
        }

        for (int i = 0; i < T; i++) for (int j = 0; j < D; j++) Q[i][j] = rng.NextF64();
        for (int i = 0; i < T; i++) for (int j = 0; j < D; j++) K[i][j] = rng.NextF64();
        for (int i = 0; i < T; i++) for (int j = 0; j < D; j++) V[i][j] = rng.NextF64();

        double invSqrtD = 1.0 / Math.Sqrt(D);

        double[] scores = new double[T];
        double[] weights = new double[T];
        double[] outVec = new double[D];

        long checksum = 0;

        var sw = Stopwatch.StartNew();

        for (int i = 0; i < T; i++)
        {
            double[] qi = Q[i];

            // 1) scores
            double maxS = -1e300;
            for (int j = 0; j < T; j++)
            {
                double[] kj = K[j];
                double s = 0.0;
                for (int k = 0; k < D; k++) s += qi[k] * kj[k];
                s *= invSqrtD;
                scores[j] = s;
                if (s > maxS) maxS = s;
            }

            // 2) softmax
            double denom = 0.0;
            for (int j = 0; j < T; j++)
            {
                double w = Math.Exp(scores[j] - maxS);
                weights[j] = w;
                denom += w;
            }
            double invDenom = 1.0 / denom;

            // 3) output
            Array.Clear(outVec, 0, D);
            for (int j = 0; j < T; j++)
            {
                double a = weights[j] * invDenom;
                double[] vj = V[j];
                for (int col = 0; col < D; col++) outVec[col] += a * vj[col];
            }

            // checksum
            for (int col = 0; col < D; col++)
            {
                long q = (long)Math.Round(outVec[col] * 1e6);
                checksum = (checksum + q) % MOD;
            }
        }

        sw.Stop();

        Console.WriteLine($"language=csharp, workload=attention, t={T}, d={D}, seconds={sw.Elapsed.TotalSeconds:F6}, checksum={checksum}");
    }
}
