#include <stdio.h>

double compute_pi(int n) {
        double pi = 4.0;
        long i;
        for (i = 3; i <= n + 2; i += 2) {
                pi = pi * (i-1)/i * (i+1)/i;
        }
        return pi;
}

int main(void) {
        long n;
        for (n=1; n <= 100000000; n *= 2) {
                printf("%ld %16.14f\n", n, compute_pi(n));
        }
        return 0;

}
