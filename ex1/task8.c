#include <stdio.h>
int main(void)
{
    printf("Hex(uppercase)\tHex(lowercase)\tDecimal\n");
    for (int i=0; i<=256; i++)
    {
        printf("%X\t%x\t%d\n", i, i, i);
    }
    return(0);
}