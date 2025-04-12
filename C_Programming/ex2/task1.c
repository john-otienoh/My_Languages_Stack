#include <stdio.h>

int _islower(int c);

int main(void)
{
    int r;

    r = _islower('H');
    putchar(r + '0');
    r = _islower('o');
    putchar(r + '0');
    r = _islower(108);
    putchar(r + '0');
    putchar('\n');
    return (0);
}

int _islower(int c)
{
    // function that checks for lowercase character.
    return (c >= 97 & c <= 122) ? (1) : (0);
}