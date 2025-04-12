#include <stdio.h>

int _abs(int);

int main(void)
{
    int r;

    r = _abs(-1);
    printf("%d\n", r);
    r = _abs(0);
    printf("%d\n", r);
    r = _abs(1);
    printf("%d\n", r);
    r = _abs(-98);
    printf("%d\n", r);
    return (0);
}

int _abs(int n)
{
    // function that computes the absolute value of an integer.
    return (n < 0) ? n * -1 : n;
}