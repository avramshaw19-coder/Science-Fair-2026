using System;
using System.Diagnostics;

class Program
{
    const int MOD = 1_000_000_007;

    static void Main(string[] args)
    {
        int N = args.Length > 0 && int.TryParse(args[0], out int v) ? v : 512;

        int[,] A = new int[N, N];
        int[,] B = new int[N, N];
        int[,] C = new int[N, N];

        for (int i = 0; i < N; i++)
            for (int j = 0; j < N; j++)
            {
                A[i, j] = (i + j) % 100;
                B[i, j] = (i * j) % 100;
            }

        var sw = Stopwatch.StartNew();

        for (int i = 0; i < N; i++)
        {
            for (int k = 0; k < N; k++)
            {
                int aik = A[i, k];
                for (int j = 0; j < N; j++)
                {
                    C[i, j] += aik * B[k, j];
                }
            }
        }

        long checksum = 0;
        for (int i = 0; i < N; i++)
            for (int j = 0; j < N; j++)
                checksum = (checksum + C[i, j]) % MOD;

        sw.Stop();

        Console.WriteLine(
            $"language=csharp, workload=matmul, n={N}, seconds={sw.Elapsed.TotalSeconds:F6}, checksum={checksum}"
        );
    }
}
