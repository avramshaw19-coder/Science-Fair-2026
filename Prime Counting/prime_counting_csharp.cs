using System;
using System.Diagnostics;

class Program
{
    static bool IsPrime(int n)
    {
        if (n < 2) return false;
        if (n % 2 == 0) return n == 2;

        int r = (int)Math.Sqrt(n);
        for (int f = 3; f <= r; f += 2)
        {
            if (n % f == 0) return false;
        }
        return true;
    }

    static int ParseN(string[] args, int defaultValue)
    {
        if (args.Length == 0) return defaultValue;
        return int.TryParse(args[0], out int n) ? n : defaultValue;
    }

    static void Main(string[] args)
    {
        int N = ParseN(args, 300000);

        var sw = Stopwatch.StartNew();

        int count = 0;
        int checksum = 0;
        const int MOD = 1_000_000_007;

        for (int x = 1; x <= N; x++)
        {
            if (IsPrime(x))
            {
                count++;
                checksum = (checksum + x) % MOD;
            }
        }

        sw.Stop();

        Console.WriteLine($"language=csharp, workload=primes, n={N}, primes={count}, seconds={sw.Elapsed.TotalSeconds:F6}, checksum={checksum}");
    }
}
