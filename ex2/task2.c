#include<stdio.h>

int print_sign(int n);

int main(void)
{
    int r;

    r = print_sign(98);
    putchar(',');
    putchar(' ');
    putchar(r + '0');
    putchar('\n');
    r = print_sign(0);
    putchar(',');
    putchar(' ');
    putchar(r + '0');
    putchar('\n');
    r = print_sign(0xff);
    putchar(',');
    putchar(' ');
    putchar(r + '0');
    putchar('\n');
    r = print_sign(-1);
    putchar(',');
    putchar(' ');
    putchar(r + '0');
    putchar('\n');
    return (0);
}

int print_sign(int n)
{
    // function that prints the sign of a number.
    if (n > 0)
    {
        printf("+");
        return (1);
    }
    else if (n < 0)
    {
        printf("-");
        return (-1);
    }
    else
    {
        printf("0");
        return (0);
    }
}