#include <stdio.h>

void print_alphabet_x10(void);

int main(void)
{
    print_alphabet_x10();
    return(0);
}

void print_alphabet_x10(void)
{
    // function that prints 10 times the alphabet, in lowercase
    // followed by a new line.
    for(int i = 0; i < 10; i++)
    {
        for (int i=97; i <= 122; i++)
        {
            putchar(i);
        }
        putchar('\n');
    }
}