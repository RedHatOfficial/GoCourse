#include <stdio.h>

int* get_pointer(void) {
    int i = 42;
    return &i;
}

int main(void) {
    int *p = get_pointer();
    printf("%p\n", p);
    printf("%d\n", *p);
    return 0;
}
