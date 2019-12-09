#include <stdlib.h>
#include <stdio.h>

#include "palette_mandmap.h"

void calc_mandelbrot(unsigned int width, unsigned int height, unsigned int maxiter, unsigned char palette[][3])
{
    puts("P3");
    printf("%d %d\n", width, height);
    puts("255");

    double cy = -1.5;
    int y;
    for (y=0; y<height; y++) {
        double cx = -2.0;
        int x;
        for (x=0; x<width; x++) {
            double zx = 0.0;
            double zy = 0.0;
            unsigned int i = 0;
            while (i < maxiter) {
                double zx2 = zx * zx;
                double zy2 = zy * zy;
                if (zx2 + zy2 > 4.0) {
                    break;
                }
                zy = 2.0 * zx * zy + cy;
                zx = zx2 - zy2 + cx;
                i++;
            }
            unsigned char *color = palette[i];
            unsigned char r = *color++;
            unsigned char g = *color++;
            unsigned char b = *color;
            printf("%d %d %d\n", r, g, b);
            cx += 3.0/width;
        }
        cy += 3.0/height;
    }
}

#define BUFFER_SIZE 1*1024*1024
char buffer[BUFFER_SIZE];

int main(int argc, char **argv)
{
    setvbuf(stdout, buffer, _IOFBF, BUFFER_SIZE);
    if (argc < 4) {
        puts("usage: ./mandelbrot width height maxiter");
        return 1;
    }
    int width = atoi(argv[1]);
    int height = atoi(argv[2]);
    int maxiter = atoi(argv[3]);
    calc_mandelbrot(width, height, maxiter, palette);
    return 0;
}

