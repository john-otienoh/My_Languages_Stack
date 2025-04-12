#include <stdio.h>
#include  <stdlib.h>

int print_last_digit(int);

int main(void)
{
    int r;

    print_last_digit(98);
    print_last_digit(0);
    r = print_last_digit(-1024);
    putchar('0' + r);
    putchar('\n');
    return (0);
}

int print_last_digit(int n)
{
    printf("%d", abs(n%10));
    return (abs(n % 10));
}