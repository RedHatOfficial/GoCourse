#include <stdio.h>
// WARNING: The get_pointer() function below contains UNDEFINED BEHAVIOR in C.
// It returns a pointer to a local variable that is destroyed when the function exits.
// This example is intentionally shown to demonstrate how Go differs from C:
// Go handles this safely using automatic memory management,
// while C results in undefined behavior if the pointer is dereferenced.


int* get_pointer(void) {
    int i = 42;
    return &i;
}

// Note: Running this program may appear to work on some systems due to
// undefined behavior, but it can crash or produce garbage output.

int main(void) {
    int *p = get_pointer();
    printf("%p\n", p);
    printf("%d\n", *p);
    return 0;
}
