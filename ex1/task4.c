#include <stdio.h>

int main(void)
{
    for (int i=97; i <= 122; i++)
    {
        if (i != 101 & i != 113)
        {
            putchar(i);
        }
    }
    putchar('\n');
    return(0);
}