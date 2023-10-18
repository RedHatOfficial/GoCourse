#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>

#include "so1.h"

int main()
{
    void *library;
    void (*hello)();

    library = dlopen("./so1.so", RTLD_LAZY);
    if (library != NULL) {
        printf("dynamic library loaded: %p\n", library);
    } else {
        puts("unable to load dynamic library");
        return 1;
    }

    hello = dlsym(library, "hello");
