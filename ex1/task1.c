#include <stdlib.h>
#include <stdio.h>
#include <time.h>

int main(void)
{
	int n;
    char *prefix = "Last digit of";

	srand(time(0));
	n = rand() - RAND_MAX / 2;

    if (n % 10 > 5)
    {
        printf("%s %d is %d and is greater than 5\n", prefix, n, (n % 10));
    }
    else if (n % 10 == 0)
    {
        printf("%s %d is %d and is 0\n", prefix, n , (n%10));
    }
    else
    {
        printf("%s %d is %d and is less than 6 and not 0\n", prefix, n, (n%10));
    }
	return (0);
}